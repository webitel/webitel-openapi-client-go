// Code generated by go-swagger; DO NOT EDIT.

package contact_groups

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

// NewContactGroupsLocateContactGroupParams creates a new ContactGroupsLocateContactGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContactGroupsLocateContactGroupParams() *ContactGroupsLocateContactGroupParams {
	return &ContactGroupsLocateContactGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContactGroupsLocateContactGroupParamsWithTimeout creates a new ContactGroupsLocateContactGroupParams object
// with the ability to set a timeout on a request.
func NewContactGroupsLocateContactGroupParamsWithTimeout(timeout time.Duration) *ContactGroupsLocateContactGroupParams {
	return &ContactGroupsLocateContactGroupParams{
		timeout: timeout,
	}
}

// NewContactGroupsLocateContactGroupParamsWithContext creates a new ContactGroupsLocateContactGroupParams object
// with the ability to set a context for a request.
func NewContactGroupsLocateContactGroupParamsWithContext(ctx context.Context) *ContactGroupsLocateContactGroupParams {
	return &ContactGroupsLocateContactGroupParams{
		Context: ctx,
	}
}

// NewContactGroupsLocateContactGroupParamsWithHTTPClient creates a new ContactGroupsLocateContactGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewContactGroupsLocateContactGroupParamsWithHTTPClient(client *http.Client) *ContactGroupsLocateContactGroupParams {
	return &ContactGroupsLocateContactGroupParams{
		HTTPClient: client,
	}
}

/*
ContactGroupsLocateContactGroupParams contains all the parameters to send to the API endpoint

	for the contact groups locate contact group operation.

	Typically these are written to a http.Request.
*/
type ContactGroupsLocateContactGroupParams struct {

	/* ContactID.

	   Contact source ID.
	*/
	ContactID string

	/* Etag.

	     Unique group link IDentifier.
	Accept: `etag` (obsolete+) or `id`.

	     Format: \w+
	*/
	Etag string

	/* Fields.

	   Fields to be retrieved into result.
	*/
	Fields []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the contact groups locate contact group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContactGroupsLocateContactGroupParams) WithDefaults() *ContactGroupsLocateContactGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the contact groups locate contact group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContactGroupsLocateContactGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the contact groups locate contact group params
func (o *ContactGroupsLocateContactGroupParams) WithTimeout(timeout time.Duration) *ContactGroupsLocateContactGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the contact groups locate contact group params
func (o *ContactGroupsLocateContactGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the contact groups locate contact group params
func (o *ContactGroupsLocateContactGroupParams) WithContext(ctx context.Context) *ContactGroupsLocateContactGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the contact groups locate contact group params
func (o *ContactGroupsLocateContactGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the contact groups locate contact group params
func (o *ContactGroupsLocateContactGroupParams) WithHTTPClient(client *http.Client) *ContactGroupsLocateContactGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the contact groups locate contact group params
func (o *ContactGroupsLocateContactGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactID adds the contactID to the contact groups locate contact group params
func (o *ContactGroupsLocateContactGroupParams) WithContactID(contactID string) *ContactGroupsLocateContactGroupParams {
	o.SetContactID(contactID)
	return o
}

// SetContactID adds the contactId to the contact groups locate contact group params
func (o *ContactGroupsLocateContactGroupParams) SetContactID(contactID string) {
	o.ContactID = contactID
}

// WithEtag adds the etag to the contact groups locate contact group params
func (o *ContactGroupsLocateContactGroupParams) WithEtag(etag string) *ContactGroupsLocateContactGroupParams {
	o.SetEtag(etag)
	return o
}

// SetEtag adds the etag to the contact groups locate contact group params
func (o *ContactGroupsLocateContactGroupParams) SetEtag(etag string) {
	o.Etag = etag
}

// WithFields adds the fields to the contact groups locate contact group params
func (o *ContactGroupsLocateContactGroupParams) WithFields(fields []string) *ContactGroupsLocateContactGroupParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the contact groups locate contact group params
func (o *ContactGroupsLocateContactGroupParams) SetFields(fields []string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *ContactGroupsLocateContactGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contact_id
	if err := r.SetPathParam("contact_id", o.ContactID); err != nil {
		return err
	}

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

// bindParamContactGroupsLocateContactGroup binds the parameter fields
func (o *ContactGroupsLocateContactGroupParams) bindParamFields(formats strfmt.Registry) []string {
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
