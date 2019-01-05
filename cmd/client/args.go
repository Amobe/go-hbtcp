package main

import (
	"flag"
)

var gProcConfig ProcConfig

// ProcConfig represent the configuration of process
type ProcConfig struct {
	// Amount represent the number of testing client
	Amount int
	// DestAddr represent the address of target TCP server
	DestAddr string
	// ConnTimeout represent the time in second.
	// Server will close the connection if there is no outgoing message over this time.
	ConnTimeout int
}

func parseArgs() {
	const (
		shortHandStr = " (shorthand)"

		shortCmdAmount = "a"
		longCmdAmount  = "amount"
		defaultAmount  = 1
		usageAmount    = "Number of testing client"

		shortCmdDestAddr = "d"
		longCmdDestAddr  = "destAddr"
		defaultDestAddr  = "127.0.0.1:20000"
		usageDestAddr    = "Dest TCP server address"
	)
	flag.IntVar(&gProcConfig.Amount, shortCmdAmount, defaultAmount, usageAmount+shortHandStr)
	flag.IntVar(&gProcConfig.Amount, longCmdAmount, defaultAmount, usageAmount)
	flag.StringVar(&gProcConfig.DestAddr, shortCmdDestAddr, defaultDestAddr, usageDestAddr+shortHandStr)
	flag.StringVar(&gProcConfig.DestAddr, longCmdDestAddr, defaultDestAddr, usageDestAddr)
	flag.Parse()
}
