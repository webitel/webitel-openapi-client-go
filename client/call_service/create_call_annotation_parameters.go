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

// NewCreateCallAnnotationParams creates a new CreateCallAnnotationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateCallAnnotationParams() *CreateCallAnnotationParams {
	return &CreateCallAnnotationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCallAnnotationParamsWithTimeout creates a new CreateCallAnnotationParams object
// with the ability to set a timeout on a request.
func NewCreateCallAnnotationParamsWithTimeout(timeout time.Duration) *CreateCallAnnotationParams {
	return &CreateCallAnnotationParams{
		timeout: timeout,
	}
}

// NewCreateCallAnnotationParamsWithContext creates a new CreateCallAnnotationParams object
// with the ability to set a context for a request.
func NewCreateCallAnnotationParamsWithContext(ctx context.Context) *CreateCallAnnotationParams {
	return &CreateCallAnnotationParams{
		Context: ctx,
	}
}

// NewCreateCallAnnotationParamsWithHTTPClient creates a new CreateCallAnnotationParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateCallAnnotationParamsWithHTTPClient(client *http.Client) *CreateCallAnnotationParams {
	return &CreateCallAnnotationParams{
		HTTPClient: client,
	}
}

/*
CreateCallAnnotationParams contains all the parameters to send to the API endpoint

	for the create call annotation operation.

	Typically these are written to a http.Request.
*/
type CreateCallAnnotationParams struct {

	// Body.
	Body *models.EngineCreateCallAnnotationRequest

	// CallID.
	CallID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create call annotation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCallAnnotationParams) WithDefaults() *CreateCallAnnotationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create call annotation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCallAnnotationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create call annotation params
func (o *CreateCallAnnotationParams) WithTimeout(timeout time.Duration) *CreateCallAnnotationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create call annotation params
func (o *CreateCallAnnotationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create call annotation params
func (o *CreateCallAnnotationParams) WithContext(ctx context.Context) *CreateCallAnnotationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create call annotation params
func (o *CreateCallAnnotationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create call annotation params
func (o *CreateCallAnnotationParams) WithHTTPClient(client *http.Client) *CreateCallAnnotationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create call annotation params
func (o *CreateCallAnnotationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create call annotation params
func (o *CreateCallAnnotationParams) WithBody(body *models.EngineCreateCallAnnotationRequest) *CreateCallAnnotationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create call annotation params
func (o *CreateCallAnnotationParams) SetBody(body *models.EngineCreateCallAnnotationRequest) {
	o.Body = body
}

// WithCallID adds the callID to the create call annotation params
func (o *CreateCallAnnotationParams) WithCallID(callID string) *CreateCallAnnotationParams {
	o.SetCallID(callID)
	return o
}

// SetCallID adds the callId to the create call annotation params
func (o *CreateCallAnnotationParams) SetCallID(callID string) {
	o.CallID = callID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCallAnnotationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
