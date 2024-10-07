// Code generated by go-swagger; DO NOT EDIT.

package preset_query_service

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

// NewDeletePresetQueryParams creates a new DeletePresetQueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePresetQueryParams() *DeletePresetQueryParams {
	return &DeletePresetQueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePresetQueryParamsWithTimeout creates a new DeletePresetQueryParams object
// with the ability to set a timeout on a request.
func NewDeletePresetQueryParamsWithTimeout(timeout time.Duration) *DeletePresetQueryParams {
	return &DeletePresetQueryParams{
		timeout: timeout,
	}
}

// NewDeletePresetQueryParamsWithContext creates a new DeletePresetQueryParams object
// with the ability to set a context for a request.
func NewDeletePresetQueryParamsWithContext(ctx context.Context) *DeletePresetQueryParams {
	return &DeletePresetQueryParams{
		Context: ctx,
	}
}

// NewDeletePresetQueryParamsWithHTTPClient creates a new DeletePresetQueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePresetQueryParamsWithHTTPClient(client *http.Client) *DeletePresetQueryParams {
	return &DeletePresetQueryParams{
		HTTPClient: client,
	}
}

/*
DeletePresetQueryParams contains all the parameters to send to the API endpoint

	for the delete preset query operation.

	Typically these are written to a http.Request.
*/
type DeletePresetQueryParams struct {

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete preset query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePresetQueryParams) WithDefaults() *DeletePresetQueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete preset query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePresetQueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete preset query params
func (o *DeletePresetQueryParams) WithTimeout(timeout time.Duration) *DeletePresetQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete preset query params
func (o *DeletePresetQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete preset query params
func (o *DeletePresetQueryParams) WithContext(ctx context.Context) *DeletePresetQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete preset query params
func (o *DeletePresetQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete preset query params
func (o *DeletePresetQueryParams) WithHTTPClient(client *http.Client) *DeletePresetQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete preset query params
func (o *DeletePresetQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete preset query params
func (o *DeletePresetQueryParams) WithID(id int32) *DeletePresetQueryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete preset query params
func (o *DeletePresetQueryParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePresetQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
