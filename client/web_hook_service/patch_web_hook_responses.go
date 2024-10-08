// Code generated by go-swagger; DO NOT EDIT.

package web_hook_service

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

// PatchWebHookReader is a Reader for the PatchWebHook structure.
type PatchWebHookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchWebHookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchWebHookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchWebHookDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchWebHookOK creates a PatchWebHookOK with default headers values
func NewPatchWebHookOK() *PatchWebHookOK {
	return &PatchWebHookOK{}
}

/*
PatchWebHookOK describes a response with status code 200, with default header values.

A successful response.
*/
type PatchWebHookOK struct {
	Payload *models.EngineWebHook
}

// IsSuccess returns true when this patch web hook Ok response has a 2xx status code
func (o *PatchWebHookOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch web hook Ok response has a 3xx status code
func (o *PatchWebHookOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch web hook Ok response has a 4xx status code
func (o *PatchWebHookOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch web hook Ok response has a 5xx status code
func (o *PatchWebHookOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch web hook Ok response a status code equal to that given
func (o *PatchWebHookOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch web hook Ok response
func (o *PatchWebHookOK) Code() int {
	return 200
}

func (o *PatchWebHookOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /hook/{id}][%d] patchWebHookOk %s", 200, payload)
}

func (o *PatchWebHookOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /hook/{id}][%d] patchWebHookOk %s", 200, payload)
}

func (o *PatchWebHookOK) GetPayload() *models.EngineWebHook {
	return o.Payload
}

func (o *PatchWebHookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineWebHook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWebHookDefault creates a PatchWebHookDefault with default headers values
func NewPatchWebHookDefault(code int) *PatchWebHookDefault {
	return &PatchWebHookDefault{
		_statusCode: code,
	}
}

/*
PatchWebHookDefault describes a response with status code -1, with default header values.

Server error
*/
type PatchWebHookDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this patch web hook default response has a 2xx status code
func (o *PatchWebHookDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch web hook default response has a 3xx status code
func (o *PatchWebHookDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch web hook default response has a 4xx status code
func (o *PatchWebHookDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch web hook default response has a 5xx status code
func (o *PatchWebHookDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch web hook default response a status code equal to that given
func (o *PatchWebHookDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the patch web hook default response
func (o *PatchWebHookDefault) Code() int {
	return o._statusCode
}

func (o *PatchWebHookDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /hook/{id}][%d] PatchWebHook default %s", o._statusCode, payload)
}

func (o *PatchWebHookDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /hook/{id}][%d] PatchWebHook default %s", o._statusCode, payload)
}

func (o *PatchWebHookDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *PatchWebHookDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
