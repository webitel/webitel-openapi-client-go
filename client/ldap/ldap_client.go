// Code generated by go-swagger; DO NOT EDIT.

package ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/webitel-openapi-client-go/models"
)

// New creates a new ldap API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ldap API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	LDAPCreateLDAPCatalog(body *models.APILDAPCatalog, opts ...ClientOption) (*LDAPCreateLDAPCatalogOK, error)
	LDAPCreateLDAPCatalogWithParams(params *LDAPCreateLDAPCatalogParams, opts ...ClientOption) (*LDAPCreateLDAPCatalogOK, error)

	LDAPCreateLDAPTemplate(catalogID string, body *models.APILDAPCreateLDAPTemplateBody, opts ...ClientOption) (*LDAPCreateLDAPTemplateOK, error)
	LDAPCreateLDAPTemplateWithParams(params *LDAPCreateLDAPTemplateParams, opts ...ClientOption) (*LDAPCreateLDAPTemplateOK, error)

	LDAPDeleteLDAPCatalog(params *LDAPDeleteLDAPCatalogParams, opts ...ClientOption) (*LDAPDeleteLDAPCatalogOK, error)

	LDAPDeleteLDAPCatalog2(body *models.APIDeleteLDAPRequest, opts ...ClientOption) (*LDAPDeleteLDAPCatalog2OK, error)
	LDAPDeleteLDAPCatalog2WithParams(params *LDAPDeleteLDAPCatalog2Params, opts ...ClientOption) (*LDAPDeleteLDAPCatalog2OK, error)

	LDAPDeleteLDAPTemplate(params *LDAPDeleteLDAPTemplateParams, opts ...ClientOption) (*LDAPDeleteLDAPTemplateOK, error)

	LDAPDeleteLDAPTemplate2(catalogID string, body *models.APILDAPDeleteLDAPTemplateBody, opts ...ClientOption) (*LDAPDeleteLDAPTemplate2OK, error)
	LDAPDeleteLDAPTemplate2WithParams(params *LDAPDeleteLDAPTemplate2Params, opts ...ClientOption) (*LDAPDeleteLDAPTemplate2OK, error)

	LDAPLDAPSearch(catalogID string, body *models.APILDAPLDAPSearchBody, opts ...ClientOption) (*LDAPLDAPSearchOK, error)
	LDAPLDAPSearchWithParams(params *LDAPLDAPSearchParams, opts ...ClientOption) (*LDAPLDAPSearchOK, error)

	LDAPLDAPSearch2(params *LDAPLDAPSearch2Params, opts ...ClientOption) (*LDAPLDAPSearch2OK, error)

	LDAPLDAPSearch3(body *models.APILDAPSearchRequest, opts ...ClientOption) (*LDAPLDAPSearch3OK, error)
	LDAPLDAPSearch3WithParams(params *LDAPLDAPSearch3Params, opts ...ClientOption) (*LDAPLDAPSearch3OK, error)

	LDAPLDAPSearch4(params *LDAPLDAPSearch4Params, opts ...ClientOption) (*LDAPLDAPSearch4OK, error)

	LDAPLocateLDAPCatalog(params *LDAPLocateLDAPCatalogParams, opts ...ClientOption) (*LDAPLocateLDAPCatalogOK, error)

	LDAPLocateLDAPTemplate(params *LDAPLocateLDAPTemplateParams, opts ...ClientOption) (*LDAPLocateLDAPTemplateOK, error)

	LDAPLocateLDAProcess(params *LDAPLocateLDAProcessParams, opts ...ClientOption) (*LDAPLocateLDAProcessOK, error)

	LDAPResyncLDAPCatalog(catalogID string, body *models.APILDAPResyncLDAPCatalogBody, opts ...ClientOption) (*LDAPResyncLDAPCatalogOK, error)
	LDAPResyncLDAPCatalogWithParams(params *LDAPResyncLDAPCatalogParams, opts ...ClientOption) (*LDAPResyncLDAPCatalogOK, error)

	LDAPSearchLDAPCatalog(params *LDAPSearchLDAPCatalogParams, opts ...ClientOption) (*LDAPSearchLDAPCatalogOK, error)

	LDAPSearchLDAPTemplate(params *LDAPSearchLDAPTemplateParams, opts ...ClientOption) (*LDAPSearchLDAPTemplateOK, error)

	LDAPSearchLDAProcess(params *LDAPSearchLDAProcessParams, opts ...ClientOption) (*LDAPSearchLDAProcessOK, error)

	LDAPUpdateLDAPCatalog(catalogID string, body *models.APILDAPUpdateLDAPCatalogBody, opts ...ClientOption) (*LDAPUpdateLDAPCatalogOK, error)
	LDAPUpdateLDAPCatalogWithParams(params *LDAPUpdateLDAPCatalogParams, opts ...ClientOption) (*LDAPUpdateLDAPCatalogOK, error)

	LDAPUpdateLDAPCatalog2(catalogID string, body *models.APILDAPUpdateLDAPCatalogBody, opts ...ClientOption) (*LDAPUpdateLDAPCatalog2OK, error)
	LDAPUpdateLDAPCatalog2WithParams(params *LDAPUpdateLDAPCatalog2Params, opts ...ClientOption) (*LDAPUpdateLDAPCatalog2OK, error)

	LDAPUpdateLDAPTemplate(params *LDAPUpdateLDAPTemplateParams, opts ...ClientOption) (*LDAPUpdateLDAPTemplateOK, error)

	LDAPUpdateLDAPTemplate2(params *LDAPUpdateLDAPTemplate2Params, opts ...ClientOption) (*LDAPUpdateLDAPTemplate2OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
LDAPCreateLDAPCatalog creates LDAP profile
*/
func (a *Client) LDAPCreateLDAPCatalog(body *models.APILDAPCatalog, opts ...ClientOption) (*LDAPCreateLDAPCatalogOK, error) {
	params := NewLDAPCreateLDAPCatalogParams().WithBody(body)
	return a.LDAPCreateLDAPCatalogWithParams(params, opts...)
}

func (a *Client) LDAPCreateLDAPCatalogWithParams(params *LDAPCreateLDAPCatalogParams, opts ...ClientOption) (*LDAPCreateLDAPCatalogOK, error) {
	if params == nil {
		params = NewLDAPCreateLDAPCatalogParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LDAP_CreateLDAPCatalog",
		Method:             "POST",
		PathPattern:        "/ldap",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LDAPCreateLDAPCatalogReader{formats: a.formats},
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
	success, ok := result.(*LDAPCreateLDAPCatalogOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LDAP_CreateLDAPCatalog: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LDAPCreateLDAPTemplate LDAP create LDAP template API
*/
func (a *Client) LDAPCreateLDAPTemplate(catalogID string, body *models.APILDAPCreateLDAPTemplateBody, opts ...ClientOption) (*LDAPCreateLDAPTemplateOK, error) {
	params := NewLDAPCreateLDAPTemplateParams().WithBody(body).WithCatalogID(catalogID)
	return a.LDAPCreateLDAPTemplateWithParams(params, opts...)
}

func (a *Client) LDAPCreateLDAPTemplateWithParams(params *LDAPCreateLDAPTemplateParams, opts ...ClientOption) (*LDAPCreateLDAPTemplateOK, error) {
	if params == nil {
		params = NewLDAPCreateLDAPTemplateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LDAP_CreateLDAPTemplate",
		Method:             "POST",
		PathPattern:        "/ldap/{catalog.id}/templates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LDAPCreateLDAPTemplateReader{formats: a.formats},
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
	success, ok := result.(*LDAPCreateLDAPTemplateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LDAP_CreateLDAPTemplate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LDAPDeleteLDAPCatalog deletes LDAP profile s
*/

func (a *Client) LDAPDeleteLDAPCatalog(params *LDAPDeleteLDAPCatalogParams, opts ...ClientOption) (*LDAPDeleteLDAPCatalogOK, error) {
	if params == nil {
		params = NewLDAPDeleteLDAPCatalogParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LDAP_DeleteLDAPCatalog",
		Method:             "DELETE",
		PathPattern:        "/ldap/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LDAPDeleteLDAPCatalogReader{formats: a.formats},
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
	success, ok := result.(*LDAPDeleteLDAPCatalogOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LDAP_DeleteLDAPCatalog: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LDAPDeleteLDAPCatalog2 deletes LDAP profile s
*/
func (a *Client) LDAPDeleteLDAPCatalog2(body *models.APIDeleteLDAPRequest, opts ...ClientOption) (*LDAPDeleteLDAPCatalog2OK, error) {
	params := NewLDAPDeleteLDAPCatalog2Params().WithBody(body)
	return a.LDAPDeleteLDAPCatalog2WithParams(params, opts...)
}

func (a *Client) LDAPDeleteLDAPCatalog2WithParams(params *LDAPDeleteLDAPCatalog2Params, opts ...ClientOption) (*LDAPDeleteLDAPCatalog2OK, error) {
	if params == nil {
		params = NewLDAPDeleteLDAPCatalog2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "LDAP_DeleteLDAPCatalog2",
		Method:             "DELETE",
		PathPattern:        "/ldap",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LDAPDeleteLDAPCatalog2Reader{formats: a.formats},
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
	success, ok := result.(*LDAPDeleteLDAPCatalog2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LDAP_DeleteLDAPCatalog2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LDAPDeleteLDAPTemplate LDAP delete LDAP template API
*/

func (a *Client) LDAPDeleteLDAPTemplate(params *LDAPDeleteLDAPTemplateParams, opts ...ClientOption) (*LDAPDeleteLDAPTemplateOK, error) {
	if params == nil {
		params = NewLDAPDeleteLDAPTemplateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LDAP_DeleteLDAPTemplate",
		Method:             "DELETE",
		PathPattern:        "/ldap/{catalog_id}/templates/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LDAPDeleteLDAPTemplateReader{formats: a.formats},
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
	success, ok := result.(*LDAPDeleteLDAPTemplateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LDAP_DeleteLDAPTemplate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LDAPDeleteLDAPTemplate2 LDAP delete LDAP template2 API
*/
func (a *Client) LDAPDeleteLDAPTemplate2(catalogID string, body *models.APILDAPDeleteLDAPTemplateBody, opts ...ClientOption) (*LDAPDeleteLDAPTemplate2OK, error) {
	params := NewLDAPDeleteLDAPTemplate2Params().WithBody(body).WithCatalogID(catalogID)
	return a.LDAPDeleteLDAPTemplate2WithParams(params, opts...)
}

func (a *Client) LDAPDeleteLDAPTemplate2WithParams(params *LDAPDeleteLDAPTemplate2Params, opts ...ClientOption) (*LDAPDeleteLDAPTemplate2OK, error) {
	if params == nil {
		params = NewLDAPDeleteLDAPTemplate2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "LDAP_DeleteLDAPTemplate2",
		Method:             "DELETE",
		PathPattern:        "/ldap/{catalog_id}/templates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LDAPDeleteLDAPTemplate2Reader{formats: a.formats},
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
	success, ok := result.(*LDAPDeleteLDAPTemplate2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LDAP_DeleteLDAPTemplate2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LDAPLDAPSearch LDAPs search performs LDAP search operation
*/
func (a *Client) LDAPLDAPSearch(catalogID string, body *models.APILDAPLDAPSearchBody, opts ...ClientOption) (*LDAPLDAPSearchOK, error) {
	params := NewLDAPLDAPSearchParams().WithBody(body).WithCatalogID(catalogID)
	return a.LDAPLDAPSearchWithParams(params, opts...)
}

func (a *Client) LDAPLDAPSearchWithParams(params *LDAPLDAPSearchParams, opts ...ClientOption) (*LDAPLDAPSearchOK, error) {
	if params == nil {
		params = NewLDAPLDAPSearchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LDAP_LDAPSearch",
		Method:             "POST",
		PathPattern:        "/ldap/{catalog_id}/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LDAPLDAPSearchReader{formats: a.formats},
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
	success, ok := result.(*LDAPLDAPSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LDAP_LDAPSearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LDAPLDAPSearch2 LDAPs search performs LDAP search operation
*/

func (a *Client) LDAPLDAPSearch2(params *LDAPLDAPSearch2Params, opts ...ClientOption) (*LDAPLDAPSearch2OK, error) {
	if params == nil {
		params = NewLDAPLDAPSearch2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "LDAP_LDAPSearch2",
		Method:             "GET",
		PathPattern:        "/ldap/{catalog_id}/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LDAPLDAPSearch2Reader{formats: a.formats},
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
	success, ok := result.(*LDAPLDAPSearch2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LDAP_LDAPSearch2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LDAPLDAPSearch3 LDAPs search performs LDAP search operation
*/
func (a *Client) LDAPLDAPSearch3(body *models.APILDAPSearchRequest, opts ...ClientOption) (*LDAPLDAPSearch3OK, error) {
	params := NewLDAPLDAPSearch3Params().WithBody(body)
	return a.LDAPLDAPSearch3WithParams(params, opts...)
}

func (a *Client) LDAPLDAPSearch3WithParams(params *LDAPLDAPSearch3Params, opts ...ClientOption) (*LDAPLDAPSearch3OK, error) {
	if params == nil {
		params = NewLDAPLDAPSearch3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "LDAP_LDAPSearch3",
		Method:             "POST",
		PathPattern:        "/ldap/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LDAPLDAPSearch3Reader{formats: a.formats},
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
	success, ok := result.(*LDAPLDAPSearch3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LDAP_LDAPSearch3: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LDAPLDAPSearch4 LDAPs search performs LDAP search operation
*/

func (a *Client) LDAPLDAPSearch4(params *LDAPLDAPSearch4Params, opts ...ClientOption) (*LDAPLDAPSearch4OK, error) {
	if params == nil {
		params = NewLDAPLDAPSearch4Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "LDAP_LDAPSearch4",
		Method:             "GET",
		PathPattern:        "/ldap/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LDAPLDAPSearch4Reader{formats: a.formats},
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
	success, ok := result.(*LDAPLDAPSearch4OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LDAP_LDAPSearch4: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LDAPLocateLDAPCatalog reads LDAP catalog profile by id
*/

func (a *Client) LDAPLocateLDAPCatalog(params *LDAPLocateLDAPCatalogParams, opts ...ClientOption) (*LDAPLocateLDAPCatalogOK, error) {
	if params == nil {
		params = NewLDAPLocateLDAPCatalogParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LDAP_LocateLDAPCatalog",
		Method:             "GET",
		PathPattern:        "/ldap/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LDAPLocateLDAPCatalogReader{formats: a.formats},
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
	success, ok := result.(*LDAPLocateLDAPCatalogOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LDAP_LocateLDAPCatalog: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LDAPLocateLDAPTemplate LDAP locate LDAP template API
*/

func (a *Client) LDAPLocateLDAPTemplate(params *LDAPLocateLDAPTemplateParams, opts ...ClientOption) (*LDAPLocateLDAPTemplateOK, error) {
	if params == nil {
		params = NewLDAPLocateLDAPTemplateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LDAP_LocateLDAPTemplate",
		Method:             "GET",
		PathPattern:        "/ldap/{catalog_id}/templates/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LDAPLocateLDAPTemplateReader{formats: a.formats},
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
	success, ok := result.(*LDAPLocateLDAPTemplateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LDAP_LocateLDAPTemplate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LDAPLocateLDAProcess resyncs LDAP response
*/

func (a *Client) LDAPLocateLDAProcess(params *LDAPLocateLDAProcessParams, opts ...ClientOption) (*LDAPLocateLDAProcessOK, error) {
	if params == nil {
		params = NewLDAPLocateLDAProcessParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LDAP_LocateLDAProcess",
		Method:             "GET",
		PathPattern:        "/ldap/{catalog_id}/rsync/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LDAPLocateLDAProcessReader{formats: a.formats},
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
	success, ok := result.(*LDAPLocateLDAProcessOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LDAP_LocateLDAProcess: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LDAPResyncLDAPCatalog resyncs LDAP directory catalog changes

(ResyncLDAPResponse) {
*/
func (a *Client) LDAPResyncLDAPCatalog(catalogID string, body *models.APILDAPResyncLDAPCatalogBody, opts ...ClientOption) (*LDAPResyncLDAPCatalogOK, error) {
	params := NewLDAPResyncLDAPCatalogParams().WithBody(body).WithCatalogID(catalogID)
	return a.LDAPResyncLDAPCatalogWithParams(params, opts...)
}

func (a *Client) LDAPResyncLDAPCatalogWithParams(params *LDAPResyncLDAPCatalogParams, opts ...ClientOption) (*LDAPResyncLDAPCatalogOK, error) {
	if params == nil {
		params = NewLDAPResyncLDAPCatalogParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LDAP_ResyncLDAPCatalog",
		Method:             "POST",
		PathPattern:        "/ldap/{catalog_id}/rsync",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LDAPResyncLDAPCatalogReader{formats: a.formats},
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
	success, ok := result.(*LDAPResyncLDAPCatalogOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LDAP_ResyncLDAPCatalog: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LDAPSearchLDAPCatalog searches for LDAP directories
*/

func (a *Client) LDAPSearchLDAPCatalog(params *LDAPSearchLDAPCatalogParams, opts ...ClientOption) (*LDAPSearchLDAPCatalogOK, error) {
	if params == nil {
		params = NewLDAPSearchLDAPCatalogParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LDAP_SearchLDAPCatalog",
		Method:             "GET",
		PathPattern:        "/ldap",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LDAPSearchLDAPCatalogReader{formats: a.formats},
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
	success, ok := result.(*LDAPSearchLDAPCatalogOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LDAP_SearchLDAPCatalog: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LDAPSearchLDAPTemplate LDAP search LDAP template API
*/

func (a *Client) LDAPSearchLDAPTemplate(params *LDAPSearchLDAPTemplateParams, opts ...ClientOption) (*LDAPSearchLDAPTemplateOK, error) {
	if params == nil {
		params = NewLDAPSearchLDAPTemplateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LDAP_SearchLDAPTemplate",
		Method:             "GET",
		PathPattern:        "/ldap/{catalog_id}/templates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LDAPSearchLDAPTemplateReader{formats: a.formats},
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
	success, ok := result.(*LDAPSearchLDAPTemplateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LDAP_SearchLDAPTemplate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LDAPSearchLDAProcess resyncs LDAP response
*/

func (a *Client) LDAPSearchLDAProcess(params *LDAPSearchLDAProcessParams, opts ...ClientOption) (*LDAPSearchLDAProcessOK, error) {
	if params == nil {
		params = NewLDAPSearchLDAProcessParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LDAP_SearchLDAProcess",
		Method:             "GET",
		PathPattern:        "/ldap/{catalog_id}/rsync",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LDAPSearchLDAProcessReader{formats: a.formats},
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
	success, ok := result.(*LDAPSearchLDAProcessOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LDAP_SearchLDAProcess: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LDAPUpdateLDAPCatalog updates LDAP profile details
*/
func (a *Client) LDAPUpdateLDAPCatalog(catalogID string, body *models.APILDAPUpdateLDAPCatalogBody, opts ...ClientOption) (*LDAPUpdateLDAPCatalogOK, error) {
	params := NewLDAPUpdateLDAPCatalogParams().WithBody(body).WithCatalogID(catalogID)
	return a.LDAPUpdateLDAPCatalogWithParams(params, opts...)
}

func (a *Client) LDAPUpdateLDAPCatalogWithParams(params *LDAPUpdateLDAPCatalogParams, opts ...ClientOption) (*LDAPUpdateLDAPCatalogOK, error) {
	if params == nil {
		params = NewLDAPUpdateLDAPCatalogParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LDAP_UpdateLDAPCatalog",
		Method:             "PUT",
		PathPattern:        "/ldap/{catalog.id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LDAPUpdateLDAPCatalogReader{formats: a.formats},
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
	success, ok := result.(*LDAPUpdateLDAPCatalogOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LDAP_UpdateLDAPCatalog: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LDAPUpdateLDAPCatalog2 updates LDAP profile details
*/
func (a *Client) LDAPUpdateLDAPCatalog2(catalogID string, body *models.APILDAPUpdateLDAPCatalogBody, opts ...ClientOption) (*LDAPUpdateLDAPCatalog2OK, error) {
	params := NewLDAPUpdateLDAPCatalog2Params().WithBody(body).WithCatalogID(catalogID)
	return a.LDAPUpdateLDAPCatalog2WithParams(params, opts...)
}

func (a *Client) LDAPUpdateLDAPCatalog2WithParams(params *LDAPUpdateLDAPCatalog2Params, opts ...ClientOption) (*LDAPUpdateLDAPCatalog2OK, error) {
	if params == nil {
		params = NewLDAPUpdateLDAPCatalog2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "LDAP_UpdateLDAPCatalog2",
		Method:             "PATCH",
		PathPattern:        "/ldap/{catalog.id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LDAPUpdateLDAPCatalog2Reader{formats: a.formats},
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
	success, ok := result.(*LDAPUpdateLDAPCatalog2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LDAP_UpdateLDAPCatalog2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LDAPUpdateLDAPTemplate LDAP update LDAP template API
*/

func (a *Client) LDAPUpdateLDAPTemplate(params *LDAPUpdateLDAPTemplateParams, opts ...ClientOption) (*LDAPUpdateLDAPTemplateOK, error) {
	if params == nil {
		params = NewLDAPUpdateLDAPTemplateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LDAP_UpdateLDAPTemplate",
		Method:             "PATCH",
		PathPattern:        "/ldap/{template.catalog.id}/templates/{template.id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LDAPUpdateLDAPTemplateReader{formats: a.formats},
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
	success, ok := result.(*LDAPUpdateLDAPTemplateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LDAP_UpdateLDAPTemplate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LDAPUpdateLDAPTemplate2 LDAP update LDAP template2 API
*/

func (a *Client) LDAPUpdateLDAPTemplate2(params *LDAPUpdateLDAPTemplate2Params, opts ...ClientOption) (*LDAPUpdateLDAPTemplate2OK, error) {
	if params == nil {
		params = NewLDAPUpdateLDAPTemplate2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "LDAP_UpdateLDAPTemplate2",
		Method:             "PUT",
		PathPattern:        "/ldap/{template.catalog.id}/templates/{template.id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LDAPUpdateLDAPTemplate2Reader{formats: a.formats},
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
	success, ok := result.(*LDAPUpdateLDAPTemplate2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LDAP_UpdateLDAPTemplate2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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
