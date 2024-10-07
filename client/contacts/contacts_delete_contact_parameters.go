// Code generated by go-swagger; DO NOT EDIT.

package contacts

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

// NewContactsDeleteContactParams creates a new ContactsDeleteContactParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContactsDeleteContactParams() *ContactsDeleteContactParams {
	return &ContactsDeleteContactParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContactsDeleteContactParamsWithTimeout creates a new ContactsDeleteContactParams object
// with the ability to set a timeout on a request.
func NewContactsDeleteContactParamsWithTimeout(timeout time.Duration) *ContactsDeleteContactParams {
	return &ContactsDeleteContactParams{
		timeout: timeout,
	}
}

// NewContactsDeleteContactParamsWithContext creates a new ContactsDeleteContactParams object
// with the ability to set a context for a request.
func NewContactsDeleteContactParamsWithContext(ctx context.Context) *ContactsDeleteContactParams {
	return &ContactsDeleteContactParams{
		Context: ctx,
	}
}

// NewContactsDeleteContactParamsWithHTTPClient creates a new ContactsDeleteContactParams object
// with the ability to set a custom HTTPClient for a request.
func NewContactsDeleteContactParamsWithHTTPClient(client *http.Client) *ContactsDeleteContactParams {
	return &ContactsDeleteContactParams{
		HTTPClient: client,
	}
}

/*
ContactsDeleteContactParams contains all the parameters to send to the API endpoint

	for the contacts delete contact operation.

	Typically these are written to a http.Request.
*/
type ContactsDeleteContactParams struct {

	/* Etag.

	   Unique ID of the latest version of a resource.
	*/
	Etag string

	/* Fields.

	   Fields to be retrieved into result of changes.
	*/
	Fields []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the contacts delete contact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContactsDeleteContactParams) WithDefaults() *ContactsDeleteContactParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the contacts delete contact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContactsDeleteContactParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the contacts delete contact params
func (o *ContactsDeleteContactParams) WithTimeout(timeout time.Duration) *ContactsDeleteContactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the contacts delete contact params
func (o *ContactsDeleteContactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the contacts delete contact params
func (o *ContactsDeleteContactParams) WithContext(ctx context.Context) *ContactsDeleteContactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the contacts delete contact params
func (o *ContactsDeleteContactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the contacts delete contact params
func (o *ContactsDeleteContactParams) WithHTTPClient(client *http.Client) *ContactsDeleteContactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the contacts delete contact params
func (o *ContactsDeleteContactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEtag adds the etag to the contacts delete contact params
func (o *ContactsDeleteContactParams) WithEtag(etag string) *ContactsDeleteContactParams {
	o.SetEtag(etag)
	return o
}

// SetEtag adds the etag to the contacts delete contact params
func (o *ContactsDeleteContactParams) SetEtag(etag string) {
	o.Etag = etag
}

// WithFields adds the fields to the contacts delete contact params
func (o *ContactsDeleteContactParams) WithFields(fields []string) *ContactsDeleteContactParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the contacts delete contact params
func (o *ContactsDeleteContactParams) SetFields(fields []string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *ContactsDeleteContactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param etag
	if err := r.SetPathParam("etag", o.Etag); err != nil {
		return err
	}

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

// bindParamContactsDeleteContact binds the parameter fields
func (o *ContactsDeleteContactParams) bindParamFields(formats strfmt.Registry) []string {
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
