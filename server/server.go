package server

import (
	"net/http"
)

type Server struct {
	srv *http.Server
	
	errChan chan error 
}

func (s *Server) Serve(addr string) error {
	s.srv.Addr = addr 
	s.srv.Handler = http.DefaultServeMux

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

func (s *Server) Shutdown() error {
	return nil
}