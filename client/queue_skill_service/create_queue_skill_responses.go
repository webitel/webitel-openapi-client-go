// Code generated by go-swagger; DO NOT EDIT.

package queue_skill_service

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

// CreateQueueSkillReader is a Reader for the CreateQueueSkill structure.
type CreateQueueSkillReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateQueueSkillReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateQueueSkillOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateQueueSkillDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateQueueSkillOK creates a CreateQueueSkillOK with default headers values
func NewCreateQueueSkillOK() *CreateQueueSkillOK {
	return &CreateQueueSkillOK{}
}

/*
CreateQueueSkillOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateQueueSkillOK struct {
	Payload *models.EngineQueueSkill
}

// IsSuccess returns true when this create queue skill Ok response has a 2xx status code
func (o *CreateQueueSkillOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create queue skill Ok response has a 3xx status code
func (o *CreateQueueSkillOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create queue skill Ok response has a 4xx status code
func (o *CreateQueueSkillOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create queue skill Ok response has a 5xx status code
func (o *CreateQueueSkillOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create queue skill Ok response a status code equal to that given
func (o *CreateQueueSkillOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create queue skill Ok response
func (o *CreateQueueSkillOK) Code() int {
	return 200
}

func (o *CreateQueueSkillOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/queues/{queue_id}/skills][%d] createQueueSkillOk %s", 200, payload)
}

func (o *CreateQueueSkillOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/queues/{queue_id}/skills][%d] createQueueSkillOk %s", 200, payload)
}

func (o *CreateQueueSkillOK) GetPayload() *models.EngineQueueSkill {
	return o.Payload
}

func (o *CreateQueueSkillOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineQueueSkill)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQueueSkillDefault creates a CreateQueueSkillDefault with default headers values
func NewCreateQueueSkillDefault(code int) *CreateQueueSkillDefault {
	return &CreateQueueSkillDefault{
		_statusCode: code,
	}
}

/*
CreateQueueSkillDefault describes a response with status code -1, with default header values.

Server error
*/
type CreateQueueSkillDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this create queue skill default response has a 2xx status code
func (o *CreateQueueSkillDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create queue skill default response has a 3xx status code
func (o *CreateQueueSkillDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create queue skill default response has a 4xx status code
func (o *CreateQueueSkillDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create queue skill default response has a 5xx status code
func (o *CreateQueueSkillDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create queue skill default response a status code equal to that given
func (o *CreateQueueSkillDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create queue skill default response
func (o *CreateQueueSkillDefault) Code() int {
	return o._statusCode
}

func (o *CreateQueueSkillDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/queues/{queue_id}/skills][%d] CreateQueueSkill default %s", o._statusCode, payload)
}

func (o *CreateQueueSkillDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/queues/{queue_id}/skills][%d] CreateQueueSkill default %s", o._statusCode, payload)
}

func (o *CreateQueueSkillDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *CreateQueueSkillDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
