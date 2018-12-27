package main

import (
	"os"

	"go-hbtcp/logger"
	"go-hbtcp/tcpConn"
)

var (
	pStats  Stats
	pLogger = logger.GetLoggerInstance()
)

func main() {
	rc := mainWithCode()
	os.Exit(rc)
}

func mainWithCode() int {
	parseArgs()

	tcpConn.StartHBServer(gProcConfig.ListenAddr)

	return 0
}
