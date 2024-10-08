// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewUsersReadUserParams creates a new UsersReadUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersReadUserParams() *UsersReadUserParams {
	return &UsersReadUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersReadUserParamsWithTimeout creates a new UsersReadUserParams object
// with the ability to set a timeout on a request.
func NewUsersReadUserParamsWithTimeout(timeout time.Duration) *UsersReadUserParams {
	return &UsersReadUserParams{
		timeout: timeout,
	}
}

// NewUsersReadUserParamsWithContext creates a new UsersReadUserParams object
// with the ability to set a context for a request.
func NewUsersReadUserParamsWithContext(ctx context.Context) *UsersReadUserParams {
	return &UsersReadUserParams{
		Context: ctx,
	}
}

// NewUsersReadUserParamsWithHTTPClient creates a new UsersReadUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersReadUserParamsWithHTTPClient(client *http.Client) *UsersReadUserParams {
	return &UsersReadUserParams{
		HTTPClient: client,
	}
}

/*
UsersReadUserParams contains all the parameters to send to the API endpoint

	for the users read user operation.

	Typically these are written to a http.Request.
*/
type UsersReadUserParams struct {

	/* Fields.

	   partial output
	*/
	Fields []string

	// ID.
	//
	// Format: int64
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users read user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersReadUserParams) WithDefaults() *UsersReadUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users read user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersReadUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users read user params
func (o *UsersReadUserParams) WithTimeout(timeout time.Duration) *UsersReadUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users read user params
func (o *UsersReadUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users read user params
func (o *UsersReadUserParams) WithContext(ctx context.Context) *UsersReadUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users read user params
func (o *UsersReadUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users read user params
func (o *UsersReadUserParams) WithHTTPClient(client *http.Client) *UsersReadUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users read user params
func (o *UsersReadUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the users read user params
func (o *UsersReadUserParams) WithFields(fields []string) *UsersReadUserParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the users read user params
func (o *UsersReadUserParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithID adds the id to the users read user params
func (o *UsersReadUserParams) WithID(id string) *UsersReadUserParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users read user params
func (o *UsersReadUserParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UsersReadUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamUsersReadUser binds the parameter fields
func (o *UsersReadUserParams) bindParamFields(formats strfmt.Registry) []string {
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
