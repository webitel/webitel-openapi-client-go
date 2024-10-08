// Code generated by go-swagger; DO NOT EDIT.

package forecast_calculation_service

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

// NewForecastCalculationServiceReadForecastCalculationParams creates a new ForecastCalculationServiceReadForecastCalculationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewForecastCalculationServiceReadForecastCalculationParams() *ForecastCalculationServiceReadForecastCalculationParams {
	return &ForecastCalculationServiceReadForecastCalculationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewForecastCalculationServiceReadForecastCalculationParamsWithTimeout creates a new ForecastCalculationServiceReadForecastCalculationParams object
// with the ability to set a timeout on a request.
func NewForecastCalculationServiceReadForecastCalculationParamsWithTimeout(timeout time.Duration) *ForecastCalculationServiceReadForecastCalculationParams {
	return &ForecastCalculationServiceReadForecastCalculationParams{
		timeout: timeout,
	}
}

// NewForecastCalculationServiceReadForecastCalculationParamsWithContext creates a new ForecastCalculationServiceReadForecastCalculationParams object
// with the ability to set a context for a request.
func NewForecastCalculationServiceReadForecastCalculationParamsWithContext(ctx context.Context) *ForecastCalculationServiceReadForecastCalculationParams {
	return &ForecastCalculationServiceReadForecastCalculationParams{
		Context: ctx,
	}
}

// NewForecastCalculationServiceReadForecastCalculationParamsWithHTTPClient creates a new ForecastCalculationServiceReadForecastCalculationParams object
// with the ability to set a custom HTTPClient for a request.
func NewForecastCalculationServiceReadForecastCalculationParamsWithHTTPClient(client *http.Client) *ForecastCalculationServiceReadForecastCalculationParams {
	return &ForecastCalculationServiceReadForecastCalculationParams{
		HTTPClient: client,
	}
}

/*
ForecastCalculationServiceReadForecastCalculationParams contains all the parameters to send to the API endpoint

	for the forecast calculation service read forecast calculation operation.

	Typically these are written to a http.Request.
*/
type ForecastCalculationServiceReadForecastCalculationParams struct {

	// Fields.
	Fields []string

	// ID.
	//
	// Format: int64
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the forecast calculation service read forecast calculation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ForecastCalculationServiceReadForecastCalculationParams) WithDefaults() *ForecastCalculationServiceReadForecastCalculationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the forecast calculation service read forecast calculation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ForecastCalculationServiceReadForecastCalculationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the forecast calculation service read forecast calculation params
func (o *ForecastCalculationServiceReadForecastCalculationParams) WithTimeout(timeout time.Duration) *ForecastCalculationServiceReadForecastCalculationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the forecast calculation service read forecast calculation params
func (o *ForecastCalculationServiceReadForecastCalculationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the forecast calculation service read forecast calculation params
func (o *ForecastCalculationServiceReadForecastCalculationParams) WithContext(ctx context.Context) *ForecastCalculationServiceReadForecastCalculationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the forecast calculation service read forecast calculation params
func (o *ForecastCalculationServiceReadForecastCalculationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the forecast calculation service read forecast calculation params
func (o *ForecastCalculationServiceReadForecastCalculationParams) WithHTTPClient(client *http.Client) *ForecastCalculationServiceReadForecastCalculationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the forecast calculation service read forecast calculation params
func (o *ForecastCalculationServiceReadForecastCalculationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the forecast calculation service read forecast calculation params
func (o *ForecastCalculationServiceReadForecastCalculationParams) WithFields(fields []string) *ForecastCalculationServiceReadForecastCalculationParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the forecast calculation service read forecast calculation params
func (o *ForecastCalculationServiceReadForecastCalculationParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithID adds the id to the forecast calculation service read forecast calculation params
func (o *ForecastCalculationServiceReadForecastCalculationParams) WithID(id string) *ForecastCalculationServiceReadForecastCalculationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the forecast calculation service read forecast calculation params
func (o *ForecastCalculationServiceReadForecastCalculationParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ForecastCalculationServiceReadForecastCalculationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamForecastCalculationServiceReadForecastCalculation binds the parameter fields
func (o *ForecastCalculationServiceReadForecastCalculationParams) bindParamFields(formats strfmt.Registry) []string {
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
