package handlers

import (
	"html/template"
	"net/http"

	"groupie-tracker/api"
)

func Locationhandler(w http.ResponseWriter, r *http.Request) {
	locations, err := api.DecodeLocations()
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error", "Error", "Reload")
		return
	}

	t, err := template.ParseFiles("templates/locations.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error", "Error", "Reload")
		return
	}

	t.Execute(w, locations)
}
