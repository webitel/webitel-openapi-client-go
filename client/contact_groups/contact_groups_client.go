// Code generated by go-swagger; DO NOT EDIT.

package contact_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new contact groups API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for contact groups API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ContactGroupsDeleteContactGroup(params *ContactGroupsDeleteContactGroupParams, opts ...ClientOption) (*ContactGroupsDeleteContactGroupOK, error)

	ContactGroupsDeleteContactGroups(params *ContactGroupsDeleteContactGroupsParams, opts ...ClientOption) (*ContactGroupsDeleteContactGroupsOK, error)

	ContactGroupsListContactGroups(params *ContactGroupsListContactGroupsParams, opts ...ClientOption) (*ContactGroupsListContactGroupsOK, error)

	ContactGroupsLocateContactGroup(params *ContactGroupsLocateContactGroupParams, opts ...ClientOption) (*ContactGroupsLocateContactGroupOK, error)

	ContactGroupsMergeContactGroups(params *ContactGroupsMergeContactGroupsParams, opts ...ClientOption) (*ContactGroupsMergeContactGroupsOK, error)

	ContactGroupsResetContactGroups(params *ContactGroupsResetContactGroupsParams, opts ...ClientOption) (*ContactGroupsResetContactGroupsOK, error)

	ContactGroupsUpdateContactGroup(params *ContactGroupsUpdateContactGroupParams, opts ...ClientOption) (*ContactGroupsUpdateContactGroupOK, error)

	ContactGroupsUpdateContactGroup2(params *ContactGroupsUpdateContactGroup2Params, opts ...ClientOption) (*ContactGroupsUpdateContactGroup2OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ContactGroupsDeleteContactGroup removes the contact s group association
*/

func (a *Client) ContactGroupsDeleteContactGroup(params *ContactGroupsDeleteContactGroupParams, opts ...ClientOption) (*ContactGroupsDeleteContactGroupOK, error) {
	if params == nil {
		params = NewContactGroupsDeleteContactGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContactGroups_DeleteContactGroup",
		Method:             "DELETE",
		PathPattern:        "/contacts/{contact_id}/groups/{etag}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ContactGroupsDeleteContactGroupReader{formats: a.formats},
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
	success, ok := result.(*ContactGroupsDeleteContactGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContactGroups_DeleteContactGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ContactGroupsDeleteContactGroups removes the contact s group s
*/

func (a *Client) ContactGroupsDeleteContactGroups(params *ContactGroupsDeleteContactGroupsParams, opts ...ClientOption) (*ContactGroupsDeleteContactGroupsOK, error) {
	if params == nil {
		params = NewContactGroupsDeleteContactGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContactGroups_DeleteContactGroups",
		Method:             "DELETE",
		PathPattern:        "/contacts/{contact_id}/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ContactGroupsDeleteContactGroupsReader{formats: a.formats},
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
	success, ok := result.(*ContactGroupsDeleteContactGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContactGroups_DeleteContactGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ContactGroupsListContactGroups lists of the contact s group s
*/

func (a *Client) ContactGroupsListContactGroups(params *ContactGroupsListContactGroupsParams, opts ...ClientOption) (*ContactGroupsListContactGroupsOK, error) {
	if params == nil {
		params = NewContactGroupsListContactGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContactGroups_ListContactGroups",
		Method:             "GET",
		PathPattern:        "/contacts/{contact_id}/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ContactGroupsListContactGroupsReader{formats: a.formats},
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
	success, ok := result.(*ContactGroupsListContactGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContactGroups_ListContactGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ContactGroupsLocateContactGroup locates the contact s group association
*/

func (a *Client) ContactGroupsLocateContactGroup(params *ContactGroupsLocateContactGroupParams, opts ...ClientOption) (*ContactGroupsLocateContactGroupOK, error) {
	if params == nil {
		params = NewContactGroupsLocateContactGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContactGroups_LocateContactGroup",
		Method:             "GET",
		PathPattern:        "/contacts/{contact_id}/groups/{etag}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ContactGroupsLocateContactGroupReader{formats: a.formats},
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
	success, ok := result.(*ContactGroupsLocateContactGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContactGroups_LocateContactGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ContactGroupsMergeContactGroups associates more group s with the contact
*/

func (a *Client) ContactGroupsMergeContactGroups(params *ContactGroupsMergeContactGroupsParams, opts ...ClientOption) (*ContactGroupsMergeContactGroupsOK, error) {
	if params == nil {
		params = NewContactGroupsMergeContactGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContactGroups_MergeContactGroups",
		Method:             "POST",
		PathPattern:        "/contacts/{contact_id}/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ContactGroupsMergeContactGroupsReader{formats: a.formats},
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
	success, ok := result.(*ContactGroupsMergeContactGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContactGroups_MergeContactGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ContactGroupsResetContactGroups resets all groups of the contact according to the input dataset
*/

func (a *Client) ContactGroupsResetContactGroups(params *ContactGroupsResetContactGroupsParams, opts ...ClientOption) (*ContactGroupsResetContactGroupsOK, error) {
	if params == nil {
		params = NewContactGroupsResetContactGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContactGroups_ResetContactGroups",
		Method:             "PUT",
		PathPattern:        "/contacts/{contact_id}/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ContactGroupsResetContactGroupsReader{formats: a.formats},
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
	success, ok := result.(*ContactGroupsResetContactGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContactGroups_ResetContactGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ContactGroupsUpdateContactGroup updates the contact s group details
*/

func (a *Client) ContactGroupsUpdateContactGroup(params *ContactGroupsUpdateContactGroupParams, opts ...ClientOption) (*ContactGroupsUpdateContactGroupOK, error) {
	if params == nil {
		params = NewContactGroupsUpdateContactGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContactGroups_UpdateContactGroup",
		Method:             "PUT",
		PathPattern:        "/contacts/{contact_id}/groups/{etag}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ContactGroupsUpdateContactGroupReader{formats: a.formats},
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
	success, ok := result.(*ContactGroupsUpdateContactGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContactGroups_UpdateContactGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ContactGroupsUpdateContactGroup2 updates the contact s group details
*/

func (a *Client) ContactGroupsUpdateContactGroup2(params *ContactGroupsUpdateContactGroup2Params, opts ...ClientOption) (*ContactGroupsUpdateContactGroup2OK, error) {
	if params == nil {
		params = NewContactGroupsUpdateContactGroup2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContactGroups_UpdateContactGroup2",
		Method:             "PATCH",
		PathPattern:        "/contacts/{contact_id}/groups/{etag}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ContactGroupsUpdateContactGroup2Reader{formats: a.formats},
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
	success, ok := result.(*ContactGroupsUpdateContactGroup2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContactGroups_UpdateContactGroup2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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
