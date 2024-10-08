// Code generated by go-swagger; DO NOT EDIT.

package managers

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

// NewManagersUpdateManagerParams creates a new ManagersUpdateManagerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewManagersUpdateManagerParams() *ManagersUpdateManagerParams {
	return &ManagersUpdateManagerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewManagersUpdateManagerParamsWithTimeout creates a new ManagersUpdateManagerParams object
// with the ability to set a timeout on a request.
func NewManagersUpdateManagerParamsWithTimeout(timeout time.Duration) *ManagersUpdateManagerParams {
	return &ManagersUpdateManagerParams{
		timeout: timeout,
	}
}

// NewManagersUpdateManagerParamsWithContext creates a new ManagersUpdateManagerParams object
// with the ability to set a context for a request.
func NewManagersUpdateManagerParamsWithContext(ctx context.Context) *ManagersUpdateManagerParams {
	return &ManagersUpdateManagerParams{
		Context: ctx,
	}
}

// NewManagersUpdateManagerParamsWithHTTPClient creates a new ManagersUpdateManagerParams object
// with the ability to set a custom HTTPClient for a request.
func NewManagersUpdateManagerParamsWithHTTPClient(client *http.Client) *ManagersUpdateManagerParams {
	return &ManagersUpdateManagerParams{
		HTTPClient: client,
	}
}

/*
ManagersUpdateManagerParams contains all the parameters to send to the API endpoint

	for the managers update manager operation.

	Typically these are written to a http.Request.
*/
type ManagersUpdateManagerParams struct {

	/* ContactID.

	   Link contact ID.
	*/
	ContactID string

	/* Etag.

	   Unique ID of the latest version of an existing resource.
	*/
	Etag string

	/* Fields.

	   Fields to be retrieved into result of changes.
	*/
	Fields []string

	/* Input.

	   NEW Update of the manager link.
	*/
	Input *models.ManagersUpdateManagerParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the managers update manager params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ManagersUpdateManagerParams) WithDefaults() *ManagersUpdateManagerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the managers update manager params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ManagersUpdateManagerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the managers update manager params
func (o *ManagersUpdateManagerParams) WithTimeout(timeout time.Duration) *ManagersUpdateManagerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the managers update manager params
func (o *ManagersUpdateManagerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the managers update manager params
func (o *ManagersUpdateManagerParams) WithContext(ctx context.Context) *ManagersUpdateManagerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the managers update manager params
func (o *ManagersUpdateManagerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the managers update manager params
func (o *ManagersUpdateManagerParams) WithHTTPClient(client *http.Client) *ManagersUpdateManagerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the managers update manager params
func (o *ManagersUpdateManagerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactID adds the contactID to the managers update manager params
func (o *ManagersUpdateManagerParams) WithContactID(contactID string) *ManagersUpdateManagerParams {
	o.SetContactID(contactID)
	return o
}

// SetContactID adds the contactId to the managers update manager params
func (o *ManagersUpdateManagerParams) SetContactID(contactID string) {
	o.ContactID = contactID
}

// WithEtag adds the etag to the managers update manager params
func (o *ManagersUpdateManagerParams) WithEtag(etag string) *ManagersUpdateManagerParams {
	o.SetEtag(etag)
	return o
}

// SetEtag adds the etag to the managers update manager params
func (o *ManagersUpdateManagerParams) SetEtag(etag string) {
	o.Etag = etag
}

// WithFields adds the fields to the managers update manager params
func (o *ManagersUpdateManagerParams) WithFields(fields []string) *ManagersUpdateManagerParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the managers update manager params
func (o *ManagersUpdateManagerParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithInput adds the input to the managers update manager params
func (o *ManagersUpdateManagerParams) WithInput(input *models.ManagersUpdateManagerParamsBody) *ManagersUpdateManagerParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the managers update manager params
func (o *ManagersUpdateManagerParams) SetInput(input *models.ManagersUpdateManagerParamsBody) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *ManagersUpdateManagerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamManagersUpdateManager binds the parameter fields
func (o *ManagersUpdateManagerParams) bindParamFields(formats strfmt.Registry) []string {
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
