// Package classification of Product API
//
// Documentation for product API
//
//     	Schemes: http
//     	Host: localhost
//     	BasePath: /
//		Version: 1.0.0
//
//     	Consumes:
//     	- application/json
//
//     	Produces:
//     	- application/json
//
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/govipul/GoMicroservices/episode_7/data"
)

// A list of products returns in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger:response noContent
type productsNoContent struct {
}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The id of the prduct to delete from db
	// in: path
	// required: true
	ID int `json:"id"`
}

//Products struct
type Products struct {
	l *log.Logger
}

//NewProducts instance
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// swagger:route GET /products products listProducts
// Returns a list of porducts
// responses:
//	200: productsResponse

// GetProducts get the list of all products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	//d, err := json.Marshal(lp)
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal data", http.StatusInternalServerError)
		return
	}
	//rw.Write(d)
}

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

type KeyProduct struct {
}

//MiddlewareProductValication middle ware layer
func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Unable to marshal json", http.StatusBadRequest)
			return
		}

		//validate the product
		err = prod.Validate()

		if err != nil {
			http.Error(rw, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
