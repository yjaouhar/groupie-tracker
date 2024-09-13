package groupie

import (
	"html/template"
	"net/http"
	"strconv"
)

func Posthandel(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	IDa, er := strconv.Atoi(id)
	if er != nil || IDa <= 0 || IDa > 52 {
		Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	if r.URL.Path != "/artist/"+id {
		Error(w, "Page Not Found", http.StatusNotFound)
		return
	}
	temp, err := template.ParseFiles("template/index2.html")
	if err != nil {
		Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	Cards.Art = Artist[IDa-1]
	Fetch(w,"locations", "/"+id)
	Fetch(w,"dates", "/"+id)
	Fetch(w,"relation", "/"+id)

	ExecuteTemplate(temp, w)

}