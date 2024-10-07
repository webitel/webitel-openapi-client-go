// Code generated by go-swagger; DO NOT EDIT.

package routing_chat_plan_service

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

// CreateChatPlanReader is a Reader for the CreateChatPlan structure.
type CreateChatPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateChatPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateChatPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateChatPlanDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateChatPlanOK creates a CreateChatPlanOK with default headers values
func NewCreateChatPlanOK() *CreateChatPlanOK {
	return &CreateChatPlanOK{}
}

/*
CreateChatPlanOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateChatPlanOK struct {
	Payload *models.EngineChatPlan
}

// IsSuccess returns true when this create chat plan Ok response has a 2xx status code
func (o *CreateChatPlanOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create chat plan Ok response has a 3xx status code
func (o *CreateChatPlanOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create chat plan Ok response has a 4xx status code
func (o *CreateChatPlanOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create chat plan Ok response has a 5xx status code
func (o *CreateChatPlanOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create chat plan Ok response a status code equal to that given
func (o *CreateChatPlanOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create chat plan Ok response
func (o *CreateChatPlanOK) Code() int {
	return 200
}

func (o *CreateChatPlanOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /routing/outbound/chat][%d] createChatPlanOk %s", 200, payload)
}

func (o *CreateChatPlanOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /routing/outbound/chat][%d] createChatPlanOk %s", 200, payload)
}

func (o *CreateChatPlanOK) GetPayload() *models.EngineChatPlan {
	return o.Payload
}

func (o *CreateChatPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineChatPlan)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateChatPlanDefault creates a CreateChatPlanDefault with default headers values
func NewCreateChatPlanDefault(code int) *CreateChatPlanDefault {
	return &CreateChatPlanDefault{
		_statusCode: code,
	}
}

/*
CreateChatPlanDefault describes a response with status code -1, with default header values.

Server error
*/
type CreateChatPlanDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this create chat plan default response has a 2xx status code
func (o *CreateChatPlanDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create chat plan default response has a 3xx status code
func (o *CreateChatPlanDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create chat plan default response has a 4xx status code
func (o *CreateChatPlanDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create chat plan default response has a 5xx status code
func (o *CreateChatPlanDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create chat plan default response a status code equal to that given
func (o *CreateChatPlanDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create chat plan default response
func (o *CreateChatPlanDefault) Code() int {
	return o._statusCode
}

func (o *CreateChatPlanDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /routing/outbound/chat][%d] CreateChatPlan default %s", o._statusCode, payload)
}

func (o *CreateChatPlanDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /routing/outbound/chat][%d] CreateChatPlan default %s", o._statusCode, payload)
}

func (o *CreateChatPlanDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *CreateChatPlanDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
