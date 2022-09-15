package job

import (
	"time"
)

type JobState int32

const (
	Pending JobState = iota
	Processing
	Successful
)

type Job struct {
	JobId       string    `json:"job_id"`
	QueueId     string    `json:"json:"queue_id"`
	Payload     any       `json:"payload"`
	PriorityIdx int32     `json:"priority_id"`
	State       JobState  `json:"job_state"`
	TTL         time.Time `json:"ttl"` //TTL determines the duration a particular job can live for before it expires
}

func NewJob(id, queueid string, payload any, priorityIdx int32) *Job {
	return &Job{
		JobId:       id,
		QueueId:     queueid,
		Payload:     payload,
		PriorityIdx: priorityIdx,
		State:       0,
		TTL:         time.Now(),
	}
}
func (j Job) Valid() bool {
	if j.JobId == "" || j.QueueId == "" {
		return false
	}
	return true
}
