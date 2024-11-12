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

// HandleArtistDetail processes the request to view detailed information about an artist.
func HandleArtistDetail(w http.ResponseWriter, r *http.Request) {
	// Extract artist ID from the URL by trimming the "/artist/" prefix
	idStr := strings.TrimPrefix(r.URL.Path, "/artist/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorPage(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}
	// Fetch all artist data from the external API
	var artists []models.Artist
	if err := database.FetchData("https://groupietrackers.herokuapp.com/api/artists", &artists); err != nil {
		log.Printf("Error fetching artists data: %v", err)
		ErrorPage(w, "Error fetching artists data", http.StatusInternalServerError)
		return
	}
	// Find the specific artist based on the ID in the request
	var artist models.Artist
	for _, a := range artists {
		if a.ID == id {
			artist = a // Store the found artist
			break
		}
	}
	if artist.ID == 0 {
		ErrorPage(w, "Artist not found", http.StatusNotFound)
		return
	}
	// Fetch the artist's location data from the external API
	var locationsResponse models.LocationsResponse
	if err := database.FetchData("https://groupietrackers.herokuapp.com/api/locations", &locationsResponse); err != nil {
		log.Printf("Error fetching locations data: %v", err)
		ErrorPage(w, "Error fetching locations data", http.StatusInternalServerError)
		return
	}
	// Find the specific location based on the artist's ID
	var location models.Location
	for _, loc := range locationsResponse.Index {
		if loc.ID == id {
			location = loc // Store the found location
			break
		}
	}
	// Fetch the artist's concert dates from the external API
	var datesResponse models.DatesResponse
	if err := database.FetchData("https://groupietrackers.herokuapp.com/api/dates", &datesResponse); err != nil {
		log.Printf("Error fetching dates data: %v", err)
		ErrorPage(w, "Error fetching dates data", http.StatusInternalServerError)
		return
	}
	// Find the specific concert date based on the artist's ID
	var date models.Date
	for _, d := range datesResponse.Index {
		if d.ID == id {
			date = d // store the found concert date
			break
		}
	}
	// Fetch the relations (dates to locations mapping) from the external API
	var relationsResponse models.RelationsResponse
	if err := database.FetchData("https://groupietrackers.herokuapp.com/api/relation", &relationsResponse); err != nil {
		log.Printf("Error fetching relations data: %v", err)
		ErrorPage(w, "Error fetching relations data", http.StatusInternalServerError)
		return
	}
	// Find the specific relation data based on the artist's ID
	var relation models.Relation
	for _, rel := range relationsResponse.Index {
		if rel.ID == id {
			relation = rel // Store the found relation
			break
		}
	}
	// Prepare the data that will be passed to the template
	data := models.ArtistDetailData{
		Artist:   artist,
		Location: location,
		Date:     date,
		Relation: relation,
	}
	// Parse the "artist.html" template file
	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		ErrorPage(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	// Execute the parsed template with the prepared data and write the response to the client
	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
		ErrorPage(w, "Error executing template", http.StatusInternalServerError)
	}
}
