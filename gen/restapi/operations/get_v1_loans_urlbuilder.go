// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// GetV1LoansURL generates an URL for the get v1 loans operation
type GetV1LoansURL struct {
	Asc     *bool
	Order   *string
	Page    *int64
	PerPage *int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetV1LoansURL) WithBasePath(bp string) *GetV1LoansURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetV1LoansURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetV1LoansURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/v1/loans"

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var asc string
	if o.Asc != nil {
		asc = swag.FormatBool(*o.Asc)
	}
	if asc != "" {
		qs.Set("asc", asc)
	}

	var order string
	if o.Order != nil {
		order = *o.Order
	}
	if order != "" {
		qs.Set("order", order)
	}

	var page string
	if o.Page != nil {
		page = swag.FormatInt64(*o.Page)
	}
	if page != "" {
		qs.Set("page", page)
	}

	var perPage string
	if o.PerPage != nil {
		perPage = swag.FormatInt64(*o.PerPage)
	}
	if perPage != "" {
		qs.Set("per_page", perPage)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetV1LoansURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetV1LoansURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetV1LoansURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetV1LoansURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetV1LoansURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetV1LoansURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
