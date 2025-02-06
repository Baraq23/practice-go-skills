package handlers

import (
	// "fmt"
	"fmt"
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
		log.Println("error parsing index.html templates.")
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
	var artist datastructures.ARTISTS
	fmt.Println("id", id)
	for _, art := range datastructures.Artists {
		if art.Id == id {
			fmt.Println("here")
			artist = art
			break
		}
	}
	err = Tmpl.ExecuteTemplate(w, "artist.html", artist)
	if err != nil {
		log.Println("Error passing artist.html templates")
		return
	}
}
