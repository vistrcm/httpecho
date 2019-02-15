package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if cerr := r.Body.Close(); cerr != nil {
			log.Printf("[WARN] can not close request body %v.", r)
		}
	}()

	w.WriteHeader(http.StatusOK)
	if _, copyErr := io.Copy(w, r.Body); copyErr != nil {
		log.Printf("[WARN] can not copy to writer %v.", copyErr)
	}
}

func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
