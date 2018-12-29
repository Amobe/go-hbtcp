package queue

import (
	"go-hbtcp/logger"
)

var (
	pLogger = logger.GetLoggerInstance()
)

const (
	queueSize int = 1000
)

// Job represent a job which can be put into the queue.
type Job interface {
	Do()
}

// JobQueue represent a queue which can executes the jobs one by one.
type JobQueue struct {
	queue    chan Job
	stopChan chan bool
	isStop   bool
}

// Handler pop the oldest job from queue and execute the job.
func (jq *JobQueue) Handler() {
	for {
		select {
		case <-jq.stopChan:
			pLogger.Info("JobQueue STOP")
			return
		case job := <-jq.queue:
			job.Do()
		}
	}
}

// InsertJob insert the job to the end of the queue.
func (jq *JobQueue) InsertJob(j Job) {
	if jq.isStop {
		pLogger.Warn("JobQueue cannot insert job into a stoped queue")
		return
	}
	if j == nil {
		pLogger.Error("JobQueue cannot insert a nil job")
		return
	}
	jq.queue <- j
}

// Stop turn off the handling process of the queue.
func (jq *JobQueue) Stop() {
	jq.stopChan <- true
	jq.isStop = true
}

// StartDispatcher create a job queue instance and return it.
// The JobQueue will turn on by default.
func StartDispatcher() JobQueue {
	queue := JobQueue{
		make(chan Job, queueSize),
		make(chan bool, 1),
		false,
	}
	go queue.Handler()
	return queue
}
