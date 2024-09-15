package main

import (
	"fmt"
	"log"
	"net/http"

	groupie "groupie-tracker/func"
)

func main() {
	groupie.Isfitch = groupie.Fetch("artists", "") // get ipa data and fitched
	http.HandleFunc("/style/style.css", groupie.Style)
	http.HandleFunc("/style/background.png.jpg", groupie.Style)
	http.HandleFunc("/", groupie.Gethandel)
	http.HandleFunc("/artist/{id}", groupie.Posthandel)
	fmt.Println("http://localhost:5461")
	log.Fatal(http.ListenAndServe(":5461", nil))

}
