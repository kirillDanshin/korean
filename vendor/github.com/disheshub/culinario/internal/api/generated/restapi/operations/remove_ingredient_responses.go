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

// RemoveIngredientOKCode is the HTTP code returned for type RemoveIngredientOK
const RemoveIngredientOKCode int = 200

/*RemoveIngredientOK The server successfully processed the request and is not returning any content.

swagger:response removeIngredientOK
*/
type RemoveIngredientOK struct {
}

// NewRemoveIngredientOK creates RemoveIngredientOK with default headers values
func NewRemoveIngredientOK() *RemoveIngredientOK {

	return &RemoveIngredientOK{}
}

// WriteResponse to the client
func (o *RemoveIngredientOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

func (o *RemoveIngredientOK) RemoveIngredientResponder() {}

/*RemoveIngredientDefault Generic error response.

swagger:response removeIngredientDefault
*/
type RemoveIngredientDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRemoveIngredientDefault creates RemoveIngredientDefault with default headers values
func NewRemoveIngredientDefault(code int) *RemoveIngredientDefault {
	if code <= 0 {
		code = 500
	}

	return &RemoveIngredientDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the remove ingredient default response
func (o *RemoveIngredientDefault) WithStatusCode(code int) *RemoveIngredientDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the remove ingredient default response
func (o *RemoveIngredientDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the remove ingredient default response
func (o *RemoveIngredientDefault) WithPayload(payload *models.Error) *RemoveIngredientDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove ingredient default response
func (o *RemoveIngredientDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveIngredientDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *RemoveIngredientDefault) RemoveIngredientResponder() {}

type RemoveIngredientNotImplementedResponder struct {
	middleware.Responder
}

func (*RemoveIngredientNotImplementedResponder) RemoveIngredientResponder() {}

func RemoveIngredientNotImplemented() RemoveIngredientResponder {
	return &RemoveIngredientNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.RemoveIngredient has not yet been implemented",
		),
	}
}

type RemoveIngredientResponder interface {
	middleware.Responder
	RemoveIngredientResponder()
}