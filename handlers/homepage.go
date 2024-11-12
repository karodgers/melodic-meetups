package handlers

import (
	"html/template"
	"log"
	"net/http"

	database "groupie-trackers/data"
	"groupie-trackers/models"
)

// The HandleHomepage function is responsible for handling the homepage by fetching data from the API server and also parsing
// the index page from the templates folder.
func HandleHomepage(w http.ResponseWriter, r *http.Request) {
	var artists []models.Artist
	if r.URL.Path != "/" {
		ErrorPage(w, "Page not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		ErrorPage(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := database.FetchData("https://groupietrackers.herokuapp.com/api/artists", &artists); err != nil {
		log.Printf("Error fetching artists data: %v", err)
		ErrorPage(w, "Error fetching artists data", http.StatusInternalServerError)
		return
	}

	data := models.PageData{
		Artists: artists,
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		ErrorPage(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
		ErrorPage(w, "Error executing template", http.StatusInternalServerError)
	}
}
