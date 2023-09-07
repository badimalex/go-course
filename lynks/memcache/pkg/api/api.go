package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/badimalex/go-course/lynks/memcache/pkg/metrics"
	"github.com/badimalex/go-course/lynks/memcache/pkg/redis"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog"
)

type Api struct {
	redis   *redis.Storage
	metrics *metrics.Metrics
	logger  zerolog.Logger
}

func New(redis *redis.Storage, m *metrics.Metrics, l zerolog.Logger) *Api {
	return &Api{
		redis:   redis,
		metrics: m,
		logger:  l,
	}
}

func (h *Api) Init(r *mux.Router) {
	r.HandleFunc("/", h.create).Methods("POST")
	r.HandleFunc("/{short}", h.get).Methods("GET")
}

func (h *Api) create(w http.ResponseWriter, r *http.Request) {
	h.logger.Info().Msg("API create method called")
	start := time.Now()

	var req struct {
		Dest  string `json:"destination"`
		Short string `json:"short"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		h.logger.Error().Err(err).Msg("Failed to decode request body")
		return
	}
	data := redis.Data{Destination: req.Dest, Short: req.Short}

	err = h.redis.Save(data)

	if err != nil {
		http.Error(w, "Error creating short URL", http.StatusInternalServerError)
		h.logger.Error().Err(err).Msg("Error creating short URL")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	h.logger.Info().Msgf("Data created successfully: %v", data)
	h.metrics.HttpDuration.WithLabelValues("/").Observe(time.Since(start).Seconds())
	h.metrics.HttpRequestsTotal.WithLabelValues("POST", "/").Inc()
}

func (h *Api) get(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	vars := mux.Vars(r)
	short := vars["short"]

	destination, err := h.redis.Load(short)
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to load URL")
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	resp := struct {
		Destination string `json:"destination"`
	}{
		Destination: destination.Destination,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

	h.metrics.HttpRequestsTotal.With(prometheus.Labels{"method": r.Method, "path": "/"}).Inc()
	h.metrics.HttpDuration.With(prometheus.Labels{"path": "/"}).Observe(time.Since(start).Seconds())

	h.logger.Info().Str("method", r.Method).Str("path", r.URL.Path).Str("duration", time.Since(start).String()).Msg("Request processed successfully")
}
