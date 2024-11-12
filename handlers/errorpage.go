package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type ErrorData struct {
	StatusCode   int
	ErrorMessage string
}

// The ErrorPage function displays error pages with the error messages and status codes.
func ErrorPage(w http.ResponseWriter, errorMessage string, statusCode int) {
	// Parse the error template
	t, err := template.ParseFiles("templates/error.html")
	if err != nil {
		log.Printf("Error parsing error template: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the status code and create error data
	w.WriteHeader(statusCode)
	errorData := ErrorData{
		StatusCode:   statusCode,
		ErrorMessage: errorMessage,
	}

	// Execute the template
	err = t.Execute(w, errorData)
	if err != nil {
		log.Printf("Error executing error template: %v\n", err)
	}
}
