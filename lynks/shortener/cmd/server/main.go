package main

import (
	"log"
	"net/http"

	"github.com/badimalex/go-course/lynks/shortener/pkg/api"
	"github.com/badimalex/go-course/lynks/shortener/pkg/storage"
	"github.com/badimalex/go-course/lynks/shortener/pkg/urls"

	"github.com/gorilla/mux"
)

const url = "postgres://dmitriybadichan:123@127.0.0.1/lynks?sslmode=disable"

func main() {
	postgresStorage, err := storage.New(url)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	err = postgresStorage.Init()
	if err != nil {
		log.Fatalf("failed to initialize database table: %v", err)
	}

	urlService := urls.New(postgresStorage)
	apiHandler := api.New(urlService, "http://localhost:8080/")

	r := mux.NewRouter()
	apiHandler.Init(r)

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
