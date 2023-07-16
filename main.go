package main

import (
	"os"

	"github.com/akamensky/argparse"
	"github.com/retro49/porter/plogger"
	_ "github.com/retro49/porter/scanner"
)

func main() {
	var logger = plogger.NewPlogger()
	parser := argparse.NewParser("porter", ARG_MANUAL)

	argHelp := parser.Flag("h", "help",
		&argparse.Options{
			Required: false,
			Help:     USAGE_HELP,
		},
	)

	argVersion := parser.Flag("v", "version",
		&argparse.Options{
			Required: false,
			Help:     USAGE_VERSION,
		},
	)

	argHost := parser.String("H", "host",
		&argparse.Options{
			Required: false,
			Help:     USAGE_HOST,
			Default:  DEFAULT_HOST,
		},
	)

	argNetwork := parser.Selector("n", "network",
		[]string{"tcp", "udp"},
		&argparse.Options{
			Required: false,
			Default:  DEFAULT_NETWORK,
			Help:     USAGE_NETWORK,
		},
	)

	argPort := parser.Int("p", "port",
		&argparse.Options{
			Required: false,
			Default:  -1,
			Help:     USAGE_PORT,
		},
	)

	argThreads := parser.Int("t", "threads",
		&argparse.Options{
			Required: false,
			Default:  DEFAULT_THREADS,
			Help:     USAGE_THREADS,
		},
	)

	argStart := parser.Int("s", "start",
		&argparse.Options{
			Required: false,
			Default:  DEFAULT_START,
			Help:     USAGE_START,
		},
	)

	argOutput := parser.String("o", "output",
		&argparse.Options{
			Default:  DEFAULT_OUTPUT,
			Required: false,
			Help:     USAGE_OUTPUT,
		},
	)

	argFormat := parser.Selector("f", "format",
		[]string{"json", "normal"},
		&argparse.Options{
			Default:  DEFAULT_FORMAT,
			Required: false,
			Help:     USAGE_FORMAT,
		},
	)

	argEnd := parser.Int("e", "end",
		&argparse.Options{
			Required: false,
			Default:  DEFAULT_END,
			Help:     USAGE_END,
		},
	)

	argStep := parser.Int("S", "step",
		&argparse.Options{
			Default:  1,
			Required: false,
			Help:     USAGE_STEP,
		},
	)

	argRange := parser.String("r", "range",
		&argparse.Options{
			Default:  DEFAULT_RANGE,
			Required: false,
			Help:     USAGE_RANGE,
		},
	)

	argSkip := parser.IntList("k", "skip",
		&argparse.Options{
			Default:  DEFAULT_SKIP,
			Required: false,
			Help:     USAGE_SKIP,
		},
	)

	parser.Parse(os.Args)

	if *argHelp {
		plogger.NewPlogger().Log("HELP", ARG_MANUAL)
	}

        if *argVersion{
            // this is a teporary version...
            // replace the version with a real version.
            logger.Debug("version", "1.0")
        }

	if *argStart < 1 || *argStart > DEFAULT_END {
		logger.Error("start", "given start port is not a valid port")
		os.Exit(2)
	}

	if *argEnd < 1 || *argEnd > DEFAULT_END {
		logger.Error("end", "given end port is not a valid port")
		os.Exit(3)
	}

	if *argStep < 1 || *argStep > DEFAULT_END {
		logger.Error("end", "given end port is not a valid port")
		os.Exit(4)
	}

	// scan single port
	if *argPort != -1 {
		if *argPort < 1 || *argPort > DEFAULT_END {
			logger.Error("port", "given port is not a valid port")
			os.Exit(5)
		}
	} else if *argRange != "" {
		// take it from here okay bud...
		// range scan... do it!

	} else {
		// a normal scan
	}

	logger.Log("host", *argHost)
	logger.Log("network", *argNetwork)
	logger.Log("port", *argPort)
	logger.Log("start", *argStart)
	logger.Log("end", *argEnd)
	logger.Log("range", *argRange)
	logger.Log("step", *argStep)
	logger.Log("format", *argFormat)
	logger.Log("output", *argOutput)
	logger.Log("threads", *argThreads)
	logger.Log("skip", *argSkip)
}
