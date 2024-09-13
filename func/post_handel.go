package groupie

import (
	"html/template"
	"net/http"
	"strconv"
)

func Posthandel(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	IDa, er := strconv.Atoi(id)
	if er != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	} else {
		if IDa <= 0 || IDa > 52 {
			http.Error(w, "mok zina ", http.StatusNotFound)
			return
		}
	}

	if r.URL.Path != "/artist/"+id {
		http.Error(w, "404 Page not found", http.StatusNotFound)
		return
	}
	temp, err := template.ParseFiles("template/index2.html")
	if err != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}

	Cards.Art = Artist[IDa-1]
	Fetch("locations", "/"+id)
	Fetch("dates", "/"+id)
	Fetch("relation", "/"+id)

	Er(temp, w)

	//	temp.Execute(w, Cards)

}
