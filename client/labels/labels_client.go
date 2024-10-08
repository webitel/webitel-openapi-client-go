// Code generated by go-swagger; DO NOT EDIT.

package labels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new labels API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for labels API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	LabelsDeleteLabels(params *LabelsDeleteLabelsParams, opts ...ClientOption) (*LabelsDeleteLabelsOK, error)

	LabelsGetLabels(params *LabelsGetLabelsParams, opts ...ClientOption) (*LabelsGetLabelsOK, error)

	LabelsListLabels(params *LabelsListLabelsParams, opts ...ClientOption) (*LabelsListLabelsOK, error)

	LabelsMergeLabels(params *LabelsMergeLabelsParams, opts ...ClientOption) (*LabelsMergeLabelsOK, error)

	LabelsResetLabels(params *LabelsResetLabelsParams, opts ...ClientOption) (*LabelsResetLabelsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
LabelsDeleteLabels removes contact labels associations
*/

func (a *Client) LabelsDeleteLabels(params *LabelsDeleteLabelsParams, opts ...ClientOption) (*LabelsDeleteLabelsOK, error) {
	if params == nil {
		params = NewLabelsDeleteLabelsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Labels_DeleteLabels",
		Method:             "DELETE",
		PathPattern:        "/contacts/{contact_id}/labels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LabelsDeleteLabelsReader{formats: a.formats},
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
	success, ok := result.(*LabelsDeleteLabelsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Labels_DeleteLabels: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LabelsGetLabels searches for contacts engaged label s
*/

func (a *Client) LabelsGetLabels(params *LabelsGetLabelsParams, opts ...ClientOption) (*LabelsGetLabelsOK, error) {
	if params == nil {
		params = NewLabelsGetLabelsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Labels_GetLabels",
		Method:             "GET",
		PathPattern:        "/contacts/labels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LabelsGetLabelsReader{formats: a.formats},
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
	success, ok := result.(*LabelsGetLabelsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Labels_GetLabels: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LabelsListLabels locates the contact s associated label s
*/

func (a *Client) LabelsListLabels(params *LabelsListLabelsParams, opts ...ClientOption) (*LabelsListLabelsOK, error) {
	if params == nil {
		params = NewLabelsListLabelsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Labels_ListLabels",
		Method:             "GET",
		PathPattern:        "/contacts/{contact_id}/labels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LabelsListLabelsReader{formats: a.formats},
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
	success, ok := result.(*LabelsListLabelsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Labels_ListLabels: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LabelsMergeLabels associates n e w labels to the contact
*/

func (a *Client) LabelsMergeLabels(params *LabelsMergeLabelsParams, opts ...ClientOption) (*LabelsMergeLabelsOK, error) {
	if params == nil {
		params = NewLabelsMergeLabelsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Labels_MergeLabels",
		Method:             "POST",
		PathPattern:        "/contacts/{contact_id}/labels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LabelsMergeLabelsReader{formats: a.formats},
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
	success, ok := result.(*LabelsMergeLabelsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Labels_MergeLabels: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LabelsResetLabels resets labels to fit the specified final set
*/

func (a *Client) LabelsResetLabels(params *LabelsResetLabelsParams, opts ...ClientOption) (*LabelsResetLabelsOK, error) {
	if params == nil {
		params = NewLabelsResetLabelsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Labels_ResetLabels",
		Method:             "PUT",
		PathPattern:        "/contacts/{contact_id}/labels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LabelsResetLabelsReader{formats: a.formats},
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
	success, ok := result.(*LabelsResetLabelsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Labels_ResetLabels: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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
