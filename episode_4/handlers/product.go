package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
)

//Products struct
type Products struct {
	l *log.Logger
}

//NewProducts instance
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}
	//handle to add new product
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}
	//Handle to update the product
	if r.Method == http.MethodPut {
		//expect the id in URL
		reg := regexp.MustCompile("/([0-9]+)")
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 || len(g[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "Invalid ID provided", http.StatusBadRequest)
			return
		}
		p.updateProduct(id, rw, r)
		return
	}
	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	//d, err := json.Marshal(lp)
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal data", http.StatusInternalServerError)
		return
	}
	//rw.Write(d)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle a PUT product")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusBadRequest)
	}
	p.l.Printf("Prod: %#v", prod)
	dataErr := data.UpdateProduct(id, prod)
	if dataErr == data.ErrProductNotFound {
		http.Error(rw, "Error while updating the product", http.StatusNotFound)
		return
	}
	if dataErr != nil {
		http.Error(rw, "Product Not found", http.StatusInternalServerError)
		return
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle a POST product")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusBadRequest)
	}
	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}
