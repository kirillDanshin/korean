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

// RmMeasureWeightOKCode is the HTTP code returned for type RmMeasureWeightOK
const RmMeasureWeightOKCode int = 200

/*RmMeasureWeightOK The server successfully processed the request and is not returning any content.

swagger:response rmMeasureWeightOK
*/
type RmMeasureWeightOK struct {
}

// NewRmMeasureWeightOK creates RmMeasureWeightOK with default headers values
func NewRmMeasureWeightOK() *RmMeasureWeightOK {

	return &RmMeasureWeightOK{}
}

// WriteResponse to the client
func (o *RmMeasureWeightOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

func (o *RmMeasureWeightOK) RmMeasureWeightResponder() {}

/*RmMeasureWeightDefault Generic error response.

swagger:response rmMeasureWeightDefault
*/
type RmMeasureWeightDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRmMeasureWeightDefault creates RmMeasureWeightDefault with default headers values
func NewRmMeasureWeightDefault(code int) *RmMeasureWeightDefault {
	if code <= 0 {
		code = 500
	}

	return &RmMeasureWeightDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the rm measure weight default response
func (o *RmMeasureWeightDefault) WithStatusCode(code int) *RmMeasureWeightDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the rm measure weight default response
func (o *RmMeasureWeightDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the rm measure weight default response
func (o *RmMeasureWeightDefault) WithPayload(payload *models.Error) *RmMeasureWeightDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the rm measure weight default response
func (o *RmMeasureWeightDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RmMeasureWeightDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *RmMeasureWeightDefault) RmMeasureWeightResponder() {}

type RmMeasureWeightNotImplementedResponder struct {
	middleware.Responder
}

func (*RmMeasureWeightNotImplementedResponder) RmMeasureWeightResponder() {}

func RmMeasureWeightNotImplemented() RmMeasureWeightResponder {
	return &RmMeasureWeightNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.RmMeasureWeight has not yet been implemented",
		),
	}
}

type RmMeasureWeightResponder interface {
	middleware.Responder
	RmMeasureWeightResponder()
}
