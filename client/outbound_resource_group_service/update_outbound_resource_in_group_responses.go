// Code generated by go-swagger; DO NOT EDIT.

package outbound_resource_group_service

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

// UpdateOutboundResourceInGroupReader is a Reader for the UpdateOutboundResourceInGroup structure.
type UpdateOutboundResourceInGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOutboundResourceInGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOutboundResourceInGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateOutboundResourceInGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateOutboundResourceInGroupOK creates a UpdateOutboundResourceInGroupOK with default headers values
func NewUpdateOutboundResourceInGroupOK() *UpdateOutboundResourceInGroupOK {
	return &UpdateOutboundResourceInGroupOK{}
}

/*
UpdateOutboundResourceInGroupOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateOutboundResourceInGroupOK struct {
	Payload *models.EngineOutboundResourceInGroup
}

// IsSuccess returns true when this update outbound resource in group Ok response has a 2xx status code
func (o *UpdateOutboundResourceInGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update outbound resource in group Ok response has a 3xx status code
func (o *UpdateOutboundResourceInGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update outbound resource in group Ok response has a 4xx status code
func (o *UpdateOutboundResourceInGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update outbound resource in group Ok response has a 5xx status code
func (o *UpdateOutboundResourceInGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update outbound resource in group Ok response a status code equal to that given
func (o *UpdateOutboundResourceInGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update outbound resource in group Ok response
func (o *UpdateOutboundResourceInGroupOK) Code() int {
	return 200
}

func (o *UpdateOutboundResourceInGroupOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/resource_group/{group_id}/resource/{id}][%d] updateOutboundResourceInGroupOk %s", 200, payload)
}

func (o *UpdateOutboundResourceInGroupOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/resource_group/{group_id}/resource/{id}][%d] updateOutboundResourceInGroupOk %s", 200, payload)
}

func (o *UpdateOutboundResourceInGroupOK) GetPayload() *models.EngineOutboundResourceInGroup {
	return o.Payload
}

func (o *UpdateOutboundResourceInGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineOutboundResourceInGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOutboundResourceInGroupDefault creates a UpdateOutboundResourceInGroupDefault with default headers values
func NewUpdateOutboundResourceInGroupDefault(code int) *UpdateOutboundResourceInGroupDefault {
	return &UpdateOutboundResourceInGroupDefault{
		_statusCode: code,
	}
}

/*
UpdateOutboundResourceInGroupDefault describes a response with status code -1, with default header values.

Server error
*/
type UpdateOutboundResourceInGroupDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this update outbound resource in group default response has a 2xx status code
func (o *UpdateOutboundResourceInGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update outbound resource in group default response has a 3xx status code
func (o *UpdateOutboundResourceInGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update outbound resource in group default response has a 4xx status code
func (o *UpdateOutboundResourceInGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update outbound resource in group default response has a 5xx status code
func (o *UpdateOutboundResourceInGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update outbound resource in group default response a status code equal to that given
func (o *UpdateOutboundResourceInGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update outbound resource in group default response
func (o *UpdateOutboundResourceInGroupDefault) Code() int {
	return o._statusCode
}

func (o *UpdateOutboundResourceInGroupDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/resource_group/{group_id}/resource/{id}][%d] UpdateOutboundResourceInGroup default %s", o._statusCode, payload)
}

func (o *UpdateOutboundResourceInGroupDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/resource_group/{group_id}/resource/{id}][%d] UpdateOutboundResourceInGroup default %s", o._statusCode, payload)
}

func (o *UpdateOutboundResourceInGroupDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *UpdateOutboundResourceInGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
