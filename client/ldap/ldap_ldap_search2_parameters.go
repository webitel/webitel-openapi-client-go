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

// NewLDAPLDAPSearch2Params creates a new LDAPLDAPSearch2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLDAPLDAPSearch2Params() *LDAPLDAPSearch2Params {
	return &LDAPLDAPSearch2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewLDAPLDAPSearch2ParamsWithTimeout creates a new LDAPLDAPSearch2Params object
// with the ability to set a timeout on a request.
func NewLDAPLDAPSearch2ParamsWithTimeout(timeout time.Duration) *LDAPLDAPSearch2Params {
	return &LDAPLDAPSearch2Params{
		timeout: timeout,
	}
}

// NewLDAPLDAPSearch2ParamsWithContext creates a new LDAPLDAPSearch2Params object
// with the ability to set a context for a request.
func NewLDAPLDAPSearch2ParamsWithContext(ctx context.Context) *LDAPLDAPSearch2Params {
	return &LDAPLDAPSearch2Params{
		Context: ctx,
	}
}

// NewLDAPLDAPSearch2ParamsWithHTTPClient creates a new LDAPLDAPSearch2Params object
// with the ability to set a custom HTTPClient for a request.
func NewLDAPLDAPSearch2ParamsWithHTTPClient(client *http.Client) *LDAPLDAPSearch2Params {
	return &LDAPLDAPSearch2Params{
		HTTPClient: client,
	}
}

/*
LDAPLDAPSearch2Params contains all the parameters to send to the API endpoint

	for the LDAP LDAP search2 operation.

	Typically these are written to a http.Request.
*/
type LDAPLDAPSearch2Params struct {

	/* Attributes.

	   AttributeSelection
	*/
	Attributes []string

	/* BaseObject.

	     ----- SearchRequest -----
	baseObject [D]istinguished[N]ame
	*/
	BaseObject *string

	/* Bind.

	     ----- BIND: Authorization -----

	authorization method e.g.: SIMPLE, SAML, NTLM, etc.
	*/
	Bind *string

	/* CatalogID.

	     ----- connection -----
	Optional. ID of the preconfigured LDAP catalog

	     Format: int64
	*/
	CatalogID string

	/* DerefAliases.

	     neverDerefAliases       (0),
	derefInSearching        (1),
	derefFindingBaseObj     (2),
	derefAlways             (3)

	     Format: int32
	*/
	DerefAliases *int32

	/* Filter.

	   Filter,
	*/
	Filter *string

	/* Password.

	   password
	*/
	Password *string

	/* Scope.

	     baseObject              (0),
	singleLevel             (1),
	wholeSubtree            (2)

	     Format: int32
	*/
	Scope *int32

	/* SizeLimit.

	   INTEGER (0 ..  maxInt),

	   Format: int64
	*/
	SizeLimit *string

	/* TimeLimit.

	   INTEGER (0 ..  maxInt),

	   Format: int64
	*/
	TimeLimit *string

	/* TLSPEM.

	     TODO: (!)

	base64

	     Format: byte
	*/
	TLSPEM *strfmt.Base64

	/* TypesOnly.

	   BOOLEAN,
	*/
	TypesOnly *bool

	/* URL.

	     Optional. URL to establish connection to LDAP catalog

	URL e.g.: [(ldap|ldapi|ldaps)://]host[:port]
	*/
	URL *string

	/* Username.

	   bind_dn
	*/
	Username *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the LDAP LDAP search2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LDAPLDAPSearch2Params) WithDefaults() *LDAPLDAPSearch2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the LDAP LDAP search2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LDAPLDAPSearch2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) WithTimeout(timeout time.Duration) *LDAPLDAPSearch2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) WithContext(ctx context.Context) *LDAPLDAPSearch2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) WithHTTPClient(client *http.Client) *LDAPLDAPSearch2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttributes adds the attributes to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) WithAttributes(attributes []string) *LDAPLDAPSearch2Params {
	o.SetAttributes(attributes)
	return o
}

// SetAttributes adds the attributes to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) SetAttributes(attributes []string) {
	o.Attributes = attributes
}

