package extConn

import (
	"time"

	"go-hbtcp/queue"
)

const (
	// requestLimitNumber represent the limit request number during a time interval
	requestLimitNumber int = 30
	// requestLimitInterval represent the time duration in millisecond
	requestLimitInterval int = 1000
)

// RequestJob is an implementation of the job interface,
// which can forward the message to the external API.
type RequestJob struct {
	ID      int
	Message string
}

// Do forward the message to the external API
// Do implement the Job.Do method
func (rj *RequestJob) Do() {
	send(rj)
}

// RequestQueue represent a job queue with a exection rate limitation
// RequestQueue inherits from queue.JobQueue
type RequestQueue struct {
	*queue.JobQueue
	requestLimitNumber   int
	requestLimitInterval int
}

// NewRequestQueue create an instance of request queue and return it
func NewRequestQueue() *RequestQueue {
	requestQueue := &RequestQueue{
		queue.NewJobQueue(),
		requestLimitNumber,
		requestLimitInterval,
	}
	return requestQueue
}

// Start turn on the handling process of the queue.
func (rq *RequestQueue) Start() {
	go rq.Handler()
}

// Handler pop the oldest job from queue and execute the job.
// The executor only allow to execute 'requestLimitNumber' request in each 'requestLimitInterval' ms.
// RequestQueue.Handler override the JobQueue.Handler method
func (rq *RequestQueue) Handler() {
	duration := time.Duration(rq.requestLimitInterval) * time.Millisecond
	ticker := time.NewTicker(duration)
	jobCounter := 0
	waiter := make(chan bool)

	go func() {
		for range ticker.C {
			pLogger.Debug("Queue limit reset by ticker")
			if jobCounter >= rq.requestLimitNumber {
				// The handling process will be blocked if the counter hit the limit.
				// Push the signal into waiter to active the handling process.
				waiter <- true
			}
			jobCounter = 0
		}
	}()

	for job := range rq.GetQueue() {
		if rq.IsStop() {
			break
		}
		job.Do()
		if jobCounter++; jobCounter >= rq.requestLimitNumber {
			// Block the handling process if the counter reach the limit
			pLogger.Info("RequestQueue reach the request limit, block the handling process")
			<-waiter
		}
	}

	ticker.Stop()
	pLogger.Info("RequestQueue EXIT")
}
