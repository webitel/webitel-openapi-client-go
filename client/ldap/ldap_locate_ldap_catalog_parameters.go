// Code generated by go-swagger; DO NOT EDIT.

package ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewLDAPLocateLDAPCatalogParams creates a new LDAPLocateLDAPCatalogParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLDAPLocateLDAPCatalogParams() *LDAPLocateLDAPCatalogParams {
	return &LDAPLocateLDAPCatalogParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLDAPLocateLDAPCatalogParamsWithTimeout creates a new LDAPLocateLDAPCatalogParams object
// with the ability to set a timeout on a request.
func NewLDAPLocateLDAPCatalogParamsWithTimeout(timeout time.Duration) *LDAPLocateLDAPCatalogParams {
	return &LDAPLocateLDAPCatalogParams{
		timeout: timeout,
	}
}

// NewLDAPLocateLDAPCatalogParamsWithContext creates a new LDAPLocateLDAPCatalogParams object
// with the ability to set a context for a request.
func NewLDAPLocateLDAPCatalogParamsWithContext(ctx context.Context) *LDAPLocateLDAPCatalogParams {
	return &LDAPLocateLDAPCatalogParams{
		Context: ctx,
	}
}

// NewLDAPLocateLDAPCatalogParamsWithHTTPClient creates a new LDAPLocateLDAPCatalogParams object
// with the ability to set a custom HTTPClient for a request.
func NewLDAPLocateLDAPCatalogParamsWithHTTPClient(client *http.Client) *LDAPLocateLDAPCatalogParams {
	return &LDAPLocateLDAPCatalogParams{
		HTTPClient: client,
	}
}

/*
LDAPLocateLDAPCatalogParams contains all the parameters to send to the API endpoint

	for the LDAP locate LDAP catalog operation.

	Typically these are written to a http.Request.
*/
type LDAPLocateLDAPCatalogParams struct {

	/* Access.

	   [M]andatory[A]ccess[C]ontrol: with access mode (action) granted!
	*/
	Access *string

	/* Fields.

	   attributes list
	*/
	Fields []string

	/* ID.

	     ----- Search Basic Filters ---------------------------

	selection: by unique identifier
	*/
	ID []string

	/* Name.

	   case-ignore substring match: ILIKE '*' - any; '?' - one
	*/
	Name *string

	/* Page.

	     ----- Select Options -------------------------

	default: 1

	     Format: int32
	*/
	Page *int32

	/* Q.

	   term-of-search: lookup[name]
	*/
	Q *string

	/* Size.

	   default: 16

	   Format: int32
	*/
	Size *int32

	/* Sort.

	   e.g.: "updated_at" - ASC; "!updated_at" - DESC;
	*/
	Sort []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the LDAP locate LDAP catalog params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LDAPLocateLDAPCatalogParams) WithDefaults() *LDAPLocateLDAPCatalogParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the LDAP locate LDAP catalog params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LDAPLocateLDAPCatalogParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the LDAP locate LDAP catalog params
func (o *LDAPLocateLDAPCatalogParams) WithTimeout(timeout time.Duration) *LDAPLocateLDAPCatalogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the LDAP locate LDAP catalog params
func (o *LDAPLocateLDAPCatalogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the LDAP locate LDAP catalog params
func (o *LDAPLocateLDAPCatalogParams) WithContext(ctx context.Context) *LDAPLocateLDAPCatalogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the LDAP locate LDAP catalog params
func (o *LDAPLocateLDAPCatalogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the LDAP locate LDAP catalog params
func (o *LDAPLocateLDAPCatalogParams) WithHTTPClient(client *http.Client) *LDAPLocateLDAPCatalogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the LDAP locate LDAP catalog params
func (o *LDAPLocateLDAPCatalogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccess adds the access to the LDAP locate LDAP catalog params
func (o *LDAPLocateLDAPCatalogParams) WithAccess(access *string) *LDAPLocateLDAPCatalogParams {
	o.SetAccess(access)
	return o
}

// SetAccess adds the access to the LDAP locate LDAP catalog params
func (o *LDAPLocateLDAPCatalogParams) SetAccess(access *string) {
	o.Access = access
}

// WithFields adds the fields to the LDAP locate LDAP catalog params
func (o *LDAPLocateLDAPCatalogParams) WithFields(fields []string) *LDAPLocateLDAPCatalogParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the LDAP locate LDAP catalog params
func (o *LDAPLocateLDAPCatalogParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithID adds the id to the LDAP locate LDAP catalog params
func (o *LDAPLocateLDAPCatalogParams) WithID(id []string) *LDAPLocateLDAPCatalogParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the LDAP locate LDAP catalog params
func (o *LDAPLocateLDAPCatalogParams) SetID(id []string) {
	o.ID = id
}

// WithName adds the name to the LDAP locate LDAP catalog params
func (o *LDAPLocateLDAPCatalogParams) WithName(name *string) *LDAPLocateLDAPCatalogParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the LDAP locate LDAP catalog params
func (o *LDAPLocateLDAPCatalogParams) SetName(name *string) {
	o.Name = name
}

// WithPage adds the page to the LDAP locate LDAP catalog params
func (o *LDAPLocateLDAPCatalogParams) WithPage(page *int32) *LDAPLocateLDAPCatalogParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the LDAP locate LDAP catalog params
func (o *LDAPLocateLDAPCatalogParams) SetPage(page *int32) {
	o.Page = page
}

// WithQ adds the q to the LDAP locate LDAP catalog params
func (o *LDAPLocateLDAPCatalogParams) WithQ(q *string) *LDAPLocateLDAPCatalogParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the LDAP locate LDAP catalog params
func (o *LDAPLocateLDAPCatalogParams) SetQ(q *string) {
	o.Q = q
}

// WithSize adds the size to the LDAP locate LDAP catalog params
func (o *LDAPLocateLDAPCatalogParams) WithSize(size *int32) *LDAPLocateLDAPCatalogParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the LDAP locate LDAP catalog params
func (o *LDAPLocateLDAPCatalogParams) SetSize(size *int32) {
	o.Size = size
}

// WithSort adds the sort to the LDAP locate LDAP catalog params
func (o *LDAPLocateLDAPCatalogParams) WithSort(sort []string) *LDAPLocateLDAPCatalogParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the LDAP locate LDAP catalog params
func (o *LDAPLocateLDAPCatalogParams) SetSort(sort []string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *LDAPLocateLDAPCatalogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Access != nil {

		// query param access
		var qrAccess string

		if o.Access != nil {
			qrAccess = *o.Access
		}
		qAccess := qrAccess
		if qAccess != "" {

			if err := r.SetQueryParam("access", qAccess); err != nil {
				return err
			}
		}
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.ID != nil {

		// binding items for id
		joinedID := o.bindParamID(reg)

		// path array param id
		// SetPathParam does not support variadic arguments, since we used JoinByFormat
		// we can send the first item in the array as it's all the items of the previous
		// array joined together
		if len(joinedID) > 0 {
			if err := r.SetPathParam("id", joinedID[0]); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.Size != nil {

		// query param size
		var qrSize int32

		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt32(qrSize)
		if qSize != "" {

			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// binding items for sort
		joinedSort := o.bindParamSort(reg)

		// query array param sort
		if err := r.SetQueryParam("sort", joinedSort...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamLDAPLocateLDAPCatalog binds the parameter fields
func (o *LDAPLocateLDAPCatalogParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "multi"
	fieldsIS := swag.JoinByFormat(fieldsIC, "multi")

	return fieldsIS
}

// bindParamLDAPLocateLDAPCatalog binds the parameter id
func (o *LDAPLocateLDAPCatalogParams) bindParamID(formats strfmt.Registry) []string {
	iDIR := o.ID

	var iDIC []string
	for _, iDIIR := range iDIR { // explode []string

		iDIIV := iDIIR // string as string
		iDIC = append(iDIC, iDIIV)
	}

	// items.CollectionFormat: "csv"
	iDIS := swag.JoinByFormat(iDIC, "csv")

	return iDIS
}

// bindParamLDAPLocateLDAPCatalog binds the parameter sort
func (o *LDAPLocateLDAPCatalogParams) bindParamSort(formats strfmt.Registry) []string {
	sortIR := o.Sort

	var sortIC []string
	for _, sortIIR := range sortIR { // explode []string

		sortIIV := sortIIR // string as string
		sortIC = append(sortIC, sortIIV)
	}

	// items.CollectionFormat: "multi"
	sortIS := swag.JoinByFormat(sortIC, "multi")

	return sortIS
}
