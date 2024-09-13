package groupie

import (
	"html/template"
	"net/http"
)

func Posthandel(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 Methode not allow", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/index2.html" {
		http.Error(w, "404 Page not found", http.StatusNotFound)
		return
	}

	temp, err := template.ParseFiles("template/index2.html")
	if err != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}
	Fetch("artists", "/4")
	temp.Execute(w, Cards)

}
