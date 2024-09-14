package main

import (
	"fmt"
	"log"
	"net/http"

	groupie "groupie-tracker/func"
)

func main() {
	groupie.Isfitch = groupie.Fetch("artists", "")
	http.HandleFunc("/style/", groupie.Style)
	http.HandleFunc("/", groupie.Gethandel)
	http.HandleFunc("/artist/{id}", groupie.Posthandel)
	fmt.Println("http://localhost:8083")
	log.Fatal(http.ListenAndServe(":8083", nil))

}
