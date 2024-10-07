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

	"github.com/webitel/webitel-openapi-client-go/models"
)

// NewContactGroupsResetContactGroupsParams creates a new ContactGroupsResetContactGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContactGroupsResetContactGroupsParams() *ContactGroupsResetContactGroupsParams {
	return &ContactGroupsResetContactGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContactGroupsResetContactGroupsParamsWithTimeout creates a new ContactGroupsResetContactGroupsParams object
// with the ability to set a timeout on a request.
func NewContactGroupsResetContactGroupsParamsWithTimeout(timeout time.Duration) *ContactGroupsResetContactGroupsParams {
	return &ContactGroupsResetContactGroupsParams{
		timeout: timeout,
	}
}

// NewContactGroupsResetContactGroupsParamsWithContext creates a new ContactGroupsResetContactGroupsParams object
// with the ability to set a context for a request.
func NewContactGroupsResetContactGroupsParamsWithContext(ctx context.Context) *ContactGroupsResetContactGroupsParams {
	return &ContactGroupsResetContactGroupsParams{
		Context: ctx,
	}
}

// NewContactGroupsResetContactGroupsParamsWithHTTPClient creates a new ContactGroupsResetContactGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewContactGroupsResetContactGroupsParamsWithHTTPClient(client *http.Client) *ContactGroupsResetContactGroupsParams {
	return &ContactGroupsResetContactGroupsParams{
		HTTPClient: client,
	}
}

/*
ContactGroupsResetContactGroupsParams contains all the parameters to send to the API endpoint

	for the contact groups reset contact groups operation.

	Typically these are written to a http.Request.
*/
type ContactGroupsResetContactGroupsParams struct {

	/* ContactID.

	   Link contact ID.
	*/
	ContactID string

	/* Fields.

	   Fields to be retrieved into result of changes.
	*/
	Fields []string

	/* Input.

	     Final set of group(s) to be linked with the contact.
	Group(s) that are already linked with the contact
	but not given in here will be removed.
	*/
	Input []*models.WebitelContactsInputContactGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the contact groups reset contact groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContactGroupsResetContactGroupsParams) WithDefaults() *ContactGroupsResetContactGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the contact groups reset contact groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContactGroupsResetContactGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the contact groups reset contact groups params
func (o *ContactGroupsResetContactGroupsParams) WithTimeout(timeout time.Duration) *ContactGroupsResetContactGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the contact groups reset contact groups params
func (o *ContactGroupsResetContactGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the contact groups reset contact groups params
func (o *ContactGroupsResetContactGroupsParams) WithContext(ctx context.Context) *ContactGroupsResetContactGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the contact groups reset contact groups params
func (o *ContactGroupsResetContactGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the contact groups reset contact groups params
func (o *ContactGroupsResetContactGroupsParams) WithHTTPClient(client *http.Client) *ContactGroupsResetContactGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the contact groups reset contact groups params
func (o *ContactGroupsResetContactGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactID adds the contactID to the contact groups reset contact groups params
func (o *ContactGroupsResetContactGroupsParams) WithContactID(contactID string) *ContactGroupsResetContactGroupsParams {
	o.SetContactID(contactID)
	return o
}

// SetContactID adds the contactId to the contact groups reset contact groups params
func (o *ContactGroupsResetContactGroupsParams) SetContactID(contactID string) {
	o.ContactID = contactID
}

// WithFields adds the fields to the contact groups reset contact groups params
func (o *ContactGroupsResetContactGroupsParams) WithFields(fields []string) *ContactGroupsResetContactGroupsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the contact groups reset contact groups params
func (o *ContactGroupsResetContactGroupsParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithInput adds the input to the contact groups reset contact groups params
func (o *ContactGroupsResetContactGroupsParams) WithInput(input []*models.WebitelContactsInputContactGroup) *ContactGroupsResetContactGroupsParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the contact groups reset contact groups params
func (o *ContactGroupsResetContactGroupsParams) SetInput(input []*models.WebitelContactsInputContactGroup) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *ContactGroupsResetContactGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contact_id
	if err := r.SetPathParam("contact_id", o.ContactID); err != nil {
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
	if o.Input != nil {
		if err := r.SetBodyParam(o.Input); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamContactGroupsResetContactGroups binds the parameter fields
func (o *ContactGroupsResetContactGroupsParams) bindParamFields(formats strfmt.Registry) []string {
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
