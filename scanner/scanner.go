package scanner

import (
	"errors"
	"fmt"
	"net"

	"github.com/retro49/porter/plogger"
)

type scanner struct {
    ports []int
    network string
    host string
}

const (
    SCANNER_NETWROK_TYPE_UDP string = "udp"
    SCANNER_NETWORK_TYPE_TCP string = "tcp"
    SCANNER_LOCAL_HOST = "127.0.0.1"
)

var ERROR_INVALID_PORT_NUMBER = errors.New("invalid port number")

// creates a new sequential scanner
func NewScanner(network, host string, ports []int) (*scanner, error){
    if network != SCANNER_NETWORK_TYPE_TCP && network != SCANNER_NETWROK_TYPE_UDP {
        panic("error network type provided, network protocol must be ether tcp or udp")
    }

    if host == "" {
        host = SCANNER_LOCAL_HOST
    }

    for _, port := range ports{
        if port < 0 || port >= 1 << 16 {
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
