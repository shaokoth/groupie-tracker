package main

import (
	"net/http"

	"groupie-tracker/handlers"
)

func main() {
	// fmt.Println(string(api.FetchAPI("https://groupietrackers.herokuapp.com/api/artists")))
	http.HandleFunc("/", handlers.Artisthandler)
	http.HandleFunc("/dates", handlers.Datehandler)
	http.HandleFunc("/locations", handlers.Locationhandler)
	http.HandleFunc("/relations", handlers.Relationthandler)

	http.ListenAndServe(":8080", nil)
}
