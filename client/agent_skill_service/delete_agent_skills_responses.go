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

// DeleteAgentSkillsReader is a Reader for the DeleteAgentSkills structure.
type DeleteAgentSkillsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAgentSkillsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAgentSkillsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteAgentSkillsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAgentSkillsOK creates a DeleteAgentSkillsOK with default headers values
func NewDeleteAgentSkillsOK() *DeleteAgentSkillsOK {
	return &DeleteAgentSkillsOK{}
}

/*
DeleteAgentSkillsOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteAgentSkillsOK struct {
	Payload *models.EngineListAgentSkill
}

// IsSuccess returns true when this delete agent skills Ok response has a 2xx status code
func (o *DeleteAgentSkillsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete agent skills Ok response has a 3xx status code
func (o *DeleteAgentSkillsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete agent skills Ok response has a 4xx status code
func (o *DeleteAgentSkillsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete agent skills Ok response has a 5xx status code
func (o *DeleteAgentSkillsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete agent skills Ok response a status code equal to that given
func (o *DeleteAgentSkillsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete agent skills Ok response
func (o *DeleteAgentSkillsOK) Code() int {
	return 200
}

func (o *DeleteAgentSkillsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/agents/{agent_id}/skills][%d] deleteAgentSkillsOk %s", 200, payload)
}

func (o *DeleteAgentSkillsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/agents/{agent_id}/skills][%d] deleteAgentSkillsOk %s", 200, payload)
}

func (o *DeleteAgentSkillsOK) GetPayload() *models.EngineListAgentSkill {
	return o.Payload
}

func (o *DeleteAgentSkillsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineListAgentSkill)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAgentSkillsDefault creates a DeleteAgentSkillsDefault with default headers values
func NewDeleteAgentSkillsDefault(code int) *DeleteAgentSkillsDefault {
	return &DeleteAgentSkillsDefault{
		_statusCode: code,
	}
}

/*
DeleteAgentSkillsDefault describes a response with status code -1, with default header values.

Server error
*/
type DeleteAgentSkillsDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this delete agent skills default response has a 2xx status code
func (o *DeleteAgentSkillsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete agent skills default response has a 3xx status code
func (o *DeleteAgentSkillsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete agent skills default response has a 4xx status code
func (o *DeleteAgentSkillsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete agent skills default response has a 5xx status code
func (o *DeleteAgentSkillsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete agent skills default response a status code equal to that given
func (o *DeleteAgentSkillsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete agent skills default response
func (o *DeleteAgentSkillsDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAgentSkillsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/agents/{agent_id}/skills][%d] DeleteAgentSkills default %s", o._statusCode, payload)
}

func (o *DeleteAgentSkillsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/agents/{agent_id}/skills][%d] DeleteAgentSkills default %s", o._statusCode, payload)
}

func (o *DeleteAgentSkillsDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *DeleteAgentSkillsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
