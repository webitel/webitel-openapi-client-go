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

	"github.com/webitel/webitel-openapi-client-go/models"
)

// NewPatchHistoryCallParams creates a new PatchHistoryCallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchHistoryCallParams() *PatchHistoryCallParams {
	return &PatchHistoryCallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchHistoryCallParamsWithTimeout creates a new PatchHistoryCallParams object
// with the ability to set a timeout on a request.
func NewPatchHistoryCallParamsWithTimeout(timeout time.Duration) *PatchHistoryCallParams {
	return &PatchHistoryCallParams{
		timeout: timeout,
	}
}

// NewPatchHistoryCallParamsWithContext creates a new PatchHistoryCallParams object
// with the ability to set a context for a request.
func NewPatchHistoryCallParamsWithContext(ctx context.Context) *PatchHistoryCallParams {
	return &PatchHistoryCallParams{
		Context: ctx,
	}
}

// NewPatchHistoryCallParamsWithHTTPClient creates a new PatchHistoryCallParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchHistoryCallParamsWithHTTPClient(client *http.Client) *PatchHistoryCallParams {
	return &PatchHistoryCallParams{
		HTTPClient: client,
	}
}

/*
PatchHistoryCallParams contains all the parameters to send to the API endpoint

	for the patch history call operation.

	Typically these are written to a http.Request.
*/
type PatchHistoryCallParams struct {

	// Body.
	Body *models.EnginePatchHistoryCallRequest

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch history call params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchHistoryCallParams) WithDefaults() *PatchHistoryCallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch history call params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchHistoryCallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch history call params
func (o *PatchHistoryCallParams) WithTimeout(timeout time.Duration) *PatchHistoryCallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch history call params
func (o *PatchHistoryCallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch history call params
func (o *PatchHistoryCallParams) WithContext(ctx context.Context) *PatchHistoryCallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch history call params
func (o *PatchHistoryCallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch history call params
func (o *PatchHistoryCallParams) WithHTTPClient(client *http.Client) *PatchHistoryCallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch history call params
func (o *PatchHistoryCallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch history call params
func (o *PatchHistoryCallParams) WithBody(body *models.EnginePatchHistoryCallRequest) *PatchHistoryCallParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch history call params
func (o *PatchHistoryCallParams) SetBody(body *models.EnginePatchHistoryCallRequest) {
	o.Body = body
}

// WithID adds the id to the patch history call params
func (o *PatchHistoryCallParams) WithID(id string) *PatchHistoryCallParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch history call params
func (o *PatchHistoryCallParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchHistoryCallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
