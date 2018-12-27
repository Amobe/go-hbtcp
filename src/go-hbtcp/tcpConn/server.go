package tcpConn

import (
	"bufio"
	"io"
	"net"
	"time"

	"go-hbtcp/logger"
)

var (
	pLogger = logger.GetLoggerInstance()
)

const (
	readBufferSize int = 2048
)

// HBConn represent the read/write handler of TCP connection.
type HBConn struct {
	Conn    *net.TCPConn
	writer  *bufio.Writer
	scanner *bufio.Scanner
}

// NewHBConn return a new instance with giving TCP connection.
func NewHBConn(conn *net.TCPConn) *HBConn {
	if conn == nil {
		pLogger.Error("nil pointer of TCPConn instance/n")
		return nil
	}

	writer := bufio.NewWriter(conn)
	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanLines)
	scanner.Buffer(make([]byte, readBufferSize), bufio.MaxScanTokenSize)

	c := &HBConn{
		Conn:    conn,
		writer:  writer,
		scanner: scanner,
	}
	return c
}

// Read analyze the incoming message from reading buffer.
// The handling process will be closed when received 'quit' message or timeout.
// The connection also close after handling process closed.
func (c *HBConn) Read() {
	quitFlag := false
	quitChan := make(chan bool, 1)
	interval := time.NewTicker(10 * time.Millisecond)
	timeout := time.NewTicker(5 * time.Second)
	incomingChan := make(chan string)

	go func() {
		for c.scanner.Scan() {
			msg := c.scanner.Text()
			if len(msg) != 0 {
				incomingChan <- msg
			}
			if msg == "quit" {
				break
			}
		}
		err := c.scanner.Err()
		if err != nil && err != io.EOF {
			pLogger.Error("broken data, %v\n", err)
		}
	}()

readLoop:
	for {
		select {
		case <-timeout.C:
			pLogger.Info("CONN TIMEOUT\n")
			quitFlag = true
			break
		case msg := <-incomingChan:
			pLogger.Info("MSG: %s\n", msg)
		default:
		}
		if quitFlag {
			quitChan <- true
			break readLoop
		}
		select {
		case <-interval.C:
		}
	}

	c.Conn.Close()
	<-quitChan
	pLogger.Info("CONN CLOSE %s\n", c.Conn.RemoteAddr().String())
}

// StartHBServer resolve the address and start a TCP server
func StartHBServer(address string) error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		pLogger.Error("%v/n", err)
	}

	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		pLogger.Error("%v/n", err)
	}
	defer tcpListener.Close()
	pLogger.Info("START LISTEN %s\n", address)

	for {
		conn, err := tcpListener.AcceptTCP()
		if err != nil {
			pLogger.Error("%v/n", err)
			return err
		}
		pLogger.Info("CONN ESTABLISH %s\n", conn.RemoteAddr().String())
		tcpConn := NewHBConn(conn)
		go tcpConn.Read()
	}

	pLogger.Info("STOP LISTEN %s\n", address)
	return nil
}
