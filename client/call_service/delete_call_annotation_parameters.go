// Code generated by go-swagger; DO NOT EDIT.

package call_service

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

// NewDeleteCallAnnotationParams creates a new DeleteCallAnnotationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteCallAnnotationParams() *DeleteCallAnnotationParams {
	return &DeleteCallAnnotationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCallAnnotationParamsWithTimeout creates a new DeleteCallAnnotationParams object
// with the ability to set a timeout on a request.
func NewDeleteCallAnnotationParamsWithTimeout(timeout time.Duration) *DeleteCallAnnotationParams {
	return &DeleteCallAnnotationParams{
		timeout: timeout,
	}
}

// NewDeleteCallAnnotationParamsWithContext creates a new DeleteCallAnnotationParams object
// with the ability to set a context for a request.
func NewDeleteCallAnnotationParamsWithContext(ctx context.Context) *DeleteCallAnnotationParams {
	return &DeleteCallAnnotationParams{
		Context: ctx,
	}
}

// NewDeleteCallAnnotationParamsWithHTTPClient creates a new DeleteCallAnnotationParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteCallAnnotationParamsWithHTTPClient(client *http.Client) *DeleteCallAnnotationParams {
	return &DeleteCallAnnotationParams{
		HTTPClient: client,
	}
}

/*
DeleteCallAnnotationParams contains all the parameters to send to the API endpoint

	for the delete call annotation operation.

	Typically these are written to a http.Request.
*/
type DeleteCallAnnotationParams struct {

	// CallID.
	CallID string

	// ID.
	//
	// Format: int64
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete call annotation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCallAnnotationParams) WithDefaults() *DeleteCallAnnotationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete call annotation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCallAnnotationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete call annotation params
func (o *DeleteCallAnnotationParams) WithTimeout(timeout time.Duration) *DeleteCallAnnotationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete call annotation params
func (o *DeleteCallAnnotationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete call annotation params
func (o *DeleteCallAnnotationParams) WithContext(ctx context.Context) *DeleteCallAnnotationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete call annotation params
func (o *DeleteCallAnnotationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete call annotation params
func (o *DeleteCallAnnotationParams) WithHTTPClient(client *http.Client) *DeleteCallAnnotationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete call annotation params
func (o *DeleteCallAnnotationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCallID adds the callID to the delete call annotation params
func (o *DeleteCallAnnotationParams) WithCallID(callID string) *DeleteCallAnnotationParams {
	o.SetCallID(callID)
	return o
}

// SetCallID adds the callId to the delete call annotation params
func (o *DeleteCallAnnotationParams) SetCallID(callID string) {
	o.CallID = callID
}

// WithID adds the id to the delete call annotation params
func (o *DeleteCallAnnotationParams) WithID(id string) *DeleteCallAnnotationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete call annotation params
func (o *DeleteCallAnnotationParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCallAnnotationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param call_id
	if err := r.SetPathParam("call_id", o.CallID); err != nil {
		return err
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
