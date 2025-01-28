package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker/pkg/api"
	"groupie-tracker/pkg/models"
)

func ArtistInfoHandler(writer http.ResponseWriter, request *http.Request,
	client api.HTTPClient,
	fetchLocation func(string) (*models.LocationResponse, error),
	fetchDates func(string) (*models.Date, error),
	fetchRelations func(string) (*models.Relation, error),
) {
	// Check HTTP method
	if request.Method != http.MethodGet {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	path := request.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) < 3 {
		http.Error(writer, "Invalid artist ID", http.StatusBadRequest)
		return
	} else if len(parts) > 3 {
		http.Error(writer, "Bad Request", http.StatusBadRequest)
		return
	}

	infoTemplate := templ.Lookup("artistsinfo.html")
	if infoTemplate == nil {
		http.Error(writer, "InfoTemplate Not Found", http.StatusInternalServerError)
		return
	}

	artistID, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(writer, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	// Fetch artists from the cache
	artists, err := api.GetArtists(client, fetchLocation, fetchDates, fetchRelations)
	if err != nil {
		http.Error(writer, "Error fetching Artists Data", http.StatusInternalServerError)
		return
	}

	// finding the artist by Id in the slice
	var foundArtist models.Artist
	found := false
	for _, artist := range *artists {
		if artist.Id == artistID {
			foundArtist = artist
			found = true
			break
		}
	}
	if !found {
		http.Error(writer, "Artist not found", http.StatusInternalServerError)
		return
	}

	// Render the artist information using a detailed template
	err = infoTemplate.Execute(writer, foundArtist)
	if err != nil {
		http.Error(writer, "Error rendering template", http.StatusInternalServerError)
	}
}
