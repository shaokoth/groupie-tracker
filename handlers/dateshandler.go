package handlers

import (
	"html/template"
	"net/http"

	"groupie-tracker/api"
)

func Datehandler(w http.ResponseWriter, r *http.Request) {
	dates, err := api.DecodeDates()
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error", "Error", "Reload")
		return
	}

	t, err := template.ParseFiles("templates/dates.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error", "Error", "Reload")
		return
	}

	t.Execute(w, dates)
}
