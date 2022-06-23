package queue

import (
	"log"
	"sync"

	"github.com/midepeter/tou/internal/job"
)

type Interface interface {
	//Get() gets job from the queue
	Get() job.Job
	Insert(t job.Job) error
	Shutdown() error
}

type queue struct {
	mu       *sync.Mutex
	elem     []job.Job
	shutdown bool
}

func (q queue) Get() job.Job {
	job := q.elem[0]

	q.elem = q.elem[1:]
	return job
}

func (q queue) Insert(t job.Job) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.elem = append(q.elem, t)
	return nil
}

func (q queue) Shutdown() error {
	if !q.shutdown {
		q.elem = nil
		log.Println("Shutting down queue")
	}

	return nil
}
