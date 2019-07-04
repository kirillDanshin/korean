// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProductCreate product create
// swagger:model ProductCreate
type ProductCreate struct {

	// apply
	// Required: true
	Apply *string `json:"apply"`

	// description
	// Required: true
	Description *string `json:"description"`

	// name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// price
	// Required: true
	Price *int64 `json:"price"`
}

// Validate validates this product create
func (m *ProductCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApply(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductCreate) validateApply(formats strfmt.Registry) error {

	if err := validate.Required("apply", "body", m.Apply); err != nil {
		return err
	}

	return nil
}

func (m *ProductCreate) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ProductCreate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 100); err != nil {
		return err
	}

	return nil
}

func (m *ProductCreate) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", m.Price); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProductCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductCreate) UnmarshalBinary(b []byte) error {
	var res ProductCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
