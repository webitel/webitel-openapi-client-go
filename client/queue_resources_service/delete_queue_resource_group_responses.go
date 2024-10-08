// Code generated by go-swagger; DO NOT EDIT.

package queue_resources_service

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

// DeleteQueueResourceGroupReader is a Reader for the DeleteQueueResourceGroup structure.
type DeleteQueueResourceGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteQueueResourceGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteQueueResourceGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteQueueResourceGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteQueueResourceGroupOK creates a DeleteQueueResourceGroupOK with default headers values
func NewDeleteQueueResourceGroupOK() *DeleteQueueResourceGroupOK {
	return &DeleteQueueResourceGroupOK{}
}

/*
DeleteQueueResourceGroupOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteQueueResourceGroupOK struct {
	Payload *models.EngineQueueResourceGroup
}

// IsSuccess returns true when this delete queue resource group Ok response has a 2xx status code
func (o *DeleteQueueResourceGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete queue resource group Ok response has a 3xx status code
func (o *DeleteQueueResourceGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete queue resource group Ok response has a 4xx status code
func (o *DeleteQueueResourceGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete queue resource group Ok response has a 5xx status code
func (o *DeleteQueueResourceGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete queue resource group Ok response a status code equal to that given
func (o *DeleteQueueResourceGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete queue resource group Ok response
func (o *DeleteQueueResourceGroupOK) Code() int {
	return 200
}

func (o *DeleteQueueResourceGroupOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/queues/{queue_id}/resource_groups/{id}][%d] deleteQueueResourceGroupOk %s", 200, payload)
}

func (o *DeleteQueueResourceGroupOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/queues/{queue_id}/resource_groups/{id}][%d] deleteQueueResourceGroupOk %s", 200, payload)
}

func (o *DeleteQueueResourceGroupOK) GetPayload() *models.EngineQueueResourceGroup {
	return o.Payload
}

func (o *DeleteQueueResourceGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineQueueResourceGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteQueueResourceGroupDefault creates a DeleteQueueResourceGroupDefault with default headers values
func NewDeleteQueueResourceGroupDefault(code int) *DeleteQueueResourceGroupDefault {
	return &DeleteQueueResourceGroupDefault{
		_statusCode: code,
	}
}

/*
DeleteQueueResourceGroupDefault describes a response with status code -1, with default header values.

Server error
*/
type DeleteQueueResourceGroupDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this delete queue resource group default response has a 2xx status code
func (o *DeleteQueueResourceGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete queue resource group default response has a 3xx status code
func (o *DeleteQueueResourceGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete queue resource group default response has a 4xx status code
func (o *DeleteQueueResourceGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete queue resource group default response has a 5xx status code
func (o *DeleteQueueResourceGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete queue resource group default response a status code equal to that given
func (o *DeleteQueueResourceGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete queue resource group default response
func (o *DeleteQueueResourceGroupDefault) Code() int {
	return o._statusCode
}

func (o *DeleteQueueResourceGroupDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/queues/{queue_id}/resource_groups/{id}][%d] DeleteQueueResourceGroup default %s", o._statusCode, payload)
}

func (o *DeleteQueueResourceGroupDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/queues/{queue_id}/resource_groups/{id}][%d] DeleteQueueResourceGroup default %s", o._statusCode, payload)
}

func (o *DeleteQueueResourceGroupDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *DeleteQueueResourceGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
