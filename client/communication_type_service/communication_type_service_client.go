// Code generated by go-swagger; DO NOT EDIT.

package communication_type_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/webitel-openapi-client-go/models"
)

// New creates a new communication type service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for communication type service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateCommunicationType(body *models.EngineCommunicationTypeRequest, opts ...ClientOption) (*CreateCommunicationTypeOK, error)
	CreateCommunicationTypeWithParams(params *CreateCommunicationTypeParams, opts ...ClientOption) (*CreateCommunicationTypeOK, error)

	DeleteCommunicationType(params *DeleteCommunicationTypeParams, opts ...ClientOption) (*DeleteCommunicationTypeOK, error)

	PatchCommunicationType(id string, body *models.EnginePatchCommunicationTypeRequest, opts ...ClientOption) (*PatchCommunicationTypeOK, error)
	PatchCommunicationTypeWithParams(params *PatchCommunicationTypeParams, opts ...ClientOption) (*PatchCommunicationTypeOK, error)

	ReadCommunicationType(params *ReadCommunicationTypeParams, opts ...ClientOption) (*ReadCommunicationTypeOK, error)

	SearchCommunicationType(params *SearchCommunicationTypeParams, opts ...ClientOption) (*SearchCommunicationTypeOK, error)

	UpdateCommunicationType(id string, body *models.EngineUpdateCommunicationTypeRequest, opts ...ClientOption) (*UpdateCommunicationTypeOK, error)
	UpdateCommunicationTypeWithParams(params *UpdateCommunicationTypeParams, opts ...ClientOption) (*UpdateCommunicationTypeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateCommunicationType creates communication type
*/
func (a *Client) CreateCommunicationType(body *models.EngineCommunicationTypeRequest, opts ...ClientOption) (*CreateCommunicationTypeOK, error) {
	params := NewCreateCommunicationTypeParams().WithBody(body)
	return a.CreateCommunicationTypeWithParams(params, opts...)
}

func (a *Client) CreateCommunicationTypeWithParams(params *CreateCommunicationTypeParams, opts ...ClientOption) (*CreateCommunicationTypeOK, error) {
	if params == nil {
		params = NewCreateCommunicationTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateCommunicationType",
		Method:             "POST",
		PathPattern:        "/call_center/communication_type",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateCommunicationTypeReader{formats: a.formats},
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
	success, ok := result.(*CreateCommunicationTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateCommunicationTypeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteCommunicationType removes communication type
*/

func (a *Client) DeleteCommunicationType(params *DeleteCommunicationTypeParams, opts ...ClientOption) (*DeleteCommunicationTypeOK, error) {
	if params == nil {
		params = NewDeleteCommunicationTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteCommunicationType",
		Method:             "DELETE",
		PathPattern:        "/call_center/communication_type/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCommunicationTypeReader{formats: a.formats},
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
	success, ok := result.(*DeleteCommunicationTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteCommunicationTypeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchCommunicationType patch communication type API
*/
func (a *Client) PatchCommunicationType(id string, body *models.EnginePatchCommunicationTypeRequest, opts ...ClientOption) (*PatchCommunicationTypeOK, error) {
	params := NewPatchCommunicationTypeParams().WithBody(body).WithID(id)
	return a.PatchCommunicationTypeWithParams(params, opts...)
}

func (a *Client) PatchCommunicationTypeWithParams(params *PatchCommunicationTypeParams, opts ...ClientOption) (*PatchCommunicationTypeOK, error) {
	if params == nil {
		params = NewPatchCommunicationTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchCommunicationType",
		Method:             "PATCH",
		PathPattern:        "/call_center/communication_type/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchCommunicationTypeReader{formats: a.formats},
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
	success, ok := result.(*PatchCommunicationTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchCommunicationTypeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReadCommunicationType communications type item
*/

func (a *Client) ReadCommunicationType(params *ReadCommunicationTypeParams, opts ...ClientOption) (*ReadCommunicationTypeOK, error) {
	if params == nil {
		params = NewReadCommunicationTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadCommunicationType",
		Method:             "GET",
		PathPattern:        "/call_center/communication_type/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReadCommunicationTypeReader{formats: a.formats},
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
	success, ok := result.(*ReadCommunicationTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadCommunicationTypeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SearchCommunicationType lists of communication type
*/

func (a *Client) SearchCommunicationType(params *SearchCommunicationTypeParams, opts ...ClientOption) (*SearchCommunicationTypeOK, error) {
	if params == nil {
		params = NewSearchCommunicationTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchCommunicationType",
		Method:             "GET",
		PathPattern:        "/call_center/communication_type",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchCommunicationTypeReader{formats: a.formats},
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
	success, ok := result.(*SearchCommunicationTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SearchCommunicationTypeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateCommunicationType updates communication type
*/
func (a *Client) UpdateCommunicationType(id string, body *models.EngineUpdateCommunicationTypeRequest, opts ...ClientOption) (*UpdateCommunicationTypeOK, error) {
	params := NewUpdateCommunicationTypeParams().WithBody(body).WithID(id)
	return a.UpdateCommunicationTypeWithParams(params, opts...)
}

func (a *Client) UpdateCommunicationTypeWithParams(params *UpdateCommunicationTypeParams, opts ...ClientOption) (*UpdateCommunicationTypeOK, error) {
	if params == nil {
		params = NewUpdateCommunicationTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateCommunicationType",
		Method:             "PUT",
		PathPattern:        "/call_center/communication_type/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateCommunicationTypeReader{formats: a.formats},
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
	success, ok := result.(*UpdateCommunicationTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateCommunicationTypeDefault)
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
