package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/siraphopgit/building-microservices-youtube/product-api/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Return a list of products
// responses:
// 	201: noContent

// DeleteProduct deletes a product from the database
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	// this will alwats convert because of the router
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	p.l.Println("Handle DELETE PRoduct", id)

	err := data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
