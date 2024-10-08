// Code generated by go-swagger; DO NOT EDIT.

package list_service

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

// NewCreateListParams creates a new CreateListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateListParams() *CreateListParams {
	return &CreateListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateListParamsWithTimeout creates a new CreateListParams object
// with the ability to set a timeout on a request.
func NewCreateListParamsWithTimeout(timeout time.Duration) *CreateListParams {
	return &CreateListParams{
		timeout: timeout,
	}
}

// NewCreateListParamsWithContext creates a new CreateListParams object
// with the ability to set a context for a request.
func NewCreateListParamsWithContext(ctx context.Context) *CreateListParams {
	return &CreateListParams{
		Context: ctx,
	}
}

// NewCreateListParamsWithHTTPClient creates a new CreateListParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateListParamsWithHTTPClient(client *http.Client) *CreateListParams {
	return &CreateListParams{
		HTTPClient: client,
	}
}

/*
CreateListParams contains all the parameters to send to the API endpoint

	for the create list operation.

	Typically these are written to a http.Request.
*/
type CreateListParams struct {

	// Body.
	Body *models.EngineCreateListRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateListParams) WithDefaults() *CreateListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create list params
func (o *CreateListParams) WithTimeout(timeout time.Duration) *CreateListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create list params
func (o *CreateListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create list params
func (o *CreateListParams) WithContext(ctx context.Context) *CreateListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create list params
func (o *CreateListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create list params
func (o *CreateListParams) WithHTTPClient(client *http.Client) *CreateListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create list params
func (o *CreateListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create list params
func (o *CreateListParams) WithBody(body *models.EngineCreateListRequest) *CreateListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create list params
func (o *CreateListParams) SetBody(body *models.EngineCreateListRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
