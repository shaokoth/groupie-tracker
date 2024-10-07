package handlers

import (
	"html/template"
	"net/http"

	"groupie-tracker/api"
)

func Artisthandler(w http.ResponseWriter, r *http.Request) {
	artists, err := api.DecodeArtists()
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error", "Error", "Reload")
		return
	}

	t, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server", "Error", "Reload")
		return
	}

	t.Execute(w, artists)
}
