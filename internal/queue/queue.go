package queue

import (
	"fmt"

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
	id       string
	elem     []string//
	size     uint64    //Size signifies the size of jobs each queue can take
	shutdown bool
}

func NewQueue(id string) *Queue {
	return &Queue{
		id:   id,
		size: 10,
		elem: make([]string, 0),
	}
}

func (q *Queue) Get() string {

	job := q.elem[len(q.elem)-1]

	if len(q.elem) > 0 {
		q.elem = q.elem[1:]
	}

	return job
}

func (q *Queue) Insert(t string) error {

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
