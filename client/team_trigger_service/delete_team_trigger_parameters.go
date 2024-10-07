// Code generated by go-swagger; DO NOT EDIT.

package team_trigger_service

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

// NewDeleteTeamTriggerParams creates a new DeleteTeamTriggerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteTeamTriggerParams() *DeleteTeamTriggerParams {
	return &DeleteTeamTriggerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTeamTriggerParamsWithTimeout creates a new DeleteTeamTriggerParams object
// with the ability to set a timeout on a request.
func NewDeleteTeamTriggerParamsWithTimeout(timeout time.Duration) *DeleteTeamTriggerParams {
	return &DeleteTeamTriggerParams{
		timeout: timeout,
	}
}

// NewDeleteTeamTriggerParamsWithContext creates a new DeleteTeamTriggerParams object
// with the ability to set a context for a request.
func NewDeleteTeamTriggerParamsWithContext(ctx context.Context) *DeleteTeamTriggerParams {
	return &DeleteTeamTriggerParams{
		Context: ctx,
	}
}

// NewDeleteTeamTriggerParamsWithHTTPClient creates a new DeleteTeamTriggerParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteTeamTriggerParamsWithHTTPClient(client *http.Client) *DeleteTeamTriggerParams {
	return &DeleteTeamTriggerParams{
		HTTPClient: client,
	}
}

/*
DeleteTeamTriggerParams contains all the parameters to send to the API endpoint

	for the delete team trigger operation.

	Typically these are written to a http.Request.
*/
type DeleteTeamTriggerParams struct {

	// ID.
	//
	// Format: int64
	ID int64

	// TeamID.
	//
	// Format: int64
	TeamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete team trigger params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTeamTriggerParams) WithDefaults() *DeleteTeamTriggerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete team trigger params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTeamTriggerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete team trigger params
func (o *DeleteTeamTriggerParams) WithTimeout(timeout time.Duration) *DeleteTeamTriggerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete team trigger params
func (o *DeleteTeamTriggerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete team trigger params
func (o *DeleteTeamTriggerParams) WithContext(ctx context.Context) *DeleteTeamTriggerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete team trigger params
func (o *DeleteTeamTriggerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete team trigger params
func (o *DeleteTeamTriggerParams) WithHTTPClient(client *http.Client) *DeleteTeamTriggerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete team trigger params
func (o *DeleteTeamTriggerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete team trigger params
func (o *DeleteTeamTriggerParams) WithID(id int64) *DeleteTeamTriggerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete team trigger params
func (o *DeleteTeamTriggerParams) SetID(id int64) {
	o.ID = id
}

// WithTeamID adds the teamID to the delete team trigger params
func (o *DeleteTeamTriggerParams) WithTeamID(teamID string) *DeleteTeamTriggerParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the delete team trigger params
func (o *DeleteTeamTriggerParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTeamTriggerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param team_id
	if err := r.SetPathParam("team_id", o.TeamID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
