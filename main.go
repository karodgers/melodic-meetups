package main

import (
	"log"
	"net/http"

	"groupie-trackers/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HandleHomepage)
	http.HandleFunc("/artist/", handlers.HandleArtistDetail)
	http.HandleFunc("/artist/dates/", handlers.HandleConcertDates)
	http.HandleFunc("/login", handlers.HandleLogin)
	http.HandleFunc("/signup", handlers.HandleSignup)
	http.HandleFunc("/about", handlers.HandleAbout)
	http.HandleFunc("/search", handlers.HandleSearchResults)
	http.HandleFunc("/geolocation", handlers.HandleGeolocation)
	http.HandleFunc("/handle-search", handlers.HandleSearch)
	http.Handle("/api/proxy", handlers.CorsMiddleware(http.HandlerFunc(handlers.ProxyHandler)))

	// Serve static files
	http.HandleFunc("/static/", handlers.ServeDir("static"))
	http.HandleFunc("/images/", handlers.ServeDir("images"))
	http.HandleFunc("/api/", handlers.ServeDir("api"))

	log.Println("Server started at http://localhost:9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
