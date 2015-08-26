package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var guides = []string{
	"Melchior",
	"Gaspar",
	"Balthazar",
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/available", available)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router))
}

func index(w http.ResponseWriter, req *http.Request) {}

func available(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, guides[rand.Intn(len(guides))])
}
