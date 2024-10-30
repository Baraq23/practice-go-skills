package decoders

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type RELATIONS struct {
	Index []RelationsObj `json:"index"`
}

type RelationsObj struct {
	Id            int `json:"id"`
	DateLocations map[string][]string `json:"datesLocations"`
}

type RelationsData RELATIONS

func DecodeRelations() *RelationsData {
	url := "https://groupietrackers.herokuapp.com/api/relations"

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error fetching locations response.")
		return nil
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Erro reading response body of ralations data.")
		return nil
	}
	log.Println("Hello", string(data)) // Log the raw JSON

	var RelationsInfo RelationsData
	err = json.Unmarshal(data, &RelationsInfo)
	if err != nil {
		log.Println("Error unmarshalling relations data", err)
		return nil
	}
	return &RelationsInfo
}
