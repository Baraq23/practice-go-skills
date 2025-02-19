package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"groupie/datastructures"
	
)

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HomeHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("HomeHandler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestArtistHandler(t *testing.T) {
	// Mock data
	datastructures.Artists = []datastructures.ARTISTS{
		{Id: 1, Name: "Artist1"},
	}
	datastructures.Dates = datastructures.DATES{
		Index: []datastructures.DatesStructItem{
			{Id: 1, Dates: []string{"2023-01-01"}},
		},
	}
	datastructures.Locations = datastructures.LOCATIONS{
		Index: []datastructures.LocationsStructItem{
			{Id: 1, Locations: []string{"Location1"}},
		},
	}

	req, err := http.NewRequest("GET", "/artist?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ArtistHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("ArtistHandler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestErrorHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/nonexistent", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ErrorHandler(w, http.StatusNotFound, "Page Not Found.")
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("ErrorHandler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}