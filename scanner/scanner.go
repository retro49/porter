package scanner

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/retro49/porter/plogger"
)

var ERROR_INVALID_PORT_NUMBER = errors.New("invalid port number")

func fromInfo(s ScanCoordinator) []int {
	ports := make([]int, 0)
blk:
	for i := s.Info.GetStart(); i <= s.Info.GetEnd(); i += s.Info.GetStep() {
		for _, skip := range s.Info.GetSkip() {
			if i == skip {
				continue blk
			}
		}
		ports = append(ports, i)
	}
	return ports
}

func fromThreads(threads int, ports []int) [][]int {
	result := make([][]int, threads)
	var i int = 0
	for _, port := range ports {
		if i > threads-1 {
			i = 0
		}
		result[i] = append(result[i], port)
		i += 1
	}
	return result
}

type ScanCoordinator struct {
	Info ScanInfo
}

func (s ScanCoordinator) StartScan() {
	logger := plogger.NewPlogger()
	ports := fromInfo(s)
	thrds := s.Info.GetThreads()
	if len(ports) < thrds {
		thrds = len(ports)
	}

	threads := fromThreads(thrds, ports)
	threadChannel := make(chan []int, thrds)
	for _, thread := range threads {
		scnr := scanner{
			network: s.Info.GetNetwork(),
			host:    s.Info.GetHost(),
			timeout: time.Duration(s.Info.GetTimeout()) * time.Second,
			ports:   thread,
		}
		go scnr.scan(threadChannel)
	}

	// retrive the scanned result
	for i := 0; i < thrds; i++ {
		openPorts := <-threadChannel
		logger.Debug("open ports", openPorts)
	}
}

type scanner struct {
	ports   []int
	network string
	host    string
	timeout time.Duration
}

func (s scanner) scan(ch chan []int) {
	openPorts := make([]int, 0)
	for _, port := range s.ports {
		address := fmt.Sprintf("%s:%d", s.host, port)
		ch := make(chan int)

		go func(add string, c chan int) {
			_, err := net.Dial(s.network, add)
			if err != nil {
				c <- -1
			} else {
				c <- port
			}
		}(address, ch)

		select {
		case <-time.After(s.timeout):
			continue
		case open := <-ch:
			if open != -1 {
				openPorts = append(openPorts, open)
			}
		}
	}
	ch <- openPorts
}
