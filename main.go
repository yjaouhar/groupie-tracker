package main

import (
	"fmt"
	"log"
	"net/http"

	groupie "groupie-tracker/func"
)

func main() {
	groupie.Isfitch = groupie.Fetch("artists", "") // get ipa data and fitched
	http.HandleFunc("/style/{file}", groupie.Style)
	http.HandleFunc("/", groupie.Gethandel)
	http.HandleFunc("/artist/{id}", groupie.Posthandel)
	fmt.Println("http://localhost:4061")
	log.Fatal(http.ListenAndServe(":4061", nil))

}
