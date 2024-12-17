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
type locationsStructItem struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	DatesLink string   `json:"dates"`
}

type LOCATIONS struct {
	Index []locationsStructItem `json:"index"`
}

// Dates API
type datesStructItem struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type DATES struct {
	Index []datesStructItem `json:"index"`
}

// Relations API
type dateLocationStruct struct {
	locDate []map[string][]string
}

type relationsStructItem struct {
	Id             int                  `json:"id"`
	DatesLocations []dateLocationStruct `json:"datesLocations"`
}

type RELATION struct {
	Index []relationsStructItem `json:"index"`
}
