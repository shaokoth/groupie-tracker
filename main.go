package main

import (
	"fmt"
	"net/http"

	"groupie-tracker/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Artisthandler)
	http.HandleFunc("/details", handlers.DetailsHandler)
	// http.HandleFunc("/locations", handlers.Locationhandler)
	// http.HandleFunc("/relations", handlers.Relationthandler)

	fmt.Println("Server running on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
