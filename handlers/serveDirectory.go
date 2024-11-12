package handlers

import (
	"net/http"
	"os"
	"path/filepath"
)

// serves files, prevents directory listings
func ServeDir(dir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filePath := filepath.Join(dir, r.URL.Path[len("/"+dir+"/"):])

		// Check if the requested path is a directory
		if info, err := os.Stat(filePath); err == nil && info.IsDir() {
			ErrorPage(w, "Access Not Allowed", http.StatusForbidden)
			return
		}
		// Serve the file
		http.StripPrefix("/"+dir+"/", http.FileServer(http.Dir(dir))).ServeHTTP(w, r)
	}
}
