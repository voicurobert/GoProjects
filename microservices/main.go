package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"github.com/voicurobert/GoProjects/microservices/handlers"
)

func oldHandlers() {
	/*
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
	*/
}

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	//hh := handlers.NewHello(l)
	//gh := handlers.NewGoodbye(l)
	ph := handlers.NewProducts(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods("GET").Subrouter()

	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.MiddlewareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidation)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan

	l.Println("Received terminate, graceful shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(tc)

	// web server, if handler is nill, the DefaultServeMux is used
	//http.ListenAndServe(":9090", sm)
}
