package main

import (
	"fmt"
	"net/http"

	"groupie-tracker/handlers"
)

func main() {
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	
	http.HandleFunc("/", handlers.PathHandler)
	http.HandleFunc("/details", handlers.DetailsHandler)
	// http.HandleFunc("/locations", handlers.Locationhandler)
	// http.HandleFunc("/relations", handlers.Relationthandler)

	fmt.Println("Server running on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
