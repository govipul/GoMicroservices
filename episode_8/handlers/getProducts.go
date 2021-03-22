package handlers

import (
	"net/http"

	"github.com/govipul/GoMicroservices/episode_8/data"
)

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
