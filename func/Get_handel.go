package groupie

import (
	"html/template"
	"net/http"
)

// Get home page 
func  Gethandel(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		Error(w, http.StatusNotFound)
		return
	}
	if !Isfitch {
		Error(w, http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodGet {
		Error(w, http.StatusMethodNotAllowed)
		return
	}

	temp, err := template.ParseFiles("template/index.html")

	if err != nil {
		Error(w, http.StatusInternalServerError)
		return
	}
	ExecuteTemplate(temp, "artist", w, nil, 0) 
}
