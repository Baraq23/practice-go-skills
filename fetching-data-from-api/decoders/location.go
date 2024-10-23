package decoders

import (
	"encoding/json"
	"io"
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

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error fetching locations response.")
		return nil
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading locations response body.")
		return nil
	}

	var locationInfo LocationData
	err = json.Unmarshal(data, &locationInfo)
	if err != nil {
		log.Println("Error unmarshalling location data")
	}

	return &locationInfo
}