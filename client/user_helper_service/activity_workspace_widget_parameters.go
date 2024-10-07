// Code generated by go-swagger; DO NOT EDIT.

package user_helper_service

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
)

// NewActivityWorkspaceWidgetParams creates a new ActivityWorkspaceWidgetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewActivityWorkspaceWidgetParams() *ActivityWorkspaceWidgetParams {
	return &ActivityWorkspaceWidgetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewActivityWorkspaceWidgetParamsWithTimeout creates a new ActivityWorkspaceWidgetParams object
// with the ability to set a timeout on a request.
func NewActivityWorkspaceWidgetParamsWithTimeout(timeout time.Duration) *ActivityWorkspaceWidgetParams {
	return &ActivityWorkspaceWidgetParams{
		timeout: timeout,
	}
}

// NewActivityWorkspaceWidgetParamsWithContext creates a new ActivityWorkspaceWidgetParams object
// with the ability to set a context for a request.
func NewActivityWorkspaceWidgetParamsWithContext(ctx context.Context) *ActivityWorkspaceWidgetParams {
	return &ActivityWorkspaceWidgetParams{
		Context: ctx,
	}
}

// NewActivityWorkspaceWidgetParamsWithHTTPClient creates a new ActivityWorkspaceWidgetParams object
// with the ability to set a custom HTTPClient for a request.
func NewActivityWorkspaceWidgetParamsWithHTTPClient(client *http.Client) *ActivityWorkspaceWidgetParams {
	return &ActivityWorkspaceWidgetParams{
		HTTPClient: client,
	}
}

/*
ActivityWorkspaceWidgetParams contains all the parameters to send to the API endpoint

	for the activity workspace widget operation.

	Typically these are written to a http.Request.
*/
type ActivityWorkspaceWidgetParams struct {

	// Fields.
	Fields []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the activity workspace widget params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActivityWorkspaceWidgetParams) WithDefaults() *ActivityWorkspaceWidgetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the activity workspace widget params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActivityWorkspaceWidgetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the activity workspace widget params
func (o *ActivityWorkspaceWidgetParams) WithTimeout(timeout time.Duration) *ActivityWorkspaceWidgetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the activity workspace widget params
func (o *ActivityWorkspaceWidgetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the activity workspace widget params
func (o *ActivityWorkspaceWidgetParams) WithContext(ctx context.Context) *ActivityWorkspaceWidgetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the activity workspace widget params
func (o *ActivityWorkspaceWidgetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the activity workspace widget params
func (o *ActivityWorkspaceWidgetParams) WithHTTPClient(client *http.Client) *ActivityWorkspaceWidgetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the activity workspace widget params
func (o *ActivityWorkspaceWidgetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the activity workspace widget params
func (o *ActivityWorkspaceWidgetParams) WithFields(fields []string) *ActivityWorkspaceWidgetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the activity workspace widget params
func (o *ActivityWorkspaceWidgetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *ActivityWorkspaceWidgetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamActivityWorkspaceWidget binds the parameter fields
func (o *ActivityWorkspaceWidgetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "multi"
	fieldsIS := swag.JoinByFormat(fieldsIC, "multi")

	return fieldsIS
}
