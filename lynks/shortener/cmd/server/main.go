package main

import (
	"log"
	"net/http"

	"github.com/badimalex/go-course/lynks/shortener/pkg/api"
	"github.com/badimalex/go-course/lynks/shortener/pkg/cache"
	"github.com/badimalex/go-course/lynks/shortener/pkg/metrics"
	"github.com/badimalex/go-course/lynks/shortener/pkg/storage"
	"github.com/badimalex/go-course/lynks/shortener/pkg/urls"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gorilla/mux"
)

const url = "postgres://postgres:postgres@127.0.0.1/url_shortener?sslmode=disable"
const root = "http://localhost:8080/"
const cachePath = "http://localhost:8081/"

func main() {
	reg := prometheus.NewRegistry()
	m := metrics.NewMetrics(reg)

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
	apiHandler := api.New(urlService, cacheService, m, root)

	r := mux.NewRouter()
	apiHandler.Init(r)
	r.Handle("/metrics", promhttp.Handler())

	log.Println("Server is running on ", root)
	log.Fatal(http.ListenAndServe(":8080", r))
}
