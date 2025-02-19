package datastructures

// Artists API
type ARTISTS struct {
	Id               int      `json:"id"`
	ImageLink        string   `json:"image"`
	Name             string   `json:"name"`
	Members          []string `json:"members"`
	CreationDate     int      `json:"creationDate"`
	FirstAlbum       string   `json:"firstAlbum"`
	LocationsLink    string   `json:"locations"`
	ConcertDatesLink string   `json:"concertDates"`
	RelationsLink    string   `json:"relations"`
}

// Locations API
type LocationsStructItem struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	DatesLink string   `json:"dates"`
}

type LOCATIONS struct {
	Index []LocationsStructItem `json:"index"`
}

// Dates API
type DatesStructItem struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type DATES struct {
	Index []DatesStructItem `json:"index"`
}

// Relations API
type RelationsStructItem struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type RELATION struct {
	Index []RelationsStructItem `json:"index"`
}
