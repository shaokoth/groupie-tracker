package handlers

import (
	"html/template"
	"net/http"

	"groupie-tracker/api"
)
// Artistshandler handles the HTTP requests for the artists' page.
// It responds to GET requests by fetching artist data and rendering the artists.html template.
func Artistshandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "Method Not Allowed", "Error", "Reload")
	}

	artists, err := api.DecodeArtists()
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error", "Error", "Reload")
		return
	}

	t, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error", "Error", "Reload")
		return
	}
	t.Execute(w, artists)
}
