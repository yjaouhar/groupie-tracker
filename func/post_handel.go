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
		Error(w, http.StatusNotFound)
		return
	}

	if r.URL.Path != "/artist/"+id {
		Error(w, http.StatusNotFound)
		return
	}
	temp, err := template.ParseFiles("template/index2.html")
	if err != nil {
		Error(w, http.StatusInternalServerError)
		return
	}

	Cards.Art = Artist[IDa-1]
	if !Fetch("locations", "/"+id) {
		Error(w, http.StatusInternalServerError)
		return
	}
	if !Fetch("dates", "/"+id) {
		Error(w, http.StatusInternalServerError)
		return
	}
	if !Fetch("relation", "/"+id) {
		Error(w, http.StatusInternalServerError)
		return
	}

	ExecuteTemplate(temp, "", w, nil)

}
