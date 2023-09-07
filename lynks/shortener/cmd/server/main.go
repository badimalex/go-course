package main

import (
	"log"
	"net/http"

	"github.com/badimalex/go-course/lynks/shortener/pkg/api"
	"github.com/badimalex/go-course/lynks/shortener/pkg/cache"
	"github.com/badimalex/go-course/lynks/shortener/pkg/storage"
	"github.com/badimalex/go-course/lynks/shortener/pkg/urls"

	"github.com/gorilla/mux"
)

const url = "postgres://dmitriybadichan:123@127.0.0.1/lynks?sslmode=disable"
const root = "http://localhost:8080/"
const cachePath = "http://localhost:8081/"

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
	cacheService := cache.New(cachePath)
	apiHandler := api.New(urlService, cacheService, root)

	r := mux.NewRouter()
	apiHandler.Init(r)

	log.Println("Server is running on ", root)
	log.Fatal(http.ListenAndServe(":8080", r))
}
