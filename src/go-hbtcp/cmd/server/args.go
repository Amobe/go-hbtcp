package main

import (
	"flag"
)

var gProcConfig ProcConfig

// ProcConfig represent the configuration of process
type ProcConfig struct {
	// ListenAddr represent the listening address of TCP server
	ListenAddr string
	// ConnTimeout represent the time in second.
	// Server will close the connection if there is no incoming message over this time.
	ConnTimeout int
}

func parseArgs() {
	const (
		shortHandStr = " (shorthand)"

		shortCmdListenAddr = "l"
		longCmdListenAddr  = "listen"
		defaultListenAddr  = "127.0.0.1:20000"
		usageListenAddr    = "TCP server listening address"

		shortCmdConnTimeout = "t"
		longCmdConnTimeout  = "timeout"
		defaultConnTimeout  = 30
		usageConnTimeout    = "TCP connection timeout in second"
	)
	flag.StringVar(&gProcConfig.ListenAddr, shortCmdListenAddr, defaultListenAddr, usageListenAddr+shortHandStr)
	flag.StringVar(&gProcConfig.ListenAddr, longCmdListenAddr, defaultListenAddr, usageListenAddr)
	flag.IntVar(&gProcConfig.ConnTimeout, shortCmdConnTimeout, defaultConnTimeout, usageConnTimeout+shortHandStr)
	flag.IntVar(&gProcConfig.ConnTimeout, longCmdConnTimeout, defaultConnTimeout, usageConnTimeout)
	flag.Parse()
}
