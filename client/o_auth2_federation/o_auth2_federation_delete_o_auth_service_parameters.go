// Code generated by go-swagger; DO NOT EDIT.

package o_auth2_federation

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

// NewOAuth2FederationDeleteOAuthServiceParams creates a new OAuth2FederationDeleteOAuthServiceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOAuth2FederationDeleteOAuthServiceParams() *OAuth2FederationDeleteOAuthServiceParams {
	return &OAuth2FederationDeleteOAuthServiceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOAuth2FederationDeleteOAuthServiceParamsWithTimeout creates a new OAuth2FederationDeleteOAuthServiceParams object
// with the ability to set a timeout on a request.
func NewOAuth2FederationDeleteOAuthServiceParamsWithTimeout(timeout time.Duration) *OAuth2FederationDeleteOAuthServiceParams {
	return &OAuth2FederationDeleteOAuthServiceParams{
		timeout: timeout,
	}
}

// NewOAuth2FederationDeleteOAuthServiceParamsWithContext creates a new OAuth2FederationDeleteOAuthServiceParams object
// with the ability to set a context for a request.
func NewOAuth2FederationDeleteOAuthServiceParamsWithContext(ctx context.Context) *OAuth2FederationDeleteOAuthServiceParams {
	return &OAuth2FederationDeleteOAuthServiceParams{
		Context: ctx,
	}
}

// NewOAuth2FederationDeleteOAuthServiceParamsWithHTTPClient creates a new OAuth2FederationDeleteOAuthServiceParams object
// with the ability to set a custom HTTPClient for a request.
func NewOAuth2FederationDeleteOAuthServiceParamsWithHTTPClient(client *http.Client) *OAuth2FederationDeleteOAuthServiceParams {
	return &OAuth2FederationDeleteOAuthServiceParams{
		HTTPClient: client,
	}
}

/*
OAuth2FederationDeleteOAuthServiceParams contains all the parameters to send to the API endpoint

	for the o auth2 federation delete o auth service operation.

	Typically these are written to a http.Request.
*/
type OAuth2FederationDeleteOAuthServiceParams struct {

	// Body.
	Body *models.APIDeleteOAuthServiceRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the o auth2 federation delete o auth service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OAuth2FederationDeleteOAuthServiceParams) WithDefaults() *OAuth2FederationDeleteOAuthServiceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the o auth2 federation delete o auth service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OAuth2FederationDeleteOAuthServiceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the o auth2 federation delete o auth service params
func (o *OAuth2FederationDeleteOAuthServiceParams) WithTimeout(timeout time.Duration) *OAuth2FederationDeleteOAuthServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the o auth2 federation delete o auth service params
func (o *OAuth2FederationDeleteOAuthServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the o auth2 federation delete o auth service params
func (o *OAuth2FederationDeleteOAuthServiceParams) WithContext(ctx context.Context) *OAuth2FederationDeleteOAuthServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the o auth2 federation delete o auth service params
func (o *OAuth2FederationDeleteOAuthServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the o auth2 federation delete o auth service params
func (o *OAuth2FederationDeleteOAuthServiceParams) WithHTTPClient(client *http.Client) *OAuth2FederationDeleteOAuthServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the o auth2 federation delete o auth service params
func (o *OAuth2FederationDeleteOAuthServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the o auth2 federation delete o auth service params
func (o *OAuth2FederationDeleteOAuthServiceParams) WithBody(body *models.APIDeleteOAuthServiceRequest) *OAuth2FederationDeleteOAuthServiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the o auth2 federation delete o auth service params
func (o *OAuth2FederationDeleteOAuthServiceParams) SetBody(body *models.APIDeleteOAuthServiceRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *OAuth2FederationDeleteOAuthServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
