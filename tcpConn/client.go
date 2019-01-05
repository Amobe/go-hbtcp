package tcpConn

import (
	"fmt"
	"net"
	"time"
)

// StartHBClient resolve the address and connect to a TCP server.
func StartHBClient(address string, timeout int) error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		pLogger.Error("%v/n", err)
		return err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		pLogger.Error("%v/n", err)
		return err
	}
	defer conn.Close()

	tcpConn := NewHBConn(conn, timeout)

	counter := 0
	ticker := time.NewTicker(15 * time.Millisecond)
	for range ticker.C {
		msg := fmt.Sprintf("Hello_%d\n", counter)
		tcpConn.Write(msg)
		counter++
	}

	return nil
}
