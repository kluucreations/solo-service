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

// Amount Amount object
// swagger:model amount
type Amount struct {

	// amount
	// Required: true
	Amount *float64 `json:"amount"`

	// iso currency code
	// Required: true
	IsoCurrencyCode *int64 `json:"iso_currency_code"`
}

// Validate validates this amount
func (m *Amount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsoCurrencyCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Amount) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *Amount) validateIsoCurrencyCode(formats strfmt.Registry) error {

	if err := validate.Required("iso_currency_code", "body", m.IsoCurrencyCode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Amount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Amount) UnmarshalBinary(b []byte) error {
	var res Amount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
