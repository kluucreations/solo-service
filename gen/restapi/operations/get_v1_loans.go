// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetV1LoansHandlerFunc turns a function with the right signature into a get v1 loans handler
type GetV1LoansHandlerFunc func(GetV1LoansParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetV1LoansHandlerFunc) Handle(params GetV1LoansParams) middleware.Responder {
	return fn(params)
}

// GetV1LoansHandler interface for that can handle valid get v1 loans params
type GetV1LoansHandler interface {
	Handle(GetV1LoansParams) middleware.Responder
}

// NewGetV1Loans creates a new http.Handler for the get v1 loans operation
func NewGetV1Loans(ctx *middleware.Context, handler GetV1LoansHandler) *GetV1Loans {
	return &GetV1Loans{Context: ctx, Handler: handler}
}

/*GetV1Loans swagger:route GET /v1/loans getV1Loans

Retrieve Loans

*/
type GetV1Loans struct {
	Context *middleware.Context
	Handler GetV1LoansHandler
}

func (o *GetV1Loans) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetV1LoansParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
