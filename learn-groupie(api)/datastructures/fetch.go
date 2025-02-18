package datastructures

import (
	"encoding/json"
	"net/http"
)

var (
	ArtistsAPI   string = "https://groupietrackers.herokuapp.com/api/artists"
	LocationsAPI string = "https://groupietrackers.herokuapp.com/api/locations"
	DatesAPI     string = "https://groupietrackers.herokuapp.com/api/dates"
	RelationAPI  string = "https://groupietrackers.herokuapp.com/api/relation"
)

var (
	Artists   []ARTISTS
	Locations LOCATIONS
	Dates     DATES
	Relations RELATION
)

var err error

func Fetch() error {
	Artists, err = FetchArtists(ArtistsAPI)
	if err != nil {

		return err
	}

	Locations, err = FetchLocations(LocationsAPI)
	if err != nil {

		return err
	}

	Dates, err = FetchDates(DatesAPI)
	if err != nil {

		return err
	}

	Relations, err = FetchRelations(RelationAPI)
	if err != nil {

		return err
	}
	return nil
}

func FetchArtists(api string) ([]ARTISTS, error) {
	resp, err := http.Get(api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	artistsSlice := []ARTISTS{}
	decode := json.NewDecoder(resp.Body)
	if err = decode.Decode(&artistsSlice); err != nil {
		return nil, err
	}
	return artistsSlice, nil
}

func FetchLocations(api string) (LOCATIONS, error) {
	resp, err := http.Get(api)
	if err != nil {
		return LOCATIONS{}, err
	}
	defer resp.Body.Close()

	locationsSlice := LOCATIONS{}
	decode := json.NewDecoder(resp.Body)
	if err = decode.Decode(&locationsSlice); err != nil {
		return LOCATIONS{}, err
	}
	return locationsSlice, nil
}

func FetchDates(api string) (DATES, error) {
	resp, err := http.Get(api)
	if err != nil {
		return DATES{}, err
	}
	defer resp.Body.Close()

	datesSlice := DATES{}
	decode := json.NewDecoder(resp.Body)
	if err = decode.Decode(&datesSlice); err != nil {
		return DATES{}, nil
	}
	return datesSlice, nil
}

func FetchRelations(api string) (RELATION, error) {
	resp, err := http.Get(api)
	if err != nil {
		return RELATION{}, err
	}
	defer resp.Body.Close()

	relationsSlice := RELATION{}
	decode := json.NewDecoder(resp.Body)
	if err = decode.Decode(&relationsSlice); err != nil {
		return RELATION{}, err
	}
	return relationsSlice, nil
}
