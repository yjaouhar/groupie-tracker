package groupie

import (
	"html/template"
	"net/http"
)

func Gethandel(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "405 Methode not allow", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, "404 page not found ", http.StatusNotFound)
	}
	temp, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "500 Internal server error", http.StatusInternalServerError)
	}
	Fetch("artists","")

	temp.Execute(w, Cards)
}
