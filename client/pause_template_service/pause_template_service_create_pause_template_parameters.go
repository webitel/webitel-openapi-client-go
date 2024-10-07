// Code generated by go-swagger; DO NOT EDIT.

package pause_template_service

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

// NewPauseTemplateServiceCreatePauseTemplateParams creates a new PauseTemplateServiceCreatePauseTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPauseTemplateServiceCreatePauseTemplateParams() *PauseTemplateServiceCreatePauseTemplateParams {
	return &PauseTemplateServiceCreatePauseTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPauseTemplateServiceCreatePauseTemplateParamsWithTimeout creates a new PauseTemplateServiceCreatePauseTemplateParams object
// with the ability to set a timeout on a request.
func NewPauseTemplateServiceCreatePauseTemplateParamsWithTimeout(timeout time.Duration) *PauseTemplateServiceCreatePauseTemplateParams {
	return &PauseTemplateServiceCreatePauseTemplateParams{
		timeout: timeout,
	}
}

// NewPauseTemplateServiceCreatePauseTemplateParamsWithContext creates a new PauseTemplateServiceCreatePauseTemplateParams object
// with the ability to set a context for a request.
func NewPauseTemplateServiceCreatePauseTemplateParamsWithContext(ctx context.Context) *PauseTemplateServiceCreatePauseTemplateParams {
	return &PauseTemplateServiceCreatePauseTemplateParams{
		Context: ctx,
	}
}

// NewPauseTemplateServiceCreatePauseTemplateParamsWithHTTPClient creates a new PauseTemplateServiceCreatePauseTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPauseTemplateServiceCreatePauseTemplateParamsWithHTTPClient(client *http.Client) *PauseTemplateServiceCreatePauseTemplateParams {
	return &PauseTemplateServiceCreatePauseTemplateParams{
		HTTPClient: client,
	}
}

/*
PauseTemplateServiceCreatePauseTemplateParams contains all the parameters to send to the API endpoint

	for the pause template service create pause template operation.

	Typically these are written to a http.Request.
*/
type PauseTemplateServiceCreatePauseTemplateParams struct {

	// Body.
	Body *models.WfmCreatePauseTemplateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pause template service create pause template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PauseTemplateServiceCreatePauseTemplateParams) WithDefaults() *PauseTemplateServiceCreatePauseTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pause template service create pause template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PauseTemplateServiceCreatePauseTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pause template service create pause template params
func (o *PauseTemplateServiceCreatePauseTemplateParams) WithTimeout(timeout time.Duration) *PauseTemplateServiceCreatePauseTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pause template service create pause template params
func (o *PauseTemplateServiceCreatePauseTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pause template service create pause template params
func (o *PauseTemplateServiceCreatePauseTemplateParams) WithContext(ctx context.Context) *PauseTemplateServiceCreatePauseTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pause template service create pause template params
func (o *PauseTemplateServiceCreatePauseTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pause template service create pause template params
func (o *PauseTemplateServiceCreatePauseTemplateParams) WithHTTPClient(client *http.Client) *PauseTemplateServiceCreatePauseTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pause template service create pause template params
func (o *PauseTemplateServiceCreatePauseTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pause template service create pause template params
func (o *PauseTemplateServiceCreatePauseTemplateParams) WithBody(body *models.WfmCreatePauseTemplateRequest) *PauseTemplateServiceCreatePauseTemplateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pause template service create pause template params
func (o *PauseTemplateServiceCreatePauseTemplateParams) SetBody(body *models.WfmCreatePauseTemplateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PauseTemplateServiceCreatePauseTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
