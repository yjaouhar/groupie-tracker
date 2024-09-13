package groupie

import (
	"html/template"
	"net/http"
)

func Error( w http.ResponseWriter, message string, status int) {
	tmp, err := template.ParseFiles("static/error.html")
	if err != nil {
		http.Error(w, "enternal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	st_temp := Err{
		Status: status,
		Message: message,
	}
	tmp.Execute(w, st_temp)
}