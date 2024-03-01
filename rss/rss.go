package rss

import (
	"encoding/json"
	"net/http"

)

type RequestParams struct {
	ChronologicalPreference string `json:"chronological_preference"`
}

func Compile(w http.ResponseWriter, r *http.Request) {
	var params RequestParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Your logic based on chronological preference goes here
	// For demonstration, let's just print the preference
	w.Write([]byte("More Articles Handler. Chronological Preference: " + params.ChronologicalPreference))
}

func MoreEventsHandler(w http.ResponseWriter, r *http.Request) {
	var params RequestParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Your logic based on chronological preference goes here
	// For demonstration, let's just print the preference
	w.Write([]byte("More Events Handler. Chronological Preference: " + params.ChronologicalPreference))
}
