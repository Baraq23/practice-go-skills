package main

import (
	"log"
	"net/http"

	// "fmt"

	"groupie/datastructures"
	"groupie/handlers"
	// "golang.org/x/tools/cmd/getgo/server"
)

var err error

func main() {
	err := datastructures.Fetch()
	if err != nil {
		log.Println(err)
		return
	}
	port := ":3000"
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/artist", handlers.ArtistHandler)
	http.HandleFunc("/static", func(w http.ResponseWriter, r *http.Request) {
		handlers.ErrorHandler(w, http.StatusForbidden, "Access forbidden.")
	})

	log.Println("Server running on http://localhost" + port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Println("Error creating server: ", err)
		return
	}
}
