// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetV1LoanLoanIDPaymentsParams creates a new GetV1LoanLoanIDPaymentsParams object
// no default values defined in spec.
func NewGetV1LoanLoanIDPaymentsParams() GetV1LoanLoanIDPaymentsParams {

	return GetV1LoanLoanIDPaymentsParams{}
}

// GetV1LoanLoanIDPaymentsParams contains all the bound params for the get v1 loan loan ID payments operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetV1LoanLoanIDPayments
type GetV1LoanLoanIDPaymentsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	LoanID int64
	/*
	  In: query
	*/
	Page *int64
	/*
	  In: query
	*/
	PerPage *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetV1LoanLoanIDPaymentsParams() beforehand.
func (o *GetV1LoanLoanIDPaymentsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rLoanID, rhkLoanID, _ := route.Params.GetOK("loan_id")
	if err := o.bindLoanID(rLoanID, rhkLoanID, route.Formats); err != nil {
		res = append(res, err)
	}

	qPage, qhkPage, _ := qs.GetOK("page")
	if err := o.bindPage(qPage, qhkPage, route.Formats); err != nil {
		res = append(res, err)
	}

	qPerPage, qhkPerPage, _ := qs.GetOK("per_page")
	if err := o.bindPerPage(qPerPage, qhkPerPage, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindLoanID binds and validates parameter LoanID from path.
func (o *GetV1LoanLoanIDPaymentsParams) bindLoanID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("loan_id", "path", "int64", raw)
	}
	o.LoanID = value

	return nil
}

// bindPage binds and validates parameter Page from query.
func (o *GetV1LoanLoanIDPaymentsParams) bindPage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("page", "query", "int64", raw)
	}
	o.Page = &value

	return nil
}

// bindPerPage binds and validates parameter PerPage from query.
func (o *GetV1LoanLoanIDPaymentsParams) bindPerPage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("per_page", "query", "int64", raw)
	}
	o.PerPage = &value

	return nil
}