// Code generated by go-swagger; DO NOT EDIT.

package region_service

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

// NewCreateRegionParams creates a new CreateRegionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRegionParams() *CreateRegionParams {
	return &CreateRegionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRegionParamsWithTimeout creates a new CreateRegionParams object
// with the ability to set a timeout on a request.
func NewCreateRegionParamsWithTimeout(timeout time.Duration) *CreateRegionParams {
	return &CreateRegionParams{
		timeout: timeout,
	}
}

// NewCreateRegionParamsWithContext creates a new CreateRegionParams object
// with the ability to set a context for a request.
func NewCreateRegionParamsWithContext(ctx context.Context) *CreateRegionParams {
	return &CreateRegionParams{
		Context: ctx,
	}
}

// NewCreateRegionParamsWithHTTPClient creates a new CreateRegionParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRegionParamsWithHTTPClient(client *http.Client) *CreateRegionParams {
	return &CreateRegionParams{
		HTTPClient: client,
	}
}

/*
CreateRegionParams contains all the parameters to send to the API endpoint

	for the create region operation.

	Typically these are written to a http.Request.
*/
type CreateRegionParams struct {

	// Body.
	Body *models.EngineCreateRegionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create region params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRegionParams) WithDefaults() *CreateRegionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create region params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRegionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create region params
func (o *CreateRegionParams) WithTimeout(timeout time.Duration) *CreateRegionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create region params
func (o *CreateRegionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create region params
func (o *CreateRegionParams) WithContext(ctx context.Context) *CreateRegionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create region params
func (o *CreateRegionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create region params
func (o *CreateRegionParams) WithHTTPClient(client *http.Client) *CreateRegionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create region params
func (o *CreateRegionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create region params
func (o *CreateRegionParams) WithBody(body *models.EngineCreateRegionRequest) *CreateRegionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create region params
func (o *CreateRegionParams) SetBody(body *models.EngineCreateRegionRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRegionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
