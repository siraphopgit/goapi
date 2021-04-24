package handlers

import (
	"net/http"

	"github.com/siraphopgit/building-microservices-youtube/product-api/data"
)

// swagger:route GET /products products listProducts
// Return a list of products
// responses:
// 	200: productsResponse

// getProducts returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
