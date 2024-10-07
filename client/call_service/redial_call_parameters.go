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

// NewRedialCallParams creates a new RedialCallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRedialCallParams() *RedialCallParams {
	return &RedialCallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRedialCallParamsWithTimeout creates a new RedialCallParams object
// with the ability to set a timeout on a request.
func NewRedialCallParamsWithTimeout(timeout time.Duration) *RedialCallParams {
	return &RedialCallParams{
		timeout: timeout,
	}
}

// NewRedialCallParamsWithContext creates a new RedialCallParams object
// with the ability to set a context for a request.
func NewRedialCallParamsWithContext(ctx context.Context) *RedialCallParams {
	return &RedialCallParams{
		Context: ctx,
	}
}

// NewRedialCallParamsWithHTTPClient creates a new RedialCallParams object
// with the ability to set a custom HTTPClient for a request.
func NewRedialCallParamsWithHTTPClient(client *http.Client) *RedialCallParams {
	return &RedialCallParams{
		HTTPClient: client,
	}
}

/*
RedialCallParams contains all the parameters to send to the API endpoint

	for the redial call operation.

	Typically these are written to a http.Request.
*/
type RedialCallParams struct {

	// Body.
	Body *models.EngineRedialCallRequest

	// CallID.
	CallID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the redial call params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RedialCallParams) WithDefaults() *RedialCallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the redial call params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RedialCallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the redial call params
func (o *RedialCallParams) WithTimeout(timeout time.Duration) *RedialCallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the redial call params
func (o *RedialCallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the redial call params
func (o *RedialCallParams) WithContext(ctx context.Context) *RedialCallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the redial call params
func (o *RedialCallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the redial call params
func (o *RedialCallParams) WithHTTPClient(client *http.Client) *RedialCallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the redial call params
func (o *RedialCallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the redial call params
func (o *RedialCallParams) WithBody(body *models.EngineRedialCallRequest) *RedialCallParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the redial call params
func (o *RedialCallParams) SetBody(body *models.EngineRedialCallRequest) {
	o.Body = body
}

// WithCallID adds the callID to the redial call params
func (o *RedialCallParams) WithCallID(callID string) *RedialCallParams {
	o.SetCallID(callID)
	return o
}

// SetCallID adds the callId to the redial call params
func (o *RedialCallParams) SetCallID(callID string) {
	o.CallID = callID
}

// WriteToRequest writes these params to a swagger request
func (o *RedialCallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param call_id
	if err := r.SetPathParam("call_id", o.CallID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
