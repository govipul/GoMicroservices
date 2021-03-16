package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/govipul/GoMicroservices/episode_7/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
//
// responses:
//	200: noContent

// DeleteProducts delete product from the data store
func (p Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, errAtoI := strconv.Atoi(vars["id"])
	if errAtoI != nil {
		http.Error(rw, "Unable to unable to parse id", http.StatusBadRequest)
		return
	}
	p.l.Println("handle a DELETE product")

	dataErr := data.DeleteProduct(id)
	if dataErr == data.ErrProductNotFound {
		http.Error(rw, "Error while updating the product", http.StatusNotFound)
		return
	}
	if dataErr != nil {
		http.Error(rw, "Product Not found", http.StatusInternalServerError)
		return
	}
}
