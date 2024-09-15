package groupie

import (
	"net/http"
)

func Style(w http.ResponseWriter, r *http.Request) {
	style := http.StripPrefix("/style/", http.FileServer(http.Dir("style")))

	style.ServeHTTP(w, r)

}
