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

// AgentTodayStatisticsReader is a Reader for the AgentTodayStatistics structure.
type AgentTodayStatisticsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AgentTodayStatisticsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAgentTodayStatisticsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAgentTodayStatisticsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAgentTodayStatisticsOK creates a AgentTodayStatisticsOK with default headers values
func NewAgentTodayStatisticsOK() *AgentTodayStatisticsOK {
	return &AgentTodayStatisticsOK{}
}

/*
AgentTodayStatisticsOK describes a response with status code 200, with default header values.

A successful response.
*/
type AgentTodayStatisticsOK struct {
	Payload *models.EngineAgentTodayStatisticsResponse
}

// IsSuccess returns true when this agent today statistics Ok response has a 2xx status code
func (o *AgentTodayStatisticsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this agent today statistics Ok response has a 3xx status code
func (o *AgentTodayStatisticsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this agent today statistics Ok response has a 4xx status code
func (o *AgentTodayStatisticsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this agent today statistics Ok response has a 5xx status code
func (o *AgentTodayStatisticsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this agent today statistics Ok response a status code equal to that given
func (o *AgentTodayStatisticsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the agent today statistics Ok response
func (o *AgentTodayStatisticsOK) Code() int {
	return 200
}

func (o *AgentTodayStatisticsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/agents/{agent_id}/statistics/today][%d] agentTodayStatisticsOk %s", 200, payload)
}

func (o *AgentTodayStatisticsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/agents/{agent_id}/statistics/today][%d] agentTodayStatisticsOk %s", 200, payload)
}

func (o *AgentTodayStatisticsOK) GetPayload() *models.EngineAgentTodayStatisticsResponse {
	return o.Payload
}

func (o *AgentTodayStatisticsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAgentTodayStatisticsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAgentTodayStatisticsDefault creates a AgentTodayStatisticsDefault with default headers values
func NewAgentTodayStatisticsDefault(code int) *AgentTodayStatisticsDefault {
	return &AgentTodayStatisticsDefault{
		_statusCode: code,
	}
}

/*
AgentTodayStatisticsDefault describes a response with status code -1, with default header values.

Server error
*/
type AgentTodayStatisticsDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this agent today statistics default response has a 2xx status code
func (o *AgentTodayStatisticsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this agent today statistics default response has a 3xx status code
func (o *AgentTodayStatisticsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this agent today statistics default response has a 4xx status code
func (o *AgentTodayStatisticsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this agent today statistics default response has a 5xx status code
func (o *AgentTodayStatisticsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this agent today statistics default response a status code equal to that given
func (o *AgentTodayStatisticsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the agent today statistics default response
func (o *AgentTodayStatisticsDefault) Code() int {
	return o._statusCode
}

func (o *AgentTodayStatisticsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/agents/{agent_id}/statistics/today][%d] AgentTodayStatistics default %s", o._statusCode, payload)
}

func (o *AgentTodayStatisticsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/agents/{agent_id}/statistics/today][%d] AgentTodayStatistics default %s", o._statusCode, payload)
}

func (o *AgentTodayStatisticsDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *AgentTodayStatisticsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
