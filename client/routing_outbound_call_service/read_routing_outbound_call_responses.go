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

// ReadRoutingOutboundCallReader is a Reader for the ReadRoutingOutboundCall structure.
type ReadRoutingOutboundCallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadRoutingOutboundCallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadRoutingOutboundCallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewReadRoutingOutboundCallDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadRoutingOutboundCallOK creates a ReadRoutingOutboundCallOK with default headers values
func NewReadRoutingOutboundCallOK() *ReadRoutingOutboundCallOK {
	return &ReadRoutingOutboundCallOK{}
}

/*
ReadRoutingOutboundCallOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadRoutingOutboundCallOK struct {
	Payload *models.EngineRoutingOutboundCall
}

// IsSuccess returns true when this read routing outbound call Ok response has a 2xx status code
func (o *ReadRoutingOutboundCallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read routing outbound call Ok response has a 3xx status code
func (o *ReadRoutingOutboundCallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read routing outbound call Ok response has a 4xx status code
func (o *ReadRoutingOutboundCallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read routing outbound call Ok response has a 5xx status code
func (o *ReadRoutingOutboundCallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read routing outbound call Ok response a status code equal to that given
func (o *ReadRoutingOutboundCallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read routing outbound call Ok response
func (o *ReadRoutingOutboundCallOK) Code() int {
	return 200
}

func (o *ReadRoutingOutboundCallOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /routing/outbound/calls/{id}][%d] readRoutingOutboundCallOk %s", 200, payload)
}

func (o *ReadRoutingOutboundCallOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /routing/outbound/calls/{id}][%d] readRoutingOutboundCallOk %s", 200, payload)
}

func (o *ReadRoutingOutboundCallOK) GetPayload() *models.EngineRoutingOutboundCall {
	return o.Payload
}

func (o *ReadRoutingOutboundCallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineRoutingOutboundCall)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadRoutingOutboundCallDefault creates a ReadRoutingOutboundCallDefault with default headers values
func NewReadRoutingOutboundCallDefault(code int) *ReadRoutingOutboundCallDefault {
	return &ReadRoutingOutboundCallDefault{
		_statusCode: code,
	}
}

/*
ReadRoutingOutboundCallDefault describes a response with status code -1, with default header values.

Server error
*/
type ReadRoutingOutboundCallDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this read routing outbound call default response has a 2xx status code
func (o *ReadRoutingOutboundCallDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read routing outbound call default response has a 3xx status code
func (o *ReadRoutingOutboundCallDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read routing outbound call default response has a 4xx status code
func (o *ReadRoutingOutboundCallDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read routing outbound call default response has a 5xx status code
func (o *ReadRoutingOutboundCallDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read routing outbound call default response a status code equal to that given
func (o *ReadRoutingOutboundCallDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read routing outbound call default response
func (o *ReadRoutingOutboundCallDefault) Code() int {
	return o._statusCode
}

func (o *ReadRoutingOutboundCallDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /routing/outbound/calls/{id}][%d] ReadRoutingOutboundCall default %s", o._statusCode, payload)
}

func (o *ReadRoutingOutboundCallDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /routing/outbound/calls/{id}][%d] ReadRoutingOutboundCall default %s", o._statusCode, payload)
}

func (o *ReadRoutingOutboundCallDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *ReadRoutingOutboundCallDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
