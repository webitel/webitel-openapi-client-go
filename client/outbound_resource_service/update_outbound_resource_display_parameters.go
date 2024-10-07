// Code generated by go-swagger; DO NOT EDIT.

package outbound_resource_service

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

// NewUpdateOutboundResourceDisplayParams creates a new UpdateOutboundResourceDisplayParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateOutboundResourceDisplayParams() *UpdateOutboundResourceDisplayParams {
	return &UpdateOutboundResourceDisplayParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOutboundResourceDisplayParamsWithTimeout creates a new UpdateOutboundResourceDisplayParams object
// with the ability to set a timeout on a request.
func NewUpdateOutboundResourceDisplayParamsWithTimeout(timeout time.Duration) *UpdateOutboundResourceDisplayParams {
	return &UpdateOutboundResourceDisplayParams{
		timeout: timeout,
	}
}

// NewUpdateOutboundResourceDisplayParamsWithContext creates a new UpdateOutboundResourceDisplayParams object
// with the ability to set a context for a request.
func NewUpdateOutboundResourceDisplayParamsWithContext(ctx context.Context) *UpdateOutboundResourceDisplayParams {
	return &UpdateOutboundResourceDisplayParams{
		Context: ctx,
	}
}

// NewUpdateOutboundResourceDisplayParamsWithHTTPClient creates a new UpdateOutboundResourceDisplayParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateOutboundResourceDisplayParamsWithHTTPClient(client *http.Client) *UpdateOutboundResourceDisplayParams {
	return &UpdateOutboundResourceDisplayParams{
		HTTPClient: client,
	}
}

/*
UpdateOutboundResourceDisplayParams contains all the parameters to send to the API endpoint

	for the update outbound resource display operation.

	Typically these are written to a http.Request.
*/
type UpdateOutboundResourceDisplayParams struct {

	// Body.
	Body *models.EngineUpdateOutboundResourceDisplayRequest

	// ID.
	//
	// Format: int64
	ID string

	// ResourceID.
	//
	// Format: int64
	ResourceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update outbound resource display params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOutboundResourceDisplayParams) WithDefaults() *UpdateOutboundResourceDisplayParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update outbound resource display params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOutboundResourceDisplayParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update outbound resource display params
func (o *UpdateOutboundResourceDisplayParams) WithTimeout(timeout time.Duration) *UpdateOutboundResourceDisplayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update outbound resource display params
func (o *UpdateOutboundResourceDisplayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update outbound resource display params
func (o *UpdateOutboundResourceDisplayParams) WithContext(ctx context.Context) *UpdateOutboundResourceDisplayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update outbound resource display params
func (o *UpdateOutboundResourceDisplayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update outbound resource display params
func (o *UpdateOutboundResourceDisplayParams) WithHTTPClient(client *http.Client) *UpdateOutboundResourceDisplayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update outbound resource display params
func (o *UpdateOutboundResourceDisplayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update outbound resource display params
func (o *UpdateOutboundResourceDisplayParams) WithBody(body *models.EngineUpdateOutboundResourceDisplayRequest) *UpdateOutboundResourceDisplayParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update outbound resource display params
func (o *UpdateOutboundResourceDisplayParams) SetBody(body *models.EngineUpdateOutboundResourceDisplayRequest) {
	o.Body = body
}

// WithID adds the id to the update outbound resource display params
func (o *UpdateOutboundResourceDisplayParams) WithID(id string) *UpdateOutboundResourceDisplayParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update outbound resource display params
func (o *UpdateOutboundResourceDisplayParams) SetID(id string) {
	o.ID = id
}

// WithResourceID adds the resourceID to the update outbound resource display params
func (o *UpdateOutboundResourceDisplayParams) WithResourceID(resourceID string) *UpdateOutboundResourceDisplayParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the update outbound resource display params
func (o *UpdateOutboundResourceDisplayParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOutboundResourceDisplayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param resource_id
	if err := r.SetPathParam("resource_id", o.ResourceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
