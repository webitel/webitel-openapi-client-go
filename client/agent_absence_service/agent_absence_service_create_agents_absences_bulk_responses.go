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

// AgentAbsenceServiceCreateAgentsAbsencesBulkReader is a Reader for the AgentAbsenceServiceCreateAgentsAbsencesBulk structure.
type AgentAbsenceServiceCreateAgentsAbsencesBulkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AgentAbsenceServiceCreateAgentsAbsencesBulkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAgentAbsenceServiceCreateAgentsAbsencesBulkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAgentAbsenceServiceCreateAgentsAbsencesBulkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAgentAbsenceServiceCreateAgentsAbsencesBulkOK creates a AgentAbsenceServiceCreateAgentsAbsencesBulkOK with default headers values
func NewAgentAbsenceServiceCreateAgentsAbsencesBulkOK() *AgentAbsenceServiceCreateAgentsAbsencesBulkOK {
	return &AgentAbsenceServiceCreateAgentsAbsencesBulkOK{}
}

/*
AgentAbsenceServiceCreateAgentsAbsencesBulkOK describes a response with status code 200, with default header values.

A successful response.
*/
type AgentAbsenceServiceCreateAgentsAbsencesBulkOK struct {
	Payload *models.WfmCreateAgentsAbsencesBulkResponse
}

// IsSuccess returns true when this agent absence service create agents absences bulk Ok response has a 2xx status code
func (o *AgentAbsenceServiceCreateAgentsAbsencesBulkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this agent absence service create agents absences bulk Ok response has a 3xx status code
func (o *AgentAbsenceServiceCreateAgentsAbsencesBulkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this agent absence service create agents absences bulk Ok response has a 4xx status code
func (o *AgentAbsenceServiceCreateAgentsAbsencesBulkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this agent absence service create agents absences bulk Ok response has a 5xx status code
func (o *AgentAbsenceServiceCreateAgentsAbsencesBulkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this agent absence service create agents absences bulk Ok response a status code equal to that given
func (o *AgentAbsenceServiceCreateAgentsAbsencesBulkOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the agent absence service create agents absences bulk Ok response
func (o *AgentAbsenceServiceCreateAgentsAbsencesBulkOK) Code() int {
	return 200
}

func (o *AgentAbsenceServiceCreateAgentsAbsencesBulkOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /wfm/agents/absences][%d] agentAbsenceServiceCreateAgentsAbsencesBulkOk %s", 200, payload)
}

func (o *AgentAbsenceServiceCreateAgentsAbsencesBulkOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /wfm/agents/absences][%d] agentAbsenceServiceCreateAgentsAbsencesBulkOk %s", 200, payload)
}

func (o *AgentAbsenceServiceCreateAgentsAbsencesBulkOK) GetPayload() *models.WfmCreateAgentsAbsencesBulkResponse {
	return o.Payload
}

func (o *AgentAbsenceServiceCreateAgentsAbsencesBulkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WfmCreateAgentsAbsencesBulkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAgentAbsenceServiceCreateAgentsAbsencesBulkDefault creates a AgentAbsenceServiceCreateAgentsAbsencesBulkDefault with default headers values
func NewAgentAbsenceServiceCreateAgentsAbsencesBulkDefault(code int) *AgentAbsenceServiceCreateAgentsAbsencesBulkDefault {
	return &AgentAbsenceServiceCreateAgentsAbsencesBulkDefault{
		_statusCode: code,
	}
}

/*
AgentAbsenceServiceCreateAgentsAbsencesBulkDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AgentAbsenceServiceCreateAgentsAbsencesBulkDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this agent absence service create agents absences bulk default response has a 2xx status code
func (o *AgentAbsenceServiceCreateAgentsAbsencesBulkDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this agent absence service create agents absences bulk default response has a 3xx status code
func (o *AgentAbsenceServiceCreateAgentsAbsencesBulkDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this agent absence service create agents absences bulk default response has a 4xx status code
func (o *AgentAbsenceServiceCreateAgentsAbsencesBulkDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this agent absence service create agents absences bulk default response has a 5xx status code
func (o *AgentAbsenceServiceCreateAgentsAbsencesBulkDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this agent absence service create agents absences bulk default response a status code equal to that given
func (o *AgentAbsenceServiceCreateAgentsAbsencesBulkDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the agent absence service create agents absences bulk default response
func (o *AgentAbsenceServiceCreateAgentsAbsencesBulkDefault) Code() int {
	return o._statusCode
}

func (o *AgentAbsenceServiceCreateAgentsAbsencesBulkDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /wfm/agents/absences][%d] AgentAbsenceService_CreateAgentsAbsencesBulk default %s", o._statusCode, payload)
}

func (o *AgentAbsenceServiceCreateAgentsAbsencesBulkDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /wfm/agents/absences][%d] AgentAbsenceService_CreateAgentsAbsencesBulk default %s", o._statusCode, payload)
}

func (o *AgentAbsenceServiceCreateAgentsAbsencesBulkDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *AgentAbsenceServiceCreateAgentsAbsencesBulkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
