// Code generated by go-swagger; DO NOT EDIT.

package call_service

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

// RedialCallReader is a Reader for the RedialCall structure.
type RedialCallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RedialCallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRedialCallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRedialCallDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRedialCallOK creates a RedialCallOK with default headers values
func NewRedialCallOK() *RedialCallOK {
	return &RedialCallOK{}
}

/*
RedialCallOK describes a response with status code 200, with default header values.

A successful response.
*/
type RedialCallOK struct {
	Payload *models.EngineCreateCallResponse
}

// IsSuccess returns true when this redial call Ok response has a 2xx status code
func (o *RedialCallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this redial call Ok response has a 3xx status code
func (o *RedialCallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this redial call Ok response has a 4xx status code
func (o *RedialCallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this redial call Ok response has a 5xx status code
func (o *RedialCallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this redial call Ok response a status code equal to that given
func (o *RedialCallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the redial call Ok response
func (o *RedialCallOK) Code() int {
	return 200
}

func (o *RedialCallOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /calls/history/{call_id}/redial][%d] redialCallOk %s", 200, payload)
}

func (o *RedialCallOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /calls/history/{call_id}/redial][%d] redialCallOk %s", 200, payload)
}

func (o *RedialCallOK) GetPayload() *models.EngineCreateCallResponse {
	return o.Payload
}

func (o *RedialCallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineCreateCallResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRedialCallDefault creates a RedialCallDefault with default headers values
func NewRedialCallDefault(code int) *RedialCallDefault {
	return &RedialCallDefault{
		_statusCode: code,
	}
}

/*
RedialCallDefault describes a response with status code -1, with default header values.

Server error
*/
type RedialCallDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this redial call default response has a 2xx status code
func (o *RedialCallDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this redial call default response has a 3xx status code
func (o *RedialCallDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this redial call default response has a 4xx status code
func (o *RedialCallDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this redial call default response has a 5xx status code
func (o *RedialCallDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this redial call default response a status code equal to that given
func (o *RedialCallDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the redial call default response
func (o *RedialCallDefault) Code() int {
	return o._statusCode
}

func (o *RedialCallDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /calls/history/{call_id}/redial][%d] RedialCall default %s", o._statusCode, payload)
}

func (o *RedialCallDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /calls/history/{call_id}/redial][%d] RedialCall default %s", o._statusCode, payload)
}

func (o *RedialCallDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *RedialCallDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
