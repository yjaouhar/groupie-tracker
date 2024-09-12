package main

import (
	"log"
	"net/http"

	groupie "groupie-tracker/func"
)

func main() {

	groupie.Fetch("relation", "/3")
	//fmt.Println(groupie.Cards.Loca)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
