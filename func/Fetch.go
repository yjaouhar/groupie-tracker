package groupie

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Fetch(s, id string) {

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
	//err = json.Unmarshal(data, &Cards.Loca)
	switch s {
	case "artists":
		err = json.Unmarshal(data, &Cards.Art)
	case "locations":
		err = json.Unmarshal(data, &Cards.Loca)
	case "dates":
		err = json.Unmarshal(data, &Cards.Conc)
	case "relation":
		err = json.Unmarshal(data, &Cards.Rela)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(Cards.Rela)
}
