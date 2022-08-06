package queue

import (
	"fmt"
	"sync"

	"github.com/midepeter/tou/internal/job"
)

type Store interface {
	//Get() gets job from the queue
	Get() job.Job

	//Inserts Item in to the database
	Insert(t job.Job) error

	//Shutdown
	Shutdown() error
}

type Queue struct {
	mu       sync.Mutex
	id       string
	elem     []job.Job //
	size     uint64    //Size signifies the size of jobs each queue can take
	shutdown bool
}

func NewQueue(id string) *Queue {
	return &Queue{
		id:   id,
		size: 10,
		elem: make([]job.Job, 0),
	}
}

func (q *Queue) Get() job.Job {
	q.mu.Lock()
	defer q.mu.Unlock()
	job := q.elem[len(q.elem)-1]

	if len(q.elem) > 0 {
		q.elem = q.elem[1:]
	}

	return job
}

func (q *Queue) Insert(t job.Job) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.elem) > int(q.size) {
		return fmt.Errorf("Unable to add job to queue")
	}

	q.elem = append(q.elem, t)
	return nil
}

func (q *Queue) Shutdown() error {
	if !q.shutdown {
		q = nil
	}

	return nil
}
