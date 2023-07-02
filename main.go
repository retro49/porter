package main

import (
    "github.com/retro49/porter/plogger"
)

var (
    ARG_OPT_PORTER_HELP string
    ARG_OPT_PORTER_HOST string
    ARG_OPT_PORTER_NETWORK string
    ARG_OPT_PORTER_START int 
    ARG_OPT_PORTER_END int 
    ARG_OPT_PORTER_RANGE string
    ARG_OPT_PORTER_SKIP int 
    ARG_OPT_PORTER_STEP int
    ARG_OPT_PORTER_PORT int
    ARG_OPT_PORTER_OUTPUT string
    ARG_OPT_PORTER_FORMAT string
)

func main(){
    logger := plogger.NewPlogger()
    logger.Log("head log", "logging")
    logger.Warn("head warn", "warning")
    logger.Error("head error", "error")
    logger.Debug("head debug", "debugging")
}
