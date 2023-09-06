package api

import (
	"encoding/json"
	"net/http"

	"github.com/badimalex/go-course/lynks/shortener/pkg/urls"

	"github.com/gorilla/mux"
)

type Api struct {
	url  *urls.Service
	root string
}

func New(urlService *urls.Service, root string) *Api {
	return &Api{
		url:  urlService,
		root: root,
	}
}

func (h *Api) Init(r *mux.Router) {
	r.HandleFunc("/", h.createShortURL).Methods("POST")
	r.HandleFunc("/{short}", h.redirectToURL).Methods("GET")
}

func (h *Api) createShortURL(w http.ResponseWriter, r *http.Request) {
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

	resp := struct {
		ShortURL    string `json:"shortUrl"`
		Destination string `json:"destination"`
	}{
		ShortURL:    h.root + url.Short,
		Destination: req.Dest,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Api) redirectToURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	short := vars["short"]

	destination, err := h.url.Get(short)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, destination.Destination, http.StatusSeeOther)
}
