package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := fmt.Fprintf(w, "hellow there")
	log.Printf("%d bites have been written on the page", b)
}
