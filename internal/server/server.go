package server

import (
	"github.com/go-chi/chi"
	v1 "github.com/goandfootball/moviereviews/internal/server/v1"
	"net/http"
)

type Server struct {
	server *http.Server
}

func New(pPort string) (*Server, error) {
	r := chi.NewRouter()

	//r.Mount("/api/v1", v1.New())
	r.Mount("/", v1.New())

	srv := &http.Server{
		Addr:    ":" + pPort,
		Handler: r,
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

func (srv *Server) Start() error {
	err := srv.server.ListenAndServe()

	if err != nil {
		return err
	}
	return nil
}
