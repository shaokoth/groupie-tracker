package handlers

import (
	"html/template"
	"net/http"

	"groupie-tracker/api"
)

<<<<<<< HEAD:handlers/artistshandler.go
func Artistshandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "Method Not Allowed", "Error", "Reload")
	}

=======
func Artisthandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "Method Not Allowed", "Error", "Reload")
	}
>>>>>>> 1fc0de3f1a7dea6bd1038a14477f9970727df077:handlers/artisthandler.go
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
