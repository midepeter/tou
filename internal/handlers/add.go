package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/midepeter/tou/internal/job"
	"github.com/midepeter/tou/internal/queue"
	"go.uber.org/zap"
)

type Handler struct {
	Queue *queue.Queue
	Log   *zap.SugaredLogger
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	var newJob job.Job
	err = json.Unmarshal(body, &newJob)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if err := newJob.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if err := h.Queue.Insert(newJob); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write(body)

	h.Log.Info("Adding a new job %#v", newJob)
}
