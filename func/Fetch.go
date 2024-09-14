package groupie

import (
	"encoding/json"
	"io"
	"net/http"
)

func Fetch( s, id string) bool{
	response, err := http.Get(Url +s + id)
	if err != nil {
		return false
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK{
		if err != nil {

			return false
		}
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return false
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
	}
	if err != nil {
		return false
	}
	return true
}
