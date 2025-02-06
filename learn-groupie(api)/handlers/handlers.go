package handlers

import (
	// "fmt"
	"groupie/datastructures"
	"html/template"
	"log"
	"net/http"
)


func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	err := tmpl.ExecuteTemplate(w, "index.html", datastructures.Artists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("error parsing templates.")
	}
}
