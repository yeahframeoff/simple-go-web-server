package main

import (
	"os"

	"github.com/withmandala/go-log"
)

var (
	errorLog *log.Logger
	infoLog  *log.Logger
)

func setupLoggers() {
	errorLog = log.New(os.Stderr)
	infoLog = log.New(os.Stdout)
}
