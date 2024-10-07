// Code generated by go-swagger; DO NOT EDIT.

package outbound_resource_service

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

// NewDeleteOutboundResourceParams creates a new DeleteOutboundResourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOutboundResourceParams() *DeleteOutboundResourceParams {
	return &DeleteOutboundResourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOutboundResourceParamsWithTimeout creates a new DeleteOutboundResourceParams object
// with the ability to set a timeout on a request.
func NewDeleteOutboundResourceParamsWithTimeout(timeout time.Duration) *DeleteOutboundResourceParams {
	return &DeleteOutboundResourceParams{
		timeout: timeout,
	}
}

// NewDeleteOutboundResourceParamsWithContext creates a new DeleteOutboundResourceParams object
// with the ability to set a context for a request.
func NewDeleteOutboundResourceParamsWithContext(ctx context.Context) *DeleteOutboundResourceParams {
	return &DeleteOutboundResourceParams{
		Context: ctx,
	}
}

// NewDeleteOutboundResourceParamsWithHTTPClient creates a new DeleteOutboundResourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOutboundResourceParamsWithHTTPClient(client *http.Client) *DeleteOutboundResourceParams {
	return &DeleteOutboundResourceParams{
		HTTPClient: client,
	}
}

/*
DeleteOutboundResourceParams contains all the parameters to send to the API endpoint

	for the delete outbound resource operation.

	Typically these are written to a http.Request.
*/
type DeleteOutboundResourceParams struct {

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

// WithDefaults hydrates default values in the delete outbound resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOutboundResourceParams) WithDefaults() *DeleteOutboundResourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete outbound resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOutboundResourceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete outbound resource params
func (o *DeleteOutboundResourceParams) WithTimeout(timeout time.Duration) *DeleteOutboundResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete outbound resource params
func (o *DeleteOutboundResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete outbound resource params
func (o *DeleteOutboundResourceParams) WithContext(ctx context.Context) *DeleteOutboundResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete outbound resource params
func (o *DeleteOutboundResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete outbound resource params
func (o *DeleteOutboundResourceParams) WithHTTPClient(client *http.Client) *DeleteOutboundResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete outbound resource params
func (o *DeleteOutboundResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the delete outbound resource params
func (o *DeleteOutboundResourceParams) WithDomainID(domainID *string) *DeleteOutboundResourceParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the delete outbound resource params
func (o *DeleteOutboundResourceParams) SetDomainID(domainID *string) {
	o.DomainID = domainID
}

// WithID adds the id to the delete outbound resource params
func (o *DeleteOutboundResourceParams) WithID(id string) *DeleteOutboundResourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete outbound resource params
func (o *DeleteOutboundResourceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOutboundResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
