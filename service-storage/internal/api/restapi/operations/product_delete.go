// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ProductDeleteHandlerFunc turns a function with the right signature into a product delete handler
type ProductDeleteHandlerFunc func(ProductDeleteParams, *string) ProductDeleteResponder

// Handle executing the request and returning a response
func (fn ProductDeleteHandlerFunc) Handle(params ProductDeleteParams, principal *string) ProductDeleteResponder {
	return fn(params, principal)
}

// ProductDeleteHandler interface for that can handle valid product delete params
type ProductDeleteHandler interface {
	Handle(ProductDeleteParams, *string) ProductDeleteResponder
}

// NewProductDelete creates a new http.Handler for the product delete operation
func NewProductDelete(ctx *middleware.Context, handler ProductDeleteHandler) *ProductDelete {
	return &ProductDelete{Context: ctx, Handler: handler}
}

/*ProductDelete swagger:route DELETE /product productDelete

Delete product

*/
type ProductDelete struct {
	Context *middleware.Context
	Handler ProductDeleteHandler
}

func (o *ProductDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewProductDeleteParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *string
	if uprinc != nil {
		principal = uprinc.(*string) // this is really a string, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
