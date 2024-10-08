// Code generated by go-swagger; DO NOT EDIT.

package agent_skill_service

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

// NewReadAgentSkillParams creates a new ReadAgentSkillParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadAgentSkillParams() *ReadAgentSkillParams {
	return &ReadAgentSkillParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadAgentSkillParamsWithTimeout creates a new ReadAgentSkillParams object
// with the ability to set a timeout on a request.
func NewReadAgentSkillParamsWithTimeout(timeout time.Duration) *ReadAgentSkillParams {
	return &ReadAgentSkillParams{
		timeout: timeout,
	}
}

// NewReadAgentSkillParamsWithContext creates a new ReadAgentSkillParams object
// with the ability to set a context for a request.
func NewReadAgentSkillParamsWithContext(ctx context.Context) *ReadAgentSkillParams {
	return &ReadAgentSkillParams{
		Context: ctx,
	}
}

// NewReadAgentSkillParamsWithHTTPClient creates a new ReadAgentSkillParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadAgentSkillParamsWithHTTPClient(client *http.Client) *ReadAgentSkillParams {
	return &ReadAgentSkillParams{
		HTTPClient: client,
	}
}

/*
ReadAgentSkillParams contains all the parameters to send to the API endpoint

	for the read agent skill operation.

	Typically these are written to a http.Request.
*/
type ReadAgentSkillParams struct {

	// AgentID.
	//
	// Format: int64
	AgentID string

	// DomainID.
	//
	// Format: int64
	DomainID *string

	// ID.
	//
	// Format: int64
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read agent skill params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadAgentSkillParams) WithDefaults() *ReadAgentSkillParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read agent skill params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadAgentSkillParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read agent skill params
func (o *ReadAgentSkillParams) WithTimeout(timeout time.Duration) *ReadAgentSkillParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read agent skill params
func (o *ReadAgentSkillParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read agent skill params
func (o *ReadAgentSkillParams) WithContext(ctx context.Context) *ReadAgentSkillParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read agent skill params
func (o *ReadAgentSkillParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read agent skill params
func (o *ReadAgentSkillParams) WithHTTPClient(client *http.Client) *ReadAgentSkillParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read agent skill params
func (o *ReadAgentSkillParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentID adds the agentID to the read agent skill params
func (o *ReadAgentSkillParams) WithAgentID(agentID string) *ReadAgentSkillParams {
	o.SetAgentID(agentID)
	return o
}

// SetAgentID adds the agentId to the read agent skill params
func (o *ReadAgentSkillParams) SetAgentID(agentID string) {
	o.AgentID = agentID
}

// WithDomainID adds the domainID to the read agent skill params
func (o *ReadAgentSkillParams) WithDomainID(domainID *string) *ReadAgentSkillParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the read agent skill params
func (o *ReadAgentSkillParams) SetDomainID(domainID *string) {
	o.DomainID = domainID
}

// WithID adds the id to the read agent skill params
func (o *ReadAgentSkillParams) WithID(id string) *ReadAgentSkillParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the read agent skill params
func (o *ReadAgentSkillParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReadAgentSkillParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agent_id
	if err := r.SetPathParam("agent_id", o.AgentID); err != nil {
		return err
	}

	if o.DomainID != nil {

		// query param domain_id
		var qrDomainID string

		if o.DomainID != nil {
			qrDomainID = *o.DomainID
		}
		qDomainID := qrDomainID
		if qDomainID != "" {

			if err := r.SetQueryParam("domain_id", qDomainID); err != nil {
				return err
			}
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
