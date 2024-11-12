package handlers

import (
	"net/http"
	"os"
	"path/filepath"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorPage(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	filePath := filepath.Join("templates", "login.html")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		ErrorPage(w, "Login page not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, filePath)
}
