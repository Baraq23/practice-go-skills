package main

import (
	"log"
	"net/http"

	"forum/handlers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HomeHandler(w, r)
	})
	log.Println("Server running on: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
