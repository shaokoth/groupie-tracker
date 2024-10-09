package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	groupyHandler("templates/details.html", w, r)
}

func DatesHandler(w http.ResponseWriter, r *http.Request) {
	groupyHandler("templates/dates.html", w, r)
}

func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	groupyHandler("templates/locations.html", w, r)
}

func ConcertsHandler(w http.ResponseWriter, r *http.Request) {
	groupyHandler("templates/concerts.html", w, r)
}

func groupyHandler(tpl string, w http.ResponseWriter, r *http.Request) {
	artist, err := Artist(r)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error", "Error", "Reload")
		return
	}

	if artist.ArtistName == "Not Found" {
		errTxt := "Paths cross empty void\nSeeking what once existed\nSilence answers all."
		ErrorHandler(w, r, http.StatusNotFound, errTxt, "404 Not Found", "Artist")
		return
	}

	t, err := template.ParseFiles(tpl)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error", "Error", "Reload")
		fmt.Println(err)
		return
	}
	t.Execute(w, artist)

}
