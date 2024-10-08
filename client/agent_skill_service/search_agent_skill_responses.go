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

// SearchAgentSkillReader is a Reader for the SearchAgentSkill structure.
type SearchAgentSkillReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchAgentSkillReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchAgentSkillOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSearchAgentSkillDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchAgentSkillOK creates a SearchAgentSkillOK with default headers values
func NewSearchAgentSkillOK() *SearchAgentSkillOK {
	return &SearchAgentSkillOK{}
}

/*
SearchAgentSkillOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchAgentSkillOK struct {
	Payload *models.EngineListAgentSkill
}

// IsSuccess returns true when this search agent skill Ok response has a 2xx status code
func (o *SearchAgentSkillOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search agent skill Ok response has a 3xx status code
func (o *SearchAgentSkillOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search agent skill Ok response has a 4xx status code
func (o *SearchAgentSkillOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search agent skill Ok response has a 5xx status code
func (o *SearchAgentSkillOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search agent skill Ok response a status code equal to that given
func (o *SearchAgentSkillOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search agent skill Ok response
func (o *SearchAgentSkillOK) Code() int {
	return 200
}

func (o *SearchAgentSkillOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/agents/{agent_id}/skills][%d] searchAgentSkillOk %s", 200, payload)
}

func (o *SearchAgentSkillOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/agents/{agent_id}/skills][%d] searchAgentSkillOk %s", 200, payload)
}

func (o *SearchAgentSkillOK) GetPayload() *models.EngineListAgentSkill {
	return o.Payload
}

func (o *SearchAgentSkillOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineListAgentSkill)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchAgentSkillDefault creates a SearchAgentSkillDefault with default headers values
func NewSearchAgentSkillDefault(code int) *SearchAgentSkillDefault {
	return &SearchAgentSkillDefault{
		_statusCode: code,
	}
}

/*
SearchAgentSkillDefault describes a response with status code -1, with default header values.

Server error
*/
type SearchAgentSkillDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this search agent skill default response has a 2xx status code
func (o *SearchAgentSkillDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this search agent skill default response has a 3xx status code
func (o *SearchAgentSkillDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this search agent skill default response has a 4xx status code
func (o *SearchAgentSkillDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this search agent skill default response has a 5xx status code
func (o *SearchAgentSkillDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this search agent skill default response a status code equal to that given
func (o *SearchAgentSkillDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the search agent skill default response
func (o *SearchAgentSkillDefault) Code() int {
	return o._statusCode
}

func (o *SearchAgentSkillDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/agents/{agent_id}/skills][%d] SearchAgentSkill default %s", o._statusCode, payload)
}

func (o *SearchAgentSkillDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/agents/{agent_id}/skills][%d] SearchAgentSkill default %s", o._statusCode, payload)
}

func (o *SearchAgentSkillDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *SearchAgentSkillDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
