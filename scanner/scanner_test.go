package scanner_test

import (
	"fmt"
	"testing"

	"github.com/retro49/porter/plogger"
	"github.com/retro49/porter/scanner"
)

func TestWellKnownPorts(t *testing.T){
    plogger.NewPlogger().Debug("scan", "started scanning...")
    ports := make([]int, 0)
    for i := 0; i < 1 << 16; i++{
        ports = append(ports, i)
    }
    scn, err := scanner.NewScanner("tcp", "192.168.56.101", ports)
    if err != nil {
        t.FailNow()
    }

    openPorts := scn.ScanWithInfo()
    for _, prt := range openPorts{
        fmt.Printf("%-6d %s  %-15s\n", prt.GetPort(), "tcp", prt.GetName())
    }
}
