// Code generated by go-swagger; DO NOT EDIT.

package agent_team_service

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

// DeleteAgentTeamReader is a Reader for the DeleteAgentTeam structure.
type DeleteAgentTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAgentTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAgentTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteAgentTeamDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAgentTeamOK creates a DeleteAgentTeamOK with default headers values
func NewDeleteAgentTeamOK() *DeleteAgentTeamOK {
	return &DeleteAgentTeamOK{}
}

/*
DeleteAgentTeamOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteAgentTeamOK struct {
	Payload *models.EngineAgentTeam
}

// IsSuccess returns true when this delete agent team Ok response has a 2xx status code
func (o *DeleteAgentTeamOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete agent team Ok response has a 3xx status code
func (o *DeleteAgentTeamOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete agent team Ok response has a 4xx status code
func (o *DeleteAgentTeamOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete agent team Ok response has a 5xx status code
func (o *DeleteAgentTeamOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete agent team Ok response a status code equal to that given
func (o *DeleteAgentTeamOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete agent team Ok response
func (o *DeleteAgentTeamOK) Code() int {
	return 200
}

func (o *DeleteAgentTeamOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/teams/{id}][%d] deleteAgentTeamOk %s", 200, payload)
}

func (o *DeleteAgentTeamOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/teams/{id}][%d] deleteAgentTeamOk %s", 200, payload)
}

func (o *DeleteAgentTeamOK) GetPayload() *models.EngineAgentTeam {
	return o.Payload
}

func (o *DeleteAgentTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAgentTeam)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAgentTeamDefault creates a DeleteAgentTeamDefault with default headers values
func NewDeleteAgentTeamDefault(code int) *DeleteAgentTeamDefault {
	return &DeleteAgentTeamDefault{
		_statusCode: code,
	}
}

/*
DeleteAgentTeamDefault describes a response with status code -1, with default header values.

Server error
*/
type DeleteAgentTeamDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this delete agent team default response has a 2xx status code
func (o *DeleteAgentTeamDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete agent team default response has a 3xx status code
func (o *DeleteAgentTeamDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete agent team default response has a 4xx status code
func (o *DeleteAgentTeamDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete agent team default response has a 5xx status code
func (o *DeleteAgentTeamDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete agent team default response a status code equal to that given
func (o *DeleteAgentTeamDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete agent team default response
func (o *DeleteAgentTeamDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAgentTeamDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/teams/{id}][%d] DeleteAgentTeam default %s", o._statusCode, payload)
}

func (o *DeleteAgentTeamDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/teams/{id}][%d] DeleteAgentTeam default %s", o._statusCode, payload)
}

func (o *DeleteAgentTeamDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *DeleteAgentTeamDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
