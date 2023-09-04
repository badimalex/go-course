package main

import (
	"analytics/internal/analytics"
	"analytics/pkg/kafka"
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	kfk, err := kafka.New(
		[]string{"localhost:29092"},
		"link-analytics",
		"analytics-group",
	)
	if err != nil {
		log.Fatal(err)
	}

	svc := analytics.New()
	kfk.Subscribe(svc.AddInfo)

	go startHTTPServer(svc)
	kfk.Consumer()
}

func startHTTPServer(svc *analytics.Service) {
	http.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		stats := svc.Stats()
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(stats); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
