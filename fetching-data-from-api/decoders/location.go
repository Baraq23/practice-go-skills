package decoders

import (
	"encoding/json"
	"log"
	"net/http"
)

type LOCATIONS struct {
	Index []LocationObj `json:"index"`
}

type LocationObj struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type LocationData LOCATIONS

func DecodeLocations() *LocationData {

	url := "https://groupietrackers.herokuapp.com/api/locations"

	data, err := http.Get(url)
	if err != nil {
		log.Println("Error fetching locations data.")
		return nil
	}

	decoder := json.NewDecoder(data.Body)

	var locationInfo LocationData

	decoder.Decode(&locationInfo)
	return &locationInfo
}