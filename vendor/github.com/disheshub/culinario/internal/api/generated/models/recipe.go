// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Recipe recipe
// swagger:model Recipe
type Recipe struct {

	// Cooking time.
	// Required: true
	CookingTime *int32 `json:"cookingTime"`

	// Description recipe.
	// Required: true
	// Min Length: 1
	Description *string `json:"description"`

	// id
	// Required: true
	ID ID `json:"id"`

	// ingridients
	// Required: true
	Ingridients []*Ingredient `json:"ingridients"`

	// Recipe name.
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// steps
	// Required: true
	Steps []*Step `json:"steps"`
}

// Validate validates this recipe
func (m *Recipe) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCookingTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngridients(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSteps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Recipe) validateCookingTime(formats strfmt.Registry) error {

	if err := validate.Required("cookingTime", "body", m.CookingTime); err != nil {
		return err
	}

	return nil
}

func (m *Recipe) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 1); err != nil {
		return err
	}

	return nil
}

func (m *Recipe) validateID(formats strfmt.Registry) error {

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *Recipe) validateIngridients(formats strfmt.Registry) error {

	if err := validate.Required("ingridients", "body", m.Ingridients); err != nil {
		return err
	}

	for i := 0; i < len(m.Ingridients); i++ {
		if swag.IsZero(m.Ingridients[i]) { // not required
			continue
		}

		if m.Ingridients[i] != nil {
			if err := m.Ingridients[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ingridients" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Recipe) validateName(formats strfmt.Registry) error {

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

func (m *Recipe) validateSteps(formats strfmt.Registry) error {

	if err := validate.Required("steps", "body", m.Steps); err != nil {
		return err
	}

	for i := 0; i < len(m.Steps); i++ {
		if swag.IsZero(m.Steps[i]) { // not required
			continue
		}

		if m.Steps[i] != nil {
			if err := m.Steps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Recipe) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Recipe) UnmarshalBinary(b []byte) error {
	var res Recipe
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}