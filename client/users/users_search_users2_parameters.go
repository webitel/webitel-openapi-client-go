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

	"github.com/webitel/webitel-openapi-client-go/models"
)

// NewUsersSearchUsers2Params creates a new UsersSearchUsers2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersSearchUsers2Params() *UsersSearchUsers2Params {
	return &UsersSearchUsers2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersSearchUsers2ParamsWithTimeout creates a new UsersSearchUsers2Params object
// with the ability to set a timeout on a request.
func NewUsersSearchUsers2ParamsWithTimeout(timeout time.Duration) *UsersSearchUsers2Params {
	return &UsersSearchUsers2Params{
		timeout: timeout,
	}
}

// NewUsersSearchUsers2ParamsWithContext creates a new UsersSearchUsers2Params object
// with the ability to set a context for a request.
func NewUsersSearchUsers2ParamsWithContext(ctx context.Context) *UsersSearchUsers2Params {
	return &UsersSearchUsers2Params{
		Context: ctx,
	}
}

// NewUsersSearchUsers2ParamsWithHTTPClient creates a new UsersSearchUsers2Params object
// with the ability to set a custom HTTPClient for a request.
func NewUsersSearchUsers2ParamsWithHTTPClient(client *http.Client) *UsersSearchUsers2Params {
	return &UsersSearchUsers2Params{
		HTTPClient: client,
	}
}

/*
UsersSearchUsers2Params contains all the parameters to send to the API endpoint

	for the users search users2 operation.

	Typically these are written to a http.Request.
*/
type UsersSearchUsers2Params struct {

	// Body.
	Body *models.APISearchUsersRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users search users2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersSearchUsers2Params) WithDefaults() *UsersSearchUsers2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users search users2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersSearchUsers2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users search users2 params
func (o *UsersSearchUsers2Params) WithTimeout(timeout time.Duration) *UsersSearchUsers2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users search users2 params
func (o *UsersSearchUsers2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users search users2 params
func (o *UsersSearchUsers2Params) WithContext(ctx context.Context) *UsersSearchUsers2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users search users2 params
func (o *UsersSearchUsers2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users search users2 params
func (o *UsersSearchUsers2Params) WithHTTPClient(client *http.Client) *UsersSearchUsers2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users search users2 params
func (o *UsersSearchUsers2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the users search users2 params
func (o *UsersSearchUsers2Params) WithBody(body *models.APISearchUsersRequest) *UsersSearchUsers2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the users search users2 params
func (o *UsersSearchUsers2Params) SetBody(body *models.APISearchUsersRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UsersSearchUsers2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
