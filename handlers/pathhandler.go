package handlers

import (
	"net/http"
)

func PathHandler(w http.ResponseWriter, r *http.Request) {
	filepath := r.URL.Path

	switch filepath {
	case "/":
		Artisthandler(w, r)
	default:
		errTxt := "Paths cross empty void\nSeeking what once existed\nSilence answers all."
		ErrorHandler(w, r, http.StatusNotFound, errTxt, "404 Not Found", "Artist")
	}
}
