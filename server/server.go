package server

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
)

type Server struct {
	srv     *http.Server
	errChan chan error
	done    chan bool
}

func NewServer() *Server {
	srv := &http.Server{}
	return &Server{
		errChan: make(chan error),
		done:    make(chan bool),
		srv:     srv,
	}
}

func (s *Server) Serve(addr string) error {
	s.srv.Addr = addr
	s.srv.Handler = http.DefaultServeMux

	http.Handle("/add", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(body)
		log.Printf("Added data %s", string(body))
	}))

	http.Handle("/inspect", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(body)

		//TODO --- Persisting the messages so it can be validated
		log.Printf("Inspection: %s added a valid data ", string(body))
	}))

	log.Println("Serving work queue server to on the addr ", addr)
	s.errChan <- s.srv.ListenAndServe()

	return <-s.errChan
}

//This is for creating a more tested and secured connection with te database
func (s *Server) ServerWithTLS(addr, certFile, keyFile string) error {
	s.srv.Addr = addr
	s.srv.Handler = http.DefaultServeMux

	s.errChan <- s.srv.ListenAndServeTLS(certFile, keyFile)

	return <-s.errChan
}

func (s *Server) Shutdown(ctx context.Context) error {
	go func() {
		s.done <- true
		if err := s.srv.Shutdown(context.Background()); err != nil {
			s.errChan <- err
		}
	}()
	return <-s.errChan
}
