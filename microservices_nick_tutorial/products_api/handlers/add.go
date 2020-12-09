package handlers

import (
	"net/http"

	"github.com/voicurobert/GoProjects/microservices_nick_tutorial/products_api/data"
)

// AddProduct func
func (p Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(prod)
}
