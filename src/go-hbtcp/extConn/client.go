package extConn

import (
	"time"

	"go-hbtcp/logger"
)

var (
	pLogger       = logger.GetLoggerInstance()
	pRequestQueue *RequestQueue
)

// send transmit the message to the external API.
func send(job *RequestJob) {
	pLogger.Info("EXT_O %s\n", job.Message)
	// simulate the execution time of sending request
	time.Sleep(100 * time.Millisecond)
}

// Init setup the global request queue and start it
func Init() {
	if pRequestQueue != nil {
		pLogger.Warn("External Client was already initialized")
		return
	}
	pRequestQueue = NewRequestQueue()
	pRequestQueue.Start()
}

// ForwardMessage package the message become a MessageJob and insert into the queue.
func ForwardMessage(msg string) {
	pRequestQueue.Insert(&RequestJob{0, msg})
}
