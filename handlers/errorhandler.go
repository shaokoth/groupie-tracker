package handlers

import (
	"html/template"
	"net/http"
)

type Error struct {
	ErrNo    int
	ErrText  string
	ErrTitle string
	ErrButton string
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, errNo int, errTxt, errTitle, errBtn string) {
	errData := Error{ErrNo: errNo, ErrText: errTxt, ErrTitle: errTitle, ErrButton: errBtn}

	tmp, err := template.ParseFiles("templates/errors.html")
	if err != nil {
		http.Error(w,"Internal Server Error", http.StatusInternalServerError)
	}

	tmp.Execute(w, errData)
}
