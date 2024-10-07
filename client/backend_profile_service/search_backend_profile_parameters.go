// Code generated by go-swagger; DO NOT EDIT.

package backend_profile_service

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

// NewSearchBackendProfileParams creates a new SearchBackendProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchBackendProfileParams() *SearchBackendProfileParams {
	return &SearchBackendProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchBackendProfileParamsWithTimeout creates a new SearchBackendProfileParams object
// with the ability to set a timeout on a request.
func NewSearchBackendProfileParamsWithTimeout(timeout time.Duration) *SearchBackendProfileParams {
	return &SearchBackendProfileParams{
		timeout: timeout,
	}
}

// NewSearchBackendProfileParamsWithContext creates a new SearchBackendProfileParams object
// with the ability to set a context for a request.
func NewSearchBackendProfileParamsWithContext(ctx context.Context) *SearchBackendProfileParams {
	return &SearchBackendProfileParams{
		Context: ctx,
	}
}

// NewSearchBackendProfileParamsWithHTTPClient creates a new SearchBackendProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchBackendProfileParamsWithHTTPClient(client *http.Client) *SearchBackendProfileParams {
	return &SearchBackendProfileParams{
		HTTPClient: client,
	}
}

/*
SearchBackendProfileParams contains all the parameters to send to the API endpoint

	for the search backend profile operation.

	Typically these are written to a http.Request.
*/
type SearchBackendProfileParams struct {

	// Fields.
	Fields []string

	// ID.
	ID []int64

	// Page.
	//
	// Format: int32
	Page *int32

	// Q.
	Q *string

	// Size.
	//
	// Format: int32
	Size *int32

	// Sort.
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search backend profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchBackendProfileParams) WithDefaults() *SearchBackendProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search backend profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchBackendProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search backend profile params
func (o *SearchBackendProfileParams) WithTimeout(timeout time.Duration) *SearchBackendProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search backend profile params
func (o *SearchBackendProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search backend profile params
func (o *SearchBackendProfileParams) WithContext(ctx context.Context) *SearchBackendProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search backend profile params
func (o *SearchBackendProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search backend profile params
func (o *SearchBackendProfileParams) WithHTTPClient(client *http.Client) *SearchBackendProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search backend profile params
func (o *SearchBackendProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the search backend profile params
func (o *SearchBackendProfileParams) WithFields(fields []string) *SearchBackendProfileParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the search backend profile params
func (o *SearchBackendProfileParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithID adds the id to the search backend profile params
func (o *SearchBackendProfileParams) WithID(id []int64) *SearchBackendProfileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the search backend profile params
func (o *SearchBackendProfileParams) SetID(id []int64) {
	o.ID = id
}

// WithPage adds the page to the search backend profile params
func (o *SearchBackendProfileParams) WithPage(page *int32) *SearchBackendProfileParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the search backend profile params
func (o *SearchBackendProfileParams) SetPage(page *int32) {
	o.Page = page
}

// WithQ adds the q to the search backend profile params
func (o *SearchBackendProfileParams) WithQ(q *string) *SearchBackendProfileParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the search backend profile params
func (o *SearchBackendProfileParams) SetQ(q *string) {
	o.Q = q
}

// WithSize adds the size to the search backend profile params
func (o *SearchBackendProfileParams) WithSize(size *int32) *SearchBackendProfileParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the search backend profile params
func (o *SearchBackendProfileParams) SetSize(size *int32) {
	o.Size = size
}

// WithSort adds the sort to the search backend profile params
func (o *SearchBackendProfileParams) WithSort(sort *string) *SearchBackendProfileParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the search backend profile params
func (o *SearchBackendProfileParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *SearchBackendProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

		// query array param id
		if err := r.SetQueryParam("id", joinedID...); err != nil {
			return err
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

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSearchBackendProfile binds the parameter fields
func (o *SearchBackendProfileParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSearchBackendProfile binds the parameter id
func (o *SearchBackendProfileParams) bindParamID(formats strfmt.Registry) []string {
	iDIR := o.ID

	var iDIC []string
	for _, iDIIR := range iDIR { // explode []int64

		iDIIV := swag.FormatInt64(iDIIR) // int64 as string
		iDIC = append(iDIC, iDIIV)
	}

	// items.CollectionFormat: "multi"
	iDIS := swag.JoinByFormat(iDIC, "multi")

	return iDIS
}
