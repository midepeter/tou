package server

import (
	"context"
	"net/http"

	"github.com/midepeter/tou/internal/handlers"
	"go.uber.org/zap"
)

type Server struct {
	srv     *http.Server
	errChan chan error
	done    chan bool
	log     *zap.SugaredLogger
	handler *handlers.Handler
}

func NewServer(h handlers.Handler, logger *zap.SugaredLogger) *Server {
	srv := &http.Server{}
	return &Server{
		errChan: make(chan error),
		done:    make(chan bool),
		srv:     srv,
		log:     logger,
		handler: &h,
	}
}

func (s *Server) Serve(addr string) error {
	s.srv.Addr = addr
	s.srv.Handler = http.DefaultServeMux

	http.HandleFunc("/add", s.handler.Add)

	http.HandleFunc("/inspect", s.handler.Inspect)

	s.log.Info("Serving work queue server to on the addr ", addr)
	s.errChan <- s.srv.ListenAndServe()

	return <-s.errChan
}

func (s *Server) Shutdown(ctx context.Context) error {
	go func() {
		s.done <- true
		if err := s.srv.Shutdown(ctx); err != nil {
			s.errChan <- err
		}
	}()
	return <-s.errChan
}
