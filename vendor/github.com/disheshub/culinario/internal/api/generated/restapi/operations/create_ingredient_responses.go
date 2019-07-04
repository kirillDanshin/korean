// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/disheshub/culinario/internal/api/generated/models"
)

// CreateIngredientOKCode is the HTTP code returned for type CreateIngredientOK
const CreateIngredientOKCode int = 200

/*CreateIngredientOK OK

swagger:response createIngredientOK
*/
type CreateIngredientOK struct {

	/*
	  In: Body
	*/
	Payload *models.ListIngredient `json:"body,omitempty"`
}

// NewCreateIngredientOK creates CreateIngredientOK with default headers values
func NewCreateIngredientOK() *CreateIngredientOK {

	return &CreateIngredientOK{}
}

// WithPayload adds the payload to the create ingredient o k response
func (o *CreateIngredientOK) WithPayload(payload *models.ListIngredient) *CreateIngredientOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create ingredient o k response
func (o *CreateIngredientOK) SetPayload(payload *models.ListIngredient) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateIngredientOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *CreateIngredientOK) CreateIngredientResponder() {}

/*CreateIngredientDefault Generic error response.

swagger:response createIngredientDefault
*/
type CreateIngredientDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateIngredientDefault creates CreateIngredientDefault with default headers values
func NewCreateIngredientDefault(code int) *CreateIngredientDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateIngredientDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create ingredient default response
func (o *CreateIngredientDefault) WithStatusCode(code int) *CreateIngredientDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create ingredient default response
func (o *CreateIngredientDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create ingredient default response
func (o *CreateIngredientDefault) WithPayload(payload *models.Error) *CreateIngredientDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create ingredient default response
func (o *CreateIngredientDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateIngredientDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *CreateIngredientDefault) CreateIngredientResponder() {}

type CreateIngredientNotImplementedResponder struct {
	middleware.Responder
}

func (*CreateIngredientNotImplementedResponder) CreateIngredientResponder() {}

func CreateIngredientNotImplemented() CreateIngredientResponder {
	return &CreateIngredientNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.CreateIngredient has not yet been implemented",
		),
	}
}

type CreateIngredientResponder interface {
	middleware.Responder
	CreateIngredientResponder()
}
