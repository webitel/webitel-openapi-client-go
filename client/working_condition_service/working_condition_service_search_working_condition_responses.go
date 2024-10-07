// Code generated by go-swagger; DO NOT EDIT.

package working_condition_service

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

// WorkingConditionServiceSearchWorkingConditionReader is a Reader for the WorkingConditionServiceSearchWorkingCondition structure.
type WorkingConditionServiceSearchWorkingConditionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkingConditionServiceSearchWorkingConditionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkingConditionServiceSearchWorkingConditionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWorkingConditionServiceSearchWorkingConditionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWorkingConditionServiceSearchWorkingConditionOK creates a WorkingConditionServiceSearchWorkingConditionOK with default headers values
func NewWorkingConditionServiceSearchWorkingConditionOK() *WorkingConditionServiceSearchWorkingConditionOK {
	return &WorkingConditionServiceSearchWorkingConditionOK{}
}

/*
WorkingConditionServiceSearchWorkingConditionOK describes a response with status code 200, with default header values.

A successful response.
*/
type WorkingConditionServiceSearchWorkingConditionOK struct {
	Payload *models.WfmSearchWorkingConditionResponse
}

// IsSuccess returns true when this working condition service search working condition Ok response has a 2xx status code
func (o *WorkingConditionServiceSearchWorkingConditionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this working condition service search working condition Ok response has a 3xx status code
func (o *WorkingConditionServiceSearchWorkingConditionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this working condition service search working condition Ok response has a 4xx status code
func (o *WorkingConditionServiceSearchWorkingConditionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this working condition service search working condition Ok response has a 5xx status code
func (o *WorkingConditionServiceSearchWorkingConditionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this working condition service search working condition Ok response a status code equal to that given
func (o *WorkingConditionServiceSearchWorkingConditionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the working condition service search working condition Ok response
func (o *WorkingConditionServiceSearchWorkingConditionOK) Code() int {
	return 200
}

func (o *WorkingConditionServiceSearchWorkingConditionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wfm/lookups/working_conditions][%d] workingConditionServiceSearchWorkingConditionOk %s", 200, payload)
}

func (o *WorkingConditionServiceSearchWorkingConditionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wfm/lookups/working_conditions][%d] workingConditionServiceSearchWorkingConditionOk %s", 200, payload)
}

func (o *WorkingConditionServiceSearchWorkingConditionOK) GetPayload() *models.WfmSearchWorkingConditionResponse {
	return o.Payload
}

func (o *WorkingConditionServiceSearchWorkingConditionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WfmSearchWorkingConditionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWorkingConditionServiceSearchWorkingConditionDefault creates a WorkingConditionServiceSearchWorkingConditionDefault with default headers values
func NewWorkingConditionServiceSearchWorkingConditionDefault(code int) *WorkingConditionServiceSearchWorkingConditionDefault {
	return &WorkingConditionServiceSearchWorkingConditionDefault{
		_statusCode: code,
	}
}

/*
WorkingConditionServiceSearchWorkingConditionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WorkingConditionServiceSearchWorkingConditionDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this working condition service search working condition default response has a 2xx status code
func (o *WorkingConditionServiceSearchWorkingConditionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this working condition service search working condition default response has a 3xx status code
func (o *WorkingConditionServiceSearchWorkingConditionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this working condition service search working condition default response has a 4xx status code
func (o *WorkingConditionServiceSearchWorkingConditionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this working condition service search working condition default response has a 5xx status code
func (o *WorkingConditionServiceSearchWorkingConditionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this working condition service search working condition default response a status code equal to that given
func (o *WorkingConditionServiceSearchWorkingConditionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the working condition service search working condition default response
func (o *WorkingConditionServiceSearchWorkingConditionDefault) Code() int {
	return o._statusCode
}

func (o *WorkingConditionServiceSearchWorkingConditionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wfm/lookups/working_conditions][%d] WorkingConditionService_SearchWorkingCondition default %s", o._statusCode, payload)
}

func (o *WorkingConditionServiceSearchWorkingConditionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wfm/lookups/working_conditions][%d] WorkingConditionService_SearchWorkingCondition default %s", o._statusCode, payload)
}

func (o *WorkingConditionServiceSearchWorkingConditionDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *WorkingConditionServiceSearchWorkingConditionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
