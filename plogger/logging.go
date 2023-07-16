package plogger 

import (
    "log"
    "os"
)

const (
    TERM_COLOR_WARN string = "\033[33m"
    TERM_COLOR_ERROR string = "\033[31m"
    TERM_COLOR_DEBUG string = "\033[32m"
    TERM_COLOR_NORMAL string = "\033[00m"
)

type plogger struct {
    Logger *log.Logger
}

func NewPlogger()plogger{
    logger := log.New(os.Stdout, "", 0)
    logger.SetFlags(3)

    return plogger{
        Logger:  logger, 
    }
}

func (p plogger)Log(key, msg any){
    p.Logger.Printf("%v: %v", key, msg)
}

func (p plogger)Debug(key, msg any){
    p.Logger.Printf("%s%v%s: %v", TERM_COLOR_DEBUG, key, TERM_COLOR_NORMAL, msg)
}

func (p plogger)Warn(key, msg any){
    p.Logger.Printf("%s%v%s: %v", TERM_COLOR_WARN, key, TERM_COLOR_NORMAL, msg)
}

func (p plogger)Error(key, msg any){
    p.Logger.Printf("%s%v%s: %v", TERM_COLOR_ERROR, key, TERM_COLOR_NORMAL, msg)
}
