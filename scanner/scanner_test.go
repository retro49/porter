package scanner_test

import (
	"testing"
	_"github.com/retro49/porter/plogger"
	"github.com/retro49/porter/scanner"
)

func TestInfo(t *testing.T){
    info := scanner.ScanInfo{
        StartPort: 1,
        EndPort: 65535,
        Step: 1,
        Skip: []int{},
        Threads: 10,
    }

    cordinator := scanner.ScanCoordinator{
        Info: info,
    }
    cordinator.StartScan()
}

