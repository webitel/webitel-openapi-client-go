// Code generated by go-swagger; DO NOT EDIT.

package catalog

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

// NewCatalogGetDialogsParams creates a new CatalogGetDialogsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCatalogGetDialogsParams() *CatalogGetDialogsParams {
	return &CatalogGetDialogsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogGetDialogsParamsWithTimeout creates a new CatalogGetDialogsParams object
// with the ability to set a timeout on a request.
func NewCatalogGetDialogsParamsWithTimeout(timeout time.Duration) *CatalogGetDialogsParams {
	return &CatalogGetDialogsParams{
		timeout: timeout,
	}
}

// NewCatalogGetDialogsParamsWithContext creates a new CatalogGetDialogsParams object
// with the ability to set a context for a request.
func NewCatalogGetDialogsParamsWithContext(ctx context.Context) *CatalogGetDialogsParams {
	return &CatalogGetDialogsParams{
		Context: ctx,
	}
}

// NewCatalogGetDialogsParamsWithHTTPClient creates a new CatalogGetDialogsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCatalogGetDialogsParamsWithHTTPClient(client *http.Client) *CatalogGetDialogsParams {
	return &CatalogGetDialogsParams{
		HTTPClient: client,
	}
}

/*
CatalogGetDialogsParams contains all the parameters to send to the API endpoint

	for the catalog get dialogs operation.

	Typically these are written to a http.Request.
*/
type CatalogGetDialogsParams struct {

	/* DateSince.

	     Since epochtime (milli).
	**Match**: greater than ..

	     Format: int64
	*/
	DateSince *string

	/* DateUntil.

	     Until epochtime (milli).
	**Match**: less or equal ..

	     Format: int64
	*/
	DateUntil *string

	/* Fields.

	   Fields [Q]uery to build result dataset record.
	*/
	Fields []string

	/* GroupString.

	   This is a request variable of the map type. The query format is "map_name[key]=value", e.g. If the map name is Age, the key type is string, and the value type is integer, the query parameter is expressed as Age["bob"]=18
	*/
	GroupString *string

	/* ID.

	     Set of unique chat IDentifier(s).
	Accept: dialog -or- member ID.
	*/
	ID []string

	/* Online.

	   Dialogs ONLY that are currently [not] active( closed: ? ).
	*/
	Online *bool

	/* Page.

	   Page number to return. **default**: 1.

	   Format: int32
	*/
	Page *int32

	/* PeerID.

	     Contact unique **ID**entifier.
	Contact **type**-specific string.
	*/
	PeerID *string

	/* PeerName.

	   Contact display **name**.
	*/
	PeerName *string

	/* PeerType.

	   Contact **type** provider.
	*/
	PeerType *string

	/* Q.

	   Search term: peer.name
	*/
	Q *string

	/* Size.

	   Page records limit. **default**: 16.

	   Format: int32
	*/
	Size *int32

	/* Sort.

	   Sort records by { fields } specification.
	*/
	Sort []string

	/* ViaID.

	     Contact unique **ID**entifier.
	Contact **type**-specific string.
	*/
	ViaID *string

	/* ViaName.

	   Contact display **name**.
	*/
	ViaName *string

	/* ViaType.

	   Contact **type** provider.
	*/
	ViaType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the catalog get dialogs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogGetDialogsParams) WithDefaults() *CatalogGetDialogsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the catalog get dialogs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogGetDialogsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the catalog get dialogs params
