// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	models "search/internal/api/models"
)

// ListProductOKCode is the HTTP code returned for type ListProductOK
const ListProductOKCode int = 200

/*ListProductOK Receive a list of products.

swagger:response listProductOK
*/
type ListProductOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Product `json:"body,omitempty"`
}

// NewListProductOK creates ListProductOK with default headers values
func NewListProductOK() *ListProductOK {

	return &ListProductOK{}
}

// WithPayload adds the payload to the list product o k response
func (o *ListProductOK) WithPayload(payload []*models.Product) *ListProductOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list product o k response
func (o *ListProductOK) SetPayload(payload []*models.Product) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListProductOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Product, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

func (o *ListProductOK) ListProductResponder() {}

/*ListProductDefault Generic error response.

swagger:response listProductDefault
*/
type ListProductDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListProductDefault creates ListProductDefault with default headers values
func NewListProductDefault(code int) *ListProductDefault {
	if code <= 0 {
		code = 500
	}

	return &ListProductDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list product default response
func (o *ListProductDefault) WithStatusCode(code int) *ListProductDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list product default response
func (o *ListProductDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list product default response
func (o *ListProductDefault) WithPayload(payload *models.Error) *ListProductDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list product default response
func (o *ListProductDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListProductDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *ListProductDefault) ListProductResponder() {}

type ListProductNotImplementedResponder struct {
	middleware.Responder
}

func (*ListProductNotImplementedResponder) ListProductResponder() {}

func ListProductNotImplemented() ListProductResponder {
	return &ListProductNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.ListProduct has not yet been implemented",
		),
	}
}

type ListProductResponder interface {
	middleware.Responder
	ListProductResponder()
}
