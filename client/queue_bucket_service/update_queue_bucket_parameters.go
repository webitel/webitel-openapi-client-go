// Code generated by go-swagger; DO NOT EDIT.

package queue_bucket_service

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

// NewUpdateQueueBucketParams creates a new UpdateQueueBucketParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateQueueBucketParams() *UpdateQueueBucketParams {
	return &UpdateQueueBucketParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateQueueBucketParamsWithTimeout creates a new UpdateQueueBucketParams object
// with the ability to set a timeout on a request.
func NewUpdateQueueBucketParamsWithTimeout(timeout time.Duration) *UpdateQueueBucketParams {
	return &UpdateQueueBucketParams{
		timeout: timeout,
	}
}

// NewUpdateQueueBucketParamsWithContext creates a new UpdateQueueBucketParams object
// with the ability to set a context for a request.
func NewUpdateQueueBucketParamsWithContext(ctx context.Context) *UpdateQueueBucketParams {
	return &UpdateQueueBucketParams{
		Context: ctx,
	}
}

// NewUpdateQueueBucketParamsWithHTTPClient creates a new UpdateQueueBucketParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateQueueBucketParamsWithHTTPClient(client *http.Client) *UpdateQueueBucketParams {
	return &UpdateQueueBucketParams{
		HTTPClient: client,
	}
}

/*
UpdateQueueBucketParams contains all the parameters to send to the API endpoint

	for the update queue bucket operation.

	Typically these are written to a http.Request.
*/
type UpdateQueueBucketParams struct {

	// Body.
	Body *models.EngineUpdateQueueBucketRequest

	// ID.
	//
	// Format: int64
	ID string

	// QueueID.
	//
	// Format: int64
	QueueID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update queue bucket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateQueueBucketParams) WithDefaults() *UpdateQueueBucketParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update queue bucket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateQueueBucketParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update queue bucket params
func (o *UpdateQueueBucketParams) WithTimeout(timeout time.Duration) *UpdateQueueBucketParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update queue bucket params
func (o *UpdateQueueBucketParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update queue bucket params
func (o *UpdateQueueBucketParams) WithContext(ctx context.Context) *UpdateQueueBucketParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update queue bucket params
func (o *UpdateQueueBucketParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update queue bucket params
func (o *UpdateQueueBucketParams) WithHTTPClient(client *http.Client) *UpdateQueueBucketParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update queue bucket params
func (o *UpdateQueueBucketParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update queue bucket params
func (o *UpdateQueueBucketParams) WithBody(body *models.EngineUpdateQueueBucketRequest) *UpdateQueueBucketParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update queue bucket params
func (o *UpdateQueueBucketParams) SetBody(body *models.EngineUpdateQueueBucketRequest) {
	o.Body = body
}

// WithID adds the id to the update queue bucket params
func (o *UpdateQueueBucketParams) WithID(id string) *UpdateQueueBucketParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update queue bucket params
func (o *UpdateQueueBucketParams) SetID(id string) {
	o.ID = id
}

// WithQueueID adds the queueID to the update queue bucket params
func (o *UpdateQueueBucketParams) WithQueueID(queueID string) *UpdateQueueBucketParams {
	o.SetQueueID(queueID)
	return o
}

// SetQueueID adds the queueId to the update queue bucket params
func (o *UpdateQueueBucketParams) SetQueueID(queueID string) {
	o.QueueID = queueID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateQueueBucketParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param queue_id
	if err := r.SetPathParam("queue_id", o.QueueID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
