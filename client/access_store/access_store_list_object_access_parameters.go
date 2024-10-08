// Code generated by go-swagger; DO NOT EDIT.

package access_store

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

// NewAccessStoreListObjectAccessParams creates a new AccessStoreListObjectAccessParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccessStoreListObjectAccessParams() *AccessStoreListObjectAccessParams {
	return &AccessStoreListObjectAccessParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccessStoreListObjectAccessParamsWithTimeout creates a new AccessStoreListObjectAccessParams object
// with the ability to set a timeout on a request.
func NewAccessStoreListObjectAccessParamsWithTimeout(timeout time.Duration) *AccessStoreListObjectAccessParams {
	return &AccessStoreListObjectAccessParams{
		timeout: timeout,
	}
}

// NewAccessStoreListObjectAccessParamsWithContext creates a new AccessStoreListObjectAccessParams object
// with the ability to set a context for a request.
func NewAccessStoreListObjectAccessParamsWithContext(ctx context.Context) *AccessStoreListObjectAccessParams {
	return &AccessStoreListObjectAccessParams{
		Context: ctx,
	}
}

// NewAccessStoreListObjectAccessParamsWithHTTPClient creates a new AccessStoreListObjectAccessParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccessStoreListObjectAccessParamsWithHTTPClient(client *http.Client) *AccessStoreListObjectAccessParams {
	return &AccessStoreListObjectAccessParams{
		HTTPClient: client,
	}
}

