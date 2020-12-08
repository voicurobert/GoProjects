package handlers

import (
	"log"
	"net/http"
)

// Goodbye struct handler
type Goodbye struct {
	l *log.Logger
}

// NewGoodbye returns a new Goodbye handler
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Byeeee"))
}
