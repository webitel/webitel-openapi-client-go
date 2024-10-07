// Code generated by go-swagger; DO NOT EDIT.

package queue_hook_service

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
	"github.com/go-openapi/swag"

	"github.com/webitel/webitel-openapi-client-go/models"
)

// NewPatchQueueHookParams creates a new PatchQueueHookParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchQueueHookParams() *PatchQueueHookParams {
	return &PatchQueueHookParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchQueueHookParamsWithTimeout creates a new PatchQueueHookParams object
// with the ability to set a timeout on a request.
func NewPatchQueueHookParamsWithTimeout(timeout time.Duration) *PatchQueueHookParams {
	return &PatchQueueHookParams{
		timeout: timeout,
	}
}

// NewPatchQueueHookParamsWithContext creates a new PatchQueueHookParams object
// with the ability to set a context for a request.
func NewPatchQueueHookParamsWithContext(ctx context.Context) *PatchQueueHookParams {
	return &PatchQueueHookParams{
		Context: ctx,
	}
}

// NewPatchQueueHookParamsWithHTTPClient creates a new PatchQueueHookParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchQueueHookParamsWithHTTPClient(client *http.Client) *PatchQueueHookParams {
	return &PatchQueueHookParams{
		HTTPClient: client,
	}
}

/*
PatchQueueHookParams contains all the parameters to send to the API endpoint

	for the patch queue hook operation.

	Typically these are written to a http.Request.
*/
type PatchQueueHookParams struct {

	// Body.
	Body *models.EnginePatchQueueHookRequest

	// ID.
	//
	// Format: int64
	ID int64

	// QueueID.
	//
	// Format: int64
	QueueID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch queue hook params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchQueueHookParams) WithDefaults() *PatchQueueHookParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch queue hook params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchQueueHookParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch queue hook params
func (o *PatchQueueHookParams) WithTimeout(timeout time.Duration) *PatchQueueHookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch queue hook params
func (o *PatchQueueHookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch queue hook params
func (o *PatchQueueHookParams) WithContext(ctx context.Context) *PatchQueueHookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch queue hook params
func (o *PatchQueueHookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch queue hook params
func (o *PatchQueueHookParams) WithHTTPClient(client *http.Client) *PatchQueueHookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch queue hook params
func (o *PatchQueueHookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch queue hook params
func (o *PatchQueueHookParams) WithBody(body *models.EnginePatchQueueHookRequest) *PatchQueueHookParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch queue hook params
func (o *PatchQueueHookParams) SetBody(body *models.EnginePatchQueueHookRequest) {
	o.Body = body
}

// WithID adds the id to the patch queue hook params
func (o *PatchQueueHookParams) WithID(id int64) *PatchQueueHookParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch queue hook params
func (o *PatchQueueHookParams) SetID(id int64) {
	o.ID = id
}

// WithQueueID adds the queueID to the patch queue hook params
func (o *PatchQueueHookParams) WithQueueID(queueID int64) *PatchQueueHookParams {
	o.SetQueueID(queueID)
	return o
}

// SetQueueID adds the queueId to the patch queue hook params
func (o *PatchQueueHookParams) SetQueueID(queueID int64) {
	o.QueueID = queueID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchQueueHookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param queue_id
	if err := r.SetPathParam("queue_id", swag.FormatInt64(o.QueueID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
