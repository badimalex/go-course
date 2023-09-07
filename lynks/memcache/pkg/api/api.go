package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/badimalex/go-course/lynks/memcache/pkg/redis"
	"github.com/gorilla/mux"
)

type Api struct {
	redis *redis.Storage
}

func New(redis *redis.Storage) *Api {
	return &Api{
		redis: redis,
	}
}

func (h *Api) Init(r *mux.Router) {
	r.HandleFunc("/", h.create).Methods("POST")
	r.HandleFunc("/{short}", h.get).Methods("GET")
}

func (h *Api) create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Dest  string `json:"destination"`
		Short string `json:"short"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	data := redis.Data{Destination: req.Dest, Short: req.Short}

	err = h.redis.Save(data)
	fmt.Println(data)
	if err != nil {
		http.Error(w, "Error creating short URL", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func (h *Api) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	short := vars["short"]

	destination, err := h.redis.Load(short)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	resp := struct {
		Destination string `json:"destination"`
	}{
		Destination: destination.Destination,
	}
	fmt.Println(resp)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
