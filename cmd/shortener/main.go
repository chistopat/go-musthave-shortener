package main

import (
	shortener "github.com/chistopat/go-musthave-shortener/internal/app"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	router := shortener.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
