package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string) error {
	s.httpServer = &http.Server{
		Addr:           fmt.Sprintf(":%s", port),
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.Shutdown(ctx)
}