// WithBaseObject adds the baseObject to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) WithBaseObject(baseObject *string) *LDAPLDAPSearch2Params {
	o.SetBaseObject(baseObject)
	return o
}

// SetBaseObject adds the baseObject to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) SetBaseObject(baseObject *string) {
	o.BaseObject = baseObject
}

// WithBind adds the bind to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) WithBind(bind *string) *LDAPLDAPSearch2Params {
	o.SetBind(bind)
	return o
}

// SetBind adds the bind to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) SetBind(bind *string) {
	o.Bind = bind
}

// WithCatalogID adds the catalogID to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) WithCatalogID(catalogID string) *LDAPLDAPSearch2Params {
	o.SetCatalogID(catalogID)
	return o
}

// SetCatalogID adds the catalogId to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) SetCatalogID(catalogID string) {
	o.CatalogID = catalogID
}

// WithDerefAliases adds the derefAliases to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) WithDerefAliases(derefAliases *int32) *LDAPLDAPSearch2Params {
	o.SetDerefAliases(derefAliases)
	return o
}

// SetDerefAliases adds the derefAliases to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) SetDerefAliases(derefAliases *int32) {
	o.DerefAliases = derefAliases
}

// WithFilter adds the filter to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) WithFilter(filter *string) *LDAPLDAPSearch2Params {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) SetFilter(filter *string) {
	o.Filter = filter
}

// WithPassword adds the password to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) WithPassword(password *string) *LDAPLDAPSearch2Params {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) SetPassword(password *string) {
	o.Password = password
}

// WithScope adds the scope to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) WithScope(scope *int32) *LDAPLDAPSearch2Params {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) SetScope(scope *int32) {
	o.Scope = scope
}

// WithSizeLimit adds the sizeLimit to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) WithSizeLimit(sizeLimit *string) *LDAPLDAPSearch2Params {
	o.SetSizeLimit(sizeLimit)
	return o
}

// SetSizeLimit adds the sizeLimit to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) SetSizeLimit(sizeLimit *string) {
	o.SizeLimit = sizeLimit
}

// WithTimeLimit adds the timeLimit to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) WithTimeLimit(timeLimit *string) *LDAPLDAPSearch2Params {
	o.SetTimeLimit(timeLimit)
	return o
}

// SetTimeLimit adds the timeLimit to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) SetTimeLimit(timeLimit *string) {
	o.TimeLimit = timeLimit
}

// WithTLSPEM adds the tLSPEM to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) WithTLSPEM(tLSPEM *strfmt.Base64) *LDAPLDAPSearch2Params {
	o.SetTLSPEM(tLSPEM)
	return o
}

// SetTLSPEM adds the tlsPEM to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) SetTLSPEM(tLSPEM *strfmt.Base64) {
	o.TLSPEM = tLSPEM
}

// WithTypesOnly adds the typesOnly to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) WithTypesOnly(typesOnly *bool) *LDAPLDAPSearch2Params {
	o.SetTypesOnly(typesOnly)
	return o
}

// SetTypesOnly adds the typesOnly to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) SetTypesOnly(typesOnly *bool) {
	o.TypesOnly = typesOnly
}

// WithURL adds the url to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) WithURL(url *string) *LDAPLDAPSearch2Params {
	o.SetURL(url)
	return o
}

// SetURL adds the url to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) SetURL(url *string) {
	o.URL = url
}

