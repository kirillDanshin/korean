// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	models "storage/internal/api/models"
)

// BrandDeleteOKCode is the HTTP code returned for type BrandDeleteOK
const BrandDeleteOKCode int = 200

/*BrandDeleteOK The server successfully processed the request and is not returning any content.

swagger:response brandDeleteOK
*/
type BrandDeleteOK struct {
}

// NewBrandDeleteOK creates BrandDeleteOK with default headers values
func NewBrandDeleteOK() *BrandDeleteOK {

	return &BrandDeleteOK{}
}

// WriteResponse to the client
func (o *BrandDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

func (o *BrandDeleteOK) BrandDeleteResponder() {}

/*BrandDeleteDefault Generic error response.

swagger:response brandDeleteDefault
*/
type BrandDeleteDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewBrandDeleteDefault creates BrandDeleteDefault with default headers values
func NewBrandDeleteDefault(code int) *BrandDeleteDefault {
	if code <= 0 {
		code = 500
	}

	return &BrandDeleteDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the brand delete default response
func (o *BrandDeleteDefault) WithStatusCode(code int) *BrandDeleteDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the brand delete default response
func (o *BrandDeleteDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the brand delete default response
func (o *BrandDeleteDefault) WithPayload(payload *models.Error) *BrandDeleteDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the brand delete default response
func (o *BrandDeleteDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BrandDeleteDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *BrandDeleteDefault) BrandDeleteResponder() {}

type BrandDeleteNotImplementedResponder struct {
	middleware.Responder
}

func (*BrandDeleteNotImplementedResponder) BrandDeleteResponder() {}

func BrandDeleteNotImplemented() BrandDeleteResponder {
	return &BrandDeleteNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.BrandDelete has not yet been implemented",
		),
	}
}

type BrandDeleteResponder interface {
	middleware.Responder
	BrandDeleteResponder()
}
