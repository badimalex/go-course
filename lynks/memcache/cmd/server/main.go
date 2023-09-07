package main

import (
	"log"
	"net/http"

	"github.com/badimalex/go-course/lynks/memcache/pkg/api"
	"github.com/badimalex/go-course/lynks/memcache/pkg/redis"

	"github.com/gorilla/mux"
)

func main() {
	redisStorage, err := redis.New("localhost:6379", "", 0)
	if err != nil {
		log.Fatalf("Failed to initialize Redis storage: %v", err)
	}

	api := api.New(redisStorage)

	r := mux.NewRouter()
	api.Init(r)

	log.Println("Server is running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
