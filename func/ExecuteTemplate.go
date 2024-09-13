package groupie

import (
	"bytes"
	"html/template"
	"net/http"
)

func ExecuteTemplate(temp *template.Template, w http.ResponseWriter) {
	var buf bytes.Buffer
	err := temp.Execute(&buf, Cards)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	temp.Execute(w, Cards)

}
