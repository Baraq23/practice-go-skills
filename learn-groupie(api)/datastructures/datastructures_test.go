package datastructures

import (
	// "groupie/datastructures"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchArtists(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id":1,"image":"image1","name":"Artist1","members":["Member1"],"creationDate":2000,"firstAlbum":"Album1","locationsLink":"link1","concertDatesLink":"link2","relationsLink":"link3"}]`))
	}))
	defer server.Close()

	// Use the mock server URL
	artists, err := FetchArtists(server.URL)
	if err != nil {
		t.Fatalf("FetchArtists() error = %v", err)
	}

	if len(artists) != 1 {
		t.Errorf("Expected 1 artist, got %d", len(artists))
	}

	if artists[0].Name != "Artist1" {
		t.Errorf("Expected artist name 'Artist1', got '%s'", artists[0].Name)
	}
}

func TestFetchLocations(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"index":[{"id":1,"locations":["Location1"],"datesLink":"link1"}]}`))
	}))
	defer server.Close()

	// Use the mock server URL
	locations, err := FetchLocations(server.URL)
	if err != nil {
		t.Fatalf("FetchLocations() error = %v", err)
	}

	if len(locations.Index) != 1 {
		t.Errorf("Expected 1 location, got %d", len(locations.Index))
	}

	if locations.Index[0].Locations[0] != "Location1" {
		t.Errorf("Expected location 'Location1', got '%s'", locations.Index[0].Locations[0])
	}
}

