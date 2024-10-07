// Code generated by go-swagger; DO NOT EDIT.

package routing_variable_service

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
)

// NewReadRoutingVariableParams creates a new ReadRoutingVariableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadRoutingVariableParams() *ReadRoutingVariableParams {
	return &ReadRoutingVariableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadRoutingVariableParamsWithTimeout creates a new ReadRoutingVariableParams object
// with the ability to set a timeout on a request.
func NewReadRoutingVariableParamsWithTimeout(timeout time.Duration) *ReadRoutingVariableParams {
	return &ReadRoutingVariableParams{
		timeout: timeout,
	}
}

// NewReadRoutingVariableParamsWithContext creates a new ReadRoutingVariableParams object
// with the ability to set a context for a request.
func NewReadRoutingVariableParamsWithContext(ctx context.Context) *ReadRoutingVariableParams {
	return &ReadRoutingVariableParams{
		Context: ctx,
	}
}

// NewReadRoutingVariableParamsWithHTTPClient creates a new ReadRoutingVariableParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadRoutingVariableParamsWithHTTPClient(client *http.Client) *ReadRoutingVariableParams {
	return &ReadRoutingVariableParams{
		HTTPClient: client,
	}
}

/*
ReadRoutingVariableParams contains all the parameters to send to the API endpoint

	for the read routing variable operation.

	Typically these are written to a http.Request.
*/
type ReadRoutingVariableParams struct {

	// DomainID.
	//
	// Format: int64
	DomainID *string

	// ID.
	//
	// Format: int64
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read routing variable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadRoutingVariableParams) WithDefaults() *ReadRoutingVariableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read routing variable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadRoutingVariableParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read routing variable params
func (o *ReadRoutingVariableParams) WithTimeout(timeout time.Duration) *ReadRoutingVariableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read routing variable params
func (o *ReadRoutingVariableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read routing variable params
func (o *ReadRoutingVariableParams) WithContext(ctx context.Context) *ReadRoutingVariableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read routing variable params
func (o *ReadRoutingVariableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read routing variable params
func (o *ReadRoutingVariableParams) WithHTTPClient(client *http.Client) *ReadRoutingVariableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read routing variable params
func (o *ReadRoutingVariableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the read routing variable params
func (o *ReadRoutingVariableParams) WithDomainID(domainID *string) *ReadRoutingVariableParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the read routing variable params
func (o *ReadRoutingVariableParams) SetDomainID(domainID *string) {
	o.DomainID = domainID
}

// WithID adds the id to the read routing variable params
func (o *ReadRoutingVariableParams) WithID(id string) *ReadRoutingVariableParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the read routing variable params
func (o *ReadRoutingVariableParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReadRoutingVariableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DomainID != nil {

		// query param domain_id
		var qrDomainID string

		if o.DomainID != nil {
			qrDomainID = *o.DomainID
		}
		qDomainID := qrDomainID
		if qDomainID != "" {

			if err := r.SetQueryParam("domain_id", qDomainID); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
