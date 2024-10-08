// Code generated by go-swagger; DO NOT EDIT.

package list_service

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

// NewDeleteListParams creates a new DeleteListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteListParams() *DeleteListParams {
	return &DeleteListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteListParamsWithTimeout creates a new DeleteListParams object
// with the ability to set a timeout on a request.
func NewDeleteListParamsWithTimeout(timeout time.Duration) *DeleteListParams {
	return &DeleteListParams{
		timeout: timeout,
	}
}

// NewDeleteListParamsWithContext creates a new DeleteListParams object
// with the ability to set a context for a request.
func NewDeleteListParamsWithContext(ctx context.Context) *DeleteListParams {
	return &DeleteListParams{
		Context: ctx,
	}
}

// NewDeleteListParamsWithHTTPClient creates a new DeleteListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteListParamsWithHTTPClient(client *http.Client) *DeleteListParams {
	return &DeleteListParams{
		HTTPClient: client,
	}
}

/*
DeleteListParams contains all the parameters to send to the API endpoint

	for the delete list operation.

	Typically these are written to a http.Request.
*/
type DeleteListParams struct {

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

// WithDefaults hydrates default values in the delete list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteListParams) WithDefaults() *DeleteListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete list params
func (o *DeleteListParams) WithTimeout(timeout time.Duration) *DeleteListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete list params
func (o *DeleteListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete list params
func (o *DeleteListParams) WithContext(ctx context.Context) *DeleteListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete list params
func (o *DeleteListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete list params
func (o *DeleteListParams) WithHTTPClient(client *http.Client) *DeleteListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete list params
func (o *DeleteListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the delete list params
func (o *DeleteListParams) WithDomainID(domainID *string) *DeleteListParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the delete list params
func (o *DeleteListParams) SetDomainID(domainID *string) {
	o.DomainID = domainID
}

// WithID adds the id to the delete list params
func (o *DeleteListParams) WithID(id string) *DeleteListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete list params
func (o *DeleteListParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
