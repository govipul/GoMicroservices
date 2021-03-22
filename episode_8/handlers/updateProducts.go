package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/govipul/GoMicroservices/episode_8/data"
)

// swagger:route PUT /products products updateProducts
// Returns a update of porduct
// responses:
//	200: productsResponse

// GetProducts get the list of all products from the data store

//UpdateProduct update the product data
func (p Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, errAtoI := strconv.Atoi(vars["id"])
	if errAtoI != nil {
		http.Error(rw, "Unable to unable to parse id", http.StatusBadRequest)
		return
	}
	p.l.Println("handle a PUT product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	dataErr := data.UpdateProduct(id, &prod)
	if dataErr == data.ErrProductNotFound {
		http.Error(rw, "Error while updating the product", http.StatusNotFound)
		return
	}
	if dataErr != nil {
		http.Error(rw, "Product Not found", http.StatusInternalServerError)
		return
	}
}
