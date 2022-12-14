package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"goinpractice.com/microservice/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		p.getProducts(rw, r)
// 		return
// 	}

// 	if r.Method == http.MethodPost {
// 		p.addNewProduct(rw, r)
// 		return
// 	}

// 	if r.Method == http.MethodPut {
// 		p.l.Println("Path is ", r.URL.Path)
// 		regx := regexp.MustCompile(`/[0-9]+`)
// 		id := regx.FindAllString(r.URL.Path, -1)
// 		if len(id) > 1 {
// 			http.Error(rw, "Only 1 id must be specified", http.StatusBadRequest)
// 			return
// 		}
// 		intid, err := strconv.Atoi(id[0][1:])
// 		if err != nil {
// 			http.Error(rw, "Id should be a number", http.StatusBadRequest)
// 			return
// 		}
// 		p.updateProduct(intid, rw, r)
// 	}

// 	//catch all
// 	rw.WriteHeader(http.StatusMethodNotAllowed)
// }

func (p *Products) GetProducts(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
	}
}

func (p *Products) AddNewProduct(rw http.ResponseWriter, h *http.Request) {
	p.l.Printf("Adding a new product")
	prod := &data.Product{}
	err := prod.FromJSON(h.Body)
	if err != nil {
		http.Error(rw, "product is not of right configuration", http.StatusBadRequest)
		return
	}
	data.AddProduct(prod)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Updating Products")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "Bad ID", http.StatusBadRequest)
		return
	}

	newProd := &data.Product{}
	err = newProd.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "product is not of right configuration", http.StatusBadRequest)
		return
	}

	oldProd, index := data.GetProductById(id)
	if oldProd.ID != newProd.ID || oldProd.ID != id || newProd.ID != id {
		panic(fmt.Sprintf("What the fuck is going on! oldProd.ID is {%d} and newProd.ID is {%d} and id is {%d}", oldProd.ID, newProd.ID, id))
	}

	data.UpdateProductAtIndex(index, newProd)
}
