// Code generated by go-swagger; DO NOT EDIT.

package domains

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/webitel-openapi-client-go/models"
)

// NewDomainsCreateDomainParams creates a new DomainsCreateDomainParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDomainsCreateDomainParams() *DomainsCreateDomainParams {
	return &DomainsCreateDomainParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDomainsCreateDomainParamsWithTimeout creates a new DomainsCreateDomainParams object
// with the ability to set a timeout on a request.
func NewDomainsCreateDomainParamsWithTimeout(timeout time.Duration) *DomainsCreateDomainParams {
	return &DomainsCreateDomainParams{
		timeout: timeout,
	}
}

// NewDomainsCreateDomainParamsWithContext creates a new DomainsCreateDomainParams object
// with the ability to set a context for a request.
func NewDomainsCreateDomainParamsWithContext(ctx context.Context) *DomainsCreateDomainParams {
	return &DomainsCreateDomainParams{
		Context: ctx,
	}
}

// NewDomainsCreateDomainParamsWithHTTPClient creates a new DomainsCreateDomainParams object
// with the ability to set a custom HTTPClient for a request.
func NewDomainsCreateDomainParamsWithHTTPClient(client *http.Client) *DomainsCreateDomainParams {
	return &DomainsCreateDomainParams{
		HTTPClient: client,
	}
}

/*
DomainsCreateDomainParams contains all the parameters to send to the API endpoint

	for the domains create domain operation.

	Typically these are written to a http.Request.
*/
type DomainsCreateDomainParams struct {

	// Body.
	Body *models.APICreateDomainRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the domains create domain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DomainsCreateDomainParams) WithDefaults() *DomainsCreateDomainParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the domains create domain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DomainsCreateDomainParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the domains create domain params
func (o *DomainsCreateDomainParams) WithTimeout(timeout time.Duration) *DomainsCreateDomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the domains create domain params
func (o *DomainsCreateDomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the domains create domain params
func (o *DomainsCreateDomainParams) WithContext(ctx context.Context) *DomainsCreateDomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the domains create domain params
func (o *DomainsCreateDomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the domains create domain params
func (o *DomainsCreateDomainParams) WithHTTPClient(client *http.Client) *DomainsCreateDomainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the domains create domain params
func (o *DomainsCreateDomainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the domains create domain params
func (o *DomainsCreateDomainParams) WithBody(body *models.APICreateDomainRequest) *DomainsCreateDomainParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the domains create domain params
func (o *DomainsCreateDomainParams) SetBody(body *models.APICreateDomainRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DomainsCreateDomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
