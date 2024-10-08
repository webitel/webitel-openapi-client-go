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

// NewPatchAgentSkillParams creates a new PatchAgentSkillParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAgentSkillParams() *PatchAgentSkillParams {
	return &PatchAgentSkillParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAgentSkillParamsWithTimeout creates a new PatchAgentSkillParams object
// with the ability to set a timeout on a request.
func NewPatchAgentSkillParamsWithTimeout(timeout time.Duration) *PatchAgentSkillParams {
	return &PatchAgentSkillParams{
		timeout: timeout,
	}
}

// NewPatchAgentSkillParamsWithContext creates a new PatchAgentSkillParams object
// with the ability to set a context for a request.
func NewPatchAgentSkillParamsWithContext(ctx context.Context) *PatchAgentSkillParams {
	return &PatchAgentSkillParams{
		Context: ctx,
	}
}

// NewPatchAgentSkillParamsWithHTTPClient creates a new PatchAgentSkillParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAgentSkillParamsWithHTTPClient(client *http.Client) *PatchAgentSkillParams {
	return &PatchAgentSkillParams{
		HTTPClient: client,
	}
}

/*
PatchAgentSkillParams contains all the parameters to send to the API endpoint

	for the patch agent skill operation.

	Typically these are written to a http.Request.
*/
type PatchAgentSkillParams struct {

	// AgentID.
	//
	// Format: int64
	AgentID string

	// Body.
	Body *models.EnginePatchAgentSkillRequest

	// ID.
	//
	// Format: int64
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch agent skill params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAgentSkillParams) WithDefaults() *PatchAgentSkillParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch agent skill params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAgentSkillParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch agent skill params
func (o *PatchAgentSkillParams) WithTimeout(timeout time.Duration) *PatchAgentSkillParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch agent skill params
func (o *PatchAgentSkillParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch agent skill params
func (o *PatchAgentSkillParams) WithContext(ctx context.Context) *PatchAgentSkillParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch agent skill params
func (o *PatchAgentSkillParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch agent skill params
func (o *PatchAgentSkillParams) WithHTTPClient(client *http.Client) *PatchAgentSkillParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch agent skill params
func (o *PatchAgentSkillParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentID adds the agentID to the patch agent skill params
func (o *PatchAgentSkillParams) WithAgentID(agentID string) *PatchAgentSkillParams {
	o.SetAgentID(agentID)
	return o
}

// SetAgentID adds the agentId to the patch agent skill params
func (o *PatchAgentSkillParams) SetAgentID(agentID string) {
	o.AgentID = agentID
}

// WithBody adds the body to the patch agent skill params
func (o *PatchAgentSkillParams) WithBody(body *models.EnginePatchAgentSkillRequest) *PatchAgentSkillParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch agent skill params
func (o *PatchAgentSkillParams) SetBody(body *models.EnginePatchAgentSkillRequest) {
	o.Body = body
}

// WithID adds the id to the patch agent skill params
func (o *PatchAgentSkillParams) WithID(id string) *PatchAgentSkillParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch agent skill params
func (o *PatchAgentSkillParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAgentSkillParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
