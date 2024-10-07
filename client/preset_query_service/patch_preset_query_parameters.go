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

	"github.com/webitel/webitel-openapi-client-go/models"
)

// NewPatchPresetQueryParams creates a new PatchPresetQueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchPresetQueryParams() *PatchPresetQueryParams {
	return &PatchPresetQueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchPresetQueryParamsWithTimeout creates a new PatchPresetQueryParams object
// with the ability to set a timeout on a request.
func NewPatchPresetQueryParamsWithTimeout(timeout time.Duration) *PatchPresetQueryParams {
	return &PatchPresetQueryParams{
		timeout: timeout,
	}
}

// NewPatchPresetQueryParamsWithContext creates a new PatchPresetQueryParams object
// with the ability to set a context for a request.
func NewPatchPresetQueryParamsWithContext(ctx context.Context) *PatchPresetQueryParams {
	return &PatchPresetQueryParams{
		Context: ctx,
	}
}

// NewPatchPresetQueryParamsWithHTTPClient creates a new PatchPresetQueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchPresetQueryParamsWithHTTPClient(client *http.Client) *PatchPresetQueryParams {
	return &PatchPresetQueryParams{
		HTTPClient: client,
	}
}

/*
PatchPresetQueryParams contains all the parameters to send to the API endpoint

	for the patch preset query operation.

	Typically these are written to a http.Request.
*/
type PatchPresetQueryParams struct {

	// Body.
	Body *models.EnginePatchPresetQueryRequest

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch preset query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchPresetQueryParams) WithDefaults() *PatchPresetQueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch preset query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchPresetQueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch preset query params
func (o *PatchPresetQueryParams) WithTimeout(timeout time.Duration) *PatchPresetQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch preset query params
func (o *PatchPresetQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch preset query params
func (o *PatchPresetQueryParams) WithContext(ctx context.Context) *PatchPresetQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch preset query params
func (o *PatchPresetQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch preset query params
func (o *PatchPresetQueryParams) WithHTTPClient(client *http.Client) *PatchPresetQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch preset query params
func (o *PatchPresetQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch preset query params
func (o *PatchPresetQueryParams) WithBody(body *models.EnginePatchPresetQueryRequest) *PatchPresetQueryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch preset query params
func (o *PatchPresetQueryParams) SetBody(body *models.EnginePatchPresetQueryRequest) {
	o.Body = body
}

// WithID adds the id to the patch preset query params
func (o *PatchPresetQueryParams) WithID(id int32) *PatchPresetQueryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch preset query params
func (o *PatchPresetQueryParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchPresetQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
