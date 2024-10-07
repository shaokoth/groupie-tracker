package handlers

import (
	"html/template"
	"net/http"
)

type Error struct {
	ErrNo   int
	ErrText string
	ErrTitle     string
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, errNo int, errTxt, errTitle string) {
	err := Error{ErrNo: errNo, ErrText: errTxt, ErrTitle: errTitle}

	tmp, _ := template.ParseFiles("templates/errors.html")

	tmp.Execute(w, err)
}
