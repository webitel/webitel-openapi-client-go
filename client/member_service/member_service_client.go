// Code generated by go-swagger; DO NOT EDIT.

package member_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/webitel-openapi-client-go/models"
)

// New creates a new member service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for member service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AttemptCallback(attemptID string, body *models.EngineAttemptCallbackRequest, opts ...ClientOption) (*AttemptCallbackOK, error)
	AttemptCallbackWithParams(params *AttemptCallbackParams, opts ...ClientOption) (*AttemptCallbackOK, error)

	AttemptResult(params *AttemptResultParams, opts ...ClientOption) (*AttemptResultOK, error)

	AttemptsRenewalResult(attemptID string, body *models.EngineAttemptRenewalResultRequest, opts ...ClientOption) (*AttemptsRenewalResultOK, error)
	AttemptsRenewalResultWithParams(params *AttemptsRenewalResultParams, opts ...ClientOption) (*AttemptsRenewalResultOK, error)

	CreateAttempt(params *CreateAttemptParams, opts ...ClientOption) (*CreateAttemptOK, error)

	CreateMember(queueID string, body *models.EngineCreateMemberRequest, opts ...ClientOption) (*CreateMemberOK, error)
	CreateMemberWithParams(params *CreateMemberParams, opts ...ClientOption) (*CreateMemberOK, error)

	CreateMemberBulk(queueID string, body *models.EngineCreateMemberBulkRequest, opts ...ClientOption) (*CreateMemberBulkOK, error)
	CreateMemberBulkWithParams(params *CreateMemberBulkParams, opts ...ClientOption) (*CreateMemberBulkOK, error)

	DeleteMember(params *DeleteMemberParams, opts ...ClientOption) (*DeleteMemberOK, error)

	DeleteMembers(queueID string, body *models.EngineDeleteMembersRequest, opts ...ClientOption) (*DeleteMembersOK, error)
	DeleteMembersWithParams(params *DeleteMembersParams, opts ...ClientOption) (*DeleteMembersOK, error)

	PatchMember(params *PatchMemberParams, opts ...ClientOption) (*PatchMemberOK, error)

	PatchMemberOne(id string, body *models.EnginePatchMemberOneRequest, opts ...ClientOption) (*PatchMemberOneOK, error)
	PatchMemberOneWithParams(params *PatchMemberOneParams, opts ...ClientOption) (*PatchMemberOneOK, error)

	ReadMember(params *ReadMemberParams, opts ...ClientOption) (*ReadMemberOK, error)

	ResetMembers(queueID string, body *models.EngineResetMembersRequest, opts ...ClientOption) (*ResetMembersOK, error)
	ResetMembersWithParams(params *ResetMembersParams, opts ...ClientOption) (*ResetMembersOK, error)

	SearchAttempts(params *SearchAttemptsParams, opts ...ClientOption) (*SearchAttemptsOK, error)

	SearchAttemptsHistory(params *SearchAttemptsHistoryParams, opts ...ClientOption) (*SearchAttemptsHistoryOK, error)

	SearchMemberAttempts(params *SearchMemberAttemptsParams, opts ...ClientOption) (*SearchMemberAttemptsOK, error)

	SearchMemberInQueue(params *SearchMemberInQueueParams, opts ...ClientOption) (*SearchMemberInQueueOK, error)

	SearchMembers(params *SearchMembersParams, opts ...ClientOption) (*SearchMembersOK, error)

	UpdateMember(params *UpdateMemberParams, opts ...ClientOption) (*UpdateMemberOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AttemptCallback attempt callback API
*/
func (a *Client) AttemptCallback(attemptID string, body *models.EngineAttemptCallbackRequest, opts ...ClientOption) (*AttemptCallbackOK, error) {
	params := NewAttemptCallbackParams().WithAttemptID(attemptID).WithBody(body)
	return a.AttemptCallbackWithParams(params, opts...)
}

func (a *Client) AttemptCallbackWithParams(params *AttemptCallbackParams, opts ...ClientOption) (*AttemptCallbackOK, error) {
	if params == nil {
		params = NewAttemptCallbackParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AttemptCallback",
		Method:             "PATCH",
		PathPattern:        "/call_center/attempts/{attempt_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AttemptCallbackReader{formats: a.formats},
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
	success, ok := result.(*AttemptCallbackOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AttemptCallbackDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AttemptResult attempt result API
*/

func (a *Client) AttemptResult(params *AttemptResultParams, opts ...ClientOption) (*AttemptResultOK, error) {
	if params == nil {
		params = NewAttemptResultParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AttemptResult",
		Method:             "PATCH",
		PathPattern:        "/call_center/queues/{queue_id}/members/{member_id}/attempts/{attempt_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AttemptResultReader{formats: a.formats},
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
	success, ok := result.(*AttemptResultOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AttemptResultDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AttemptsRenewalResult attempts renewal result API
*/
func (a *Client) AttemptsRenewalResult(attemptID string, body *models.EngineAttemptRenewalResultRequest, opts ...ClientOption) (*AttemptsRenewalResultOK, error) {
	params := NewAttemptsRenewalResultParams().WithAttemptID(attemptID).WithBody(body)
	return a.AttemptsRenewalResultWithParams(params, opts...)
}

func (a *Client) AttemptsRenewalResultWithParams(params *AttemptsRenewalResultParams, opts ...ClientOption) (*AttemptsRenewalResultOK, error) {
	if params == nil {
		params = NewAttemptsRenewalResultParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AttemptsRenewalResult",
		Method:             "PATCH",
		PathPattern:        "/call_center/queues/attempts/{attempt_id}/renewal",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AttemptsRenewalResultReader{formats: a.formats},
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
	success, ok := result.(*AttemptsRenewalResultOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AttemptsRenewalResultDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CreateAttempt offlines queue
*/

func (a *Client) CreateAttempt(params *CreateAttemptParams, opts ...ClientOption) (*CreateAttemptOK, error) {
	if params == nil {
		params = NewCreateAttemptParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateAttempt",
		Method:             "POST",
		PathPattern:        "/call_center/queues/{queue_id}/members/{member_id}/attempts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateAttemptReader{formats: a.formats},
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
	success, ok := result.(*CreateAttemptOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateAttemptDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CreateMember creates member
*/
func (a *Client) CreateMember(queueID string, body *models.EngineCreateMemberRequest, opts ...ClientOption) (*CreateMemberOK, error) {
	params := NewCreateMemberParams().WithBody(body).WithQueueID(queueID)
	return a.CreateMemberWithParams(params, opts...)
}

func (a *Client) CreateMemberWithParams(params *CreateMemberParams, opts ...ClientOption) (*CreateMemberOK, error) {
	if params == nil {
		params = NewCreateMemberParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateMember",
		Method:             "POST",
		PathPattern:        "/call_center/queues/{queue_id}/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateMemberReader{formats: a.formats},
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
	success, ok := result.(*CreateMemberOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateMemberDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CreateMemberBulk creates member
*/
func (a *Client) CreateMemberBulk(queueID string, body *models.EngineCreateMemberBulkRequest, opts ...ClientOption) (*CreateMemberBulkOK, error) {
	params := NewCreateMemberBulkParams().WithBody(body).WithQueueID(queueID)
	return a.CreateMemberBulkWithParams(params, opts...)
}

func (a *Client) CreateMemberBulkWithParams(params *CreateMemberBulkParams, opts ...ClientOption) (*CreateMemberBulkOK, error) {
	if params == nil {
		params = NewCreateMemberBulkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateMemberBulk",
		Method:             "POST",
		PathPattern:        "/call_center/queues/{queue_id}/members/bulk",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateMemberBulkReader{formats: a.formats},
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
	success, ok := result.(*CreateMemberBulkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateMemberBulkDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteMember deletes member
*/

func (a *Client) DeleteMember(params *DeleteMemberParams, opts ...ClientOption) (*DeleteMemberOK, error) {
	if params == nil {
		params = NewDeleteMemberParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteMember",
		Method:             "DELETE",
		PathPattern:        "/call_center/queues/{queue_id}/members/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteMemberReader{formats: a.formats},
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
	success, ok := result.(*DeleteMemberOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteMemberDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteMembers deletes members
*/
func (a *Client) DeleteMembers(queueID string, body *models.EngineDeleteMembersRequest, opts ...ClientOption) (*DeleteMembersOK, error) {
	params := NewDeleteMembersParams().WithBody(body).WithQueueID(queueID)
	return a.DeleteMembersWithParams(params, opts...)
}

func (a *Client) DeleteMembersWithParams(params *DeleteMembersParams, opts ...ClientOption) (*DeleteMembersOK, error) {
	if params == nil {
		params = NewDeleteMembersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteMembers",
		Method:             "DELETE",
		PathPattern:        "/call_center/queues/{queue_id}/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteMembersReader{formats: a.formats},
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
	success, ok := result.(*DeleteMembersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteMembersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchMember patches member
*/

func (a *Client) PatchMember(params *PatchMemberParams, opts ...ClientOption) (*PatchMemberOK, error) {
	if params == nil {
		params = NewPatchMemberParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchMember",
		Method:             "PATCH",
		PathPattern:        "/call_center/queues/{queue_id}/members/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchMemberReader{formats: a.formats},
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
	success, ok := result.(*PatchMemberOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchMemberDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchMemberOne patch member one API
*/
func (a *Client) PatchMemberOne(id string, body *models.EnginePatchMemberOneRequest, opts ...ClientOption) (*PatchMemberOneOK, error) {
	params := NewPatchMemberOneParams().WithBody(body).WithID(id)
	return a.PatchMemberOneWithParams(params, opts...)
}

func (a *Client) PatchMemberOneWithParams(params *PatchMemberOneParams, opts ...ClientOption) (*PatchMemberOneOK, error) {
	if params == nil {
		params = NewPatchMemberOneParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchMemberOne",
		Method:             "PATCH",
		PathPattern:        "/call_center/members/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchMemberOneReader{formats: a.formats},
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
	success, ok := result.(*PatchMemberOneOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchMemberOneDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReadMember reads queue routing
*/

func (a *Client) ReadMember(params *ReadMemberParams, opts ...ClientOption) (*ReadMemberOK, error) {
	if params == nil {
		params = NewReadMemberParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadMember",
		Method:             "GET",
		PathPattern:        "/call_center/queues/{queue_id}/members/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReadMemberReader{formats: a.formats},
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
	success, ok := result.(*ReadMemberOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadMemberDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ResetMembers resets members
*/
func (a *Client) ResetMembers(queueID string, body *models.EngineResetMembersRequest, opts ...ClientOption) (*ResetMembersOK, error) {
	params := NewResetMembersParams().WithBody(body).WithQueueID(queueID)
	return a.ResetMembersWithParams(params, opts...)
}

func (a *Client) ResetMembersWithParams(params *ResetMembersParams, opts ...ClientOption) (*ResetMembersOK, error) {
	if params == nil {
		params = NewResetMembersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ResetMembers",
		Method:             "PATCH",
		PathPattern:        "/call_center/queues/{queue_id}/members/reset",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResetMembersReader{formats: a.formats},
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
	success, ok := result.(*ResetMembersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ResetMembersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SearchAttempts searches attempts
*/

func (a *Client) SearchAttempts(params *SearchAttemptsParams, opts ...ClientOption) (*SearchAttemptsOK, error) {
	if params == nil {
		params = NewSearchAttemptsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchAttempts",
		Method:             "GET",
		PathPattern:        "/call_center/queues/attempts/active",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchAttemptsReader{formats: a.formats},
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
	success, ok := result.(*SearchAttemptsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SearchAttemptsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SearchAttemptsHistory searches member attempts
*/

func (a *Client) SearchAttemptsHistory(params *SearchAttemptsHistoryParams, opts ...ClientOption) (*SearchAttemptsHistoryOK, error) {
	if params == nil {
		params = NewSearchAttemptsHistoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchAttemptsHistory",
		Method:             "GET",
		PathPattern:        "/call_center/queues/attempts/history",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchAttemptsHistoryReader{formats: a.formats},
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
	success, ok := result.(*SearchAttemptsHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SearchAttemptsHistoryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SearchMemberAttempts searches member attempts
*/

func (a *Client) SearchMemberAttempts(params *SearchMemberAttemptsParams, opts ...ClientOption) (*SearchMemberAttemptsOK, error) {
	if params == nil {
		params = NewSearchMemberAttemptsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchMemberAttempts",
		Method:             "GET",
		PathPattern:        "/call_center/queues/{queue_id}/members/{member_id}/attempts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchMemberAttemptsReader{formats: a.formats},
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
	success, ok := result.(*SearchMemberAttemptsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SearchMemberAttemptsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SearchMemberInQueue lists of member
*/

func (a *Client) SearchMemberInQueue(params *SearchMemberInQueueParams, opts ...ClientOption) (*SearchMemberInQueueOK, error) {
	if params == nil {
		params = NewSearchMemberInQueueParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchMemberInQueue",
		Method:             "GET",
		PathPattern:        "/call_center/queues/{queue_id}/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchMemberInQueueReader{formats: a.formats},
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
	success, ok := result.(*SearchMemberInQueueOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SearchMemberInQueueDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SearchMembers searches of member
*/

func (a *Client) SearchMembers(params *SearchMembersParams, opts ...ClientOption) (*SearchMembersOK, error) {
	if params == nil {
		params = NewSearchMembersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchMembers",
		Method:             "GET",
		PathPattern:        "/call_center/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchMembersReader{formats: a.formats},
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
	success, ok := result.(*SearchMembersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SearchMembersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateMember updates member
*/

func (a *Client) UpdateMember(params *UpdateMemberParams, opts ...ClientOption) (*UpdateMemberOK, error) {
	if params == nil {
		params = NewUpdateMemberParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateMember",
		Method:             "PUT",
		PathPattern:        "/call_center/queues/{queue_id}/members/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateMemberReader{formats: a.formats},
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
	success, ok := result.(*UpdateMemberOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateMemberDefault)
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
