// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/webitel-openapi-client-go/models"
)

// New creates a new groups API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for groups API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GroupsCreateGroup(body *models.WebitelContactsCreateGroupRequest, opts ...ClientOption) (*GroupsCreateGroupOK, error)
	GroupsCreateGroupWithParams(params *GroupsCreateGroupParams, opts ...ClientOption) (*GroupsCreateGroupOK, error)

	GroupsDeleteGroup(id string, opts ...ClientOption) (*GroupsDeleteGroupOK, error)
	GroupsDeleteGroupWithParams(params *GroupsDeleteGroupParams, opts ...ClientOption) (*GroupsDeleteGroupOK, error)

	GroupsListGroups(params *GroupsListGroupsParams, opts ...ClientOption) (*GroupsListGroupsOK, error)

	GroupsLocateGroup(params *GroupsLocateGroupParams, opts ...ClientOption) (*GroupsLocateGroupOK, error)

	GroupsUpdateGroup(id string, body *models.WebitelContactsGroupsUpdateGroupBody, opts ...ClientOption) (*GroupsUpdateGroupOK, error)
	GroupsUpdateGroupWithParams(params *GroupsUpdateGroupParams, opts ...ClientOption) (*GroupsUpdateGroupOK, error)

	GroupsUpdateGroup2(id string, body *models.WebitelContactsGroupsUpdateGroupBody, opts ...ClientOption) (*GroupsUpdateGroup2OK, error)
	GroupsUpdateGroup2WithParams(params *GroupsUpdateGroup2Params, opts ...ClientOption) (*GroupsUpdateGroup2OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GroupsCreateGroup creates a new group
*/
func (a *Client) GroupsCreateGroup(body *models.WebitelContactsCreateGroupRequest, opts ...ClientOption) (*GroupsCreateGroupOK, error) {
	params := NewGroupsCreateGroupParams().WithBody(body)
	return a.GroupsCreateGroupWithParams(params, opts...)
}

func (a *Client) GroupsCreateGroupWithParams(params *GroupsCreateGroupParams, opts ...ClientOption) (*GroupsCreateGroupOK, error) {
	if params == nil {
		params = NewGroupsCreateGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Groups_CreateGroup",
		Method:             "POST",
		PathPattern:        "/contacts/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GroupsCreateGroupReader{formats: a.formats},
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
	success, ok := result.(*GroupsCreateGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Groups_CreateGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GroupsDeleteGroup deletes a group
*/
func (a *Client) GroupsDeleteGroup(id string, opts ...ClientOption) (*GroupsDeleteGroupOK, error) {
	params := NewGroupsDeleteGroupParams().WithID(id)
	return a.GroupsDeleteGroupWithParams(params, opts...)
}

func (a *Client) GroupsDeleteGroupWithParams(params *GroupsDeleteGroupParams, opts ...ClientOption) (*GroupsDeleteGroupOK, error) {
	if params == nil {
		params = NewGroupsDeleteGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Groups_DeleteGroup",
		Method:             "DELETE",
		PathPattern:        "/contacts/groups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GroupsDeleteGroupReader{formats: a.formats},
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
	success, ok := result.(*GroupsDeleteGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Groups_DeleteGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GroupsListGroups retrieves a list of groups or search groups
*/

func (a *Client) GroupsListGroups(params *GroupsListGroupsParams, opts ...ClientOption) (*GroupsListGroupsOK, error) {
	if params == nil {
		params = NewGroupsListGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Groups_ListGroups",
		Method:             "GET",
		PathPattern:        "/contacts/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GroupsListGroupsReader{formats: a.formats},
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
	success, ok := result.(*GroupsListGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Groups_ListGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GroupsLocateGroup locates a group by ID
*/

func (a *Client) GroupsLocateGroup(params *GroupsLocateGroupParams, opts ...ClientOption) (*GroupsLocateGroupOK, error) {
	if params == nil {
		params = NewGroupsLocateGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Groups_LocateGroup",
		Method:             "GET",
		PathPattern:        "/contacts/groups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GroupsLocateGroupReader{formats: a.formats},
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
	success, ok := result.(*GroupsLocateGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Groups_LocateGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GroupsUpdateGroup updates an existing group
*/
func (a *Client) GroupsUpdateGroup(id string, body *models.WebitelContactsGroupsUpdateGroupBody, opts ...ClientOption) (*GroupsUpdateGroupOK, error) {
	params := NewGroupsUpdateGroupParams().WithBody(body).WithID(id)
	return a.GroupsUpdateGroupWithParams(params, opts...)
}

func (a *Client) GroupsUpdateGroupWithParams(params *GroupsUpdateGroupParams, opts ...ClientOption) (*GroupsUpdateGroupOK, error) {
	if params == nil {
		params = NewGroupsUpdateGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Groups_UpdateGroup",
		Method:             "PUT",
		PathPattern:        "/contacts/groups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GroupsUpdateGroupReader{formats: a.formats},
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
	success, ok := result.(*GroupsUpdateGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Groups_UpdateGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GroupsUpdateGroup2 updates an existing group
*/
func (a *Client) GroupsUpdateGroup2(id string, body *models.WebitelContactsGroupsUpdateGroupBody, opts ...ClientOption) (*GroupsUpdateGroup2OK, error) {
	params := NewGroupsUpdateGroup2Params().WithBody(body).WithID(id)
	return a.GroupsUpdateGroup2WithParams(params, opts...)
}

func (a *Client) GroupsUpdateGroup2WithParams(params *GroupsUpdateGroup2Params, opts ...ClientOption) (*GroupsUpdateGroup2OK, error) {
	if params == nil {
		params = NewGroupsUpdateGroup2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "Groups_UpdateGroup2",
		Method:             "PATCH",
		PathPattern:        "/contacts/groups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GroupsUpdateGroup2Reader{formats: a.formats},
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
	success, ok := result.(*GroupsUpdateGroup2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Groups_UpdateGroup2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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
