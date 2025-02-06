package handlers

import (
	// "fmt"
	"groupie/datastructures"
	"html/template"
	"log"
	"net/http"
)


var Tmpl = template.Must(template.ParseGlob("templates/*.html"))
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	err := Tmpl.ExecuteTemplate(w, "index.html", datastructures.Artists)
	if err != nil {
		log.Println("error parsing index.html templates.")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	err:= Tmpl.ExecuteTemplate(w, "artist.html", datastructures.Artists[0])
	if err != nil {
		log.Println("Error passing artist.html templates")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
