package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	groupieHandler("templates/details.html", w, r)
}

func DatesHandler(w http.ResponseWriter, r *http.Request) {
	groupieHandler("templates/dates.html", w, r)
}

func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	groupieHandler("templates/locations.html", w, r)
}

func ConcertsHandler(w http.ResponseWriter, r *http.Request) {
	groupieHandler("templates/concerts.html", w, r)
}

func groupieHandler(tpl string, w http.ResponseWriter, r *http.Request) {
	artist, err := Artist(r)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error", "Error", "Reload")
		return
	}

	if artist.ArtistName == "Not Found" {
		errTxt := "Oops! The page you are looking for does not exist\n"
		ErrorHandler(w, r, http.StatusNotFound, errTxt, "404 Not Found", "Artists")
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
