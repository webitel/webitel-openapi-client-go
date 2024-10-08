// Code generated by go-swagger; DO NOT EDIT.

package agent_absence_service

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

// AgentAbsenceServiceUpdateAgentAbsenceReader is a Reader for the AgentAbsenceServiceUpdateAgentAbsence structure.
type AgentAbsenceServiceUpdateAgentAbsenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AgentAbsenceServiceUpdateAgentAbsenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAgentAbsenceServiceUpdateAgentAbsenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAgentAbsenceServiceUpdateAgentAbsenceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAgentAbsenceServiceUpdateAgentAbsenceOK creates a AgentAbsenceServiceUpdateAgentAbsenceOK with default headers values
func NewAgentAbsenceServiceUpdateAgentAbsenceOK() *AgentAbsenceServiceUpdateAgentAbsenceOK {
	return &AgentAbsenceServiceUpdateAgentAbsenceOK{}
}

/*
AgentAbsenceServiceUpdateAgentAbsenceOK describes a response with status code 200, with default header values.

A successful response.
*/
type AgentAbsenceServiceUpdateAgentAbsenceOK struct {
	Payload *models.WfmUpdateAgentAbsenceResponse
}

// IsSuccess returns true when this agent absence service update agent absence Ok response has a 2xx status code
func (o *AgentAbsenceServiceUpdateAgentAbsenceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this agent absence service update agent absence Ok response has a 3xx status code
func (o *AgentAbsenceServiceUpdateAgentAbsenceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this agent absence service update agent absence Ok response has a 4xx status code
func (o *AgentAbsenceServiceUpdateAgentAbsenceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this agent absence service update agent absence Ok response has a 5xx status code
func (o *AgentAbsenceServiceUpdateAgentAbsenceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this agent absence service update agent absence Ok response a status code equal to that given
func (o *AgentAbsenceServiceUpdateAgentAbsenceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the agent absence service update agent absence Ok response
func (o *AgentAbsenceServiceUpdateAgentAbsenceOK) Code() int {
	return 200
}

func (o *AgentAbsenceServiceUpdateAgentAbsenceOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /wfm/agents/{item.agent.id}/absences/{item.absence.id}][%d] agentAbsenceServiceUpdateAgentAbsenceOk %s", 200, payload)
}

func (o *AgentAbsenceServiceUpdateAgentAbsenceOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /wfm/agents/{item.agent.id}/absences/{item.absence.id}][%d] agentAbsenceServiceUpdateAgentAbsenceOk %s", 200, payload)
}

func (o *AgentAbsenceServiceUpdateAgentAbsenceOK) GetPayload() *models.WfmUpdateAgentAbsenceResponse {
	return o.Payload
}

func (o *AgentAbsenceServiceUpdateAgentAbsenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WfmUpdateAgentAbsenceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAgentAbsenceServiceUpdateAgentAbsenceDefault creates a AgentAbsenceServiceUpdateAgentAbsenceDefault with default headers values
func NewAgentAbsenceServiceUpdateAgentAbsenceDefault(code int) *AgentAbsenceServiceUpdateAgentAbsenceDefault {
	return &AgentAbsenceServiceUpdateAgentAbsenceDefault{
		_statusCode: code,
	}
}

/*
AgentAbsenceServiceUpdateAgentAbsenceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AgentAbsenceServiceUpdateAgentAbsenceDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this agent absence service update agent absence default response has a 2xx status code
func (o *AgentAbsenceServiceUpdateAgentAbsenceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this agent absence service update agent absence default response has a 3xx status code
func (o *AgentAbsenceServiceUpdateAgentAbsenceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this agent absence service update agent absence default response has a 4xx status code
func (o *AgentAbsenceServiceUpdateAgentAbsenceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this agent absence service update agent absence default response has a 5xx status code
func (o *AgentAbsenceServiceUpdateAgentAbsenceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this agent absence service update agent absence default response a status code equal to that given
func (o *AgentAbsenceServiceUpdateAgentAbsenceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the agent absence service update agent absence default response
func (o *AgentAbsenceServiceUpdateAgentAbsenceDefault) Code() int {
	return o._statusCode
}

func (o *AgentAbsenceServiceUpdateAgentAbsenceDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /wfm/agents/{item.agent.id}/absences/{item.absence.id}][%d] AgentAbsenceService_UpdateAgentAbsence default %s", o._statusCode, payload)
}

func (o *AgentAbsenceServiceUpdateAgentAbsenceDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /wfm/agents/{item.agent.id}/absences/{item.absence.id}][%d] AgentAbsenceService_UpdateAgentAbsence default %s", o._statusCode, payload)
}

func (o *AgentAbsenceServiceUpdateAgentAbsenceDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *AgentAbsenceServiceUpdateAgentAbsenceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
