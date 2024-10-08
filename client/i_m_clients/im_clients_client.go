// Code generated by go-swagger; DO NOT EDIT.

package i_m_clients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new i m clients API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for i m clients API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	IMClientsDeleteIMClient(id string, contactID string, opts ...ClientOption) (*IMClientsDeleteIMClientOK, error)
	IMClientsDeleteIMClientWithParams(params *IMClientsDeleteIMClientParams, opts ...ClientOption) (*IMClientsDeleteIMClientOK, error)

	IMClientsListIMClients(params *IMClientsListIMClientsParams, opts ...ClientOption) (*IMClientsListIMClientsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
IMClientsDeleteIMClient i m clients delete i m client API
*/
func (a *Client) IMClientsDeleteIMClient(id string, contactID string, opts ...ClientOption) (*IMClientsDeleteIMClientOK, error) {
	params := NewIMClientsDeleteIMClientParams().WithContactID(contactID).WithID(id)
	return a.IMClientsDeleteIMClientWithParams(params, opts...)
}

func (a *Client) IMClientsDeleteIMClientWithParams(params *IMClientsDeleteIMClientParams, opts ...ClientOption) (*IMClientsDeleteIMClientOK, error) {
	if params == nil {
		params = NewIMClientsDeleteIMClientParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "IMClients_DeleteIMClient",
		Method:             "DELETE",
		PathPattern:        "/contacts/{contact_id}/imclients/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IMClientsDeleteIMClientReader{formats: a.formats},
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
	success, ok := result.(*IMClientsDeleteIMClientOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for IMClients_DeleteIMClient: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
IMClientsListIMClients searches i m client links
*/

func (a *Client) IMClientsListIMClients(params *IMClientsListIMClientsParams, opts ...ClientOption) (*IMClientsListIMClientsOK, error) {
	if params == nil {
		params = NewIMClientsListIMClientsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "IMClients_ListIMClients",
		Method:             "GET",
		PathPattern:        "/contacts/{contact_id}/imclients",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IMClientsListIMClientsReader{formats: a.formats},
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
	success, ok := result.(*IMClientsListIMClientsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for IMClients_ListIMClients: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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
