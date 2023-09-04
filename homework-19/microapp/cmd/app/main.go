package main

import (
	"microapp/internal/api"
	"microapp/internal/storage"
	"net/http"

	"github.com/segmentio/kafka-go"
)

func main() {
	srv := New()
	srv.run()
}

type Server struct {
	api *api.API
}

func New() *Server {
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:29092"},
		Topic:   "link-analytics",
	})
	store := storage.New(kafkaWriter)
	s := Server{
		api: api.New(store),
	}
	return &s
}

func (s *Server) run() {
	http.ListenAndServe(":8080", s.api.Router)
}
