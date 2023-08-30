package webapp

import (
	"encoding/json"
	"go-course/homework-12/pkg/crawler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type WebApp struct {
	idx  map[string][]int
	docs []crawler.Document
}

func New(idx map[string][]int, docs []crawler.Document) *WebApp {
	return &WebApp{
		idx:  idx,
		docs: docs,
	}
}

func (s *WebApp) indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.idx)
}

func (s *WebApp) docsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.docs)
}

func (w *WebApp) Start() {
	r := mux.NewRouter()

	r.HandleFunc("/index", w.indexHandler).Methods("GET")
	r.HandleFunc("/docs", w.docsHandler).Methods("GET")

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
