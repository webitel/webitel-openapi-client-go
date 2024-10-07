// Code generated by go-swagger; DO NOT EDIT.

package classes

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
)

// NewClassesReadClassParams creates a new ClassesReadClassParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClassesReadClassParams() *ClassesReadClassParams {
	return &ClassesReadClassParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClassesReadClassParamsWithTimeout creates a new ClassesReadClassParams object
// with the ability to set a timeout on a request.
func NewClassesReadClassParamsWithTimeout(timeout time.Duration) *ClassesReadClassParams {
	return &ClassesReadClassParams{
		timeout: timeout,
	}
}

// NewClassesReadClassParamsWithContext creates a new ClassesReadClassParams object
// with the ability to set a context for a request.
func NewClassesReadClassParamsWithContext(ctx context.Context) *ClassesReadClassParams {
	return &ClassesReadClassParams{
		Context: ctx,
	}
}

// NewClassesReadClassParamsWithHTTPClient creates a new ClassesReadClassParams object
// with the ability to set a custom HTTPClient for a request.
func NewClassesReadClassParamsWithHTTPClient(client *http.Client) *ClassesReadClassParams {
	return &ClassesReadClassParams{
		HTTPClient: client,
	}
}

/*
ClassesReadClassParams contains all the parameters to send to the API endpoint

	for the classes read class operation.

	Typically these are written to a http.Request.
*/
type ClassesReadClassParams struct {

	/* Class.

	   [filter]: like '%class%'
	*/
	Class *string

	// Domain.
	Domain *string

	/* ID.

	   [filter]: obj.id = id

	   Format: int64
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the classes read class params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClassesReadClassParams) WithDefaults() *ClassesReadClassParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the classes read class params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClassesReadClassParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the classes read class params
func (o *ClassesReadClassParams) WithTimeout(timeout time.Duration) *ClassesReadClassParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the classes read class params
func (o *ClassesReadClassParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the classes read class params
func (o *ClassesReadClassParams) WithContext(ctx context.Context) *ClassesReadClassParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the classes read class params
func (o *ClassesReadClassParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the classes read class params
func (o *ClassesReadClassParams) WithHTTPClient(client *http.Client) *ClassesReadClassParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the classes read class params
func (o *ClassesReadClassParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClass adds the class to the classes read class params
func (o *ClassesReadClassParams) WithClass(class *string) *ClassesReadClassParams {
	o.SetClass(class)
	return o
}

// SetClass adds the class to the classes read class params
func (o *ClassesReadClassParams) SetClass(class *string) {
	o.Class = class
}

// WithDomain adds the domain to the classes read class params
func (o *ClassesReadClassParams) WithDomain(domain *string) *ClassesReadClassParams {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the classes read class params
func (o *ClassesReadClassParams) SetDomain(domain *string) {
	o.Domain = domain
}

// WithID adds the id to the classes read class params
func (o *ClassesReadClassParams) WithID(id string) *ClassesReadClassParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the classes read class params
func (o *ClassesReadClassParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ClassesReadClassParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Class != nil {

		// query param class
		var qrClass string

		if o.Class != nil {
			qrClass = *o.Class
		}
		qClass := qrClass
		if qClass != "" {

			if err := r.SetQueryParam("class", qClass); err != nil {
				return err
			}
		}
	}

	if o.Domain != nil {

		// query param domain
		var qrDomain string

		if o.Domain != nil {
			qrDomain = *o.Domain
		}
		qDomain := qrDomain
		if qDomain != "" {

			if err := r.SetQueryParam("domain", qDomain); err != nil {
				return err
			}
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
