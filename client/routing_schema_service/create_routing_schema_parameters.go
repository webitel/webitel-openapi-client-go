// Code generated by go-swagger; DO NOT EDIT.

package routing_schema_service

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

// NewCreateRoutingSchemaParams creates a new CreateRoutingSchemaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRoutingSchemaParams() *CreateRoutingSchemaParams {
	return &CreateRoutingSchemaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRoutingSchemaParamsWithTimeout creates a new CreateRoutingSchemaParams object
// with the ability to set a timeout on a request.
func NewCreateRoutingSchemaParamsWithTimeout(timeout time.Duration) *CreateRoutingSchemaParams {
	return &CreateRoutingSchemaParams{
		timeout: timeout,
	}
}

// NewCreateRoutingSchemaParamsWithContext creates a new CreateRoutingSchemaParams object
// with the ability to set a context for a request.
func NewCreateRoutingSchemaParamsWithContext(ctx context.Context) *CreateRoutingSchemaParams {
	return &CreateRoutingSchemaParams{
		Context: ctx,
	}
}

// NewCreateRoutingSchemaParamsWithHTTPClient creates a new CreateRoutingSchemaParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRoutingSchemaParamsWithHTTPClient(client *http.Client) *CreateRoutingSchemaParams {
	return &CreateRoutingSchemaParams{
		HTTPClient: client,
	}
}

/*
CreateRoutingSchemaParams contains all the parameters to send to the API endpoint

	for the create routing schema operation.

	Typically these are written to a http.Request.
*/
type CreateRoutingSchemaParams struct {

	// Body.
	Body *models.EngineCreateRoutingSchemaRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create routing schema params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRoutingSchemaParams) WithDefaults() *CreateRoutingSchemaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create routing schema params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRoutingSchemaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create routing schema params
func (o *CreateRoutingSchemaParams) WithTimeout(timeout time.Duration) *CreateRoutingSchemaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create routing schema params
func (o *CreateRoutingSchemaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create routing schema params
func (o *CreateRoutingSchemaParams) WithContext(ctx context.Context) *CreateRoutingSchemaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create routing schema params
func (o *CreateRoutingSchemaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create routing schema params
func (o *CreateRoutingSchemaParams) WithHTTPClient(client *http.Client) *CreateRoutingSchemaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create routing schema params
func (o *CreateRoutingSchemaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create routing schema params
func (o *CreateRoutingSchemaParams) WithBody(body *models.EngineCreateRoutingSchemaRequest) *CreateRoutingSchemaParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create routing schema params
func (o *CreateRoutingSchemaParams) SetBody(body *models.EngineCreateRoutingSchemaRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRoutingSchemaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
