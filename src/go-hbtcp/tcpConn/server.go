package tcpConn

import (
	"bufio"
	"io"
	"net"
	"time"

	"go-hbtcp/extConn"
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
	Conn       *net.TCPConn
	writer     *bufio.Writer
	scanner    *bufio.Scanner
	timeout    int
	closeTimer *time.Timer
}

// NewHBConn return a new instance with giving TCP connection.
func NewHBConn(conn *net.TCPConn, timeout int) *HBConn {
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
		timeout: timeout,
	}
	return c
}

// Read analyze the incoming message from reading buffer.
// The handling process will be closed when received 'quit' message or timeout.
// The connection also close after handling process closed.
func (c *HBConn) Read() {
	quitFlag := false
	quitChan := make(chan bool, 1)
	incomingChan := make(chan string)
	interval := time.NewTicker(10 * time.Millisecond)
	var closeTimer *time.Timer

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

	closeTimer = c.resetCloseTimer()
readLoop:
	for {
		select {
		case <-closeTimer.C:
			pLogger.Info("CONN TIMEOUT\n")
			quitFlag = true
			break
		case msg := <-incomingChan:
			c.resetCloseTimer()
			pLogger.Info("TCP_I %s\n", msg)
			extConn.ForwardMessage(msg)
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

	c.close()
	interval.Stop()
	<-quitChan
	pLogger.Info("CONN CLOSE %s\n", c.Conn.RemoteAddr().String())
}

// resetCloseTimer reset the time and return the timer.
// The timer will be create at first if not exists.
func (c *HBConn) resetCloseTimer() *time.Timer {
	duration := time.Duration(c.timeout) * time.Second
	if c.closeTimer == nil {
		pLogger.Info("CONN NEW TIMER %ds\n", c.timeout)
		c.closeTimer = time.NewTimer(duration)
	} else {
		pLogger.Info("CONN RESET TIMER %ds\n", c.timeout)
		c.closeTimer.Reset(duration)
	}
	return c.closeTimer
}

// close terminate the connection and stop the close timer if exists.
func (c *HBConn) close() {
	c.Conn.Close()
	if c.closeTimer != nil {
		c.closeTimer.Stop()
	}
}

// StartHBServer resolve the address and start a TCP server.
func StartHBServer(address string, timeout int) error {
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
		tcpConn := NewHBConn(conn, timeout)
		go tcpConn.Read()
	}

	pLogger.Info("STOP LISTEN %s\n", address)
	return nil
}
