package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", homehandler)

	log.Println("http://localhost:3001")

	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Println("error starting server:", err)
	}
}

func homehandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
