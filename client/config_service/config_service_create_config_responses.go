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

// ConfigServiceCreateConfigReader is a Reader for the ConfigServiceCreateConfig structure.
type ConfigServiceCreateConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConfigServiceCreateConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConfigServiceCreateConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConfigServiceCreateConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConfigServiceCreateConfigOK creates a ConfigServiceCreateConfigOK with default headers values
func NewConfigServiceCreateConfigOK() *ConfigServiceCreateConfigOK {
	return &ConfigServiceCreateConfigOK{}
}

/*
ConfigServiceCreateConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type ConfigServiceCreateConfigOK struct {
	Payload *models.LoggerConfig
}

// IsSuccess returns true when this config service create config Ok response has a 2xx status code
func (o *ConfigServiceCreateConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this config service create config Ok response has a 3xx status code
func (o *ConfigServiceCreateConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this config service create config Ok response has a 4xx status code
func (o *ConfigServiceCreateConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this config service create config Ok response has a 5xx status code
func (o *ConfigServiceCreateConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this config service create config Ok response a status code equal to that given
func (o *ConfigServiceCreateConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the config service create config Ok response
func (o *ConfigServiceCreateConfigOK) Code() int {
	return 200
}

func (o *ConfigServiceCreateConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /logger/config][%d] configServiceCreateConfigOk %s", 200, payload)
}

func (o *ConfigServiceCreateConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /logger/config][%d] configServiceCreateConfigOk %s", 200, payload)
}

func (o *ConfigServiceCreateConfigOK) GetPayload() *models.LoggerConfig {
	return o.Payload
}

func (o *ConfigServiceCreateConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LoggerConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigServiceCreateConfigDefault creates a ConfigServiceCreateConfigDefault with default headers values
func NewConfigServiceCreateConfigDefault(code int) *ConfigServiceCreateConfigDefault {
	return &ConfigServiceCreateConfigDefault{
		_statusCode: code,
	}
}

/*
ConfigServiceCreateConfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ConfigServiceCreateConfigDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// IsSuccess returns true when this config service create config default response has a 2xx status code
func (o *ConfigServiceCreateConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this config service create config default response has a 3xx status code
func (o *ConfigServiceCreateConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this config service create config default response has a 4xx status code
func (o *ConfigServiceCreateConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this config service create config default response has a 5xx status code
func (o *ConfigServiceCreateConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this config service create config default response a status code equal to that given
func (o *ConfigServiceCreateConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the config service create config default response
func (o *ConfigServiceCreateConfigDefault) Code() int {
	return o._statusCode
}

func (o *ConfigServiceCreateConfigDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /logger/config][%d] ConfigService_CreateConfig default %s", o._statusCode, payload)
}

func (o *ConfigServiceCreateConfigDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /logger/config][%d] ConfigService_CreateConfig default %s", o._statusCode, payload)
}

func (o *ConfigServiceCreateConfigDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *ConfigServiceCreateConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
