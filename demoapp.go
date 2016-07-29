package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm alright")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there screencast viewers!")
}

func main() {
	http.HandleFunc("/healthcheck", healthHandler)
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
	
