package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// registers a function to a path on a DefaultServeMux (http handler)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello world")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Ooops", http.StatusBadRequest)
			//rw.WriteHeader(http.StatusBadRequest)
			//rw.Write([]byte("Ooops"))
			return
		}
		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye world")
	})

	// web server, if handler is nill, the DefaultServeMux is used
	http.ListenAndServe(":9090", nil)
}
