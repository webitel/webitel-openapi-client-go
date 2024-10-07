// Code generated by go-swagger; DO NOT EDIT.

package variables

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

// NewVariablesMergeVariablesParams creates a new VariablesMergeVariablesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVariablesMergeVariablesParams() *VariablesMergeVariablesParams {
	return &VariablesMergeVariablesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVariablesMergeVariablesParamsWithTimeout creates a new VariablesMergeVariablesParams object
// with the ability to set a timeout on a request.
func NewVariablesMergeVariablesParamsWithTimeout(timeout time.Duration) *VariablesMergeVariablesParams {
	return &VariablesMergeVariablesParams{
		timeout: timeout,
	}
}

// NewVariablesMergeVariablesParamsWithContext creates a new VariablesMergeVariablesParams object
// with the ability to set a context for a request.
func NewVariablesMergeVariablesParamsWithContext(ctx context.Context) *VariablesMergeVariablesParams {
	return &VariablesMergeVariablesParams{
		Context: ctx,
	}
}

// NewVariablesMergeVariablesParamsWithHTTPClient creates a new VariablesMergeVariablesParams object
// with the ability to set a custom HTTPClient for a request.
func NewVariablesMergeVariablesParamsWithHTTPClient(client *http.Client) *VariablesMergeVariablesParams {
	return &VariablesMergeVariablesParams{
		HTTPClient: client,
	}
}

/*
VariablesMergeVariablesParams contains all the parameters to send to the API endpoint

	for the variables merge variables operation.

	Typically these are written to a http.Request.
*/
type VariablesMergeVariablesParams struct {

	/* ContactID.

	   Link contact ID.
	*/
	ContactID string

	/* Fields.

	   Fields to be retrieved as a result.
	*/
	Fields []string

	/* Input.

	     Fixed object of unique variables to associate with a Contact.
	Each individual key of an object represents a separate variable.
	Variable{key} already linked with the Contact and listed here will be updated.
	*/
	Input []*models.WebitelContactsInputVariable

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the variables merge variables params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VariablesMergeVariablesParams) WithDefaults() *VariablesMergeVariablesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the variables merge variables params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VariablesMergeVariablesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the variables merge variables params
func (o *VariablesMergeVariablesParams) WithTimeout(timeout time.Duration) *VariablesMergeVariablesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the variables merge variables params
func (o *VariablesMergeVariablesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the variables merge variables params
func (o *VariablesMergeVariablesParams) WithContext(ctx context.Context) *VariablesMergeVariablesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the variables merge variables params
func (o *VariablesMergeVariablesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the variables merge variables params
func (o *VariablesMergeVariablesParams) WithHTTPClient(client *http.Client) *VariablesMergeVariablesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the variables merge variables params
func (o *VariablesMergeVariablesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactID adds the contactID to the variables merge variables params
func (o *VariablesMergeVariablesParams) WithContactID(contactID string) *VariablesMergeVariablesParams {
	o.SetContactID(contactID)
	return o
}

// SetContactID adds the contactId to the variables merge variables params
func (o *VariablesMergeVariablesParams) SetContactID(contactID string) {
	o.ContactID = contactID
}

// WithFields adds the fields to the variables merge variables params
func (o *VariablesMergeVariablesParams) WithFields(fields []string) *VariablesMergeVariablesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the variables merge variables params
func (o *VariablesMergeVariablesParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithInput adds the input to the variables merge variables params
func (o *VariablesMergeVariablesParams) WithInput(input []*models.WebitelContactsInputVariable) *VariablesMergeVariablesParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the variables merge variables params
func (o *VariablesMergeVariablesParams) SetInput(input []*models.WebitelContactsInputVariable) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *VariablesMergeVariablesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamVariablesMergeVariables binds the parameter fields
func (o *VariablesMergeVariablesParams) bindParamFields(formats strfmt.Registry) []string {
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
