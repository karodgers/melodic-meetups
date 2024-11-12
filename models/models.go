package models

// Artist represents the structure of an artist with various details.
type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

// structure holding location details.
type LocationsResponse struct {
	Index []Location `json:"index"`
}

// structure of an artist's concert locations and associated dates.
type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

// response structure that holds concert dates information.
type DatesResponse struct {
	Index []Date `json:"index"` // List of dates (wrapped in a structure called "index" from JSON)
}

// structure for an artist's concert dates.
type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// response structure that holds relation details between dates and locations.
type RelationsResponse struct {
	Index []Relation `json:"index"`
}

// maps dates to their respective locations.
type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// rendering a page with multiple artists' data.
type PageData struct {
	Artists []Artist
	Query   string
}

// renders page with detailed information about a specific artist.
type ArtistDetailData struct {
	Artist   Artist
	Location Location
	Date     Date
	Relation Relation
}
