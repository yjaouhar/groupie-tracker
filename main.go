package main

import (
	"fmt"
	"log"
	"net/http"

	groupie "groupie-tracker/func"
)

func main() {
	groupie.Fetch("artists", "")
	http.HandleFunc("/style/", groupie.Style)
	http.HandleFunc("/", groupie.Gethandel)
	http.HandleFunc("/artist/{id}", groupie.Posthandel)
	fmt.Println("http://localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", nil))

}
