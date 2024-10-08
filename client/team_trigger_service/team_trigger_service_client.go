// Code generated by go-swagger; DO NOT EDIT.

package team_trigger_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/webitel-openapi-client-go/models"
)

// New creates a new team trigger service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for team trigger service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateTeamTrigger(teamID string, body *models.EngineCreateTeamTriggerRequest, opts ...ClientOption) (*CreateTeamTriggerOK, error)
	CreateTeamTriggerWithParams(params *CreateTeamTriggerParams, opts ...ClientOption) (*CreateTeamTriggerOK, error)

	DeleteTeamTrigger(teamID string, id int64, opts ...ClientOption) (*DeleteTeamTriggerOK, error)
	DeleteTeamTriggerWithParams(params *DeleteTeamTriggerParams, opts ...ClientOption) (*DeleteTeamTriggerOK, error)

	PatchTeamTrigger(params *PatchTeamTriggerParams, opts ...ClientOption) (*PatchTeamTriggerOK, error)

	ReadTeamTrigger(teamID string, id int64, opts ...ClientOption) (*ReadTeamTriggerOK, error)
	ReadTeamTriggerWithParams(params *ReadTeamTriggerParams, opts ...ClientOption) (*ReadTeamTriggerOK, error)

	RunTeamTrigger(triggerID int32, body *models.EngineRunTeamTriggerRequest, opts ...ClientOption) (*RunTeamTriggerOK, error)
	RunTeamTriggerWithParams(params *RunTeamTriggerParams, opts ...ClientOption) (*RunTeamTriggerOK, error)

	SearchAgentTrigger(params *SearchAgentTriggerParams, opts ...ClientOption) (*SearchAgentTriggerOK, error)

	SearchTeamTrigger(params *SearchTeamTriggerParams, opts ...ClientOption) (*SearchTeamTriggerOK, error)

	UpdateTeamTrigger(params *UpdateTeamTriggerParams, opts ...ClientOption) (*UpdateTeamTriggerOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateTeamTrigger create team trigger API
*/
func (a *Client) CreateTeamTrigger(teamID string, body *models.EngineCreateTeamTriggerRequest, opts ...ClientOption) (*CreateTeamTriggerOK, error) {
	params := NewCreateTeamTriggerParams().WithBody(body).WithTeamID(teamID)
	return a.CreateTeamTriggerWithParams(params, opts...)
}

func (a *Client) CreateTeamTriggerWithParams(params *CreateTeamTriggerParams, opts ...ClientOption) (*CreateTeamTriggerOK, error) {
	if params == nil {
		params = NewCreateTeamTriggerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateTeamTrigger",
		Method:             "POST",
		PathPattern:        "/call_center/teams/{team_id}/triggers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateTeamTriggerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateTeamTriggerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateTeamTriggerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteTeamTrigger delete team trigger API
*/
func (a *Client) DeleteTeamTrigger(teamID string, id int64, opts ...ClientOption) (*DeleteTeamTriggerOK, error) {
	params := NewDeleteTeamTriggerParams().WithID(id).WithTeamID(teamID)
	return a.DeleteTeamTriggerWithParams(params, opts...)
}

func (a *Client) DeleteTeamTriggerWithParams(params *DeleteTeamTriggerParams, opts ...ClientOption) (*DeleteTeamTriggerOK, error) {
	if params == nil {
		params = NewDeleteTeamTriggerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteTeamTrigger",
		Method:             "DELETE",
		PathPattern:        "/call_center/teams/{team_id}/triggers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteTeamTriggerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteTeamTriggerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteTeamTriggerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchTeamTrigger patch team trigger API
*/

func (a *Client) PatchTeamTrigger(params *PatchTeamTriggerParams, opts ...ClientOption) (*PatchTeamTriggerOK, error) {
	if params == nil {
		params = NewPatchTeamTriggerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchTeamTrigger",
		Method:             "PATCH",
		PathPattern:        "/call_center/teams/{team_id}/triggers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchTeamTriggerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchTeamTriggerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchTeamTriggerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReadTeamTrigger read team trigger API
*/
func (a *Client) ReadTeamTrigger(teamID string, id int64, opts ...ClientOption) (*ReadTeamTriggerOK, error) {
	params := NewReadTeamTriggerParams().WithID(id).WithTeamID(teamID)
	return a.ReadTeamTriggerWithParams(params, opts...)
}

func (a *Client) ReadTeamTriggerWithParams(params *ReadTeamTriggerParams, opts ...ClientOption) (*ReadTeamTriggerOK, error) {
	if params == nil {
		params = NewReadTeamTriggerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadTeamTrigger",
		Method:             "GET",
		PathPattern:        "/call_center/teams/{team_id}/triggers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReadTeamTriggerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadTeamTriggerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadTeamTriggerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RunTeamTrigger run team trigger API
*/
func (a *Client) RunTeamTrigger(triggerID int32, body *models.EngineRunTeamTriggerRequest, opts ...ClientOption) (*RunTeamTriggerOK, error) {
	params := NewRunTeamTriggerParams().WithBody(body).WithTriggerID(triggerID)
	return a.RunTeamTriggerWithParams(params, opts...)
}

func (a *Client) RunTeamTriggerWithParams(params *RunTeamTriggerParams, opts ...ClientOption) (*RunTeamTriggerOK, error) {
	if params == nil {
		params = NewRunTeamTriggerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RunTeamTrigger",
		Method:             "POST",
		PathPattern:        "/call_center/teams/triggers/{trigger_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RunTeamTriggerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RunTeamTriggerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RunTeamTriggerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SearchAgentTrigger search agent trigger API
*/

func (a *Client) SearchAgentTrigger(params *SearchAgentTriggerParams, opts ...ClientOption) (*SearchAgentTriggerOK, error) {
	if params == nil {
		params = NewSearchAgentTriggerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchAgentTrigger",
		Method:             "GET",
		PathPattern:        "/call_center/teams/triggers/agent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchAgentTriggerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SearchAgentTriggerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SearchAgentTriggerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SearchTeamTrigger search team trigger API
*/

func (a *Client) SearchTeamTrigger(params *SearchTeamTriggerParams, opts ...ClientOption) (*SearchTeamTriggerOK, error) {
	if params == nil {
		params = NewSearchTeamTriggerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchTeamTrigger",
		Method:             "GET",
		PathPattern:        "/call_center/teams/{team_id}/triggers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchTeamTriggerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SearchTeamTriggerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SearchTeamTriggerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateTeamTrigger update team trigger API
*/

func (a *Client) UpdateTeamTrigger(params *UpdateTeamTriggerParams, opts ...ClientOption) (*UpdateTeamTriggerOK, error) {
	if params == nil {
		params = NewUpdateTeamTriggerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateTeamTrigger",
		Method:             "PUT",
		PathPattern:        "/call_center/teams/{team_id}/triggers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateTeamTriggerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateTeamTriggerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateTeamTriggerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

// WithAuthInfo changes the transport on the client
func WithAuthInfo(authInfo runtime.ClientAuthInfoWriter) ClientOption {
	return func(op *runtime.ClientOperation) {
		op.AuthInfo = authInfo
	}
}
