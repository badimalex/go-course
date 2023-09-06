package api

import (
	"encoding/json"
	"net/http"

	"github.com/badimalex/go-course/lynks/shortener/pkg/urls"

	"github.com/gorilla/mux"
)

type APIHandler struct {
	urlService *urls.Service
	root       string
}

func NewAPIHandler(urlService *urls.Service, root string) *APIHandler {
	return &APIHandler{
		urlService: urlService,
		root:       root,
	}
}

func (h *APIHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", h.createShortURL).Methods("POST")
	r.HandleFunc("/{shortURL}", h.redirectToURL).Methods("GET")
}

func (h *APIHandler) createShortURL(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Destination string `json:"destination"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shortURL, err := h.urlService.Create(requestData.Destination)
	if err != nil {
		http.Error(w, "Error creating short URL", http.StatusInternalServerError)
		return
	}

	responseData := struct {
		ShortURL    string `json:"shortUrl"`
		Destination string `json:"destination"`
	}{
		ShortURL:    h.root + shortURL.Short,
		Destination: requestData.Destination,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}

func (h *APIHandler) redirectToURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shortURL"]

	destination, err := h.urlService.Get(shortURL)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, destination.Destination, http.StatusSeeOther)
}
