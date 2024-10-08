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

// NewUsersUpdateUserParams creates a new UsersUpdateUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersUpdateUserParams() *UsersUpdateUserParams {
	return &UsersUpdateUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersUpdateUserParamsWithTimeout creates a new UsersUpdateUserParams object
// with the ability to set a timeout on a request.
func NewUsersUpdateUserParamsWithTimeout(timeout time.Duration) *UsersUpdateUserParams {
	return &UsersUpdateUserParams{
		timeout: timeout,
	}
}

// NewUsersUpdateUserParamsWithContext creates a new UsersUpdateUserParams object
// with the ability to set a context for a request.
func NewUsersUpdateUserParamsWithContext(ctx context.Context) *UsersUpdateUserParams {
	return &UsersUpdateUserParams{
		Context: ctx,
	}
}

// NewUsersUpdateUserParamsWithHTTPClient creates a new UsersUpdateUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersUpdateUserParamsWithHTTPClient(client *http.Client) *UsersUpdateUserParams {
	return &UsersUpdateUserParams{
		HTTPClient: client,
	}
}

/*
UsersUpdateUserParams contains all the parameters to send to the API endpoint

	for the users update user operation.

	Typically these are written to a http.Request.
*/
type UsersUpdateUserParams struct {

	// Body.
	Body *models.APIUsersUpdateUserBody

	/* UserID.

	   Object ID

	   Format: int64
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users update user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersUpdateUserParams) WithDefaults() *UsersUpdateUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users update user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersUpdateUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users update user params
func (o *UsersUpdateUserParams) WithTimeout(timeout time.Duration) *UsersUpdateUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users update user params
func (o *UsersUpdateUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users update user params
func (o *UsersUpdateUserParams) WithContext(ctx context.Context) *UsersUpdateUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users update user params
func (o *UsersUpdateUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users update user params
func (o *UsersUpdateUserParams) WithHTTPClient(client *http.Client) *UsersUpdateUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users update user params
func (o *UsersUpdateUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the users update user params
func (o *UsersUpdateUserParams) WithBody(body *models.APIUsersUpdateUserBody) *UsersUpdateUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the users update user params
func (o *UsersUpdateUserParams) SetBody(body *models.APIUsersUpdateUserBody) {
	o.Body = body
}

// WithUserID adds the userID to the users update user params
func (o *UsersUpdateUserParams) WithUserID(userID string) *UsersUpdateUserParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the users update user params
func (o *UsersUpdateUserParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UsersUpdateUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param user.id
	if err := r.SetPathParam("user.id", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
