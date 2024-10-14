package handlers

import (
	"net/http"
)
// PathHandler handles incoming HTTP requests based on the request URL path.
// It directs requests to the appropriate handler function or returns a 404 error if the path is not recognized.
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
