// Code generated by go-swagger; DO NOT EDIT.

package email_profile_service

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

// NewPatchEmailProfileParams creates a new PatchEmailProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchEmailProfileParams() *PatchEmailProfileParams {
	return &PatchEmailProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchEmailProfileParamsWithTimeout creates a new PatchEmailProfileParams object
// with the ability to set a timeout on a request.
func NewPatchEmailProfileParamsWithTimeout(timeout time.Duration) *PatchEmailProfileParams {
	return &PatchEmailProfileParams{
		timeout: timeout,
	}
}

// NewPatchEmailProfileParamsWithContext creates a new PatchEmailProfileParams object
// with the ability to set a context for a request.
func NewPatchEmailProfileParamsWithContext(ctx context.Context) *PatchEmailProfileParams {
	return &PatchEmailProfileParams{
		Context: ctx,
	}
}

// NewPatchEmailProfileParamsWithHTTPClient creates a new PatchEmailProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchEmailProfileParamsWithHTTPClient(client *http.Client) *PatchEmailProfileParams {
	return &PatchEmailProfileParams{
		HTTPClient: client,
	}
}

/*
PatchEmailProfileParams contains all the parameters to send to the API endpoint

	for the patch email profile operation.

	Typically these are written to a http.Request.
*/
type PatchEmailProfileParams struct {

	// Body.
	Body *models.EnginePatchEmailProfileRequest

	// ID.
	//
	// Format: int64
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch email profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchEmailProfileParams) WithDefaults() *PatchEmailProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch email profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchEmailProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch email profile params
func (o *PatchEmailProfileParams) WithTimeout(timeout time.Duration) *PatchEmailProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch email profile params
func (o *PatchEmailProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch email profile params
func (o *PatchEmailProfileParams) WithContext(ctx context.Context) *PatchEmailProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch email profile params
func (o *PatchEmailProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch email profile params
func (o *PatchEmailProfileParams) WithHTTPClient(client *http.Client) *PatchEmailProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch email profile params
func (o *PatchEmailProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch email profile params
func (o *PatchEmailProfileParams) WithBody(body *models.EnginePatchEmailProfileRequest) *PatchEmailProfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch email profile params
func (o *PatchEmailProfileParams) SetBody(body *models.EnginePatchEmailProfileRequest) {
	o.Body = body
}

// WithID adds the id to the patch email profile params
func (o *PatchEmailProfileParams) WithID(id string) *PatchEmailProfileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch email profile params
func (o *PatchEmailProfileParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchEmailProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
