package decoders

import (
	"encoding/json"
	"log"
	"net/http"
)

type DATES struct {
	Index []DateObj `json:"index"`
}

type DateObj struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type DatesData DATES

func DecodeDates() *DatesData {
	url := "https://groupietrackers.herokuapp.com/api/dates"

	data, err := http.Get(url)
	if err != nil {
		log.Println("Error fetching locations data.")
		return nil
	}

	decoder := json.NewDecoder(data.Body)

	var DateInfo DatesData

	decoder.Decode(&DateInfo)
	return &DateInfo
}
