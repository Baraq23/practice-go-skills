package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie-tracker/pkg/api"
	"groupie-tracker/pkg/handlers"
)

func main() {

	err := handlers.Init()
	if err != nil {
		fmt.Println("Error loading templates")
		return
	}

	client := &api.Client{}
	// Routing
	http.HandleFunc("/",
		func(writer http.ResponseWriter, request *http.Request) {
			handlers.Index(writer, request, client, api.FetchLocation, api.FetchDates, api.FetchRelations)
		})
	http.HandleFunc("/artist/",
		func(writer http.ResponseWriter, request *http.Request) {
			handlers.ArtistInfoHandler(writer, request, client, api.FetchLocation, api.FetchDates, api.FetchRelations)
		})

	fileServer := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Creating the server
	log.Println("server running on http://localhost:5000")
	err = http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println("Error Creating a server")
		return
	}
}
