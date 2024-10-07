// Code generated by go-swagger; DO NOT EDIT.

package config_service

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

// ConfigServiceReadSystemObjectsReader is a Reader for the ConfigServiceReadSystemObjects structure.
type ConfigServiceReadSystemObjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConfigServiceReadSystemObjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConfigServiceReadSystemObjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConfigServiceReadSystemObjectsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConfigServiceReadSystemObjectsOK creates a ConfigServiceReadSystemObjectsOK with default headers values
func NewConfigServiceReadSystemObjectsOK() *ConfigServiceReadSystemObjectsOK {
	return &ConfigServiceReadSystemObjectsOK{}
}

/*
ConfigServiceReadSystemObjectsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ConfigServiceReadSystemObjectsOK struct {
	Payload *models.LoggerSystemObjects
}

// IsSuccess returns true when this config service read system objects Ok response has a 2xx status code
func (o *ConfigServiceReadSystemObjectsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this config service read system objects Ok response has a 3xx status code
func (o *ConfigServiceReadSystemObjectsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this config service read system objects Ok response has a 4xx status code
func (o *ConfigServiceReadSystemObjectsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this config service read system objects Ok response has a 5xx status code
func (o *ConfigServiceReadSystemObjectsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this config service read system objects Ok response a status code equal to that given
func (o *ConfigServiceReadSystemObjectsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the config service read system objects Ok response
func (o *ConfigServiceReadSystemObjectsOK) Code() int {
	return 200
}

func (o *ConfigServiceReadSystemObjectsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /logger/available_objects][%d] configServiceReadSystemObjectsOk %s", 200, payload)
}

func (o *ConfigServiceReadSystemObjectsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /logger/available_objects][%d] configServiceReadSystemObjectsOk %s", 200, payload)
}

func (o *ConfigServiceReadSystemObjectsOK) GetPayload() *models.LoggerSystemObjects {
	return o.Payload
}

func (o *ConfigServiceReadSystemObjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LoggerSystemObjects)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigServiceReadSystemObjectsDefault creates a ConfigServiceReadSystemObjectsDefault with default headers values
func NewConfigServiceReadSystemObjectsDefault(code int) *ConfigServiceReadSystemObjectsDefault {
	return &ConfigServiceReadSystemObjectsDefault{
		_statusCode: code,
	}
}

/*
ConfigServiceReadSystemObjectsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ConfigServiceReadSystemObjectsDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// IsSuccess returns true when this config service read system objects default response has a 2xx status code
func (o *ConfigServiceReadSystemObjectsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this config service read system objects default response has a 3xx status code
func (o *ConfigServiceReadSystemObjectsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this config service read system objects default response has a 4xx status code
func (o *ConfigServiceReadSystemObjectsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this config service read system objects default response has a 5xx status code
func (o *ConfigServiceReadSystemObjectsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this config service read system objects default response a status code equal to that given
func (o *ConfigServiceReadSystemObjectsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the config service read system objects default response
func (o *ConfigServiceReadSystemObjectsDefault) Code() int {
	return o._statusCode
}

func (o *ConfigServiceReadSystemObjectsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /logger/available_objects][%d] ConfigService_ReadSystemObjects default %s", o._statusCode, payload)
}

func (o *ConfigServiceReadSystemObjectsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /logger/available_objects][%d] ConfigService_ReadSystemObjects default %s", o._statusCode, payload)
}

func (o *ConfigServiceReadSystemObjectsDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *ConfigServiceReadSystemObjectsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
