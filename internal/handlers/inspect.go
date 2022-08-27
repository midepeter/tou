package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

func (h *Handler) Inspect(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(body)

	//TODO --- Persisting the messages so it can be validated
	//What i want to be able to do here is retrieve the persisited store and is returned back to the user

	job := h.Queue.Pop()

	log.Printf("Inspection: the %#v ", job)
}
