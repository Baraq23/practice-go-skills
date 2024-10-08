package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Index)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Println("Error creating server")
		return
	}
	log.Println("http://localhost:9000")


}


func Index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println("Error parsing indext.html file")
		return
	}
	data := Decoder()
	temp.Execute(w, data)
}
