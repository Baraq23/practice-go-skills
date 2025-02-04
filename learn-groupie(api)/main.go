package main

import (
	// "fmt"
	"log"
	"net/http"

	// "groupie/datastructures"

	// "golang.org/x/tools/cmd/getgo/server"
)

var err error

func main() {
	// err := datastructures.Fetch()
	if err != nil {
		log.Println(err)
		return
	}
	port := ":3000"
	// fmt.Println("\nRelations\n", datastructures.Relations)
	// fmt.Println("\nDates\n", datastructures.Dates)
	// fmt.Println("\nLocations", datastructures.Locations)
	// fmt.Println("\nArtists\n", datastructures.Artists)

	log.Println("Server running on http://localhost"+port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Println("Error creating server: ", err)
		return
	}
}
