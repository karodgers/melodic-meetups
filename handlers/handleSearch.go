package handlers

import (
	"net/http"
	"net/url"
	"strings"
)

// HandleSearch handles the search requests and redirects to the results page.
func HandleSearch(w http.ResponseWriter, r *http.Request) {
	// Extract search query from request parameters
	query := r.URL.Query().Get("query")
	query = strings.ToLower(query)

	if query == "" {
		ErrorPage(w, "Search query is required", http.StatusBadRequest)
		return
	}

	// Redirect to the search results page with the query
	http.Redirect(w, r, "/search?query="+url.QueryEscape(query), http.StatusSeeOther)
}
