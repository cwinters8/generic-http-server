package main

import (
	"fmt"
	"log"
	"net/http"

	ratelimiter "github.com/cwinters8/rate-limiter"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello, world!")
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
