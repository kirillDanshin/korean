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

// UpdateUserInfoOKCode is the HTTP code returned for type UpdateUserInfoOK
const UpdateUserInfoOKCode int = 200

/*UpdateUserInfoOK OK

swagger:response updateUserInfoOK
*/
type UpdateUserInfoOK struct {

	/*
	  In: Body
	*/
	Payload *models.UserInfo `json:"body,omitempty"`
}

// NewUpdateUserInfoOK creates UpdateUserInfoOK with default headers values
func NewUpdateUserInfoOK() *UpdateUserInfoOK {

	return &UpdateUserInfoOK{}
}

// WithPayload adds the payload to the update user info o k response
func (o *UpdateUserInfoOK) WithPayload(payload *models.UserInfo) *UpdateUserInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user info o k response
func (o *UpdateUserInfoOK) SetPayload(payload *models.UserInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *UpdateUserInfoOK) UpdateUserInfoResponder() {}

/*UpdateUserInfoDefault Generic error response.

swagger:response updateUserInfoDefault
*/
type UpdateUserInfoDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateUserInfoDefault creates UpdateUserInfoDefault with default headers values
func NewUpdateUserInfoDefault(code int) *UpdateUserInfoDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateUserInfoDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update user info default response
func (o *UpdateUserInfoDefault) WithStatusCode(code int) *UpdateUserInfoDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update user info default response
func (o *UpdateUserInfoDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update user info default response
func (o *UpdateUserInfoDefault) WithPayload(payload *models.Error) *UpdateUserInfoDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user info default response
func (o *UpdateUserInfoDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserInfoDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *UpdateUserInfoDefault) UpdateUserInfoResponder() {}

type UpdateUserInfoNotImplementedResponder struct {
	middleware.Responder
}

func (*UpdateUserInfoNotImplementedResponder) UpdateUserInfoResponder() {}

func UpdateUserInfoNotImplemented() UpdateUserInfoResponder {
	return &UpdateUserInfoNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.UpdateUserInfo has not yet been implemented",
		),
	}
}

type UpdateUserInfoResponder interface {
	middleware.Responder
	UpdateUserInfoResponder()
}