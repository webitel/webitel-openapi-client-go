// Code generated by go-swagger; DO NOT EDIT.

package system_setting_service

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

// ReadSystemSettingReader is a Reader for the ReadSystemSetting structure.
type ReadSystemSettingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadSystemSettingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadSystemSettingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewReadSystemSettingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadSystemSettingOK creates a ReadSystemSettingOK with default headers values
func NewReadSystemSettingOK() *ReadSystemSettingOK {
	return &ReadSystemSettingOK{}
}

/*
ReadSystemSettingOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadSystemSettingOK struct {
	Payload *models.EngineSystemSetting
}

// IsSuccess returns true when this read system setting Ok response has a 2xx status code
func (o *ReadSystemSettingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read system setting Ok response has a 3xx status code
func (o *ReadSystemSettingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read system setting Ok response has a 4xx status code
func (o *ReadSystemSettingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read system setting Ok response has a 5xx status code
func (o *ReadSystemSettingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read system setting Ok response a status code equal to that given
func (o *ReadSystemSettingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read system setting Ok response
func (o *ReadSystemSettingOK) Code() int {
	return 200
}

func (o *ReadSystemSettingOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /settings/{id}][%d] readSystemSettingOk %s", 200, payload)
}

func (o *ReadSystemSettingOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /settings/{id}][%d] readSystemSettingOk %s", 200, payload)
}

func (o *ReadSystemSettingOK) GetPayload() *models.EngineSystemSetting {
	return o.Payload
}

func (o *ReadSystemSettingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineSystemSetting)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadSystemSettingDefault creates a ReadSystemSettingDefault with default headers values
func NewReadSystemSettingDefault(code int) *ReadSystemSettingDefault {
	return &ReadSystemSettingDefault{
		_statusCode: code,
	}
}

/*
ReadSystemSettingDefault describes a response with status code -1, with default header values.

Server error
*/
type ReadSystemSettingDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this read system setting default response has a 2xx status code
func (o *ReadSystemSettingDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read system setting default response has a 3xx status code
func (o *ReadSystemSettingDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read system setting default response has a 4xx status code
func (o *ReadSystemSettingDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read system setting default response has a 5xx status code
func (o *ReadSystemSettingDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read system setting default response a status code equal to that given
func (o *ReadSystemSettingDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read system setting default response
func (o *ReadSystemSettingDefault) Code() int {
	return o._statusCode
}

func (o *ReadSystemSettingDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /settings/{id}][%d] ReadSystemSetting default %s", o._statusCode, payload)
}

func (o *ReadSystemSettingDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /settings/{id}][%d] ReadSystemSetting default %s", o._statusCode, payload)
}

func (o *ReadSystemSettingDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *ReadSystemSettingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
