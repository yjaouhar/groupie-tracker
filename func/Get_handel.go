package groupie

import (
	"html/template"
	"net/http"
)

func Gethandel(w http.ResponseWriter, r *http.Request) {
	Fetch(w, "artists", "")
	if r.URL.Path != "/" {
		Error(w, "Page Not Found", http.StatusNotFound)
	}

	if r.Method != http.MethodGet {
		Error(w, "Methode not allowed", http.StatusMethodNotAllowed)
		return
	}

	temp, err := template.ParseFiles("template/index.html")

	if err != nil {
		Error(w, "Internal server error", http.StatusInternalServerError)
	}

	temp.Execute(w, Artist)
}
