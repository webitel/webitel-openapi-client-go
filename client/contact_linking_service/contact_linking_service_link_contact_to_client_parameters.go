// Code generated by go-swagger; DO NOT EDIT.

package contact_linking_service

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
)

// NewContactLinkingServiceLinkContactToClientParams creates a new ContactLinkingServiceLinkContactToClientParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContactLinkingServiceLinkContactToClientParams() *ContactLinkingServiceLinkContactToClientParams {
	return &ContactLinkingServiceLinkContactToClientParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContactLinkingServiceLinkContactToClientParamsWithTimeout creates a new ContactLinkingServiceLinkContactToClientParams object
// with the ability to set a timeout on a request.
func NewContactLinkingServiceLinkContactToClientParamsWithTimeout(timeout time.Duration) *ContactLinkingServiceLinkContactToClientParams {
	return &ContactLinkingServiceLinkContactToClientParams{
		timeout: timeout,
	}
}

// NewContactLinkingServiceLinkContactToClientParamsWithContext creates a new ContactLinkingServiceLinkContactToClientParams object
// with the ability to set a context for a request.
func NewContactLinkingServiceLinkContactToClientParamsWithContext(ctx context.Context) *ContactLinkingServiceLinkContactToClientParams {
	return &ContactLinkingServiceLinkContactToClientParams{
		Context: ctx,
	}
}

// NewContactLinkingServiceLinkContactToClientParamsWithHTTPClient creates a new ContactLinkingServiceLinkContactToClientParams object
// with the ability to set a custom HTTPClient for a request.
func NewContactLinkingServiceLinkContactToClientParamsWithHTTPClient(client *http.Client) *ContactLinkingServiceLinkContactToClientParams {
	return &ContactLinkingServiceLinkContactToClientParams{
		HTTPClient: client,
	}
}

/*
ContactLinkingServiceLinkContactToClientParams contains all the parameters to send to the API endpoint

	for the contact linking service link contact to client operation.

	Typically these are written to a http.Request.
*/
type ContactLinkingServiceLinkContactToClientParams struct {

	// ContactID.
	ContactID *string

	// ConversationID.
	ConversationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the contact linking service link contact to client params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContactLinkingServiceLinkContactToClientParams) WithDefaults() *ContactLinkingServiceLinkContactToClientParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the contact linking service link contact to client params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContactLinkingServiceLinkContactToClientParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the contact linking service link contact to client params
func (o *ContactLinkingServiceLinkContactToClientParams) WithTimeout(timeout time.Duration) *ContactLinkingServiceLinkContactToClientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the contact linking service link contact to client params
func (o *ContactLinkingServiceLinkContactToClientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the contact linking service link contact to client params
func (o *ContactLinkingServiceLinkContactToClientParams) WithContext(ctx context.Context) *ContactLinkingServiceLinkContactToClientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the contact linking service link contact to client params
func (o *ContactLinkingServiceLinkContactToClientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the contact linking service link contact to client params
func (o *ContactLinkingServiceLinkContactToClientParams) WithHTTPClient(client *http.Client) *ContactLinkingServiceLinkContactToClientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the contact linking service link contact to client params
func (o *ContactLinkingServiceLinkContactToClientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactID adds the contactID to the contact linking service link contact to client params
func (o *ContactLinkingServiceLinkContactToClientParams) WithContactID(contactID *string) *ContactLinkingServiceLinkContactToClientParams {
	o.SetContactID(contactID)
	return o
}

// SetContactID adds the contactId to the contact linking service link contact to client params
func (o *ContactLinkingServiceLinkContactToClientParams) SetContactID(contactID *string) {
	o.ContactID = contactID
}

// WithConversationID adds the conversationID to the contact linking service link contact to client params
func (o *ContactLinkingServiceLinkContactToClientParams) WithConversationID(conversationID string) *ContactLinkingServiceLinkContactToClientParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the contact linking service link contact to client params
func (o *ContactLinkingServiceLinkContactToClientParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WriteToRequest writes these params to a swagger request
func (o *ContactLinkingServiceLinkContactToClientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContactID != nil {

		// query param contact_id
		var qrContactID string

		if o.ContactID != nil {
			qrContactID = *o.ContactID
		}
		qContactID := qrContactID
		if qContactID != "" {

			if err := r.SetQueryParam("contact_id", qContactID); err != nil {
				return err
			}
		}
	}

	// path param conversation_id
	if err := r.SetPathParam("conversation_id", o.ConversationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
