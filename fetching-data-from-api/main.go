package main

import (
	"html/template"
	"knowing-api/decoders"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/loc", Locations)
	http.HandleFunc("/dates", Dates)

	log.Println("http://localhost:9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Println("Error creating server")
		return
	}


}

func Dates(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/dates.html")

	if err != nil {
		log.Println("Error parsing dates.html file")
		return
	}
	
	data := decoders.DecodeDates()
	temp.Execute(w, data)
}

func Locations(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/locations.html")

	if err != nil {
		log.Println("Error parsing locations.html file")
		return
	}
	
	data := decoders.DecodeLocations()
	temp.Execute(w, data)
}


func Index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println("Error parsing index.html file")
		return
	}
	data := decoders.Decoder()
	temp.Execute(w, data)
}


