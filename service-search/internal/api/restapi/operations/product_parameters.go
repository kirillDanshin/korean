// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewProductParams creates a new ProductParams object
// no default values defined in spec.
func NewProductParams() ProductParams {

	return ProductParams{}
}

// ProductParams contains all the bound params for the product operation
// typically these are obtained from a http.Request
//
// swagger:parameters product
type ProductParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Track ID.
	  Required: true
	  In: path
	*/
	ProductID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewProductParams() beforehand.
func (o *ProductParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rProductID, rhkProductID, _ := route.Params.GetOK("productID")
	if err := o.bindProductID(rProductID, rhkProductID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindProductID binds and validates parameter ProductID from path.
func (o *ProductParams) bindProductID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("productID", "path", "int64", raw)
	}
	o.ProductID = value

	return nil
}