// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// NewUploadAvatarParams creates a new UploadAvatarParams object
// no default values defined in spec.
func NewUploadAvatarParams() UploadAvatarParams {

	return UploadAvatarParams{}
}

// UploadAvatarParams contains all the bound params for the upload avatar operation
// typically these are obtained from a http.Request
//
// swagger:parameters uploadAvatar
type UploadAvatarParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: formData
	*/
	UpFile io.ReadCloser
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUploadAvatarParams() beforehand.
func (o *UploadAvatarParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}

	upFile, upFileHeader, err := r.FormFile("upFile")
	if err != nil {
		res = append(res, errors.New(400, "reading file %q failed: %v", "upFile", err))
	} else if err := o.bindUpFile(upFile, upFileHeader); err != nil {
		// Required: true
		res = append(res, err)
	} else {
		o.UpFile = &runtime.File{Data: upFile, Header: upFileHeader}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindUpFile binds file parameter UpFile.
//
// The only supported validations on files are MinLength and MaxLength
func (o *UploadAvatarParams) bindUpFile(file multipart.File, header *multipart.FileHeader) error {
	return nil
}
