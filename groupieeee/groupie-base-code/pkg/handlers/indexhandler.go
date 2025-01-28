package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"groupie-tracker/pkg/api"
	"groupie-tracker/pkg/models"
)

var templ *template.Template

func Init() error {
	var err error
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	templatePath := filepath.Join(filepath.Dir(cwd), "groupie-tracker", "web", "templates", "*.html")
	templ, err = template.ParseGlob(templatePath)
	if err != nil {
		log.Printf("Error loading templates: %v", err) // Log the error
		return err
	}
	log.Println("Templates initialized successfully")
	return nil
}

func Index(writer http.ResponseWriter, request *http.Request,
	client api.HTTPClient,
	fetchLocation func(string) (*models.LocationResponse, error),
	fetchDates func(string) (*models.Date, error),
	fetchRelations func(string) (*models.Relation, error),
) {
	path := request.URL.Path
	if !(path == "/" || path == "/artists"){
		http.Error(writer, "Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Println("This is the request path >>", path)
	
	// look up the index.html
	if templ == nil {
		http.Error(writer, "Templates not initialized", http.StatusInternalServerError)
		return
	}
	indexTemplate := templ.Lookup("index.html")

	if indexTemplate == nil {
		http.Error(writer, "Template Not Found", http.StatusInternalServerError)
		return
	}
	// Artists data
	artistsData, err := api.GetArtists(client, fetchLocation, fetchDates, fetchRelations)
	if err != nil || artistsData == nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the template
	err = indexTemplate.Execute(writer, artistsData)
	if err != nil {
		http.Error(writer, "Error Rendering Template", http.StatusInternalServerError)
		return
	}
}