/*
AccessStoreListObjectAccessParams contains all the parameters to send to the API endpoint

	for the access store list object access operation.

	Typically these are written to a http.Request.
*/
type AccessStoreListObjectAccessParams struct {

	/* Fields.

	     ----- Select Options -------------------------

	attributes list
	*/
	Fields []string

	/* Granted.

	   [xrwdxrwd] [R]ecord-[b]ased-[A]ccess-[C]ontrol level check(!)
	*/
	Granted []string

	/* Grantee.

	   [optional] [TO] subject role(s) id; user -or- role
	*/
	Grantee []string

	/* Grantor.

	     --- filters ---

	[optional] [FROM] each rule owner is any of role(s) id; user -or- role
	*/
	Grantor []string

	/* ObjectID.

	   identifier

	   Format: int64
	*/
	ObjectID string

	/* ObjectName.

	   display name
	*/
	ObjectName string

	/* Page.

	   default: 1

	   Format: int32
	*/
	Page *int32

	/* Q.

	   [optional] [TO] subject.name ILIKE ?q=; user -or- role
	*/
	Q *string

	/* Size.

	     pagedResultsControl

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

// WithDefaults hydrates default values in the access store list object access params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccessStoreListObjectAccessParams) WithDefaults() *AccessStoreListObjectAccessParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the access store list object access params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccessStoreListObjectAccessParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the access store list object access params
func (o *AccessStoreListObjectAccessParams) WithTimeout(timeout time.Duration) *AccessStoreListObjectAccessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the access store list object access params
func (o *AccessStoreListObjectAccessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the access store list object access params
func (o *AccessStoreListObjectAccessParams) WithContext(ctx context.Context) *AccessStoreListObjectAccessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the access store list object access params
func (o *AccessStoreListObjectAccessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the access store list object access params
func (o *AccessStoreListObjectAccessParams) WithHTTPClient(client *http.Client) *AccessStoreListObjectAccessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the access store list object access params
func (o *AccessStoreListObjectAccessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the access store list object access params
func (o *AccessStoreListObjectAccessParams) WithFields(fields []string) *AccessStoreListObjectAccessParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the access store list object access params
func (o *AccessStoreListObjectAccessParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithGranted adds the granted to the access store list object access params
func (o *AccessStoreListObjectAccessParams) WithGranted(granted []string) *AccessStoreListObjectAccessParams {
	o.SetGranted(granted)
	return o
}

// SetGranted adds the granted to the access store list object access params
func (o *AccessStoreListObjectAccessParams) SetGranted(granted []string) {
	o.Granted = granted
}

// WithGrantee adds the grantee to the access store list object access params
func (o *AccessStoreListObjectAccessParams) WithGrantee(grantee []string) *AccessStoreListObjectAccessParams {
	o.SetGrantee(grantee)
	return o
}

// SetGrantee adds the grantee to the access store list object access params
func (o *AccessStoreListObjectAccessParams) SetGrantee(grantee []string) {
	o.Grantee = grantee
}

// WithGrantor adds the grantor to the access store list object access params
func (o *AccessStoreListObjectAccessParams) WithGrantor(grantor []string) *AccessStoreListObjectAccessParams {
	o.SetGrantor(grantor)
	return o
}

// SetGrantor adds the grantor to the access store list object access params
func (o *AccessStoreListObjectAccessParams) SetGrantor(grantor []string) {
	o.Grantor = grantor
}

// WithObjectID adds the objectID to the access store list object access params
func (o *AccessStoreListObjectAccessParams) WithObjectID(objectID string) *AccessStoreListObjectAccessParams {
	o.SetObjectID(objectID)
	return o
}

// SetObjectID adds the objectId to the access store list object access params
func (o *AccessStoreListObjectAccessParams) SetObjectID(objectID string) {
	o.ObjectID = objectID
}

// WithObjectName adds the objectName to the access store list object access params
func (o *AccessStoreListObjectAccessParams) WithObjectName(objectName string) *AccessStoreListObjectAccessParams {
	o.SetObjectName(objectName)
	return o
}

// SetObjectName adds the objectName to the access store list object access params
func (o *AccessStoreListObjectAccessParams) SetObjectName(objectName string) {
	o.ObjectName = objectName
}

// WithPage adds the page to the access store list object access params
func (o *AccessStoreListObjectAccessParams) WithPage(page *int32) *AccessStoreListObjectAccessParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the access store list object access params
func (o *AccessStoreListObjectAccessParams) SetPage(page *int32) {
	o.Page = page
}

// WithQ adds the q to the access store list object access params
func (o *AccessStoreListObjectAccessParams) WithQ(q *string) *AccessStoreListObjectAccessParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the access store list object access params
func (o *AccessStoreListObjectAccessParams) SetQ(q *string) {
	o.Q = q
}

// WithSize adds the size to the access store list object access params
func (o *AccessStoreListObjectAccessParams) WithSize(size *int32) *AccessStoreListObjectAccessParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the access store list object access params
func (o *AccessStoreListObjectAccessParams) SetSize(size *int32) {
	o.Size = size
}

// WithSort adds the sort to the access store list object access params
func (o *AccessStoreListObjectAccessParams) WithSort(sort []string) *AccessStoreListObjectAccessParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the access store list object access params
func (o *AccessStoreListObjectAccessParams) SetSort(sort []string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *AccessStoreListObjectAccessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Granted != nil {

		// binding items for granted
		joinedGranted := o.bindParamGranted(reg)

		// query array param granted
		if err := r.SetQueryParam("granted", joinedGranted...); err != nil {
			return err
		}
	}

	if o.Grantee != nil {

		// binding items for grantee
		joinedGrantee := o.bindParamGrantee(reg)

		// query array param grantee
		if err := r.SetQueryParam("grantee", joinedGrantee...); err != nil {
			return err
		}
	}

	if o.Grantor != nil {

		// binding items for grantor
		joinedGrantor := o.bindParamGrantor(reg)

		// query array param grantor
		if err := r.SetQueryParam("grantor", joinedGrantor...); err != nil {
			return err
		}
	}

	// path param object.id
	if err := r.SetPathParam("object.id", o.ObjectID); err != nil {
		return err
	}

	// path param object.name
	if err := r.SetPathParam("object.name", o.ObjectName); err != nil {
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

// bindParamAccessStoreListObjectAccess binds the parameter fields
func (o *AccessStoreListObjectAccessParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamAccessStoreListObjectAccess binds the parameter granted
func (o *AccessStoreListObjectAccessParams) bindParamGranted(formats strfmt.Registry) []string {
	grantedIR := o.Granted

	var grantedIC []string
	for _, grantedIIR := range grantedIR { // explode []string

		grantedIIV := grantedIIR // string as string
		grantedIC = append(grantedIC, grantedIIV)
	}

	// items.CollectionFormat: "multi"
	grantedIS := swag.JoinByFormat(grantedIC, "multi")

	return grantedIS
}

// bindParamAccessStoreListObjectAccess binds the parameter grantee
func (o *AccessStoreListObjectAccessParams) bindParamGrantee(formats strfmt.Registry) []string {
	granteeIR := o.Grantee

	var granteeIC []string
	for _, granteeIIR := range granteeIR { // explode []string

		granteeIIV := granteeIIR // string as string
		granteeIC = append(granteeIC, granteeIIV)
	}

	// items.CollectionFormat: "multi"
	granteeIS := swag.JoinByFormat(granteeIC, "multi")

	return granteeIS
}

// bindParamAccessStoreListObjectAccess binds the parameter grantor
func (o *AccessStoreListObjectAccessParams) bindParamGrantor(formats strfmt.Registry) []string {
	grantorIR := o.Grantor

	var grantorIC []string
	for _, grantorIIR := range grantorIR { // explode []string

		grantorIIV := grantorIIR // string as string
		grantorIC = append(grantorIC, grantorIIV)
	}

	// items.CollectionFormat: "multi"
	grantorIS := swag.JoinByFormat(grantorIC, "multi")

	return grantorIS
}

// bindParamAccessStoreListObjectAccess binds the parameter sort
func (o *AccessStoreListObjectAccessParams) bindParamSort(formats strfmt.Registry) []string {
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
