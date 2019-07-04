// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// BrandPOSTHandlerFunc turns a function with the right signature into a brand p o s t handler
type BrandPOSTHandlerFunc func(BrandPOSTParams, *int) BrandPOSTResponder

// Handle executing the request and returning a response
func (fn BrandPOSTHandlerFunc) Handle(params BrandPOSTParams, principal *int) BrandPOSTResponder {
	return fn(params, principal)
}

// BrandPOSTHandler interface for that can handle valid brand p o s t params
type BrandPOSTHandler interface {
	Handle(BrandPOSTParams, *int) BrandPOSTResponder
}

// NewBrandPOST creates a new http.Handler for the brand p o s t operation
func NewBrandPOST(ctx *middleware.Context, handler BrandPOSTHandler) *BrandPOST {
	return &BrandPOST{Context: ctx, Handler: handler}
}

/*BrandPOST swagger:route POST /brand brandPOST

Create new brand

*/
type BrandPOST struct {
	Context *middleware.Context
	Handler BrandPOSTHandler
}

func (o *BrandPOST) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewBrandPOSTParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *int
	if uprinc != nil {
		principal = uprinc.(*int) // this is really a int, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
