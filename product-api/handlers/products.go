package handlers

import (
	"log"
	"net/http"
	"github.com/AppretzelLogic/go-microservices/product-api/data"
)

// Products is a http.Handler
type ProductServiceHandler struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProductServiceHandler(l *log.Logger) *ProductServiceHandler {
	return &ProductServiceHandler{l}
}

// ServeHTTP is the main entry point for the handler and satisfies the http.Handler interface
func (ph *ProductServiceHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// handle GET
	if r.Method == http.MethodGet {
		ph.getProductHandler(rw, r)
		return
	}
	// handle POST
	if r.Method == http.MethodPost {
		ph.addProductHandler(rw, r)
		return
	}
	// handle PUT
	if r.Method == http.MethodPut {
		ph.updateProductHandler(rw, r)
		return
	}
	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// getProducts returns the products from the data store
func (ph *ProductServiceHandler) getProductHandler(rw http.ResponseWriter, r *http.Request) {
	// fetch the products from the datastore
	lp := data.GetProducts()
	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (ph *ProductServiceHandler) addProductHandler(rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

}

func (ph *ProductServiceHandler) updateProductHandler(rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	var id int
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

}
