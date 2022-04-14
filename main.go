package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	ratelimiter "github.com/cwinters8/rate-limiter"
)

func hello(w http.ResponseWriter, req *http.Request) {
	msg, err := json.Marshal(struct {
		Message   string    `json:"message"`
		Timestamp time.Time `json:"timestamp"`
	}{
		Message:   "Hello, world!",
		Timestamp: time.Now(),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(msg)
}

func setup() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", hello)

	limiter := ratelimiter.NewLimiter(10000, 1)
	return http.ListenAndServe(":8080", limiter.Limit(mux))
}

func main() {
	err := setup()
	if err != nil {
		log.Fatal(err)
	}
}
