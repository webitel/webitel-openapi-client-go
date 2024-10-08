// Code generated by go-swagger; DO NOT EDIT.

package agent_pause_cause_service

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

// NewCreateAgentPauseCauseParams creates a new CreateAgentPauseCauseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAgentPauseCauseParams() *CreateAgentPauseCauseParams {
	return &CreateAgentPauseCauseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAgentPauseCauseParamsWithTimeout creates a new CreateAgentPauseCauseParams object
// with the ability to set a timeout on a request.
func NewCreateAgentPauseCauseParamsWithTimeout(timeout time.Duration) *CreateAgentPauseCauseParams {
	return &CreateAgentPauseCauseParams{
		timeout: timeout,
	}
}

// NewCreateAgentPauseCauseParamsWithContext creates a new CreateAgentPauseCauseParams object
// with the ability to set a context for a request.
func NewCreateAgentPauseCauseParamsWithContext(ctx context.Context) *CreateAgentPauseCauseParams {
	return &CreateAgentPauseCauseParams{
		Context: ctx,
	}
}

// NewCreateAgentPauseCauseParamsWithHTTPClient creates a new CreateAgentPauseCauseParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAgentPauseCauseParamsWithHTTPClient(client *http.Client) *CreateAgentPauseCauseParams {
	return &CreateAgentPauseCauseParams{
		HTTPClient: client,
	}
}

/*
CreateAgentPauseCauseParams contains all the parameters to send to the API endpoint

	for the create agent pause cause operation.

	Typically these are written to a http.Request.
*/
type CreateAgentPauseCauseParams struct {

	// Body.
	Body *models.EngineCreateAgentPauseCauseRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create agent pause cause params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAgentPauseCauseParams) WithDefaults() *CreateAgentPauseCauseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create agent pause cause params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAgentPauseCauseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create agent pause cause params
func (o *CreateAgentPauseCauseParams) WithTimeout(timeout time.Duration) *CreateAgentPauseCauseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create agent pause cause params
func (o *CreateAgentPauseCauseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create agent pause cause params
func (o *CreateAgentPauseCauseParams) WithContext(ctx context.Context) *CreateAgentPauseCauseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create agent pause cause params
func (o *CreateAgentPauseCauseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create agent pause cause params
func (o *CreateAgentPauseCauseParams) WithHTTPClient(client *http.Client) *CreateAgentPauseCauseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create agent pause cause params
func (o *CreateAgentPauseCauseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create agent pause cause params
func (o *CreateAgentPauseCauseParams) WithBody(body *models.EngineCreateAgentPauseCauseRequest) *CreateAgentPauseCauseParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create agent pause cause params
func (o *CreateAgentPauseCauseParams) SetBody(body *models.EngineCreateAgentPauseCauseRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAgentPauseCauseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
