package handlers

import (
	"net/http"
	"os"
	"path/filepath"
)

// The HandleAbout function is used to get the about page
func HandleAbout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorPage(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	filePath := filepath.Join("templates", "about.html")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		ErrorPage(w, "About page not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, filePath)
}
