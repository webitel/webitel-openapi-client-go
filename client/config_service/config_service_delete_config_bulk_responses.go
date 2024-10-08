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

// ConfigServiceDeleteConfigBulkReader is a Reader for the ConfigServiceDeleteConfigBulk structure.
type ConfigServiceDeleteConfigBulkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConfigServiceDeleteConfigBulkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConfigServiceDeleteConfigBulkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConfigServiceDeleteConfigBulkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConfigServiceDeleteConfigBulkOK creates a ConfigServiceDeleteConfigBulkOK with default headers values
func NewConfigServiceDeleteConfigBulkOK() *ConfigServiceDeleteConfigBulkOK {
	return &ConfigServiceDeleteConfigBulkOK{}
}

/*
ConfigServiceDeleteConfigBulkOK describes a response with status code 200, with default header values.

A successful response.
*/
type ConfigServiceDeleteConfigBulkOK struct {
	Payload models.LoggerEmpty
}

// IsSuccess returns true when this config service delete config bulk Ok response has a 2xx status code
func (o *ConfigServiceDeleteConfigBulkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this config service delete config bulk Ok response has a 3xx status code
func (o *ConfigServiceDeleteConfigBulkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this config service delete config bulk Ok response has a 4xx status code
func (o *ConfigServiceDeleteConfigBulkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this config service delete config bulk Ok response has a 5xx status code
func (o *ConfigServiceDeleteConfigBulkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this config service delete config bulk Ok response a status code equal to that given
func (o *ConfigServiceDeleteConfigBulkOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the config service delete config bulk Ok response
func (o *ConfigServiceDeleteConfigBulkOK) Code() int {
	return 200
}

func (o *ConfigServiceDeleteConfigBulkOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /logger/config][%d] configServiceDeleteConfigBulkOk %s", 200, payload)
}

func (o *ConfigServiceDeleteConfigBulkOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /logger/config][%d] configServiceDeleteConfigBulkOk %s", 200, payload)
}

func (o *ConfigServiceDeleteConfigBulkOK) GetPayload() models.LoggerEmpty {
	return o.Payload
}

func (o *ConfigServiceDeleteConfigBulkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigServiceDeleteConfigBulkDefault creates a ConfigServiceDeleteConfigBulkDefault with default headers values
func NewConfigServiceDeleteConfigBulkDefault(code int) *ConfigServiceDeleteConfigBulkDefault {
	return &ConfigServiceDeleteConfigBulkDefault{
		_statusCode: code,
	}
}

/*
ConfigServiceDeleteConfigBulkDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ConfigServiceDeleteConfigBulkDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// IsSuccess returns true when this config service delete config bulk default response has a 2xx status code
func (o *ConfigServiceDeleteConfigBulkDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this config service delete config bulk default response has a 3xx status code
func (o *ConfigServiceDeleteConfigBulkDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this config service delete config bulk default response has a 4xx status code
func (o *ConfigServiceDeleteConfigBulkDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this config service delete config bulk default response has a 5xx status code
func (o *ConfigServiceDeleteConfigBulkDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this config service delete config bulk default response a status code equal to that given
func (o *ConfigServiceDeleteConfigBulkDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the config service delete config bulk default response
func (o *ConfigServiceDeleteConfigBulkDefault) Code() int {
	return o._statusCode
}

func (o *ConfigServiceDeleteConfigBulkDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /logger/config][%d] ConfigService_DeleteConfigBulk default %s", o._statusCode, payload)
}

func (o *ConfigServiceDeleteConfigBulkDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /logger/config][%d] ConfigService_DeleteConfigBulk default %s", o._statusCode, payload)
}

func (o *ConfigServiceDeleteConfigBulkDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *ConfigServiceDeleteConfigBulkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
