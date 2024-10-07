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

	"github.com/webitel/webitel-openapi-client-go/models"
)

// NewCreateAgentSkillsParams creates a new CreateAgentSkillsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAgentSkillsParams() *CreateAgentSkillsParams {
	return &CreateAgentSkillsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAgentSkillsParamsWithTimeout creates a new CreateAgentSkillsParams object
// with the ability to set a timeout on a request.
func NewCreateAgentSkillsParamsWithTimeout(timeout time.Duration) *CreateAgentSkillsParams {
	return &CreateAgentSkillsParams{
		timeout: timeout,
	}
}

// NewCreateAgentSkillsParamsWithContext creates a new CreateAgentSkillsParams object
// with the ability to set a context for a request.
func NewCreateAgentSkillsParamsWithContext(ctx context.Context) *CreateAgentSkillsParams {
	return &CreateAgentSkillsParams{
		Context: ctx,
	}
}

// NewCreateAgentSkillsParamsWithHTTPClient creates a new CreateAgentSkillsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAgentSkillsParamsWithHTTPClient(client *http.Client) *CreateAgentSkillsParams {
	return &CreateAgentSkillsParams{
		HTTPClient: client,
	}
}

/*
CreateAgentSkillsParams contains all the parameters to send to the API endpoint

	for the create agent skills operation.

	Typically these are written to a http.Request.
*/
type CreateAgentSkillsParams struct {

	// AgentID.
	//
	// Format: int64
	AgentID string

	// Body.
	Body *models.EngineCreateAgentSkillsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create agent skills params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAgentSkillsParams) WithDefaults() *CreateAgentSkillsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create agent skills params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAgentSkillsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create agent skills params
func (o *CreateAgentSkillsParams) WithTimeout(timeout time.Duration) *CreateAgentSkillsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create agent skills params
func (o *CreateAgentSkillsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create agent skills params
func (o *CreateAgentSkillsParams) WithContext(ctx context.Context) *CreateAgentSkillsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create agent skills params
func (o *CreateAgentSkillsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create agent skills params
func (o *CreateAgentSkillsParams) WithHTTPClient(client *http.Client) *CreateAgentSkillsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create agent skills params
func (o *CreateAgentSkillsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentID adds the agentID to the create agent skills params
func (o *CreateAgentSkillsParams) WithAgentID(agentID string) *CreateAgentSkillsParams {
	o.SetAgentID(agentID)
	return o
}

// SetAgentID adds the agentId to the create agent skills params
func (o *CreateAgentSkillsParams) SetAgentID(agentID string) {
	o.AgentID = agentID
}

// WithBody adds the body to the create agent skills params
func (o *CreateAgentSkillsParams) WithBody(body *models.EngineCreateAgentSkillsRequest) *CreateAgentSkillsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create agent skills params
func (o *CreateAgentSkillsParams) SetBody(body *models.EngineCreateAgentSkillsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAgentSkillsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agent_id
	if err := r.SetPathParam("agent_id", o.AgentID); err != nil {
		return err
	}
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
