package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello world")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oooooooo", http.StatusBadRequest)
			/*
				rw.WriteHeader(http.StatusBadRequest)
				rw.Write([]byte("Ooooooooo"))
			*/
			return
		}

		fmt.Fprintf(rw, "Hello %s world \n", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye world")
	})

	http.ListenAndServe(":9090", nil)
}
