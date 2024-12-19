package main

import (
	"fmt"

	"groupie/datastructures"
)

func main() {
	// artists, err := datastructures.FetchArtists(datastructures.ArtistsAPI)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// locations, err := datastructures.FetchLocations(datastructures.LocationsAPI)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// dates, err := datastructures.FetchDates(datastructures.DatesAPI)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	relations, err := datastructures.FetchRelations(datastructures.RelationAPI)
	if err != nil {
		fmt.Println(err)
		return
	}

	// listArtists(artists)
	// listLocations(locations)
	// listDates(dates)
	listRelations(relations)
}

func listArtists(content []datastructures.ARTISTS) {
	fmt.Println("PRINTING ARTISTS")
	count := 1
	for _, co := range content {
		fmt.Println(count)
		fmt.Println(co)
		fmt.Println()
		count++

	}
}

func listLocations(content datastructures.LOCATIONS) {
	fmt.Println("PRINTING LOCATIONS")
	count := 1
	fmt.Println(content)
	fmt.Println()
	count++
}

func listDates(content datastructures.DATES) {
	fmt.Println("PRINTING DATES")
	count := 1
	fmt.Println(content)
	fmt.Println()
	count++
}

func listRelations(content datastructures.RELATION) {
	fmt.Println("PRINTING RELATIONS")
	count := 1
	fmt.Println(content)
	fmt.Println()
	count++
}
