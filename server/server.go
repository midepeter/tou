package server

import (
	"context"
	"log"
	"net/http"
)

type Server struct {
	srv *http.Server	
	errChan chan error 
	done chan bool
}

func NewServer() *Server {
	srv := &http.Server{}
	return &Server{
		errChan: make(chan error),
		done: make(chan bool),
		srv: srv,
	}
}

func (s *Server) Serve(addr string) error {
	s.srv.Addr = addr 
	s.srv.Handler = http.DefaultServeMux

	log.Println("Serving work queue server to on the addr: ", addr)
	s.errChan <- s.srv.ListenAndServe()

	return <- s.errChan
} 

//This is for creating a more tested and secured connection with te database
func (s *Server) ServerWithTLS(addr, certFile, keyFile string) error {
	s.srv.Addr = addr 
	s.srv.Handler = http.DefaultServeMux

	s.errChan <- s.srv.ListenAndServeTLS(certFile, keyFile)

	return <- s.errChan
}

func (s *Server) Shutdown(ctx context.Context) error {
	go func() {
		s.done <- true
		if err := s.srv.Shutdown(ctx); err != nil {
			s.errChan <- err
		}
	}()
	<- s.done
	return <- s.errChan
}