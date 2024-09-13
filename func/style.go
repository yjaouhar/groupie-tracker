package groupie

import (
	"net/http"
	"strings"
)

func Style(w http.ResponseWriter, r *http.Request) {

	style := http.StripPrefix("/style/", http.FileServer(http.Dir("style")))
	if strings.HasSuffix(r.URL.Path, "/") {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
	style.ServeHTTP(w, r)

}
