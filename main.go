package main

import (
	"bytes"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
	if r.Body != nil {
		// Print body content
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		body := buf.String()
		log.Printf("Body: %s", body)
	}
	// Print headers
	for name, headers := range r.Header {
		for _, h := range headers {
			log.Printf("Header: %v: %v", name, h)
		}
	}
	// Print query parameters
	query := r.URL.Query()
	for key, value := range query {
		log.Printf("Query: %v: %v", key, value)
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", handler)

	log.Println("API Responder Service started on port :8888")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
