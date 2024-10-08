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

// NewCreateEmailProfileParams creates a new CreateEmailProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateEmailProfileParams() *CreateEmailProfileParams {
	return &CreateEmailProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateEmailProfileParamsWithTimeout creates a new CreateEmailProfileParams object
// with the ability to set a timeout on a request.
func NewCreateEmailProfileParamsWithTimeout(timeout time.Duration) *CreateEmailProfileParams {
	return &CreateEmailProfileParams{
		timeout: timeout,
	}
}

// NewCreateEmailProfileParamsWithContext creates a new CreateEmailProfileParams object
// with the ability to set a context for a request.
func NewCreateEmailProfileParamsWithContext(ctx context.Context) *CreateEmailProfileParams {
	return &CreateEmailProfileParams{
		Context: ctx,
	}
}

// NewCreateEmailProfileParamsWithHTTPClient creates a new CreateEmailProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateEmailProfileParamsWithHTTPClient(client *http.Client) *CreateEmailProfileParams {
	return &CreateEmailProfileParams{
		HTTPClient: client,
	}
}

/*
CreateEmailProfileParams contains all the parameters to send to the API endpoint

	for the create email profile operation.

	Typically these are written to a http.Request.
*/
type CreateEmailProfileParams struct {

	// Body.
	Body *models.EngineCreateEmailProfileRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create email profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateEmailProfileParams) WithDefaults() *CreateEmailProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create email profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateEmailProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create email profile params
func (o *CreateEmailProfileParams) WithTimeout(timeout time.Duration) *CreateEmailProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create email profile params
func (o *CreateEmailProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create email profile params
func (o *CreateEmailProfileParams) WithContext(ctx context.Context) *CreateEmailProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create email profile params
func (o *CreateEmailProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create email profile params
func (o *CreateEmailProfileParams) WithHTTPClient(client *http.Client) *CreateEmailProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create email profile params
func (o *CreateEmailProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create email profile params
func (o *CreateEmailProfileParams) WithBody(body *models.EngineCreateEmailProfileRequest) *CreateEmailProfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create email profile params
func (o *CreateEmailProfileParams) SetBody(body *models.EngineCreateEmailProfileRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateEmailProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