func (o *CatalogGetDialogsParams) WithTimeout(timeout time.Duration) *CatalogGetDialogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog get dialogs params
func (o *CatalogGetDialogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog get dialogs params
func (o *CatalogGetDialogsParams) WithContext(ctx context.Context) *CatalogGetDialogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog get dialogs params
func (o *CatalogGetDialogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog get dialogs params
func (o *CatalogGetDialogsParams) WithHTTPClient(client *http.Client) *CatalogGetDialogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog get dialogs params
func (o *CatalogGetDialogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDateSince adds the dateSince to the catalog get dialogs params
func (o *CatalogGetDialogsParams) WithDateSince(dateSince *string) *CatalogGetDialogsParams {
	o.SetDateSince(dateSince)
	return o
}

// SetDateSince adds the dateSince to the catalog get dialogs params
func (o *CatalogGetDialogsParams) SetDateSince(dateSince *string) {
	o.DateSince = dateSince
}

// WithDateUntil adds the dateUntil to the catalog get dialogs params
func (o *CatalogGetDialogsParams) WithDateUntil(dateUntil *string) *CatalogGetDialogsParams {
	o.SetDateUntil(dateUntil)
	return o
}

// SetDateUntil adds the dateUntil to the catalog get dialogs params
func (o *CatalogGetDialogsParams) SetDateUntil(dateUntil *string) {
	o.DateUntil = dateUntil
}

// WithFields adds the fields to the catalog get dialogs params
func (o *CatalogGetDialogsParams) WithFields(fields []string) *CatalogGetDialogsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the catalog get dialogs params
func (o *CatalogGetDialogsParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithGroupString adds the groupString to the catalog get dialogs params
func (o *CatalogGetDialogsParams) WithGroupString(groupString *string) *CatalogGetDialogsParams {
	o.SetGroupString(groupString)
	return o
}

// SetGroupString adds the groupString to the catalog get dialogs params
func (o *CatalogGetDialogsParams) SetGroupString(groupString *string) {
	o.GroupString = groupString
}

// WithID adds the id to the catalog get dialogs params
func (o *CatalogGetDialogsParams) WithID(id []string) *CatalogGetDialogsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the catalog get dialogs params
func (o *CatalogGetDialogsParams) SetID(id []string) {
	o.ID = id
}

// WithOnline adds the online to the catalog get dialogs params
func (o *CatalogGetDialogsParams) WithOnline(online *bool) *CatalogGetDialogsParams {
	o.SetOnline(online)
	return o
}

// SetOnline adds the online to the catalog get dialogs params
func (o *CatalogGetDialogsParams) SetOnline(online *bool) {
	o.Online = online
}

// WithPage adds the page to the catalog get dialogs params
func (o *CatalogGetDialogsParams) WithPage(page *int32) *CatalogGetDialogsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the catalog get dialogs params
func (o *CatalogGetDialogsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPeerID adds the peerID to the catalog get dialogs params
func (o *CatalogGetDialogsParams) WithPeerID(peerID *string) *CatalogGetDialogsParams {
	o.SetPeerID(peerID)
	return o
}

// SetPeerID adds the peerId to the catalog get dialogs params
func (o *CatalogGetDialogsParams) SetPeerID(peerID *string) {
	o.PeerID = peerID
}

// WithPeerName adds the peerName to the catalog get dialogs params
func (o *CatalogGetDialogsParams) WithPeerName(peerName *string) *CatalogGetDialogsParams {
	o.SetPeerName(peerName)
	return o
}

// SetPeerName adds the peerName to the catalog get dialogs params
func (o *CatalogGetDialogsParams) SetPeerName(peerName *string) {
	o.PeerName = peerName
}

// WithPeerType adds the peerType to the catalog get dialogs params
func (o *CatalogGetDialogsParams) WithPeerType(peerType *string) *CatalogGetDialogsParams {
	o.SetPeerType(peerType)
	return o
}

// SetPeerType adds the peerType to the catalog get dialogs params
func (o *CatalogGetDialogsParams) SetPeerType(peerType *string) {
	o.PeerType = peerType
}

// WithQ adds the q to the catalog get dialogs params
func (o *CatalogGetDialogsParams) WithQ(q *string) *CatalogGetDialogsParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the catalog get dialogs params
func (o *CatalogGetDialogsParams) SetQ(q *string) {
	o.Q = q
}

// WithSize adds the size to the catalog get dialogs params
func (o *CatalogGetDialogsParams) WithSize(size *int32) *CatalogGetDialogsParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the catalog get dialogs params
func (o *CatalogGetDialogsParams) SetSize(size *int32) {
	o.Size = size
}

// WithSort adds the sort to the catalog get dialogs params
func (o *CatalogGetDialogsParams) WithSort(sort []string) *CatalogGetDialogsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the catalog get dialogs params
func (o *CatalogGetDialogsParams) SetSort(sort []string) {
	o.Sort = sort
}

// WithViaID adds the viaID to the catalog get dialogs params
func (o *CatalogGetDialogsParams) WithViaID(viaID *string) *CatalogGetDialogsParams {
	o.SetViaID(viaID)
	return o
}

// SetViaID adds the viaId to the catalog get dialogs params
func (o *CatalogGetDialogsParams) SetViaID(viaID *string) {
	o.ViaID = viaID
}

// WithViaName adds the viaName to the catalog get dialogs params
func (o *CatalogGetDialogsParams) WithViaName(viaName *string) *CatalogGetDialogsParams {
	o.SetViaName(viaName)
	return o
}

// SetViaName adds the viaName to the catalog get dialogs params
func (o *CatalogGetDialogsParams) SetViaName(viaName *string) {
	o.ViaName = viaName
}

// WithViaType adds the viaType to the catalog get dialogs params
func (o *CatalogGetDialogsParams) WithViaType(viaType *string) *CatalogGetDialogsParams {
	o.SetViaType(viaType)
	return o
}

// SetViaType adds the viaType to the catalog get dialogs params
func (o *CatalogGetDialogsParams) SetViaType(viaType *string) {
	o.ViaType = viaType
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogGetDialogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DateSince != nil {

		// query param date.since
		var qrDateSince string

		if o.DateSince != nil {
			qrDateSince = *o.DateSince
		}
		qDateSince := qrDateSince
		if qDateSince != "" {

			if err := r.SetQueryParam("date.since", qDateSince); err != nil {
				return err
			}
		}
	}

	if o.DateUntil != nil {

		// query param date.until
		var qrDateUntil string

		if o.DateUntil != nil {
			qrDateUntil = *o.DateUntil
		}
		qDateUntil := qrDateUntil
		if qDateUntil != "" {

			if err := r.SetQueryParam("date.until", qDateUntil); err != nil {
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

	if o.GroupString != nil {

		// query param group[string]
		var qrGroupString string

		if o.GroupString != nil {
			qrGroupString = *o.GroupString
		}
		qGroupString := qrGroupString
		if qGroupString != "" {

			if err := r.SetQueryParam("group[string]", qGroupString); err != nil {
				return err
			}
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

	if o.Online != nil {

		// query param online
		var qrOnline bool

		if o.Online != nil {
			qrOnline = *o.Online
		}
		qOnline := swag.FormatBool(qrOnline)
		if qOnline != "" {

			if err := r.SetQueryParam("online", qOnline); err != nil {
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

	if o.PeerID != nil {

		// query param peer.id
		var qrPeerID string

		if o.PeerID != nil {
			qrPeerID = *o.PeerID
		}
		qPeerID := qrPeerID
		if qPeerID != "" {

			if err := r.SetQueryParam("peer.id", qPeerID); err != nil {
				return err
			}
		}
	}

	if o.PeerName != nil {

		// query param peer.name
		var qrPeerName string

		if o.PeerName != nil {
			qrPeerName = *o.PeerName
		}
		qPeerName := qrPeerName
		if qPeerName != "" {

			if err := r.SetQueryParam("peer.name", qPeerName); err != nil {
				return err
			}
		}
	}

	if o.PeerType != nil {

		// query param peer.type
		var qrPeerType string

		if o.PeerType != nil {
			qrPeerType = *o.PeerType
		}
		qPeerType := qrPeerType
		if qPeerType != "" {

			if err := r.SetQueryParam("peer.type", qPeerType); err != nil {
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

	if o.ViaID != nil {

		// query param via.id
		var qrViaID string

		if o.ViaID != nil {
			qrViaID = *o.ViaID
		}
		qViaID := qrViaID
		if qViaID != "" {

			if err := r.SetQueryParam("via.id", qViaID); err != nil {
				return err
			}
		}
	}

	if o.ViaName != nil {

		// query param via.name
		var qrViaName string

		if o.ViaName != nil {
			qrViaName = *o.ViaName
		}
		qViaName := qrViaName
		if qViaName != "" {

			if err := r.SetQueryParam("via.name", qViaName); err != nil {
				return err
			}
		}
	}

	if o.ViaType != nil {

		// query param via.type
		var qrViaType string

		if o.ViaType != nil {
			qrViaType = *o.ViaType
		}
		qViaType := qrViaType
		if qViaType != "" {

			if err := r.SetQueryParam("via.type", qViaType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamCatalogGetDialogs binds the parameter fields
func (o *CatalogGetDialogsParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamCatalogGetDialogs binds the parameter id
func (o *CatalogGetDialogsParams) bindParamID(formats strfmt.Registry) []string {
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

// bindParamCatalogGetDialogs binds the parameter sort
func (o *CatalogGetDialogsParams) bindParamSort(formats strfmt.Registry) []string {
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
