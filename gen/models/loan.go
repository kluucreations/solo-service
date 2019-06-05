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

// Loan loan
// swagger:model loan
type Loan struct {

	// amount
	// Required: true
	Amount *Amount `json:"amount"`

	// borrower
	// Required: true
	Borrower *User `json:"borrower"`

	// created at
	// Required: true
	CreatedAt *string `json:"created_at"`

	// description
	Description string `json:"description,omitempty"`

	// end at
	EndAt string `json:"end_at,omitempty"`

	// guid
	// Required: true
	GUID *string `json:"guid"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// start at
	StartAt string `json:"start_at,omitempty"`

	// updated at
	// Required: true
	UpdatedAt *string `json:"updated_at"`
}

// Validate validates this loan
func (m *Loan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBorrower(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Loan) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	if m.Amount != nil {
		if err := m.Amount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("amount")
			}
			return err
		}
	}

	return nil
}

func (m *Loan) validateBorrower(formats strfmt.Registry) error {

	if err := validate.Required("borrower", "body", m.Borrower); err != nil {
		return err
	}

	if m.Borrower != nil {
		if err := m.Borrower.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("borrower")
			}
			return err
		}
	}

	return nil
}

func (m *Loan) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *Loan) validateGUID(formats strfmt.Registry) error {

	if err := validate.Required("guid", "body", m.GUID); err != nil {
		return err
	}

	return nil
}

func (m *Loan) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Loan) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Loan) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Loan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Loan) UnmarshalBinary(b []byte) error {
	var res Loan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
