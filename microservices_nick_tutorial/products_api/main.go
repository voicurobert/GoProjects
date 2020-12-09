package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/nicholasjackson/env"

	gohandlers "github.com/gorilla/handlers"
	"github.com/voicurobert/GoProjects/microservices_nick_tutorial/products_api/data"
	"github.com/voicurobert/GoProjects/microservices_nick_tutorial/products_api/handlers"
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

	env.Parse()

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	//hh := handlers.NewHello(l)
	//gh := handlers.NewGoodbye(l)
	v := data.NewValidation()

	// create the handlers
	ph := handlers.NewProducts(l, v)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	// handlers for API
	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/products", ph.ListAll)
	getR.HandleFunc("/products/{id:[0-9]+}", ph.ListSingle)

	putR := sm.Methods(http.MethodPut).Subrouter()
	putR.HandleFunc("/products", ph.Update)
	putR.Use(ph.MiddlewareValidateProduct)

	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/products", ph.Create)
	postR.Use(ph.MiddlewareValidateProduct)

	deleteR := sm.Methods(http.MethodDelete).Subrouter()
	deleteR.HandleFunc("/products/{id:[0-9]+}", ph.Delete)

	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getR.Handle("/docs", sh)
	getR.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	corsHandler := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:9090"}))

	s := &http.Server{
		Addr:         ":9090",
		Handler:      corsHandler(sm),
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
