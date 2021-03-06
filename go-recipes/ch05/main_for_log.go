package main

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Package level variables, which are pointers to log.Logger.
var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

// initLog initializes log.Logger objects
func initLog(traceHandle, infoHandle, warningHandle, errorHandle io.Writer) {
	// Flags for defineing the logging properties, to log.New
	flag := log.Ldate | log.Ltime | log.Lshortfile
	// Create log.Logger objects
	Trace = log.New(traceHandle, "trace: ", flag)
	Info = log.New(infoHandle, "info: ", flag)
	Warning = log.New(warningHandle, "warning: ", flag)
	Error = log.New(errorHandle, "error: ", flag)
}

func main() {
	initLog(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	Trace.Println("Main started")
	loop()
	err := errors.New("Sample Error")
	Error.Println(err.Error())
	Trace.Println("Main completed")
}

func loop() {
	Trace.Println("Loop started")
	for i := 0; i < 10; i++ {
		Info.Println("Counter value is:", i)
	}
	Warning.Println("The counter variable is not being used")
	Trace.Println("Loop completed")
}
