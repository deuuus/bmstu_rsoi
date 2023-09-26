package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr: ":" + port,
		Handler: handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func Run(handler http.Handler) {
	var server Server
  
	server.httpServer = &http.Server{
	  Handler: handler,
	}
  
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", "8080"))
	if err != nil {
	  log.Fatal(err)
	}
  
	server.httpServer.Serve(listener)
  }

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}