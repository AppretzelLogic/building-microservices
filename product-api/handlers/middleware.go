package handlers

import (
	"context"
	"net/http"

	"github.com/AppretzelLogic/go-microservices/product-api/data"
)

// MiddlewareValidateProduct validates the product in the request and calls next if ok

type KeyProduct struct{}

func (p *ProductServiceHandler) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
			return
		}
		// validate the product
		//errs := p.v.Validate(prod)
		//if len(errs) != 0 {
		// return the validation messages as an array
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
		//}//

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)
		// call the next handler, which can be another middleware in the chain, or the final handler
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
