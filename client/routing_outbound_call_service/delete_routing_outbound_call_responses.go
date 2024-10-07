// Code generated by go-swagger; DO NOT EDIT.

package routing_outbound_call_service

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

// DeleteRoutingOutboundCallReader is a Reader for the DeleteRoutingOutboundCall structure.
type DeleteRoutingOutboundCallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoutingOutboundCallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRoutingOutboundCallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteRoutingOutboundCallDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteRoutingOutboundCallOK creates a DeleteRoutingOutboundCallOK with default headers values
func NewDeleteRoutingOutboundCallOK() *DeleteRoutingOutboundCallOK {
	return &DeleteRoutingOutboundCallOK{}
}

/*
DeleteRoutingOutboundCallOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteRoutingOutboundCallOK struct {
	Payload *models.EngineRoutingOutboundCall
}

// IsSuccess returns true when this delete routing outbound call Ok response has a 2xx status code
func (o *DeleteRoutingOutboundCallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete routing outbound call Ok response has a 3xx status code
func (o *DeleteRoutingOutboundCallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing outbound call Ok response has a 4xx status code
func (o *DeleteRoutingOutboundCallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing outbound call Ok response has a 5xx status code
func (o *DeleteRoutingOutboundCallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing outbound call Ok response a status code equal to that given
func (o *DeleteRoutingOutboundCallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete routing outbound call Ok response
func (o *DeleteRoutingOutboundCallOK) Code() int {
	return 200
}

func (o *DeleteRoutingOutboundCallOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /routing/outbound/calls/{id}][%d] deleteRoutingOutboundCallOk %s", 200, payload)
}

func (o *DeleteRoutingOutboundCallOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /routing/outbound/calls/{id}][%d] deleteRoutingOutboundCallOk %s", 200, payload)
}

func (o *DeleteRoutingOutboundCallOK) GetPayload() *models.EngineRoutingOutboundCall {
	return o.Payload
}

func (o *DeleteRoutingOutboundCallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineRoutingOutboundCall)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingOutboundCallDefault creates a DeleteRoutingOutboundCallDefault with default headers values
func NewDeleteRoutingOutboundCallDefault(code int) *DeleteRoutingOutboundCallDefault {
	return &DeleteRoutingOutboundCallDefault{
		_statusCode: code,
	}
}

/*
DeleteRoutingOutboundCallDefault describes a response with status code -1, with default header values.

Server error
*/
type DeleteRoutingOutboundCallDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this delete routing outbound call default response has a 2xx status code
func (o *DeleteRoutingOutboundCallDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete routing outbound call default response has a 3xx status code
func (o *DeleteRoutingOutboundCallDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete routing outbound call default response has a 4xx status code
func (o *DeleteRoutingOutboundCallDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete routing outbound call default response has a 5xx status code
func (o *DeleteRoutingOutboundCallDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete routing outbound call default response a status code equal to that given
func (o *DeleteRoutingOutboundCallDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete routing outbound call default response
func (o *DeleteRoutingOutboundCallDefault) Code() int {
	return o._statusCode
}

func (o *DeleteRoutingOutboundCallDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /routing/outbound/calls/{id}][%d] DeleteRoutingOutboundCall default %s", o._statusCode, payload)
}

func (o *DeleteRoutingOutboundCallDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /routing/outbound/calls/{id}][%d] DeleteRoutingOutboundCall default %s", o._statusCode, payload)
}

func (o *DeleteRoutingOutboundCallDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *DeleteRoutingOutboundCallDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
