package main

import (
	"log"
	"net/http"

	"groupie-tracker/handlers"
)

func main() {
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", handlers.PathHandler)
	http.HandleFunc("/details", handlers.DetailsHandler)
	http.HandleFunc("/dates", handlers.DatesHandler)
	http.HandleFunc("/locations", handlers.LocationsHandler)
	http.HandleFunc("/concerts", handlers.ConcertsHandler)

	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
