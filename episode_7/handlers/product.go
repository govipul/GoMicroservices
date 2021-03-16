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
	"log"

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
