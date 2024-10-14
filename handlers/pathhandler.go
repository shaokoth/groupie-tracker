package handlers

import (
	"net/http"
)

func PathHandler(w http.ResponseWriter, r *http.Request) {
	filepath := r.URL.Path
	switch filepath {
	case "/":
		Artistshandler(w, r)
	default:
		errTxt := "Oops! The page you are looking for does not exist\n."
		ErrorHandler(w, r, http.StatusNotFound, errTxt, "404 Not Found", "Artists")
	}
}
