// Code generated by go-swagger; DO NOT EDIT.

package pause_template_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/webitel/webitel-openapi-client-go/models"
)

// PauseTemplateServiceReadPauseTemplateReader is a Reader for the PauseTemplateServiceReadPauseTemplate structure.
type PauseTemplateServiceReadPauseTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PauseTemplateServiceReadPauseTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPauseTemplateServiceReadPauseTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPauseTemplateServiceReadPauseTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPauseTemplateServiceReadPauseTemplateOK creates a PauseTemplateServiceReadPauseTemplateOK with default headers values
func NewPauseTemplateServiceReadPauseTemplateOK() *PauseTemplateServiceReadPauseTemplateOK {
	return &PauseTemplateServiceReadPauseTemplateOK{}
}

/*
PauseTemplateServiceReadPauseTemplateOK describes a response with status code 200, with default header values.

A successful response.
*/
type PauseTemplateServiceReadPauseTemplateOK struct {
	Payload *models.WfmReadPauseTemplateResponse
}

// IsSuccess returns true when this pause template service read pause template Ok response has a 2xx status code
func (o *PauseTemplateServiceReadPauseTemplateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pause template service read pause template Ok response has a 3xx status code
func (o *PauseTemplateServiceReadPauseTemplateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pause template service read pause template Ok response has a 4xx status code
func (o *PauseTemplateServiceReadPauseTemplateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pause template service read pause template Ok response has a 5xx status code
func (o *PauseTemplateServiceReadPauseTemplateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pause template service read pause template Ok response a status code equal to that given
func (o *PauseTemplateServiceReadPauseTemplateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pause template service read pause template Ok response
func (o *PauseTemplateServiceReadPauseTemplateOK) Code() int {
	return 200
}

func (o *PauseTemplateServiceReadPauseTemplateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wfm/lookups/pause_templates/{id}][%d] pauseTemplateServiceReadPauseTemplateOk %s", 200, payload)
}

func (o *PauseTemplateServiceReadPauseTemplateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wfm/lookups/pause_templates/{id}][%d] pauseTemplateServiceReadPauseTemplateOk %s", 200, payload)
}

func (o *PauseTemplateServiceReadPauseTemplateOK) GetPayload() *models.WfmReadPauseTemplateResponse {
	return o.Payload
}

func (o *PauseTemplateServiceReadPauseTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WfmReadPauseTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPauseTemplateServiceReadPauseTemplateDefault creates a PauseTemplateServiceReadPauseTemplateDefault with default headers values
func NewPauseTemplateServiceReadPauseTemplateDefault(code int) *PauseTemplateServiceReadPauseTemplateDefault {
	return &PauseTemplateServiceReadPauseTemplateDefault{
		_statusCode: code,
	}
}

/*
PauseTemplateServiceReadPauseTemplateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PauseTemplateServiceReadPauseTemplateDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this pause template service read pause template default response has a 2xx status code
func (o *PauseTemplateServiceReadPauseTemplateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this pause template service read pause template default response has a 3xx status code
func (o *PauseTemplateServiceReadPauseTemplateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this pause template service read pause template default response has a 4xx status code
func (o *PauseTemplateServiceReadPauseTemplateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this pause template service read pause template default response has a 5xx status code
func (o *PauseTemplateServiceReadPauseTemplateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this pause template service read pause template default response a status code equal to that given
func (o *PauseTemplateServiceReadPauseTemplateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the pause template service read pause template default response
func (o *PauseTemplateServiceReadPauseTemplateDefault) Code() int {
	return o._statusCode
}

func (o *PauseTemplateServiceReadPauseTemplateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wfm/lookups/pause_templates/{id}][%d] PauseTemplateService_ReadPauseTemplate default %s", o._statusCode, payload)
}

func (o *PauseTemplateServiceReadPauseTemplateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wfm/lookups/pause_templates/{id}][%d] PauseTemplateService_ReadPauseTemplate default %s", o._statusCode, payload)
}

func (o *PauseTemplateServiceReadPauseTemplateDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *PauseTemplateServiceReadPauseTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
