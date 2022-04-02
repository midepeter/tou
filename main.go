package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//get jobs (I can expose this as an endpoint with a post method Where you can get jobs)
//I enqueue the jobs
//A worker is set to always search for the job and executes them

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/jobs", getJobs)

	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

type Job struct {
	ID   string
	From string
	To   string
	Body string
}

func getJobs(w http.ResponseWriter, r *http.Request) {
	var (
		err    error
		job    Job
		writer io.Writer
		q      queue
	)
	if r.Method != "POST" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Invalid method"))
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("No available jobs"))
	}

	err = json.Unmarshal(body, &job)
	if err != nil {
		log.Fatalln(err)
	}

	writer = os.Stdout
	q.Enqueue(job)
	fmt.Fprintf(writer, "The jobs are %s", job)
}

type queue struct{}

var (
	Jobqueue = make([]Job, 5)
)

//Enqueue should return always a nil error
func (q queue) Enqueue(job Job) error {
	var err error

	if job.ID == " " {
		panic(fmt.Sprintf("Job cannot be empty %s", err))
	}

	Jobqueue = append(Jobqueue, job)
	fmt.Println("The queue actually", Jobqueue)
	return nil
}

func (q queue) Dequeue(job Job) (*[]Job, error) {
	var (
		err         error
		newJobqueue []Job
	)

	for k, v := range Jobqueue {
		if v.ID == job.ID {
			newJobqueue = append(newJobqueue, Jobqueue[:k]...)
			newJobqueue = append(newJobqueue, Jobqueue[k+1:]...)
		} else {
			return nil, err
		}
	}
	return &newJobqueue, nil
}

//The worker checks into the queue to see if there are availabe jobs to excute- It should be checking at 2 minutes intervals
type worker struct{}

func (w worker) Worker(job Job) (string, error) {
	result := fmt.Sprintf(
		`From: %v
		To: %v
		Body: %v`,
		job.From, job.To, job.Body,
	)
	return result, nil
}
