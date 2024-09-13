package groupie

import (
	"encoding/json"
	"io"
	"net/http"
)

func Fetch(w http.ResponseWriter, s, id string) {
	respons, err := http.Get(Url + s + id)
	if err != nil {
		Error(w, "Failed to make GET request", http.StatusInternalServerError)
		return
	}
	defer respons.Body.Close()

	data, err := io.ReadAll(respons.Body)
	if err != nil {
		Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}
	if s == "artists" && !Fetched {
		Fetched = true
		err = json.Unmarshal(data, &Artist)
	} else if s == "locations" {
		err = json.Unmarshal(data, &Cards.Loca)
	} else if s == "dates" {
		err = json.Unmarshal(data, &Cards.Conc)
	} else if s == "relation" {
		err = json.Unmarshal(data, &Cards.Rela)
	} else {
		return
	}
	if err != nil {
		Error(w, "Failed to unmarshal JSON response", http.StatusBadRequest)
		return
	}
}
