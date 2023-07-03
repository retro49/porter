package main

import (
	_ "github.com/retro49/porter/plogger"
	"github.com/akamensky/argparse"
)

func main() {
    // argparse is the best...
    parser := argparse.NewParser("porter", ARG_MANUAL)
    arg_porter_opt_help := parser.String("h", "help", &argparse.Options{Required: false, Help: ARG_PORTER_USAGE_HELP})
    arg_porter_opt_host := parser.String("H", "host", &argparse.Options{Required: false, Default: ARG_PORTER_DEFAULT_HOST, Help: ARG_PORTER_USAGE_HOST})
    arg_porter_opt_network := parser.Selector("n", "network", []string{"tcp", "udp", "ip"}, &argparse.Options{Required: false, Default: "tcp", Help: ARG_PORTER_USAGE_NETWORK})
    arg_porter_opt_start := parser.Int("s", "start", &argparse.Options{Required: false, Default: ARG_PORTER_DEFAULT_START, Help: ARG_PORTER_USAGE_START})
    arg_porter_opt_end := parser.Int("e", "end", &argparse.Options{Required: false, Default: ARG_PORTER_DEFAULT_END, Help: ARG_PORTER_USAGE_END})
    arg_porter_opt_range := parser.String("r", "range", &argparse.Options{Required: false, Default: "1", Help: ARG_PORTER_USAGE_RANGE})
    arg_porter_opt_skip := parser.IntList("k", "skip", &argparse.Options{Required: false, Default: []int{}, Help: ARG_PORTER_USAGE_SKIP})
    arg_porter_opt_step := parser.Int("S", "step", &argparse.Options{Required: false, Default: 1, Help: ARG_PORTER_USAGE_STEP})
    arg_poter_opt_port := parser.Int("p", "port", &argparse.Options{Required: false, Default: 1, Help: ARG_PORTER_USAGE_PORT})
    arg_porter_opt_output := parser.String("o", "output", &argparse.Options{Required: false, Help: ARG_PORTER_USAGE_OUTPUT, Default: "stdout"})
    arg_porter_opt_format := parser.Selector("f", "format", []string{"normal", "json"}, &argparse.Options{Required: false, Default: "normal", Help: ARG_PORTER_USAGE_FORMAT})
    arg_porter_opt_scan_mode := parser.Selector("m", "mode", []string{"normal", "fast"}, &argparse.Options{Required: false, Default: "normal", Help: ""})
}
