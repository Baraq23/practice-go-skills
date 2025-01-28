package models

// Artists data
type Artist struct{
	Id int `json:"id"`
	Image string `json:"image"`
	Name string `json:"name"`
	Members []string `json:"members"`
	CreationDate int `json:"creationDate"`
	FirstAlbum string `json:"firstAlbum"`
	LocationURL string `json:"locations"`
	ConcertDatesURL string `json:"concertDates"`
	RelationsURL string `json:"relations"`
	Locations []string
	ConcertDates []string 
	Relations map[string][]string
	
}

type Artists []Artist

// Artists location data
type LocationResponse struct{
	Id int `json:"id"`
	Location []string `json:"locations"`
	Dates string `json:"dates"`
}
type Locations []LocationResponse

// Artists dates data
type Date struct{
	Id int `json:"id"`
	Dates []string `json:"dates"`
}
type Dates []Date


// Relationships data
type Relation struct{
	Id int `json:"id"`
	DatesLocation map[string][]string `json:"datesLocations"`
}

type Relations []Relation