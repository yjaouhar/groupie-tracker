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
	http.HandleFunc("/index1", groupie.Posthandel)
	fmt.Println("http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))

}
