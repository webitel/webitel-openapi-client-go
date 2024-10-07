// Code generated by go-swagger; DO NOT EDIT.

package skill_service

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

// DeleteSkillAgentReader is a Reader for the DeleteSkillAgent structure.
type DeleteSkillAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSkillAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSkillAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteSkillAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSkillAgentOK creates a DeleteSkillAgentOK with default headers values
func NewDeleteSkillAgentOK() *DeleteSkillAgentOK {
	return &DeleteSkillAgentOK{}
}

/*
DeleteSkillAgentOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteSkillAgentOK struct {
	Payload *models.EngineDeleteSkillAgentResponse
}

// IsSuccess returns true when this delete skill agent Ok response has a 2xx status code
func (o *DeleteSkillAgentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete skill agent Ok response has a 3xx status code
func (o *DeleteSkillAgentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete skill agent Ok response has a 4xx status code
func (o *DeleteSkillAgentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete skill agent Ok response has a 5xx status code
func (o *DeleteSkillAgentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete skill agent Ok response a status code equal to that given
func (o *DeleteSkillAgentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete skill agent Ok response
func (o *DeleteSkillAgentOK) Code() int {
	return 200
}

func (o *DeleteSkillAgentOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/skills/{skill_id}/agents][%d] deleteSkillAgentOk %s", 200, payload)
}

func (o *DeleteSkillAgentOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/skills/{skill_id}/agents][%d] deleteSkillAgentOk %s", 200, payload)
}

func (o *DeleteSkillAgentOK) GetPayload() *models.EngineDeleteSkillAgentResponse {
	return o.Payload
}

func (o *DeleteSkillAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineDeleteSkillAgentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSkillAgentDefault creates a DeleteSkillAgentDefault with default headers values
func NewDeleteSkillAgentDefault(code int) *DeleteSkillAgentDefault {
	return &DeleteSkillAgentDefault{
		_statusCode: code,
	}
}

/*
DeleteSkillAgentDefault describes a response with status code -1, with default header values.

Server error
*/
type DeleteSkillAgentDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this delete skill agent default response has a 2xx status code
func (o *DeleteSkillAgentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete skill agent default response has a 3xx status code
func (o *DeleteSkillAgentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete skill agent default response has a 4xx status code
func (o *DeleteSkillAgentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete skill agent default response has a 5xx status code
func (o *DeleteSkillAgentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete skill agent default response a status code equal to that given
func (o *DeleteSkillAgentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete skill agent default response
func (o *DeleteSkillAgentDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSkillAgentDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/skills/{skill_id}/agents][%d] DeleteSkillAgent default %s", o._statusCode, payload)
}

func (o *DeleteSkillAgentDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/skills/{skill_id}/agents][%d] DeleteSkillAgent default %s", o._statusCode, payload)
}

func (o *DeleteSkillAgentDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *DeleteSkillAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
