// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/disheshub/culinario/internal/app"
	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateRecipeHandlerFunc turns a function with the right signature into a create recipe handler
type CreateRecipeHandlerFunc func(CreateRecipeParams, *app.Session) CreateRecipeResponder

// Handle executing the request and returning a response
func (fn CreateRecipeHandlerFunc) Handle(params CreateRecipeParams, principal *app.Session) CreateRecipeResponder {
	return fn(params, principal)
}

// CreateRecipeHandler interface for that can handle valid create recipe params
type CreateRecipeHandler interface {
	Handle(CreateRecipeParams, *app.Session) CreateRecipeResponder
}

// NewCreateRecipe creates a new http.Handler for the create recipe operation
func NewCreateRecipe(ctx *middleware.Context, handler CreateRecipeHandler) *CreateRecipe {
	return &CreateRecipe{Context: ctx, Handler: handler}
}

/*CreateRecipe swagger:route POST /recipe createRecipe

CreateRecipe create recipe API

*/
type CreateRecipe struct {
	Context *middleware.Context
	Handler CreateRecipeHandler
}

func (o *CreateRecipe) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateRecipeParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *app.Session
	if uprinc != nil {
		principal = uprinc.(*app.Session) // this is really a app.Session, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}