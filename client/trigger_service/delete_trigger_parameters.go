// Code generated by go-swagger; DO NOT EDIT.

package trigger_service

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
	"github.com/go-openapi/swag"
)

// NewDeleteTriggerParams creates a new DeleteTriggerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteTriggerParams() *DeleteTriggerParams {
	return &DeleteTriggerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTriggerParamsWithTimeout creates a new DeleteTriggerParams object
// with the ability to set a timeout on a request.
func NewDeleteTriggerParamsWithTimeout(timeout time.Duration) *DeleteTriggerParams {
	return &DeleteTriggerParams{
		timeout: timeout,
	}
}

// NewDeleteTriggerParamsWithContext creates a new DeleteTriggerParams object
// with the ability to set a context for a request.
func NewDeleteTriggerParamsWithContext(ctx context.Context) *DeleteTriggerParams {
	return &DeleteTriggerParams{
		Context: ctx,
	}
}

// NewDeleteTriggerParamsWithHTTPClient creates a new DeleteTriggerParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteTriggerParamsWithHTTPClient(client *http.Client) *DeleteTriggerParams {
	return &DeleteTriggerParams{
		HTTPClient: client,
	}
}

/*
DeleteTriggerParams contains all the parameters to send to the API endpoint

	for the delete trigger operation.

	Typically these are written to a http.Request.
*/
type DeleteTriggerParams struct {

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete trigger params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTriggerParams) WithDefaults() *DeleteTriggerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete trigger params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTriggerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete trigger params
func (o *DeleteTriggerParams) WithTimeout(timeout time.Duration) *DeleteTriggerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete trigger params
func (o *DeleteTriggerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete trigger params
func (o *DeleteTriggerParams) WithContext(ctx context.Context) *DeleteTriggerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete trigger params
func (o *DeleteTriggerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete trigger params
func (o *DeleteTriggerParams) WithHTTPClient(client *http.Client) *DeleteTriggerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete trigger params
func (o *DeleteTriggerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete trigger params
func (o *DeleteTriggerParams) WithID(id int32) *DeleteTriggerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete trigger params
func (o *DeleteTriggerParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTriggerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
