package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
}

type Artists []Artist

func Decoder() *Artists {
	data, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Println("Error fetching Artists data")
		return nil
	}
	decoder := json.NewDecoder(data.Body)

	var artistsData Artists

	decoder.Decode(&artistsData)

	return &artistsData
}
