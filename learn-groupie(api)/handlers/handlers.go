package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
	"strconv"

	"groupie/datastructures"
)

var Tmpl *template.Template

func init() {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)
	templatesDir := filepath.Join(currentDir, "../templates")
	Tmpl = template.Must(template.ParseGlob(filepath.Join(templatesDir, "*.html")))
}


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
	if err != nil || id > len(datastructures.Artists) || id <= 0 {
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

	for i, art := range datastructures.Artists {
		if art.Id == id {
			ThisArtist = art
			ArtistDates = datastructures.Dates.Index[i].Dates
			ArtistLocations = datastructures.Locations.Index[i].Locations
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
