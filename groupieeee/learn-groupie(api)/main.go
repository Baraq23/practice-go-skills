package main

import (
	"fmt"

	"groupie/datastructures"
)

var (
	Artists []datastructures.ARTISTS
	Locations datastructures.LOCATIONS
	Dates datastructures.DATES
	Relations datastructures.RELATION

)

var err error

func main() {
	Artists, err = datastructures.FetchArtists(datastructures.ArtistsAPI)
	if err != nil {
		fmt.Println(err)
		return
	}

	Locations, err = datastructures.FetchLocations(datastructures.LocationsAPI)
	if err != nil {
		fmt.Println(err)
		return
	}

	Dates, err = datastructures.FetchDates(datastructures.DatesAPI)
	if err != nil {
		fmt.Println(err)
		return
	}

	Relations, err = datastructures.FetchRelations(datastructures.RelationAPI)
	if err != nil {
		fmt.Println(err)
		return
	}


	fmt.Println(Relations)
}
