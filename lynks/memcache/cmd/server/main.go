package main

import (
	"log"
	"net/http"

	"github.com/badimalex/go-course/lynks/memcache/pkg/redis"
	"github.com/badimalex/go-course/lynks/shortener/pkg/api"
	"github.com/badimalex/go-course/lynks/shortener/pkg/urls"

	"github.com/gorilla/mux"
)

func main() {
	redisStorage, err := redis.NewRedisStorage("localhost:6379", "", 0)
	if err != nil {
		log.Fatalf("Failed to initialize Redis storage: %v", err)
	}

	urls := urls.New(redisStorage)
	api := api.New(urls, "http://localhost:8081/")

	r := mux.NewRouter()
	api.Init(r)

	log.Println("Server is running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
