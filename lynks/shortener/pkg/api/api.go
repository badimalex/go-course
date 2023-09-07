package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/badimalex/go-course/lynks/shortener/pkg/cache"
	"github.com/badimalex/go-course/lynks/shortener/pkg/metrics"
	"github.com/badimalex/go-course/lynks/shortener/pkg/urls"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/gorilla/mux"
)

type Api struct {
	url     *urls.Service
	cache   *cache.Service
	metrics *metrics.Metrics
	root    string
}

func New(urls *urls.Service, cache *cache.Service, metrics *metrics.Metrics, root string) *Api {
	return &Api{
		url:     urls,
		cache:   cache,
		root:    root,
		metrics: metrics,
	}
}

func (h *Api) Init(r *mux.Router) {
	r.HandleFunc("/", h.create).Methods("POST")
	r.HandleFunc("/{short}", h.get).Methods("GET")
}

func (h *Api) create(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var req struct {
		Dest string `json:"destination"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	url, err := h.url.Create(req.Dest)
	if err != nil {
		http.Error(w, "Error creating short URL", http.StatusInternalServerError)
		return
	}

	err = h.cache.Create(req.Dest, url.Short)
	if err != nil {
		http.Error(w, "Error saving to cache", http.StatusInternalServerError)
		return
	}

	resp := struct {
		ShortURL    string `json:"shortUrl"`
		Destination string `json:"destination"`
	}{
		ShortURL:    h.root + url.Short,
		Destination: req.Dest,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

	h.metrics.HttpDuration.WithLabelValues("/").Observe(time.Since(start).Seconds())
	h.metrics.HttpRequestsTotal.WithLabelValues("POST", "/").Inc()
}

func (h *Api) get(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	vars := mux.Vars(r)
	short := vars["short"]

	c, err := h.cache.Get(short)
	if err == nil {
		http.Redirect(w, r, c.Destination, http.StatusSeeOther)
		return
	}

	u, err := h.url.Get(short)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, u.Destination, http.StatusSeeOther)

	h.metrics.HttpRequestsTotal.With(prometheus.Labels{"method": r.Method, "path": "/"}).Inc()
	h.metrics.HttpDuration.With(prometheus.Labels{"path": "/"}).Observe(time.Since(start).Seconds())
}
