package handlers

import (
	"net/http"

	"github.com/govipul/GoMicroservices/episode_8/data"
)

// swagger:route POST /products/ products addProduct
//
// responses:
//	200: noContent

//AddProduct add the products
func (p Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle a POST product")

	/*prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusBadRequest)
		return
	}*/
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(&prod)
}
