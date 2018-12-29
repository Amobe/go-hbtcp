package extConn

import (
	"go-hbtcp/logger"
	"go-hbtcp/queue"
	"time"
)

var (
	pLogger   = logger.GetLoggerInstance()
	pJobQueue = queue.StartDispatcher()
)

// MessageJob is an implementation of the job interface,
// which can forward the message to the external API.
type MessageJob struct {
	ID      int
	Message string
}

// Do forward the message to the external API
// Do implement the Job.Do method
func (mj *MessageJob) Do() {
	pLogger.Info("JOB %d EXC %s\n", mj.ID, mj.Message)
	send(mj.Message)
}

// send transmit the message to the external API.
func send(msg string) {
	pLogger.Info("EXT_O %s\n", msg)
	time.Sleep(2 * time.Second)
}

// ForwardMessage package the message become a MessageJob and insert into the queue.
func ForwardMessage(msg string) {
	pJobQueue.InsertJob(&MessageJob{0, msg})
}
