package groupie

import (
	"html/template"
	"net/http"
)

func Gethandel(w http.ResponseWriter, r *http.Request) {
	Fetch(w, "artists", "")
	if r.URL.Path != "/" {
		http.Error(w, "404 page not found ", http.StatusNotFound)
	}

	if r.Method != http.MethodGet {
		http.Error(w, "405 Methode not allowed", http.StatusMethodNotAllowed)
		return
	}

	temp, err := template.ParseFiles("template/index.html")

	if err != nil {
		http.Error(w, "500 Internal server error", http.StatusInternalServerError)
	}

	temp.Execute(w, Artist)
}
