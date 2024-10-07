// Code generated by go-swagger; DO NOT EDIT.

package agent_team_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/webitel-openapi-client-go/models"
)

// New creates a new agent team service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for agent team service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateAgentTeam(body *models.EngineCreateAgentTeamRequest, opts ...ClientOption) (*CreateAgentTeamOK, error)
	CreateAgentTeamWithParams(params *CreateAgentTeamParams, opts ...ClientOption) (*CreateAgentTeamOK, error)

	DeleteAgentTeam(params *DeleteAgentTeamParams, opts ...ClientOption) (*DeleteAgentTeamOK, error)

	ReadAgentTeam(params *ReadAgentTeamParams, opts ...ClientOption) (*ReadAgentTeamOK, error)

	SearchAgentTeam(params *SearchAgentTeamParams, opts ...ClientOption) (*SearchAgentTeamOK, error)

	UpdateAgentTeam(id string, body *models.EngineUpdateAgentTeamRequest, opts ...ClientOption) (*UpdateAgentTeamOK, error)
	UpdateAgentTeamWithParams(params *UpdateAgentTeamParams, opts ...ClientOption) (*UpdateAgentTeamOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateAgentTeam creates agent team
*/
func (a *Client) CreateAgentTeam(body *models.EngineCreateAgentTeamRequest, opts ...ClientOption) (*CreateAgentTeamOK, error) {
	params := NewCreateAgentTeamParams().WithBody(body)
	return a.CreateAgentTeamWithParams(params, opts...)
}

func (a *Client) CreateAgentTeamWithParams(params *CreateAgentTeamParams, opts ...ClientOption) (*CreateAgentTeamOK, error) {
	if params == nil {
		params = NewCreateAgentTeamParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateAgentTeam",
		Method:             "POST",
		PathPattern:        "/call_center/teams",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateAgentTeamReader{formats: a.formats},
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
	success, ok := result.(*CreateAgentTeamOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateAgentTeamDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteAgentTeam removes agent team
*/

func (a *Client) DeleteAgentTeam(params *DeleteAgentTeamParams, opts ...ClientOption) (*DeleteAgentTeamOK, error) {
	if params == nil {
		params = NewDeleteAgentTeamParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteAgentTeam",
		Method:             "DELETE",
		PathPattern:        "/call_center/teams/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAgentTeamReader{formats: a.formats},
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
	success, ok := result.(*DeleteAgentTeamOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteAgentTeamDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReadAgentTeam agents team item
*/

func (a *Client) ReadAgentTeam(params *ReadAgentTeamParams, opts ...ClientOption) (*ReadAgentTeamOK, error) {
	if params == nil {
		params = NewReadAgentTeamParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadAgentTeam",
		Method:             "GET",
		PathPattern:        "/call_center/teams/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReadAgentTeamReader{formats: a.formats},
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
	success, ok := result.(*ReadAgentTeamOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadAgentTeamDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SearchAgentTeam lists of agent team
*/

func (a *Client) SearchAgentTeam(params *SearchAgentTeamParams, opts ...ClientOption) (*SearchAgentTeamOK, error) {
	if params == nil {
		params = NewSearchAgentTeamParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchAgentTeam",
		Method:             "GET",
		PathPattern:        "/call_center/teams",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchAgentTeamReader{formats: a.formats},
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
	success, ok := result.(*SearchAgentTeamOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SearchAgentTeamDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateAgentTeam updates agent team
*/
func (a *Client) UpdateAgentTeam(id string, body *models.EngineUpdateAgentTeamRequest, opts ...ClientOption) (*UpdateAgentTeamOK, error) {
	params := NewUpdateAgentTeamParams().WithBody(body).WithID(id)
	return a.UpdateAgentTeamWithParams(params, opts...)
}

func (a *Client) UpdateAgentTeamWithParams(params *UpdateAgentTeamParams, opts ...ClientOption) (*UpdateAgentTeamOK, error) {
	if params == nil {
		params = NewUpdateAgentTeamParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateAgentTeam",
		Method:             "PUT",
		PathPattern:        "/call_center/teams/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAgentTeamReader{formats: a.formats},
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
	success, ok := result.(*UpdateAgentTeamOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateAgentTeamDefault)
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
