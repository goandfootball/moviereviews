package server

import (
	"net/http"
)

type Server struct {
	server *http.Server
}

func New(pPort string, pRouter http.Handler) (*Server, error) {
	srv := &http.Server{
		Addr:    ":" + pPort,
		Handler: pRouter,
	}

	server := Server{
		server: srv,
	}

	return &server, nil
}

func (srv *Server) Close() error {
	//TODO: add resource closure.
	return nil
}

func (srv *Server) Start() error{
	err := srv.server.ListenAndServe()

	if err != nil {
		return err
	}
	return nil
}
