package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router))
}

func index(w http.ResponseWriter, req *http.Request) {
	name := guide()
	fmt.Fprintf(w, "Your guide will be: %s", name)
}

func guide() string {
	resp, err := http.Get("http://company.cfapps.pez.pivotal.io/available")
	if err != nil {
		return ""
	}

	guide, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return ""
	}

	return string(guide)
}
