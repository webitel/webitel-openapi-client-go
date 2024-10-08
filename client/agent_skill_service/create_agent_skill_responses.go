// Code generated by go-swagger; DO NOT EDIT.

package agent_skill_service

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

// CreateAgentSkillReader is a Reader for the CreateAgentSkill structure.
type CreateAgentSkillReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAgentSkillReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAgentSkillOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateAgentSkillDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateAgentSkillOK creates a CreateAgentSkillOK with default headers values
func NewCreateAgentSkillOK() *CreateAgentSkillOK {
	return &CreateAgentSkillOK{}
}

/*
CreateAgentSkillOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateAgentSkillOK struct {
	Payload *models.EngineAgentSkill
}

// IsSuccess returns true when this create agent skill Ok response has a 2xx status code
func (o *CreateAgentSkillOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create agent skill Ok response has a 3xx status code
func (o *CreateAgentSkillOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create agent skill Ok response has a 4xx status code
func (o *CreateAgentSkillOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create agent skill Ok response has a 5xx status code
func (o *CreateAgentSkillOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create agent skill Ok response a status code equal to that given
func (o *CreateAgentSkillOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create agent skill Ok response
func (o *CreateAgentSkillOK) Code() int {
	return 200
}

func (o *CreateAgentSkillOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/agents/{agent_id}/skills][%d] createAgentSkillOk %s", 200, payload)
}

func (o *CreateAgentSkillOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/agents/{agent_id}/skills][%d] createAgentSkillOk %s", 200, payload)
}

func (o *CreateAgentSkillOK) GetPayload() *models.EngineAgentSkill {
	return o.Payload
}

func (o *CreateAgentSkillOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAgentSkill)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAgentSkillDefault creates a CreateAgentSkillDefault with default headers values
func NewCreateAgentSkillDefault(code int) *CreateAgentSkillDefault {
	return &CreateAgentSkillDefault{
		_statusCode: code,
	}
}

/*
CreateAgentSkillDefault describes a response with status code -1, with default header values.

Server error
*/
type CreateAgentSkillDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this create agent skill default response has a 2xx status code
func (o *CreateAgentSkillDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create agent skill default response has a 3xx status code
func (o *CreateAgentSkillDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create agent skill default response has a 4xx status code
func (o *CreateAgentSkillDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create agent skill default response has a 5xx status code
func (o *CreateAgentSkillDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create agent skill default response a status code equal to that given
func (o *CreateAgentSkillDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create agent skill default response
func (o *CreateAgentSkillDefault) Code() int {
	return o._statusCode
}

func (o *CreateAgentSkillDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/agents/{agent_id}/skills][%d] CreateAgentSkill default %s", o._statusCode, payload)
}

func (o *CreateAgentSkillDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/agents/{agent_id}/skills][%d] CreateAgentSkill default %s", o._statusCode, payload)
}

func (o *CreateAgentSkillDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *CreateAgentSkillDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
