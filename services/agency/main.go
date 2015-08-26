package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gorilla/mux"
)

func main() {
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/stream", hystrixStreamHandler.ServeHTTP)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router))
}

func index(w http.ResponseWriter, req *http.Request) {
	hystrix.Do("assign_guide", assignGuide(w), fallback(w))
}

func assignGuide(w io.Writer) func() error {
	return func() error {
		resp, err := http.Get("http://company.cfapps.pez.pivotal.io/available")
		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("Error requesting guide. HTTP %d", resp.StatusCode)
		}

		guide, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return err
		}

		fmt.Fprintf(w, "Your guide will be: %s", guide)
		return nil
	}
}

func fallback(w io.Writer) func(error) error {
	return func(error) error {
		fmt.Fprint(w, "We will get back to you shortly.")
		return nil
	}
}
