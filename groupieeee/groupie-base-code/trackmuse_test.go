package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"

	"groupie-tracker/pkg/handlers"
	"groupie-tracker/pkg/models"
)

type MockHTTPClient struct {
	Response *http.Response
	Err      error
}

func (m *MockHTTPClient) Get(url string) (*http.Response, error) {
	return m.Response, m.Err
}

var templ *template.Template

func TestGetArtists(t *testing.T) {
	if err := handlers.Init(); err != nil {
		t.Fatalf("Failed to parse templates: %v", err)
	}

	// Mock HTTP client
	mockClient := &MockHTTPClient{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(strings.NewReader(`[
			{
			"id":14,
			"name":"Phil Collins"
				}
			]`)),
		},
		Err: nil,
	}
	mockFetchLocation := func(url string) (*models.LocationResponse, error) {
		return &models.LocationResponse{Location: []string{
			"nevada-usa",
			"arizona-usa",
			"california-usa",
		}}, nil
	}

	mockFetchDates := func(url string) (*models.Date, error) {
		return &models.Date{Dates: []string{
			"19-10-2019",
			"18-10-2019",
			"17-10-2019",
		}}, nil
	}

	mockFetchRelations := func(url string) (*models.Relation, error) {
		return &models.Relation{DatesLocation: map[string][]string{
			"arizona-usa":    {"18-10-2019"},
			"california-usa": {"17-10-2019"},
			"nevada-usa":     {"19-10-2019"},
		}}, nil
	}

	// Create a response recorder to capture the output
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	// Call the Index handler with the mock client
	handlers.Index(rr, req, mockClient, mockFetchLocation, mockFetchDates, mockFetchRelations)

	// Optionally check the response body
	body := rr.Body.String()
	// log.Println("This is the response body =>", body)
	expectedName := "Phil Collins"
	if !strings.Contains(body, expectedName) {
		t.Errorf("handler did not return expected artist name: got %v want %v", body, expectedName)
	}
}
