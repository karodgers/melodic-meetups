package handlers

import (
	"html/template"
	"net/http"
	"net/url"
)

// HandleGeolocation processes the geolocation of a concert
func HandleGeolocation(w http.ResponseWriter, r *http.Request) {
	// Parse the location from query parameters
	location := r.URL.Query().Get("location")
	if location == "" {
		ErrorPage(w, "Location is missing", http.StatusBadRequest)
		return
	}

	// URL-decode the location
	location, err := url.QueryUnescape(location)
	if err != nil {
		ErrorPage(w, "Invalid location", http.StatusBadRequest)
		return
	}

	// Hardcoded Google Maps API key
	apiKey := "guT4g7UQnNdnfC4FixSO7DqpjJBDRFsL7v89YXkUMvk"

	// Render the map template
	tmpl := template.Must(template.ParseFiles("templates/geolocation.html"))
	data := struct {
		Location string
		APIKey   string
	}{
		Location: location,
		APIKey:   apiKey,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		ErrorPage(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
