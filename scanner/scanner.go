package scanner

import (
	"errors"
	"fmt"
	"net"
	_ "net"
	"time"

	"github.com/retro49/porter/plogger"
)

const (
	SCANNER_NETWROK_TYPE_UDP string = "udp"
	SCANNER_NETWORK_TYPE_TCP string = "tcp"
	SCANNER_LOCAL_HOST              = "127.0.0.1"
)

var ERROR_INVALID_PORT_NUMBER = errors.New("invalid port number")

/*
// creates a new sequential scanner
func NewScanner(network, host string, ports []int) (*scanner, error){
    if network != SCANNER_NETWORK_TYPE_TCP && network != SCANNER_NETWROK_TYPE_UDP {
        panic("error network type provided, network protocol must be ether tcp or udp")
    }

    if host == "" {
        host = SCANNER_LOCAL_HOST
    }

    for _, port := range ports{
        if port < 0 || port >= (1 << 16) {
            return nil, ERROR_INVALID_PORT_NUMBER
        }
    }

    return &scanner{
        ports: ports,
        host: host,
        network: network,
    }, nil
}

func (s scanner)scanPorts(response chan any){
    ports := make([]int, 0)
    for _, port := range s.ports {
        _, err := net.Dial(s.network, fmt.Sprintf("%s:%d", s.host, port))
        if err == nil{
            ports = append(ports, port)
        }
    }
    response<-ports
}

// scans the network
func (s scanner)ScanWithInfo() []portInfo{
    portScannerChannel := make(chan interface{})
    jsonLoaderChannel := make(chan interface{})
    go s.scanPorts(portScannerChannel)
    go LoadPortInfo(jsonLoaderChannel)

    jsonResponse := <- jsonLoaderChannel
    portResponse := <- portScannerChannel

    if jsonResponse == nil {
        plogger.NewPlogger().Error("error json response" , "nil json result")
    }

    jsonPortInfo := jsonResponse.(map[string]map[string]string)
    scannedPorts := portResponse.([]int)
    result := make([]portInfo, 0)

    // load into result
    for _, scannedPort := range scannedPorts{
        var key string = fmt.Sprintf("%d/%s", scannedPort, s.network)
        var pf portInfo
        if info, found := jsonPortInfo[key]; !found{
            pf = NewPortInfo("", "", scannedPort)
        } else {
            pf = NewPortInfo(info["name"], info["description"], scannedPort)
        }
        result = append(result, pf)
    }

    return result
}
*/

func fromInfo(s ScanCoordinator) []int {
	var ports []int
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
	var amt int
	var total int = len(ports)
	if total < threads {
		amt = total
	} else {
		amt = threads
	}
	result := make([][]int, amt)
	var i int = 0
	for _, port := range ports {
		if i > amt-1 {
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
	// create the ports from the provided info
	ports := fromInfo(s)
	threads := fromThreads(s.Info.GetThreads(), ports)
	threadChannel := make(chan []int, s.Info.GetThreads())
	for _, thread := range threads {
		// create a new scanner
		scnr := scanner{
			network: s.Info.GetNetwork(),
			host:    s.Info.GetHost(),
			timeout: time.Duration(s.Info.GetTimeout()),
			ports:   thread,
		}
		go scnr.scan(threadChannel)
	}
	// retrive the scanned result
	for i := 0; i < s.Info.GetThreads(); i++ {
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
