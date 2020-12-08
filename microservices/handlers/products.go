package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/voicurobert/GoProjects/microservices/data"
)

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The ID from the Product to delete from the database
	//in: path
	//required: true
	ID int `json:"id"`
}

// KeyProduct is a key used for the Product object in the context
type KeyProduct struct{}

// Products handler
type Products struct {
	l *log.Logger
	v *data.Validation
}

// ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("Invalid Path, path should be /products/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// NewProducts creates a new Products handler
func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}

// getProductID returns the product ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getProductID(r *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}
