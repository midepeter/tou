package handlers

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/midepeter/tou/internal/queue"
)

type Handler struct {
	Queue *queue.Queue
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	// var newJob job.Job
	// err = json.Unmarshal(body, &newJob)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// }

	// if err := newJob.Validate(); err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// }

	if err := h.Queue.Insert(string(body)); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write(body)

	log.Printf("Added data %s", string(body))
}
