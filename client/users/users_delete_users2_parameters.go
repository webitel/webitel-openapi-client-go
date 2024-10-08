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

// NewUsersDeleteUsers2Params creates a new UsersDeleteUsers2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersDeleteUsers2Params() *UsersDeleteUsers2Params {
	return &UsersDeleteUsers2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersDeleteUsers2ParamsWithTimeout creates a new UsersDeleteUsers2Params object
// with the ability to set a timeout on a request.
func NewUsersDeleteUsers2ParamsWithTimeout(timeout time.Duration) *UsersDeleteUsers2Params {
	return &UsersDeleteUsers2Params{
		timeout: timeout,
	}
}

// NewUsersDeleteUsers2ParamsWithContext creates a new UsersDeleteUsers2Params object
// with the ability to set a context for a request.
func NewUsersDeleteUsers2ParamsWithContext(ctx context.Context) *UsersDeleteUsers2Params {
	return &UsersDeleteUsers2Params{
		Context: ctx,
	}
}

// NewUsersDeleteUsers2ParamsWithHTTPClient creates a new UsersDeleteUsers2Params object
// with the ability to set a custom HTTPClient for a request.
func NewUsersDeleteUsers2ParamsWithHTTPClient(client *http.Client) *UsersDeleteUsers2Params {
	return &UsersDeleteUsers2Params{
		HTTPClient: client,
	}
}

/*
UsersDeleteUsers2Params contains all the parameters to send to the API endpoint

	for the users delete users2 operation.

	Typically these are written to a http.Request.
*/
type UsersDeleteUsers2Params struct {

	/* ID.

	   ONE /users/{id}

	   Format: int64
	*/
	ID *string

	// Permanent.
	Permanent *bool

	/* Selection.

	   MANY /users .ids=[id,...]
	*/
	Selection []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users delete users2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersDeleteUsers2Params) WithDefaults() *UsersDeleteUsers2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users delete users2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersDeleteUsers2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users delete users2 params
func (o *UsersDeleteUsers2Params) WithTimeout(timeout time.Duration) *UsersDeleteUsers2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users delete users2 params
func (o *UsersDeleteUsers2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users delete users2 params
func (o *UsersDeleteUsers2Params) WithContext(ctx context.Context) *UsersDeleteUsers2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users delete users2 params
func (o *UsersDeleteUsers2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users delete users2 params
func (o *UsersDeleteUsers2Params) WithHTTPClient(client *http.Client) *UsersDeleteUsers2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users delete users2 params
func (o *UsersDeleteUsers2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the users delete users2 params
func (o *UsersDeleteUsers2Params) WithID(id *string) *UsersDeleteUsers2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the users delete users2 params
func (o *UsersDeleteUsers2Params) SetID(id *string) {
	o.ID = id
}

// WithPermanent adds the permanent to the users delete users2 params
func (o *UsersDeleteUsers2Params) WithPermanent(permanent *bool) *UsersDeleteUsers2Params {
	o.SetPermanent(permanent)
	return o
}

// SetPermanent adds the permanent to the users delete users2 params
func (o *UsersDeleteUsers2Params) SetPermanent(permanent *bool) {
	o.Permanent = permanent
}

// WithSelection adds the selection to the users delete users2 params
func (o *UsersDeleteUsers2Params) WithSelection(selection []string) *UsersDeleteUsers2Params {
	o.SetSelection(selection)
	return o
}

// SetSelection adds the selection to the users delete users2 params
func (o *UsersDeleteUsers2Params) SetSelection(selection []string) {
	o.Selection = selection
}

// WriteToRequest writes these params to a swagger request
func (o *UsersDeleteUsers2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.Permanent != nil {

		// query param permanent
		var qrPermanent bool

		if o.Permanent != nil {
			qrPermanent = *o.Permanent
		}
		qPermanent := swag.FormatBool(qrPermanent)
		if qPermanent != "" {

			if err := r.SetQueryParam("permanent", qPermanent); err != nil {
				return err
			}
		}
	}
	if o.Selection != nil {
		if err := r.SetBodyParam(o.Selection); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
