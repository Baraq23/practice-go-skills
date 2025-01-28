package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"groupie-tracker/pkg/models"
)

// caching mechanism
var (
	CachedArtists  *models.Artists
	artistsFetched bool
)
type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

type Client struct{}

func (c *Client) Get(url string) (*http.Response, error) {
	return http.Get(url)
}

func GetArtists(client HTTPClient,fetchLocation func(string) (*models.LocationResponse, error),
fetchDates func(string) (*models.Date, error),
fetchRelations func(string) (*models.Relation, error)) (*models.Artists, error) {
	// If we have already fetched the artists return the cached
	if artistsFetched {
		return CachedArtists, nil
	}
	// Fetching artists from the API
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Println("Error accessing the artists\n", err)
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Printf("Received non-OK HTTP status: %s\n", response.Status)
		return nil, fmt.Errorf("Non-Ok HTTP status %s", response.Status)
	}

	var artists models.Artists
	// Creating the decoder
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&artists)
	if err != nil {
		log.Println("Error decoding Artist json data.\n", err)
		return nil, err
	}

	// Loop through each artist and fetch their additional data (locations)
	for i := range artists {
		// Fetch locations using the artist's LocationURL
		locationResponse, err := FetchLocation(artists[i].LocationURL)
		if err != nil {
			log.Println("Error fetching locations for artist", artists[i].Id)
			continue // Skip this artist if location fetching fails
		}

		// Now, we only care about the `locations` field from LocationResponse
		artists[i].Locations = locationResponse.Location

		// Fetching dates data
		dates, err := FetchDates(artists[i].ConcertDatesURL)
		if err != nil {
			log.Println("Error fetching dates for artist", artists[i].Id)
			continue // Skip this artist if dates fetching fails
		}
		// Clean up the dates
		for j, date := range dates.Dates {
			dates.Dates[j] = strings.TrimPrefix(date, "*") // Remove leading asterisk
		}
		artists[i].ConcertDates = dates.Dates

		// Fetching relations data
		relations, err := FetchRelations(artists[i].RelationsURL)
		// log.Println(relations.DatesLocation)
		if err != nil {
			log.Println("Error fetching relations for artist", artists[i].Id)
			continue // Skip this artist if relations fetching fails
		}
		artists[i].Relations = relations.DatesLocation
		// log.Println("Assigned Relations:", artists[i].Relations)
	}
	// cache the fetched data
	CachedArtists = &artists
	artistsFetched = true

	return &artists, nil
}

func FetchLocation(url string) (*models.LocationResponse, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Println("Error accessing the artists\n", err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Printf("Received non-OK HTTP status: %s\n", response.Status)
		return nil, fmt.Errorf("Non-Ok HTTP status %s", response.Status)
	}

	// Log the response body for debugging
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, err
	}
	// log.Println("Location response body:", string(body))

	var LocationResponse models.LocationResponse
	err = json.Unmarshal(body, &LocationResponse)
	if err != nil {
		log.Println("Error decoding locations json data.\n", err)
		return nil, err
	}
	return &LocationResponse, nil
}

func FetchDates(url string) (*models.Date, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Println("Error accessing dates data\n", err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Printf("Received non-OK HTTP status: %s\n", response.Status)
		return nil, fmt.Errorf("Non-Ok HTTP status %s", response.Status)
	}
	// Log the response body for debugging
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, err
	}
	var datesData models.Date
	err = json.Unmarshal(body, &datesData)
	if err != nil {
		log.Println("Error decoding dates json data.\n", err)
		return nil, err
	}
	return &datesData, nil
}

func FetchRelations(url string) (*models.Relation, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Println("Error accessing relations data\n", err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Printf("Received non-OK HTTP status: %s\n", response.Status)
		return nil, fmt.Errorf("Non-Ok HTTP status %s", response.Status)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, err
	}

	var relationsData models.Relation
	err = json.Unmarshal(body, &relationsData)
	if err != nil {
		log.Println("Error decoding relations json data.\n", err)
		return nil, err
	}

	return &relationsData, nil
}
