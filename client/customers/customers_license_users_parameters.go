// Code generated by go-swagger; DO NOT EDIT.

package customers

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

// NewCustomersLicenseUsersParams creates a new CustomersLicenseUsersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomersLicenseUsersParams() *CustomersLicenseUsersParams {
	return &CustomersLicenseUsersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomersLicenseUsersParamsWithTimeout creates a new CustomersLicenseUsersParams object
// with the ability to set a timeout on a request.
func NewCustomersLicenseUsersParamsWithTimeout(timeout time.Duration) *CustomersLicenseUsersParams {
	return &CustomersLicenseUsersParams{
		timeout: timeout,
	}
}

// NewCustomersLicenseUsersParamsWithContext creates a new CustomersLicenseUsersParams object
// with the ability to set a context for a request.
func NewCustomersLicenseUsersParamsWithContext(ctx context.Context) *CustomersLicenseUsersParams {
	return &CustomersLicenseUsersParams{
		Context: ctx,
	}
}

// NewCustomersLicenseUsersParamsWithHTTPClient creates a new CustomersLicenseUsersParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomersLicenseUsersParamsWithHTTPClient(client *http.Client) *CustomersLicenseUsersParams {
	return &CustomersLicenseUsersParams{
		HTTPClient: client,
	}
}

/*
CustomersLicenseUsersParams contains all the parameters to send to the API endpoint

	for the customers license users operation.

	Typically these are written to a http.Request.
*/
type CustomersLicenseUsersParams struct {

	/* Fields.

	   set of output fields
	*/
	Fields []string

	/* ID.

	     ----- Search Filters ----- //

	REQUIRED Product ID
	*/
	ID string

	/* Page.

	   default: 1

	   Format: int32
	*/
	Page *int32

	/* Q.

	   filter: term of search (username)
	*/
	Q *string

	/* Size.

	     ----- Search Options ----- //

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

// WithDefaults hydrates default values in the customers license users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomersLicenseUsersParams) WithDefaults() *CustomersLicenseUsersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customers license users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomersLicenseUsersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customers license users params
func (o *CustomersLicenseUsersParams) WithTimeout(timeout time.Duration) *CustomersLicenseUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customers license users params
func (o *CustomersLicenseUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customers license users params
func (o *CustomersLicenseUsersParams) WithContext(ctx context.Context) *CustomersLicenseUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customers license users params
func (o *CustomersLicenseUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customers license users params
func (o *CustomersLicenseUsersParams) WithHTTPClient(client *http.Client) *CustomersLicenseUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customers license users params
func (o *CustomersLicenseUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the customers license users params
func (o *CustomersLicenseUsersParams) WithFields(fields []string) *CustomersLicenseUsersParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the customers license users params
func (o *CustomersLicenseUsersParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithID adds the id to the customers license users params
func (o *CustomersLicenseUsersParams) WithID(id string) *CustomersLicenseUsersParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customers license users params
func (o *CustomersLicenseUsersParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the customers license users params
func (o *CustomersLicenseUsersParams) WithPage(page *int32) *CustomersLicenseUsersParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the customers license users params
func (o *CustomersLicenseUsersParams) SetPage(page *int32) {
	o.Page = page
}

// WithQ adds the q to the customers license users params
func (o *CustomersLicenseUsersParams) WithQ(q *string) *CustomersLicenseUsersParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the customers license users params
func (o *CustomersLicenseUsersParams) SetQ(q *string) {
	o.Q = q
}

// WithSize adds the size to the customers license users params
func (o *CustomersLicenseUsersParams) WithSize(size *int32) *CustomersLicenseUsersParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the customers license users params
func (o *CustomersLicenseUsersParams) SetSize(size *int32) {
	o.Size = size
}

// WithSort adds the sort to the customers license users params
func (o *CustomersLicenseUsersParams) WithSort(sort []string) *CustomersLicenseUsersParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the customers license users params
func (o *CustomersLicenseUsersParams) SetSort(sort []string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *CustomersLicenseUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
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

// bindParamCustomersLicenseUsers binds the parameter fields
func (o *CustomersLicenseUsersParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamCustomersLicenseUsers binds the parameter sort
func (o *CustomersLicenseUsersParams) bindParamSort(formats strfmt.Registry) []string {
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
