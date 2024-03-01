package lib

import (
	"fmt"
	"io"
	"net/http"

	"github.com/hossam1231/logger-go-pkg"
)

func GeoSearchHandler(w http.ResponseWriter, r *http.Request) {
	// Get the location from the query parameters
	location := r.URL.Query().Get("location")

	// Your Mapbox API token
	mapboxToken := "pk.eyJ1IjoibW9zcXVlaWN1IiwiYSI6ImNsaHhhb3MxczBzN2YzZ3BnNHRkMW9rdHIifQ.csS9yZXj5lD3BIw-Kcw6TQ"

	// Construct the URL for Mapbox API
	url := fmt.Sprintf("https://api.mapbox.com/geocoding/v5/mapbox.places/%s.json?proximity=ip&access_token=%s", location, mapboxToken)

	// Send a GET request to Mapbox API
	resp, err := http.Get(url)
	if err != nil {
		// Log the error
		logger.Error(fmt.Sprint(err))
		// Respond with an internal server error
		http.Error(w, "Failed to fetch data from Mapbox API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")
	// Write the response status code
	w.WriteHeader(resp.StatusCode)

	// Copy the response body to the response writer
	if _, err := io.Copy(w, resp.Body); err != nil {
		// Respond with an internal server error if copying fails
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
