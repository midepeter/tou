package main

import (
	"fmt"
	"net/http"

	"github.com/midepeter/tou/internal/job"
)

func main() {
	job := NewJob{}

	//A producer should be able to send the job to a queue
	//A cosumer should be able to listen for jobs to process on the queue
	producer := Producer{}

	//so the producer will have to listen to the queues for any available job and process it immediately
	consumer := Consumer{}
}

type producer struct{}

func (p producer) SendJob(job job.Job, url string) (bool, error) {
	if job == nil {
		return false, fmt.Errorf("Invalid Job! Please provide a valid job")
	}

	http.NewRequest("POST", url)
}
