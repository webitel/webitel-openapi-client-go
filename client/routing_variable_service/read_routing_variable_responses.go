// Code generated by go-swagger; DO NOT EDIT.

package routing_variable_service

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

// ReadRoutingVariableReader is a Reader for the ReadRoutingVariable structure.
type ReadRoutingVariableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadRoutingVariableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadRoutingVariableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewReadRoutingVariableDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadRoutingVariableOK creates a ReadRoutingVariableOK with default headers values
func NewReadRoutingVariableOK() *ReadRoutingVariableOK {
	return &ReadRoutingVariableOK{}
}

/*
ReadRoutingVariableOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadRoutingVariableOK struct {
	Payload *models.EngineRoutingVariable
}

// IsSuccess returns true when this read routing variable Ok response has a 2xx status code
func (o *ReadRoutingVariableOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read routing variable Ok response has a 3xx status code
func (o *ReadRoutingVariableOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read routing variable Ok response has a 4xx status code
func (o *ReadRoutingVariableOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read routing variable Ok response has a 5xx status code
func (o *ReadRoutingVariableOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read routing variable Ok response a status code equal to that given
func (o *ReadRoutingVariableOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read routing variable Ok response
func (o *ReadRoutingVariableOK) Code() int {
	return 200
}

func (o *ReadRoutingVariableOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /routing/variables/{id}][%d] readRoutingVariableOk %s", 200, payload)
}

func (o *ReadRoutingVariableOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /routing/variables/{id}][%d] readRoutingVariableOk %s", 200, payload)
}

func (o *ReadRoutingVariableOK) GetPayload() *models.EngineRoutingVariable {
	return o.Payload
}

func (o *ReadRoutingVariableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineRoutingVariable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadRoutingVariableDefault creates a ReadRoutingVariableDefault with default headers values
func NewReadRoutingVariableDefault(code int) *ReadRoutingVariableDefault {
	return &ReadRoutingVariableDefault{
		_statusCode: code,
	}
}

/*
ReadRoutingVariableDefault describes a response with status code -1, with default header values.

Server error
*/
type ReadRoutingVariableDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this read routing variable default response has a 2xx status code
func (o *ReadRoutingVariableDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read routing variable default response has a 3xx status code
func (o *ReadRoutingVariableDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read routing variable default response has a 4xx status code
func (o *ReadRoutingVariableDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read routing variable default response has a 5xx status code
func (o *ReadRoutingVariableDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read routing variable default response a status code equal to that given
func (o *ReadRoutingVariableDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read routing variable default response
func (o *ReadRoutingVariableDefault) Code() int {
	return o._statusCode
}

func (o *ReadRoutingVariableDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /routing/variables/{id}][%d] ReadRoutingVariable default %s", o._statusCode, payload)
}

func (o *ReadRoutingVariableDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /routing/variables/{id}][%d] ReadRoutingVariable default %s", o._statusCode, payload)
}

func (o *ReadRoutingVariableDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *ReadRoutingVariableDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
