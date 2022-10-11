package handlers

import (
	"log"
	"net/http"

	"goinpractice.com/microservice/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addNewProduct(rw, r)
		return
	}

	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
	}
}

func (p *Products) addNewProduct(rw http.ResponseWriter, h *http.Request) {
	p.l.Printf("Adding a new product")
	prod := &data.Product{}
	err := prod.FromJSON(h.Body)
	if err != nil {
		http.Error(rw, "product is not of right configuration", http.StatusBadRequest)
		return
	}
	data.AddProduct(prod)
}
