package handlers

import (
	"net/http"
	"github.com/AppretzelLogic/go-microservices/product-api/data"
	"github.com/gorilla/mux"
)

// Products is a http.Handler
type ProductServiceHandler struct {}

// NewProducts creates a products handler with the given logger
func (ph *ProductServiceHandler) NewProductServiceHandler(w http.ResponseWriter, r *http.Request) {
	// fetch the products from the datastore
	lp := data.GetProducts()
	// serialize the list to JSON
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
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
// ServeHttp is the main entry point for the handler and satisfies the http.Handler interface
// interface
func (p *ProductServiceHandler) ServeAPI(endpoint string) error {
	// handle GET
	handler := &ProductServiceHandler{}
	r := mux.NewRouter()
	ProductServiceHandler := r.PathPrefix("/products").Subrouter()
	ProductServiceHandler.Methods("GET").Path("").HandlerFunc(handler.NewProductServiceHandler)
	ProductServiceHandler.Methods("POST").Path("").HandlerFunc(handler.addProductHandler)
	ProductServiceHandler.Methods("PUT").Path("").HandlerFunc(handler.updateProductHandler)
	ProductServiceHandler.Methods("GET").Path("/{id:[0-9]+}").HandlerFunc(handler.getProductHandler)
	return http.ListenAndServe(endpoint, r)
}

