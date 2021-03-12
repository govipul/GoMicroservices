package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/govipul/GoMicroservices/episode_5/data"
)

//Products struct
type Products struct {
	l *log.Logger
}

//NewProducts instance
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

//GetProducts get the list of all products
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
func (p Products) MiddlewareProductValication(next http.Handler) http.Handler {
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
