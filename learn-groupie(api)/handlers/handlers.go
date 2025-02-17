package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"groupie/datastructures"
)

var Tmpl = template.Must(template.ParseGlob("templates/*.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound, "Page Not Found.")
		return
	}
	err := Tmpl.ExecuteTemplate(w, "index.html", datastructures.Artists)
	if err != nil {
		log.Println("error parsing index.html templates:", err)
		// w.WriteHeader(http.StatusInternalServerError)
		ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}

func ErrorHandler(w http.ResponseWriter, statusCode uint, message string) {
	w.WriteHeader(int(statusCode))
	data := struct {
		Status  uint
		Message string
	}{
		Status:  statusCode,
		Message: message,
	}
	err := Tmpl.ExecuteTemplate(w, "error-page.html", data)
	if err != nil {
		log.Println("error parsing error-page.html:", err)
		// http.Error(w, message, int(statusCode))
		return
	}
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	strId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(strId)

	if err != nil || id > len(datastructures.Artists) || id < 0 {
		log.Printf("Error %d: Not found.\n", http.StatusNotFound)
		ErrorHandler(w, http.StatusNotFound, fmt.Sprintf("Artist of Id: %s Does Not Exist.", strId))
		return
	}

	type ArtistInfo struct {
		Artist        datastructures.ARTISTS
		DatesInfo     []string
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
		Artist:        ThisArtist,
		DatesInfo:     ArtistDates,
		LocationsInfo: ArtistLocations,
	}

	err = Tmpl.ExecuteTemplate(w, "artist.html", data)
	if err != nil {
		log.Println("Error passing artist.html templates:", err)
		ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error.")
		return
	}
}
