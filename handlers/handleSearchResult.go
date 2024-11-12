package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	database "groupie-trackers/data"
	"groupie-trackers/models"
)

// LocationData represents the structure of individual location data
type LocationData struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

// LocationsResponse represents the full response from the locations API
type LocationsResponse struct {
	Index []LocationData `json:"index"`
}

func HandleSearchResults(w http.ResponseWriter, r *http.Request) {
	// Extract search query from request parameters
	query := r.URL.Query().Get("query")
	query = strings.ToLower(query)

	// Fetch all artists
	var artists []models.Artist
	if err := database.FetchData("https://groupietrackers.herokuapp.com/api/artists", &artists); err != nil {
		log.Printf("Error fetching artists data: %v", err)
		ErrorPage(w, "Error fetching artists data", http.StatusInternalServerError)
		return
	}

	// Fetch locations data with correct structure
	var locationsResponse LocationsResponse
	if err := database.FetchData("https://groupietrackers.herokuapp.com/api/locations", &locationsResponse); err != nil {
		log.Printf("Error fetching locations data: %v", err)
		ErrorPage(w, "Error fetching locations data", http.StatusInternalServerError)
		return
	}

	// Create a map of artist IDs to their locations for quick lookup
	artistLocations := make(map[int][]string)
	for _, loc := range locationsResponse.Index {
		artistLocations[loc.ID] = loc.Locations
	}

	// Find matching artists or members
	var matchedArtists []models.Artist
	addedArtists := make(map[int]bool) // To prevent duplicate artists in results

	for _, artist := range artists {
		// Check if artist should be added based on existing criteria
		if shouldAddArtist(artist, query) && !addedArtists[artist.ID] {
			matchedArtists = append(matchedArtists, artist)
			addedArtists[artist.ID] = true
			continue
		}

		// Check locations for this artist
		if locations, exists := artistLocations[artist.ID]; exists {
			for _, location := range locations {
				if strings.Contains(strings.ToLower(location), query) && !addedArtists[artist.ID] {
					matchedArtists = append(matchedArtists, artist)
					addedArtists[artist.ID] = true
					break
				}
			}
		}
	}

	// Prepare the data for the template
	data := models.PageData{
		Artists: matchedArtists,
		Query:   query,
	}

	// Parse the template and display the search results
	tmpl, err := template.ParseFiles("templates/search_results.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		ErrorPage(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
		ErrorPage(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

// shouldAddArtist checks if an artist matches the search query based on the original criteria
func shouldAddArtist(artist models.Artist, query string) bool {
	if strings.Contains(strings.ToLower(artist.Name), query) ||
		strings.Contains(strings.ToLower(artist.Locations), query) ||
		strings.Contains(strings.ToLower(artist.FirstAlbum), query) ||
		strings.Contains(strconv.Itoa(artist.CreationDate), query) {
		return true
	}

	for _, member := range artist.Members {
		if strings.Contains(strings.ToLower(member), query) {
			return true
		}
	}

	return false
}
