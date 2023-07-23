package scanner

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/retro49/porter/plogger"
)

var ERROR_INVALID_PORT_NUMBER = errors.New("invalid port number")

func fromInfo(s ScanCoordinator) []int {
	ports := make([]int, 0)

skp:
	for i := s.Info.GetStart(); i <= s.Info.GetEnd(); i += s.Info.GetStep() {
		for _, skip := range s.Info.GetSkip() {
			if i == skip {
				continue skp
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

func (s ScanCoordinator) Write(buff []byte) {
	if s.Info.GetOutput() == "" {
		fmt.Printf("%-20s%-20s%-20s\n", "port", "name", "description")
		fmt.Printf("%-20s%-20s%-20s\n", "----", "----", "-----------")
		os.Stdout.Write(buff)
	} else {
		cwd, err := os.Getwd()
		if err != nil {
			panic("could not find current working directory")
		}
		if cwd[len(cwd)-1] != os.PathSeparator {
			cwd += string(os.PathSeparator)
		}
		of := cwd + s.Info.GetOutput()
		os.WriteFile(of, buff, 0770)
	}
}

func (s ScanCoordinator) ParseInfo(info []portInfo) []byte {
	if s.Info.GetFormat() == "json" {
		// b, err := json.Marshal(info)
		b, err := json.MarshalIndent(info, "", "")
		if err != nil {
			plogger.NewPlogger().Error("decoding", err)
			panic("error occured while decoding...")
		}
		return b
	}

	var buff string
	for _, p := range info {
		buff += fmt.Sprintf("%-20s%-20s%-20s\n", p.GetPort(), p.GetName(), p.GetDescription())
	}

	return []byte(buff)
}

func (s ScanCoordinator) Retrive(
	threads int,
	scanChannel chan []int,
	jsonResult *map[string]map[string]string,
) {
	var out []portInfo = make([]portInfo, 0)
	for i := 0; i < threads; i++ {
		openPorts := <-scanChannel
		if len(openPorts) != 0 {
			strPorts := ToStringArr(openPorts, s.Info.GetNetwork())
			for _, sprt := range strPorts {
				jout := InJson(jsonResult, sprt)
				out = append(out, jout)
			}
		}
	}

	s.Write(s.ParseInfo(out))
}

func (s ScanCoordinator) StartScan() {
	ports := fromInfo(s)
	thrds := s.Info.GetThreads()
	if len(ports) < thrds {
		thrds = len(ports)
	}

	threads := fromThreads(thrds, ports)
	threadChannel := make(chan []int, thrds)

	jsonChannel := make(chan any)
	go LoadPortInfo(jsonChannel)

	for _, thread := range threads {
		scnr := scanner{
			network: s.Info.GetNetwork(),
			host:    s.Info.GetHost(),
			timeout: time.Duration(time.Duration(s.Info.GetTimeout()) * time.Second),
			ports:   thread,
		}
		go scnr.scan(threadChannel)
	}

	jsonChannelResult := <-jsonChannel
	jr := jsonChannelResult.(*map[string]map[string]string) // json result
	s.Retrive(thrds, threadChannel, jr)
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
