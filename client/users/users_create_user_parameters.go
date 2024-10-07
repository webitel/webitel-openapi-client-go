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

// NewUsersCreateUserParams creates a new UsersCreateUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersCreateUserParams() *UsersCreateUserParams {
	return &UsersCreateUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersCreateUserParamsWithTimeout creates a new UsersCreateUserParams object
// with the ability to set a timeout on a request.
func NewUsersCreateUserParamsWithTimeout(timeout time.Duration) *UsersCreateUserParams {
	return &UsersCreateUserParams{
		timeout: timeout,
	}
}

// NewUsersCreateUserParamsWithContext creates a new UsersCreateUserParams object
// with the ability to set a context for a request.
func NewUsersCreateUserParamsWithContext(ctx context.Context) *UsersCreateUserParams {
	return &UsersCreateUserParams{
		Context: ctx,
	}
}

// NewUsersCreateUserParamsWithHTTPClient creates a new UsersCreateUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersCreateUserParamsWithHTTPClient(client *http.Client) *UsersCreateUserParams {
	return &UsersCreateUserParams{
		HTTPClient: client,
	}
}

/*
UsersCreateUserParams contains all the parameters to send to the API endpoint

	for the users create user operation.

	Typically these are written to a http.Request.
*/
type UsersCreateUserParams struct {

	// Body.
	Body *models.APICreateUserRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users create user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersCreateUserParams) WithDefaults() *UsersCreateUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users create user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersCreateUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users create user params
func (o *UsersCreateUserParams) WithTimeout(timeout time.Duration) *UsersCreateUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users create user params
func (o *UsersCreateUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users create user params
func (o *UsersCreateUserParams) WithContext(ctx context.Context) *UsersCreateUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users create user params
func (o *UsersCreateUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users create user params
func (o *UsersCreateUserParams) WithHTTPClient(client *http.Client) *UsersCreateUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users create user params
func (o *UsersCreateUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the users create user params
func (o *UsersCreateUserParams) WithBody(body *models.APICreateUserRequest) *UsersCreateUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the users create user params
func (o *UsersCreateUserParams) SetBody(body *models.APICreateUserRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UsersCreateUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
