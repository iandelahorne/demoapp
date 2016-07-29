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
	fmt.Fprintf(w, "San Francisco is a nice city")
}

func main() {
	http.HandleFunc("/healthcheck", healthHandler)
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
	
