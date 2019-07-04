// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// BrandListHandlerFunc turns a function with the right signature into a brand list handler
type BrandListHandlerFunc func(BrandListParams) BrandListResponder

// Handle executing the request and returning a response
func (fn BrandListHandlerFunc) Handle(params BrandListParams) BrandListResponder {
	return fn(params)
}

// BrandListHandler interface for that can handle valid brand list params
type BrandListHandler interface {
	Handle(BrandListParams) BrandListResponder
}

// NewBrandList creates a new http.Handler for the brand list operation
func NewBrandList(ctx *middleware.Context, handler BrandListHandler) *BrandList {
	return &BrandList{Context: ctx, Handler: handler}
}

/*BrandList swagger:route GET /brand brandList

Get list brand

*/
type BrandList struct {
	Context *middleware.Context
	Handler BrandListHandler
}

func (o *BrandList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewBrandListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
