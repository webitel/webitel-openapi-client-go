// Code generated by go-swagger; DO NOT EDIT.

package agent_service

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

// CreateAgentReader is a Reader for the CreateAgent structure.
type CreateAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateAgentOK creates a CreateAgentOK with default headers values
func NewCreateAgentOK() *CreateAgentOK {
	return &CreateAgentOK{}
}

/*
CreateAgentOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateAgentOK struct {
	Payload *models.EngineAgent
}

// IsSuccess returns true when this create agent Ok response has a 2xx status code
func (o *CreateAgentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create agent Ok response has a 3xx status code
func (o *CreateAgentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create agent Ok response has a 4xx status code
func (o *CreateAgentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create agent Ok response has a 5xx status code
func (o *CreateAgentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create agent Ok response a status code equal to that given
func (o *CreateAgentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create agent Ok response
func (o *CreateAgentOK) Code() int {
	return 200
}

func (o *CreateAgentOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/agents][%d] createAgentOk %s", 200, payload)
}

func (o *CreateAgentOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/agents][%d] createAgentOk %s", 200, payload)
}

func (o *CreateAgentOK) GetPayload() *models.EngineAgent {
	return o.Payload
}

func (o *CreateAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAgent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAgentDefault creates a CreateAgentDefault with default headers values
func NewCreateAgentDefault(code int) *CreateAgentDefault {
	return &CreateAgentDefault{
		_statusCode: code,
	}
}

/*
CreateAgentDefault describes a response with status code -1, with default header values.

Server error
*/
type CreateAgentDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this create agent default response has a 2xx status code
func (o *CreateAgentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create agent default response has a 3xx status code
func (o *CreateAgentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create agent default response has a 4xx status code
func (o *CreateAgentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create agent default response has a 5xx status code
func (o *CreateAgentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create agent default response a status code equal to that given
func (o *CreateAgentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create agent default response
func (o *CreateAgentDefault) Code() int {
	return o._statusCode
}

func (o *CreateAgentDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/agents][%d] CreateAgent default %s", o._statusCode, payload)
}

func (o *CreateAgentDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/agents][%d] CreateAgent default %s", o._statusCode, payload)
}

func (o *CreateAgentDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *CreateAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
