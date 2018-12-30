package extConn

import (
	"time"

	"go-hbtcp/admin"
	"go-hbtcp/logger"
	"go-hbtcp/queue"
)

var (
	pLogger       = logger.GetLoggerInstance()
	pRequestQueue *RequestQueue
	pStats        = admin.GetProcStatsInstance()
)

// send transmit the message to the external API.
func send(job *RequestJob) {
	pLogger.Info("EXT_O %s\n", job.Message)
	pStats.IncServerOutPktAcc()
	pStats.DecReguestQueuePadding()
	// simulate the execution time of sending request
	time.Sleep(20 * time.Millisecond)
}

// Init setup the global request queue and start it
func Init() {
	if pRequestQueue != nil {
		pLogger.Warn("External Client was already initialized")
		return
	}
	pStats.SetReguestQueueSize(queue.QueueSize)
	pRequestQueue = NewRequestQueue()
	pRequestQueue.Start()
}

// ForwardMessage package the message become a MessageJob and insert into the queue.
func ForwardMessage(msg string) {
	pRequestQueue.Insert(&RequestJob{0, msg})
	pStats.IncReguestQueuePadding()
}
