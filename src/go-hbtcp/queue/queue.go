package queue

import (
	"errors"

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

// Queue represent a queue contains many jobs which are waiting for executing.
type Queue interface {
	Start()
	Stop()
	Handler()
	Insert(j Job) error
}

// JobQueue represent a queue which can executes the jobs one by one.
type JobQueue struct {
	queue    chan Job
	stopChan chan bool
	isStop   bool
}

// NewJobQueue create a job queue instance and return it.
func NewJobQueue() *JobQueue {
	queue := &JobQueue{
		make(chan Job, queueSize),
		make(chan bool, 1),
		false,
	}
	return queue
}

// Start turn on the handling process of the queue.
func (jq *JobQueue) Start() {
	go jq.Handler()
}

// Stop turn off the handling process of the queue.
func (jq *JobQueue) Stop() {
	jq.stopChan <- true
	jq.isStop = true
	close(jq.queue)
}

// Handler pop the oldest job from queue and execute the job.
func (jq *JobQueue) Handler() {
	for {
		select {
		case <-jq.stopChan:
			pLogger.Info("JobQueue STOP, leave %d jobs", len(jq.queue))
			return
		case job, ok := <-jq.queue:
			if !ok {
				jq.queue = nil
			} else {
				job.Do()
			}
		}
	}
}

// Insert append the job to the end of the queue.
func (jq *JobQueue) Insert(j Job) error {
	if jq.isStop {
		pLogger.Warn("JobQueue cannot insert job into a stopped queue")
		return errors.New("cannot insert job into a stopped queue")
	}
	if j == nil {
		pLogger.Error("JobQueue cannot insert a nil job")
		return errors.New("cannot insert nil job into a queue")
	}
	jq.queue <- j
	return nil
}

// GetQueue return the job channel.
func (jq *JobQueue) GetQueue() chan Job {
	return jq.queue
}

// IsStop return the status of stopping flag.
func (jq *JobQueue) IsStop() bool {
	return jq.isStop
}