// WithUsername adds the username to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) WithUsername(username *string) *LDAPLDAPSearch2Params {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the LDAP LDAP search2 params
func (o *LDAPLDAPSearch2Params) SetUsername(username *string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *LDAPLDAPSearch2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Attributes != nil {

		// binding items for attributes
		joinedAttributes := o.bindParamAttributes(reg)

		// query array param attributes
		if err := r.SetQueryParam("attributes", joinedAttributes...); err != nil {
			return err
		}
	}

	if o.BaseObject != nil {

		// query param baseObject
		var qrBaseObject string

		if o.BaseObject != nil {
			qrBaseObject = *o.BaseObject
		}
		qBaseObject := qrBaseObject
		if qBaseObject != "" {

			if err := r.SetQueryParam("baseObject", qBaseObject); err != nil {
				return err
			}
		}
	}

	if o.Bind != nil {

		// query param bind
		var qrBind string

		if o.Bind != nil {
			qrBind = *o.Bind
		}
		qBind := qrBind
		if qBind != "" {

			if err := r.SetQueryParam("bind", qBind); err != nil {
				return err
			}
		}
	}

	// path param catalog_id
	if err := r.SetPathParam("catalog_id", o.CatalogID); err != nil {
		return err
	}

	if o.DerefAliases != nil {

		// query param derefAliases
		var qrDerefAliases int32

		if o.DerefAliases != nil {
			qrDerefAliases = *o.DerefAliases
		}
		qDerefAliases := swag.FormatInt32(qrDerefAliases)
		if qDerefAliases != "" {

			if err := r.SetQueryParam("derefAliases", qDerefAliases); err != nil {
				return err
			}
		}
	}

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	if o.Password != nil {

		// query param password
		var qrPassword string

		if o.Password != nil {
			qrPassword = *o.Password
		}
		qPassword := qrPassword
		if qPassword != "" {

			if err := r.SetQueryParam("password", qPassword); err != nil {
				return err
			}
		}
	}

	if o.Scope != nil {

		// query param scope
		var qrScope int32

		if o.Scope != nil {
			qrScope = *o.Scope
		}
		qScope := swag.FormatInt32(qrScope)
		if qScope != "" {

			if err := r.SetQueryParam("scope", qScope); err != nil {
				return err
			}
		}
	}

	if o.SizeLimit != nil {

		// query param sizeLimit
		var qrSizeLimit string

		if o.SizeLimit != nil {
			qrSizeLimit = *o.SizeLimit
		}
		qSizeLimit := qrSizeLimit
		if qSizeLimit != "" {

			if err := r.SetQueryParam("sizeLimit", qSizeLimit); err != nil {
				return err
			}
		}
	}

	if o.TimeLimit != nil {

		// query param timeLimit
		var qrTimeLimit string

		if o.TimeLimit != nil {
			qrTimeLimit = *o.TimeLimit
		}
		qTimeLimit := qrTimeLimit
		if qTimeLimit != "" {

			if err := r.SetQueryParam("timeLimit", qTimeLimit); err != nil {
				return err
			}
		}
	}

	if o.TLSPEM != nil {

		// query param tls.PEM
		var qrTLSPEM strfmt.Base64

		if o.TLSPEM != nil {
			qrTLSPEM = *o.TLSPEM
		}
		qTLSPEM := qrTLSPEM.String()
		if qTLSPEM != "" {

			if err := r.SetQueryParam("tls.PEM", qTLSPEM); err != nil {
				return err
			}
		}
	}

	if o.TypesOnly != nil {

		// query param typesOnly
		var qrTypesOnly bool

		if o.TypesOnly != nil {
			qrTypesOnly = *o.TypesOnly
		}
		qTypesOnly := swag.FormatBool(qrTypesOnly)
		if qTypesOnly != "" {

			if err := r.SetQueryParam("typesOnly", qTypesOnly); err != nil {
				return err
			}
		}
	}

	if o.URL != nil {

		// query param url
		var qrURL string

		if o.URL != nil {
			qrURL = *o.URL
		}
		qURL := qrURL
		if qURL != "" {

			if err := r.SetQueryParam("url", qURL); err != nil {
				return err
			}
		}
	}

	if o.Username != nil {

		// query param username
		var qrUsername string

		if o.Username != nil {
			qrUsername = *o.Username
		}
		qUsername := qrUsername
		if qUsername != "" {

			if err := r.SetQueryParam("username", qUsername); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamLDAPLDAPSearch2 binds the parameter attributes
func (o *LDAPLDAPSearch2Params) bindParamAttributes(formats strfmt.Registry) []string {
	attributesIR := o.Attributes

	var attributesIC []string
	for _, attributesIIR := range attributesIR { // explode []string

		attributesIIV := attributesIIR // string as string
		attributesIC = append(attributesIC, attributesIIV)
	}

	// items.CollectionFormat: "multi"
	attributesIS := swag.JoinByFormat(attributesIC, "multi")

	return attributesIS
}
