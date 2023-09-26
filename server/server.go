package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
)

type Server struct {
	httpServer *http.Server
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