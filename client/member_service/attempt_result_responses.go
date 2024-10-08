// Code generated by go-swagger; DO NOT EDIT.

package member_service

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

// AttemptResultReader is a Reader for the AttemptResult structure.
type AttemptResultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AttemptResultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAttemptResultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAttemptResultDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAttemptResultOK creates a AttemptResultOK with default headers values
func NewAttemptResultOK() *AttemptResultOK {
	return &AttemptResultOK{}
}

/*
AttemptResultOK describes a response with status code 200, with default header values.

A successful response.
*/
type AttemptResultOK struct {
	Payload *models.EngineAttemptResultResponse
}

// IsSuccess returns true when this attempt result Ok response has a 2xx status code
func (o *AttemptResultOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this attempt result Ok response has a 3xx status code
func (o *AttemptResultOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attempt result Ok response has a 4xx status code
func (o *AttemptResultOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this attempt result Ok response has a 5xx status code
func (o *AttemptResultOK) IsServerError() bool {
	return false
}

// IsCode returns true when this attempt result Ok response a status code equal to that given
func (o *AttemptResultOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the attempt result Ok response
func (o *AttemptResultOK) Code() int {
	return 200
}

func (o *AttemptResultOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /call_center/queues/{queue_id}/members/{member_id}/attempts/{attempt_id}][%d] attemptResultOk %s", 200, payload)
}

func (o *AttemptResultOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /call_center/queues/{queue_id}/members/{member_id}/attempts/{attempt_id}][%d] attemptResultOk %s", 200, payload)
}

func (o *AttemptResultOK) GetPayload() *models.EngineAttemptResultResponse {
	return o.Payload
}

func (o *AttemptResultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAttemptResultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttemptResultDefault creates a AttemptResultDefault with default headers values
func NewAttemptResultDefault(code int) *AttemptResultDefault {
	return &AttemptResultDefault{
		_statusCode: code,
	}
}

/*
AttemptResultDefault describes a response with status code -1, with default header values.

Server error
*/
type AttemptResultDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this attempt result default response has a 2xx status code
func (o *AttemptResultDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this attempt result default response has a 3xx status code
func (o *AttemptResultDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this attempt result default response has a 4xx status code
func (o *AttemptResultDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this attempt result default response has a 5xx status code
func (o *AttemptResultDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this attempt result default response a status code equal to that given
func (o *AttemptResultDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the attempt result default response
func (o *AttemptResultDefault) Code() int {
	return o._statusCode
}

func (o *AttemptResultDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /call_center/queues/{queue_id}/members/{member_id}/attempts/{attempt_id}][%d] AttemptResult default %s", o._statusCode, payload)
}

func (o *AttemptResultDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /call_center/queues/{queue_id}/members/{member_id}/attempts/{attempt_id}][%d] AttemptResult default %s", o._statusCode, payload)
}

func (o *AttemptResultDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *AttemptResultDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
