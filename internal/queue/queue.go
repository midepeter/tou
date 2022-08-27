package queue

import (
	"sync"

	"container/heap"

	"github.com/midepeter/tou/internal/job"
)

type Queue struct {
	mu   sync.Mutex
	id   string
	elem []job.Job //
}

func NewQueue(id string) *Queue {

	q := &Queue{
		id:   id,
		elem: make([]job.Job, 10),
	}

	heap.Init(q)
	return q
}

func (q *Queue) Len() int {
	return len(q.elem)
}

func (q *Queue) Less(i, j int) bool {
	if q.elem[i].PriorityIdx < q.elem[j].PriorityIdx {
		return true
	}

	return false
}

func (q *Queue) Swap(i, j int) {
	q.elem[i], q.elem[j] = q.elem[j], q.elem[i]
}

func (q *Queue) Push(x any) {
	elem, ok := x.(job.Job)
	if !ok {
		panic("Unable to assert the type")
	}
	q.elem = append(q.elem, elem)
}

func (q *Queue) Pop() any {
	popVal := q.elem[0]
	q.elem = q.elem[1:]

	return popVal
}
