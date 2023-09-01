package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-course/homework-13/pkg/crawler"
	"go-course/homework-13/pkg/retriever"

	"github.com/gorilla/mux"
)

type API struct {
	router    *mux.Router
	retriever *retriever.Retriever
	data      []crawler.Document
}

func New(retriever *retriever.Retriever, data []crawler.Document) *API {
	a := &API{
		router:    mux.NewRouter(),
		retriever: retriever,
		data:      data,
	}
	a.endpoints()
	a.router.Use(jsonMiddleware)
	return a
}

func (a *API) endpoints() {
	a.router.HandleFunc("/search/{query}", a.search).Methods("GET")
	a.router.HandleFunc("/documents", a.createDocument).Methods("POST")
	a.router.HandleFunc("/documents/{id}", a.getDocument).Methods("GET")
	a.router.HandleFunc("/documents/{id}", a.updateDocument).Methods("PUT")
	a.router.HandleFunc("/documents/{id}", a.deleteDocument).Methods("DELETE")
}

func (api *API) search(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	docs := api.retriever.Find(params["query"])

	json.NewEncoder(w).Encode(docs)
}

func (a *API) createDocument(w http.ResponseWriter, r *http.Request) {
	var doc crawler.Document
	_ = json.NewDecoder(r.Body).Decode(&doc)
	a.data = append(a.data, doc)
	w.WriteHeader(http.StatusCreated)
}

func (a *API) getDocument(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, doc := range a.data {
		if doc.ID == id {
			json.NewEncoder(w).Encode(doc)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func (a *API) updateDocument(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var updatedDoc crawler.Document
	_ = json.NewDecoder(r.Body).Decode(&updatedDoc)

	for index, doc := range a.data {
		if doc.ID == id {
			a.data[index] = updatedDoc
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func (a *API) deleteDocument(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, doc := range a.data {
		if doc.ID == id {
			a.data = append(a.data[:index], a.data[index+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func (a *API) Serve(addr string) {
	http.ListenAndServe(addr, a.router)
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
