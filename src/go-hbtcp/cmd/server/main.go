package main

import (
	"go-hbtcp/admin"
	"os"

	"go-hbtcp/extConn"
	"go-hbtcp/logger"
	"go-hbtcp/tcpConn"
)

var (
	pLogger = logger.GetLoggerInstance()
	pStats  = admin.GetProcStatsInstance()
)

func main() {
	rc := mainWithCode()
	os.Exit(rc)
}

func mainWithCode() int {
	parseArgs()

	admin := admin.NewAdminServer(gProcConfig.AdminAddr, pStats)
	go admin.Start()

	extConn.Init()
	tcpConn.StartHBServer(gProcConfig.ListenAddr, gProcConfig.ConnTimeout)

	return 0
}
