package handlers

import (
	"html/template"
	"log"
	"net/http"

	database "groupie-trackers/data"
	"groupie-trackers/models"
)

// HandleConcertDates processes the request to view concert dates for an artist.
func HandleConcertDates(w http.ResponseWriter, r *http.Request) {
	// Fetch the artist's concert dates from the external API
	var datesResponse models.DatesResponse
	if err := database.FetchData("https://groupietrackers.herokuapp.com/api/dates", &datesResponse); err != nil {
		log.Printf("Error fetching dates data: %v", err)
		ErrorPage(w, "Error fetching dates data", http.StatusInternalServerError)
		return
	}

	// Prepare a slice to hold the concert dates
	var concertDates []models.Date
	concertDates = append(concertDates, datesResponse.Index...)

	// Prepare the data for the template
	data := struct {
		Dates []models.Date
	}{
		Dates: concertDates,
	}

	// Parse the "dates.html" template file
	tmpl, err := template.ParseFiles("templates/dates.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		ErrorPage(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	// Execute the parsed template with the prepared data
	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
		ErrorPage(w, "Error executing template", http.StatusInternalServerError)
	}
}
