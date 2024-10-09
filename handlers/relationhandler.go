package handlers

import (
	"html/template"
	"net/http"

	"groupie-tracker/api"
)

func Relationthandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "Method Not Allowed", "Error", "Reload")
	}
	relations, err := api.DecodeRelations()
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error", "Error", "Reload")
		return
	}

	t, err := template.ParseFiles("templates/relations.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error", "Error", "Reload")
		return
	}

	t.Execute(w, relations)
}
