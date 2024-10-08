// Code generated by go-swagger; DO NOT EDIT.

package team_trigger_service

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

// ReadTeamTriggerReader is a Reader for the ReadTeamTrigger structure.
type ReadTeamTriggerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadTeamTriggerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadTeamTriggerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewReadTeamTriggerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadTeamTriggerOK creates a ReadTeamTriggerOK with default headers values
func NewReadTeamTriggerOK() *ReadTeamTriggerOK {
	return &ReadTeamTriggerOK{}
}

/*
ReadTeamTriggerOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadTeamTriggerOK struct {
	Payload *models.EngineTeamTrigger
}

// IsSuccess returns true when this read team trigger Ok response has a 2xx status code
func (o *ReadTeamTriggerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read team trigger Ok response has a 3xx status code
func (o *ReadTeamTriggerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read team trigger Ok response has a 4xx status code
func (o *ReadTeamTriggerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read team trigger Ok response has a 5xx status code
func (o *ReadTeamTriggerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read team trigger Ok response a status code equal to that given
func (o *ReadTeamTriggerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read team trigger Ok response
func (o *ReadTeamTriggerOK) Code() int {
	return 200
}

func (o *ReadTeamTriggerOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/teams/{team_id}/triggers/{id}][%d] readTeamTriggerOk %s", 200, payload)
}

func (o *ReadTeamTriggerOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/teams/{team_id}/triggers/{id}][%d] readTeamTriggerOk %s", 200, payload)
}

func (o *ReadTeamTriggerOK) GetPayload() *models.EngineTeamTrigger {
	return o.Payload
}

func (o *ReadTeamTriggerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineTeamTrigger)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadTeamTriggerDefault creates a ReadTeamTriggerDefault with default headers values
func NewReadTeamTriggerDefault(code int) *ReadTeamTriggerDefault {
	return &ReadTeamTriggerDefault{
		_statusCode: code,
	}
}

/*
ReadTeamTriggerDefault describes a response with status code -1, with default header values.

Server error
*/
type ReadTeamTriggerDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this read team trigger default response has a 2xx status code
func (o *ReadTeamTriggerDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read team trigger default response has a 3xx status code
func (o *ReadTeamTriggerDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read team trigger default response has a 4xx status code
func (o *ReadTeamTriggerDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read team trigger default response has a 5xx status code
func (o *ReadTeamTriggerDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read team trigger default response a status code equal to that given
func (o *ReadTeamTriggerDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read team trigger default response
func (o *ReadTeamTriggerDefault) Code() int {
	return o._statusCode
}

func (o *ReadTeamTriggerDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/teams/{team_id}/triggers/{id}][%d] ReadTeamTrigger default %s", o._statusCode, payload)
}

func (o *ReadTeamTriggerDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/teams/{team_id}/triggers/{id}][%d] ReadTeamTrigger default %s", o._statusCode, payload)
}

func (o *ReadTeamTriggerDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *ReadTeamTriggerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
