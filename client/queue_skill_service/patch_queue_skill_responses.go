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

// PatchQueueSkillReader is a Reader for the PatchQueueSkill structure.
type PatchQueueSkillReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchQueueSkillReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchQueueSkillOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchQueueSkillDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchQueueSkillOK creates a PatchQueueSkillOK with default headers values
func NewPatchQueueSkillOK() *PatchQueueSkillOK {
	return &PatchQueueSkillOK{}
}

/*
PatchQueueSkillOK describes a response with status code 200, with default header values.

A successful response.
*/
type PatchQueueSkillOK struct {
	Payload *models.EngineQueueSkill
}

// IsSuccess returns true when this patch queue skill Ok response has a 2xx status code
func (o *PatchQueueSkillOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch queue skill Ok response has a 3xx status code
func (o *PatchQueueSkillOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch queue skill Ok response has a 4xx status code
func (o *PatchQueueSkillOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch queue skill Ok response has a 5xx status code
func (o *PatchQueueSkillOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch queue skill Ok response a status code equal to that given
func (o *PatchQueueSkillOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch queue skill Ok response
func (o *PatchQueueSkillOK) Code() int {
	return 200
}

func (o *PatchQueueSkillOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /call_center/queues/{queue_id}/skills/{id}][%d] patchQueueSkillOk %s", 200, payload)
}

func (o *PatchQueueSkillOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /call_center/queues/{queue_id}/skills/{id}][%d] patchQueueSkillOk %s", 200, payload)
}

func (o *PatchQueueSkillOK) GetPayload() *models.EngineQueueSkill {
	return o.Payload
}

func (o *PatchQueueSkillOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineQueueSkill)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchQueueSkillDefault creates a PatchQueueSkillDefault with default headers values
func NewPatchQueueSkillDefault(code int) *PatchQueueSkillDefault {
	return &PatchQueueSkillDefault{
		_statusCode: code,
	}
}

/*
PatchQueueSkillDefault describes a response with status code -1, with default header values.

Server error
*/
type PatchQueueSkillDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this patch queue skill default response has a 2xx status code
func (o *PatchQueueSkillDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch queue skill default response has a 3xx status code
func (o *PatchQueueSkillDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch queue skill default response has a 4xx status code
func (o *PatchQueueSkillDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch queue skill default response has a 5xx status code
func (o *PatchQueueSkillDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch queue skill default response a status code equal to that given
func (o *PatchQueueSkillDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the patch queue skill default response
func (o *PatchQueueSkillDefault) Code() int {
	return o._statusCode
}

func (o *PatchQueueSkillDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /call_center/queues/{queue_id}/skills/{id}][%d] PatchQueueSkill default %s", o._statusCode, payload)
}

func (o *PatchQueueSkillDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /call_center/queues/{queue_id}/skills/{id}][%d] PatchQueueSkill default %s", o._statusCode, payload)
}

func (o *PatchQueueSkillDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *PatchQueueSkillDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
