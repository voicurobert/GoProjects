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
	muxDispatcher = mux.NewRouter()
)

//NewMuxRouter returns a new muxRouter
func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(rw http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func(rw http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP server running on port %v", port)
	err := http.ListenAndServe(port, muxDispatcher)
	if err != nil {
		return
	}
}
