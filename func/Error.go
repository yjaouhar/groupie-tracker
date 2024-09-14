package groupie

import (
	"html/template"
	"net/http"
)

func Error(w http.ResponseWriter, status int) {
	tmp, err := template.ParseFiles("template/error.html")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	
	st_temp := Err{
		Status:  status,
		Message: http.StatusText(status),
	}
	ExecuteTemplate(tmp, "err", w, st_temp, status)
}
