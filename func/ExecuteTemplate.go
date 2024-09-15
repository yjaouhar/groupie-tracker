package groupie

import (
	"bytes"
	"html/template"
	"net/http"
)

// execute template
func ExecuteTemplate(temp *template.Template, s string, w http.ResponseWriter, info interface{}, status int) {
	var buf bytes.Buffer
	var err error
	if s == "artist" {
		err = temp.Execute(&buf, Artist)
		if err != nil {
			Error(w, http.StatusInternalServerError)
			return
		}
		temp.Execute(w, Artist)
	} else if s == "err" {
		err = temp.Execute(&buf, info)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(status)
		temp.Execute(w, info)
	} else {
		err = temp.Execute(&buf, Cards)
		if err != nil {
			Error(w, http.StatusInternalServerError)
			return
		}
		temp.Execute(w, Cards)
	}

}
