package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//MuxRouter struct implements Router interface
type muxRouter struct {
}

var (
	muxDispacher = mux.NewRouter()
)

//NewMuxRouter return a new muxRouter
func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(rw http.ResponseWriter, r *http.Request)) {
	muxDispacher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func(rw http.ResponseWriter, r *http.Request)) {
	muxDispacher.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP server running on port %v", port)
	http.ListenAndServe(port, muxDispacher)
}
