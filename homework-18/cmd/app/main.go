package main

import (
	"microapp/internal/api"
	"microapp/internal/storage"
	"net/http"
)

func main() {
	srv := New()
	srv.run()
}

type Server struct {
	api *api.API
}

func New() *Server {
	store := storage.New()
	s := Server{
		api: api.New(store),
	}
	return &s
}

func (s *Server) run() {
	http.ListenAndServe(":8080", s.api.Router)
}
