package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"groupie/datastructures"
)

var Tmpl = template.Must(template.ParseGlob("templates/*.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := Tmpl.ExecuteTemplate(w, "index.html", datastructures.Artists)

	if err != nil {
		log.Println("error parsing index.html templates:", err)
		return
	}
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	strId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(strId)

	if err != nil || id > len(datastructures.Artists) || id < 0 {
		log.Printf("Error %d: Not found.\n", http.StatusNotFound)
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	type ArtistInfo struct {
		Artist datastructures.ARTISTS
		DatesInfo []string
		LocationsInfo []string
	}

	var ThisArtist datastructures.ARTISTS
	var ArtistDates []string
	var ArtistLocations []string

	for _, art := range datastructures.Artists {
		if art.Id == id {
			ThisArtist = art
			ArtistDates = datastructures.Dates.Index[id].Dates
			ArtistLocations = datastructures.Locations.Index[id].Locations
			break
		}
	}

	data := ArtistInfo{
		Artist: ThisArtist,
		DatesInfo: ArtistDates,
		LocationsInfo: ArtistLocations,
	}

	err = Tmpl.ExecuteTemplate(w, "artist.html", data)
	if err != nil {
		log.Println("Error passing artist.html templates:", err)
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
}
