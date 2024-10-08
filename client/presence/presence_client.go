// Code generated by go-swagger; DO NOT EDIT.

package presence

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/webitel-openapi-client-go/models"
)

// New creates a new presence API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for presence API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PresenceSetStatus(publishID string, body *models.APIPresenceSetStatusBody, opts ...ClientOption) (*PresenceSetStatusOK, error)
	PresenceSetStatusWithParams(params *PresenceSetStatusParams, opts ...ClientOption) (*PresenceSetStatusOK, error)

	PresenceSetStatus2(body *models.APISetStatusRequest, opts ...ClientOption) (*PresenceSetStatus2OK, error)
	PresenceSetStatus2WithParams(params *PresenceSetStatus2Params, opts ...ClientOption) (*PresenceSetStatus2OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
PresenceSetStatus sets presence notify error

(UserPresence) {
*/
func (a *Client) PresenceSetStatus(publishID string, body *models.APIPresenceSetStatusBody, opts ...ClientOption) (*PresenceSetStatusOK, error) {
	params := NewPresenceSetStatusParams().WithBody(body).WithPublishID(publishID)
	return a.PresenceSetStatusWithParams(params, opts...)
}

func (a *Client) PresenceSetStatusWithParams(params *PresenceSetStatusParams, opts ...ClientOption) (*PresenceSetStatusOK, error) {
	if params == nil {
		params = NewPresenceSetStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Presence_SetStatus",
		Method:             "PATCH",
		PathPattern:        "/users/{publish.id}/presence",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PresenceSetStatusReader{formats: a.formats},
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
	success, ok := result.(*PresenceSetStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Presence_SetStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PresenceSetStatus2 sets presence notify error

(UserPresence) {
*/
func (a *Client) PresenceSetStatus2(body *models.APISetStatusRequest, opts ...ClientOption) (*PresenceSetStatus2OK, error) {
	params := NewPresenceSetStatus2Params().WithBody(body)
	return a.PresenceSetStatus2WithParams(params, opts...)
}

func (a *Client) PresenceSetStatus2WithParams(params *PresenceSetStatus2Params, opts ...ClientOption) (*PresenceSetStatus2OK, error) {
	if params == nil {
		params = NewPresenceSetStatus2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "Presence_SetStatus2",
		Method:             "PATCH",
		PathPattern:        "/presence",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PresenceSetStatus2Reader{formats: a.formats},
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
	success, ok := result.(*PresenceSetStatus2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Presence_SetStatus2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
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
