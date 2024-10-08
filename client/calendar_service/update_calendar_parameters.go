// Code generated by go-swagger; DO NOT EDIT.

package calendar_service

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

// NewUpdateCalendarParams creates a new UpdateCalendarParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateCalendarParams() *UpdateCalendarParams {
	return &UpdateCalendarParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCalendarParamsWithTimeout creates a new UpdateCalendarParams object
// with the ability to set a timeout on a request.
func NewUpdateCalendarParamsWithTimeout(timeout time.Duration) *UpdateCalendarParams {
	return &UpdateCalendarParams{
		timeout: timeout,
	}
}

// NewUpdateCalendarParamsWithContext creates a new UpdateCalendarParams object
// with the ability to set a context for a request.
func NewUpdateCalendarParamsWithContext(ctx context.Context) *UpdateCalendarParams {
	return &UpdateCalendarParams{
		Context: ctx,
	}
}

// NewUpdateCalendarParamsWithHTTPClient creates a new UpdateCalendarParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateCalendarParamsWithHTTPClient(client *http.Client) *UpdateCalendarParams {
	return &UpdateCalendarParams{
		HTTPClient: client,
	}
}

/*
UpdateCalendarParams contains all the parameters to send to the API endpoint

	for the update calendar operation.

	Typically these are written to a http.Request.
*/
type UpdateCalendarParams struct {

	// Body.
	Body *models.EngineUpdateCalendarRequest

	// ID.
	//
	// Format: int64
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update calendar params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCalendarParams) WithDefaults() *UpdateCalendarParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update calendar params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCalendarParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update calendar params
func (o *UpdateCalendarParams) WithTimeout(timeout time.Duration) *UpdateCalendarParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update calendar params
func (o *UpdateCalendarParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update calendar params
func (o *UpdateCalendarParams) WithContext(ctx context.Context) *UpdateCalendarParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update calendar params
func (o *UpdateCalendarParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update calendar params
func (o *UpdateCalendarParams) WithHTTPClient(client *http.Client) *UpdateCalendarParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update calendar params
func (o *UpdateCalendarParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update calendar params
func (o *UpdateCalendarParams) WithBody(body *models.EngineUpdateCalendarRequest) *UpdateCalendarParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update calendar params
func (o *UpdateCalendarParams) SetBody(body *models.EngineUpdateCalendarRequest) {
	o.Body = body
}

// WithID adds the id to the update calendar params
func (o *UpdateCalendarParams) WithID(id string) *UpdateCalendarParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update calendar params
func (o *UpdateCalendarParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCalendarParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
