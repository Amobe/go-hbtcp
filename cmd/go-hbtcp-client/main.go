package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/amobe/go-hbtcp/logger"
	"github.com/amobe/go-hbtcp/tcpConn"
)

var (
	pLogger = logger.GetLoggerInstance()
)

func main() {
	rc := mainWithCode()
	os.Exit(rc)
}

func mainWithCode() int {
	sigs := make(chan os.Signal, 1)
	stopChan := make(chan bool)
	parseArgs()

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		stopChan <- true
	}()

	for i := 0; i < gProcConfig.Amount; i++ {
		go tcpConn.StartHBClient(gProcConfig.DestAddr, 0)
	}

	<-stopChan
	return 0
}
