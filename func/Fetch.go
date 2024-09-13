package groupie

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Fetch(w http.ResponseWriter, s, id string) {
	respons, err := http.Get(Url + s + id)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := io.ReadAll(respons.Body)
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
		return
	}
}
