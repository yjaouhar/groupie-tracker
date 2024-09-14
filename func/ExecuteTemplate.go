package groupie

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

func ExecuteTemplate(temp *template.Template, s string, w http.ResponseWriter, mok interface{}) {
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
		fmt.Println(mok)
		err = temp.Execute(&buf, mok)
		if err != nil {
			http.Error(w, "mok", http.StatusInternalServerError)
			//w.Write([]byte("Internal server error"))
			return
		}
		temp.Execute(w, mok)
	} else {
		err = temp.Execute(&buf, Cards)
		if err != nil {
			Error(w, http.StatusInternalServerError)
			return
		}
		temp.Execute(w, Cards)
	}

}
