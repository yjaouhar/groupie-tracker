package groupie

import (
	"net/http"
	"os"
)

func Style(w http.ResponseWriter, r *http.Request) {
	f := r.PathValue("file")
	style := http.StripPrefix("/style/", http.FileServer(http.Dir("style")))
	_, err := os.ReadFile("style/" + f)

	if err != nil {
		Error(w, http.StatusNotFound)
		return
	}

	style.ServeHTTP(w, r)
	
}
