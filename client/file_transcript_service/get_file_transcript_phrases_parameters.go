// Code generated by go-swagger; DO NOT EDIT.

package file_transcript_service

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

// NewGetFileTranscriptPhrasesParams creates a new GetFileTranscriptPhrasesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFileTranscriptPhrasesParams() *GetFileTranscriptPhrasesParams {
	return &GetFileTranscriptPhrasesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFileTranscriptPhrasesParamsWithTimeout creates a new GetFileTranscriptPhrasesParams object
// with the ability to set a timeout on a request.
func NewGetFileTranscriptPhrasesParamsWithTimeout(timeout time.Duration) *GetFileTranscriptPhrasesParams {
	return &GetFileTranscriptPhrasesParams{
		timeout: timeout,
	}
}

// NewGetFileTranscriptPhrasesParamsWithContext creates a new GetFileTranscriptPhrasesParams object
// with the ability to set a context for a request.
func NewGetFileTranscriptPhrasesParamsWithContext(ctx context.Context) *GetFileTranscriptPhrasesParams {
	return &GetFileTranscriptPhrasesParams{
		Context: ctx,
	}
}

// NewGetFileTranscriptPhrasesParamsWithHTTPClient creates a new GetFileTranscriptPhrasesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFileTranscriptPhrasesParamsWithHTTPClient(client *http.Client) *GetFileTranscriptPhrasesParams {
	return &GetFileTranscriptPhrasesParams{
		HTTPClient: client,
	}
}

/*
GetFileTranscriptPhrasesParams contains all the parameters to send to the API endpoint

	for the get file transcript phrases operation.

	Typically these are written to a http.Request.
*/
type GetFileTranscriptPhrasesParams struct {

	// ID.
	//
	// Format: int64
	ID string

	// Page.
	//
	// Format: int32
	Page *int32

	// Size.
	//
	// Format: int32
	Size *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get file transcript phrases params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFileTranscriptPhrasesParams) WithDefaults() *GetFileTranscriptPhrasesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get file transcript phrases params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFileTranscriptPhrasesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get file transcript phrases params
func (o *GetFileTranscriptPhrasesParams) WithTimeout(timeout time.Duration) *GetFileTranscriptPhrasesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get file transcript phrases params
func (o *GetFileTranscriptPhrasesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get file transcript phrases params
func (o *GetFileTranscriptPhrasesParams) WithContext(ctx context.Context) *GetFileTranscriptPhrasesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get file transcript phrases params
func (o *GetFileTranscriptPhrasesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get file transcript phrases params
func (o *GetFileTranscriptPhrasesParams) WithHTTPClient(client *http.Client) *GetFileTranscriptPhrasesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get file transcript phrases params
func (o *GetFileTranscriptPhrasesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get file transcript phrases params
func (o *GetFileTranscriptPhrasesParams) WithID(id string) *GetFileTranscriptPhrasesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get file transcript phrases params
func (o *GetFileTranscriptPhrasesParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the get file transcript phrases params
func (o *GetFileTranscriptPhrasesParams) WithPage(page *int32) *GetFileTranscriptPhrasesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get file transcript phrases params
func (o *GetFileTranscriptPhrasesParams) SetPage(page *int32) {
	o.Page = page
}

// WithSize adds the size to the get file transcript phrases params
func (o *GetFileTranscriptPhrasesParams) WithSize(size *int32) *GetFileTranscriptPhrasesParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get file transcript phrases params
func (o *GetFileTranscriptPhrasesParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetFileTranscriptPhrasesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
