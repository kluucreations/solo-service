// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/kluucreations/solo-service/gen/models"
)

// GetV1LoanLoanIDPaymentsOKCode is the HTTP code returned for type GetV1LoanLoanIDPaymentsOK
const GetV1LoanLoanIDPaymentsOKCode int = 200

/*GetV1LoanLoanIDPaymentsOK get v1 loan loan Id payments o k

swagger:response getV1LoanLoanIdPaymentsOK
*/
type GetV1LoanLoanIDPaymentsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Payment `json:"body,omitempty"`
}

// NewGetV1LoanLoanIDPaymentsOK creates GetV1LoanLoanIDPaymentsOK with default headers values
func NewGetV1LoanLoanIDPaymentsOK() *GetV1LoanLoanIDPaymentsOK {

	return &GetV1LoanLoanIDPaymentsOK{}
}

// WithPayload adds the payload to the get v1 loan loan Id payments o k response
func (o *GetV1LoanLoanIDPaymentsOK) WithPayload(payload []*models.Payment) *GetV1LoanLoanIDPaymentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 loan loan Id payments o k response
func (o *GetV1LoanLoanIDPaymentsOK) SetPayload(payload []*models.Payment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1LoanLoanIDPaymentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Payment, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetV1LoanLoanIDPaymentsForbiddenCode is the HTTP code returned for type GetV1LoanLoanIDPaymentsForbidden
const GetV1LoanLoanIDPaymentsForbiddenCode int = 403

/*GetV1LoanLoanIDPaymentsForbidden get v1 loan loan Id payments forbidden

swagger:response getV1LoanLoanIdPaymentsForbidden
*/
type GetV1LoanLoanIDPaymentsForbidden struct {
}

// NewGetV1LoanLoanIDPaymentsForbidden creates GetV1LoanLoanIDPaymentsForbidden with default headers values
func NewGetV1LoanLoanIDPaymentsForbidden() *GetV1LoanLoanIDPaymentsForbidden {

	return &GetV1LoanLoanIDPaymentsForbidden{}
}

// WriteResponse to the client
func (o *GetV1LoanLoanIDPaymentsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// GetV1LoanLoanIDPaymentsNotFoundCode is the HTTP code returned for type GetV1LoanLoanIDPaymentsNotFound
const GetV1LoanLoanIDPaymentsNotFoundCode int = 404

/*GetV1LoanLoanIDPaymentsNotFound get v1 loan loan Id payments not found

swagger:response getV1LoanLoanIdPaymentsNotFound
*/
type GetV1LoanLoanIDPaymentsNotFound struct {
}

// NewGetV1LoanLoanIDPaymentsNotFound creates GetV1LoanLoanIDPaymentsNotFound with default headers values
func NewGetV1LoanLoanIDPaymentsNotFound() *GetV1LoanLoanIDPaymentsNotFound {

	return &GetV1LoanLoanIDPaymentsNotFound{}
}

// WriteResponse to the client
func (o *GetV1LoanLoanIDPaymentsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}