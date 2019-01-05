package queue

import (
	"testing"
	"time"
)

type MockJob struct {
	isFinish bool
}

func NewMockJob() *MockJob {
	return &MockJob{false}
}

func (m *MockJob) Do() {
	m.isFinish = true
}

type MockJobQueue struct {
	*JobQueue
}

func (m *MockJobQueue) Stop() {
	// waiting handler finished the job
	for len(m.queue) > 0 {
		continue
	}
	m.JobQueue.Stop()
}

func TestNewJobQueue(t *testing.T) {
	jobQueue := NewJobQueue()
	if jobQueue == nil {
		t.Fatalf("jobQueue is nil")
	}
	queue := jobQueue.GetQueue()
	if queue == nil {
		t.Fatalf("jobQueue.queue is nil, initialize fail")
	}
	if len(queue) > 0 {
		t.Errorf("jobQueue.queue is not empty, len: %d", len(queue))
	}
	if cap(queue) != QueueSize {
		t.Errorf("jobQueue.queue capability is not current, expect cap: %d, actaul cap: %d",
			QueueSize, cap(jobQueue.queue))
	}
	if jobQueue.stopChan == nil {
		t.Fatalf("jobQueue.stopChan is nil, initialize fail")
	}
	if jobQueue.IsStop() {
		t.Fatalf("jobQueue.isStop not reset to false, initialize fail")
	}
}

func TestJobQueue(t *testing.T) {
	var err error
	jobQueue := NewJobQueue()
	mockJob1 := NewMockJob()

	jobQueue.Insert(mockJob1)
	if len(jobQueue.queue) != 1 {
		t.Fatalf("jobQueue.queue contains an incorrect num of elements, expect: 1, actual: %d",
			len(jobQueue.queue))
	}
	jobQueue.Start()

	jobWaiter := time.NewTimer(1 * time.Second)
waitLoop:
	for {
		if mockJob1.isFinish {
			break waitLoop
		}
		select {
		case <-jobWaiter.C:
			t.Errorf("Timeout, The job didn't execute")
			break waitLoop
		default:
		}
	}

	err = jobQueue.Insert(nil)
	if err == nil {
		t.Errorf("Should not able to insert a nil job into a queue")
	}
	if err.Error() != "cannot insert nil job into a queue" {
		t.Errorf("incorrect error: %s", err.Error())
	}

	jobQueue.Stop()
	mockJob2 := NewMockJob()
	err = jobQueue.Insert(mockJob2)
	if err == nil {
		t.Errorf("Should not able to insert a job into a stopped queue")
	}
	if err.Error() != "cannot insert job into a stopped queue" {
		t.Errorf("incorrect error: %s", err.Error())
	}
}

func BenchmarkJobQueue(b *testing.B) {
	jobQueue := NewJobQueue()
	mockJobQueue := MockJobQueue{jobQueue}
	mockJobQueue.Start()

	// benchmark the handling process
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mockJobQueue.Insert(NewMockJob())
	}

	mockJobQueue.Stop()
}
