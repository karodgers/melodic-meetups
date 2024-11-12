package database

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FetchData sends an HTTP GET request to the given URL and decodes the response JSON into the provided target.
// The 'target' is an interface{}, so any kind of data structure can be passed to hold the decoded data.
func FetchData(url string, target interface{}) error {
	// Send an HTTP GET request to the specified URL
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch data from URL %s: %w", url, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	// Decode the response body from JSON format and store it in the target variable
	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		return fmt.Errorf("failed to decode JSON: %w", err)
	}
	
	return nil
}
