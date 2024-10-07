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

// SetVariablesCallReader is a Reader for the SetVariablesCall structure.
type SetVariablesCallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetVariablesCallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetVariablesCallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSetVariablesCallDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSetVariablesCallOK creates a SetVariablesCallOK with default headers values
func NewSetVariablesCallOK() *SetVariablesCallOK {
	return &SetVariablesCallOK{}
}

/*
SetVariablesCallOK describes a response with status code 200, with default header values.

A successful response.
*/
type SetVariablesCallOK struct {
	Payload models.EngineSetVariablesCallResponse
}

// IsSuccess returns true when this set variables call Ok response has a 2xx status code
func (o *SetVariablesCallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set variables call Ok response has a 3xx status code
func (o *SetVariablesCallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set variables call Ok response has a 4xx status code
func (o *SetVariablesCallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this set variables call Ok response has a 5xx status code
func (o *SetVariablesCallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this set variables call Ok response a status code equal to that given
func (o *SetVariablesCallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the set variables call Ok response
func (o *SetVariablesCallOK) Code() int {
	return 200
}

func (o *SetVariablesCallOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /calls/active/{id}/variables][%d] setVariablesCallOk %s", 200, payload)
}

func (o *SetVariablesCallOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /calls/active/{id}/variables][%d] setVariablesCallOk %s", 200, payload)
}

func (o *SetVariablesCallOK) GetPayload() models.EngineSetVariablesCallResponse {
	return o.Payload
}

func (o *SetVariablesCallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetVariablesCallDefault creates a SetVariablesCallDefault with default headers values
func NewSetVariablesCallDefault(code int) *SetVariablesCallDefault {
	return &SetVariablesCallDefault{
		_statusCode: code,
	}
}

/*
SetVariablesCallDefault describes a response with status code -1, with default header values.

Server error
*/
type SetVariablesCallDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this set variables call default response has a 2xx status code
func (o *SetVariablesCallDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this set variables call default response has a 3xx status code
func (o *SetVariablesCallDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this set variables call default response has a 4xx status code
func (o *SetVariablesCallDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this set variables call default response has a 5xx status code
func (o *SetVariablesCallDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this set variables call default response a status code equal to that given
func (o *SetVariablesCallDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the set variables call default response
func (o *SetVariablesCallDefault) Code() int {
	return o._statusCode
}

func (o *SetVariablesCallDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /calls/active/{id}/variables][%d] SetVariablesCall default %s", o._statusCode, payload)
}

func (o *SetVariablesCallDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /calls/active/{id}/variables][%d] SetVariablesCall default %s", o._statusCode, payload)
}

func (o *SetVariablesCallDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *SetVariablesCallDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
