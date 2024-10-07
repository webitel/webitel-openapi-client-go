// Code generated by go-swagger; DO NOT EDIT.

package timezones

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

// NewTimezonesListTimezonesParams creates a new TimezonesListTimezonesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTimezonesListTimezonesParams() *TimezonesListTimezonesParams {
	return &TimezonesListTimezonesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTimezonesListTimezonesParamsWithTimeout creates a new TimezonesListTimezonesParams object
// with the ability to set a timeout on a request.
func NewTimezonesListTimezonesParamsWithTimeout(timeout time.Duration) *TimezonesListTimezonesParams {
	return &TimezonesListTimezonesParams{
		timeout: timeout,
	}
}

// NewTimezonesListTimezonesParamsWithContext creates a new TimezonesListTimezonesParams object
// with the ability to set a context for a request.
func NewTimezonesListTimezonesParamsWithContext(ctx context.Context) *TimezonesListTimezonesParams {
	return &TimezonesListTimezonesParams{
		Context: ctx,
	}
}

// NewTimezonesListTimezonesParamsWithHTTPClient creates a new TimezonesListTimezonesParams object
// with the ability to set a custom HTTPClient for a request.
func NewTimezonesListTimezonesParamsWithHTTPClient(client *http.Client) *TimezonesListTimezonesParams {
	return &TimezonesListTimezonesParams{
		HTTPClient: client,
	}
}

/*
TimezonesListTimezonesParams contains all the parameters to send to the API endpoint

	for the timezones list timezones operation.

	Typically these are written to a http.Request.
*/
type TimezonesListTimezonesParams struct {

	/* ContactID.

	   Contact ID associated with.
	*/
	ContactID string

	/* Fields.

	   Fields to be retrieved as a result.
	*/
	Fields []string

	/* ID.

	   Record(s) with unique ID only.
	*/
	ID []string

	/* Page.

	   Page number of result dataset records. offset = (page*size)

	   Format: int32
	*/
	Page *int32

	/* Primary.

	   Primary timezone only.
	*/
	Primary *bool

	/* Q.

	     Search term: location name;
	`?` - matches any one character
	`*` - matches 0 or more characters
	*/
	Q *string

	/* Size.

	   Size count of records on result page. limit = (size++)

	   Format: int32
	*/
	Size *int32

	/* Sort.

	   Sort the result according to fields.
	*/
	Sort []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the timezones list timezones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TimezonesListTimezonesParams) WithDefaults() *TimezonesListTimezonesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the timezones list timezones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TimezonesListTimezonesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the timezones list timezones params
func (o *TimezonesListTimezonesParams) WithTimeout(timeout time.Duration) *TimezonesListTimezonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the timezones list timezones params
func (o *TimezonesListTimezonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the timezones list timezones params
func (o *TimezonesListTimezonesParams) WithContext(ctx context.Context) *TimezonesListTimezonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the timezones list timezones params
func (o *TimezonesListTimezonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the timezones list timezones params
func (o *TimezonesListTimezonesParams) WithHTTPClient(client *http.Client) *TimezonesListTimezonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the timezones list timezones params
func (o *TimezonesListTimezonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactID adds the contactID to the timezones list timezones params
func (o *TimezonesListTimezonesParams) WithContactID(contactID string) *TimezonesListTimezonesParams {
	o.SetContactID(contactID)
	return o
}

// SetContactID adds the contactId to the timezones list timezones params
func (o *TimezonesListTimezonesParams) SetContactID(contactID string) {
	o.ContactID = contactID
}

// WithFields adds the fields to the timezones list timezones params
func (o *TimezonesListTimezonesParams) WithFields(fields []string) *TimezonesListTimezonesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the timezones list timezones params
func (o *TimezonesListTimezonesParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithID adds the id to the timezones list timezones params
func (o *TimezonesListTimezonesParams) WithID(id []string) *TimezonesListTimezonesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the timezones list timezones params
func (o *TimezonesListTimezonesParams) SetID(id []string) {
	o.ID = id
}

// WithPage adds the page to the timezones list timezones params
func (o *TimezonesListTimezonesParams) WithPage(page *int32) *TimezonesListTimezonesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the timezones list timezones params
func (o *TimezonesListTimezonesParams) SetPage(page *int32) {
	o.Page = page
}

// WithPrimary adds the primary to the timezones list timezones params
func (o *TimezonesListTimezonesParams) WithPrimary(primary *bool) *TimezonesListTimezonesParams {
	o.SetPrimary(primary)
	return o
}

// SetPrimary adds the primary to the timezones list timezones params
func (o *TimezonesListTimezonesParams) SetPrimary(primary *bool) {
	o.Primary = primary
}

// WithQ adds the q to the timezones list timezones params
func (o *TimezonesListTimezonesParams) WithQ(q *string) *TimezonesListTimezonesParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the timezones list timezones params
func (o *TimezonesListTimezonesParams) SetQ(q *string) {
	o.Q = q
}

// WithSize adds the size to the timezones list timezones params
func (o *TimezonesListTimezonesParams) WithSize(size *int32) *TimezonesListTimezonesParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the timezones list timezones params
func (o *TimezonesListTimezonesParams) SetSize(size *int32) {
	o.Size = size
}

// WithSort adds the sort to the timezones list timezones params
func (o *TimezonesListTimezonesParams) WithSort(sort []string) *TimezonesListTimezonesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the timezones list timezones params
func (o *TimezonesListTimezonesParams) SetSort(sort []string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *TimezonesListTimezonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contact_id
	if err := r.SetPathParam("contact_id", o.ContactID); err != nil {
		return err
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

	if o.Primary != nil {

		// query param primary
		var qrPrimary bool

		if o.Primary != nil {
			qrPrimary = *o.Primary
		}
		qPrimary := swag.FormatBool(qrPrimary)
		if qPrimary != "" {

			if err := r.SetQueryParam("primary", qPrimary); err != nil {
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

// bindParamTimezonesListTimezones binds the parameter fields
func (o *TimezonesListTimezonesParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamTimezonesListTimezones binds the parameter id
func (o *TimezonesListTimezonesParams) bindParamID(formats strfmt.Registry) []string {
	iDIR := o.ID

	var iDIC []string
	for _, iDIIR := range iDIR { // explode []string

		iDIIV := iDIIR // string as string
		iDIC = append(iDIC, iDIIV)
	}

	// items.CollectionFormat: "multi"
	iDIS := swag.JoinByFormat(iDIC, "multi")

	return iDIS
}

// bindParamTimezonesListTimezones binds the parameter sort
func (o *TimezonesListTimezonesParams) bindParamSort(formats strfmt.Registry) []string {
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
