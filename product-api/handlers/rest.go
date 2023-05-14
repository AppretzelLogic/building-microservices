package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// ServeHttp is the main entry point for the handler and satisfies the http.Handler interface
// interface
func (ph *ProductServiceHandler) ServeAPI(endpoint string) error {
	// handle GET
	handlers := &ProductServiceHandler{
		l: &log.Logger{},
	}
	r := mux.NewRouter()
	r.Handle("/", ph)
	ProductServiceHandler := r.PathPrefix("/products").Subrouter()
	ProductServiceHandler.Methods("POST").Path("").HandlerFunc(handlers.addProductHandler)
	ProductServiceHandler.Methods("PUT").Path("").HandlerFunc(handlers.updateProductHandler)
	ProductServiceHandler.Methods("GET").Path("/{id:[0-9]+}").HandlerFunc(handlers.getProductHandler)
	return http.ListenAndServe(endpoint, r)
}