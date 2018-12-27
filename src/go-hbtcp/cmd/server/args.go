package main

import (
	"flag"
)

var gProcConfig ProcConfig

// ProcConfig represent the configuration of process
type ProcConfig struct {
	// listenAddr represent the listening address of TCP server
	ListenAddr string
}

func parseArgs() {
	const (
		shortHandStr = " (shorthand)"

		shortCmdListenAddr = "l"
		longCmdListenAddr  = "listen"
		defaultListenAddr  = "127.0.0.1:20000"
		usageListenAddr    = "TCP server listening address"
	)
	flag.StringVar(&gProcConfig.ListenAddr, shortCmdListenAddr, defaultListenAddr, usageListenAddr+shortHandStr)
	flag.StringVar(&gProcConfig.ListenAddr, longCmdListenAddr, defaultListenAddr, usageListenAddr)
	flag.Parse()
}
