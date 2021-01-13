package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type chiRouter struct {
}

var (
	chiDispacher = chi.NewRouter()
)

//NewChiRouter constructor
func NewChiRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) GET(uri string, f func(rw http.ResponseWriter, r *http.Request)) {
	chiDispacher.Get(uri, f)
}

func (*chiRouter) POST(uri string, f func(rw http.ResponseWriter, r *http.Request)) {
	chiDispacher.Post(uri, f)
}

func (*chiRouter) SERVE(port string) {
	fmt.Printf("Chi HTTP server running on port %v", port)
	http.ListenAndServe(port, chiDispacher)
}