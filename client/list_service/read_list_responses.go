// Code generated by go-swagger; DO NOT EDIT.

package list_service

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

// ReadListReader is a Reader for the ReadList structure.
type ReadListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewReadListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadListOK creates a ReadListOK with default headers values
func NewReadListOK() *ReadListOK {
	return &ReadListOK{}
}

/*
ReadListOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadListOK struct {
	Payload *models.EngineList
}

// IsSuccess returns true when this read list Ok response has a 2xx status code
func (o *ReadListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read list Ok response has a 3xx status code
func (o *ReadListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read list Ok response has a 4xx status code
func (o *ReadListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read list Ok response has a 5xx status code
func (o *ReadListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read list Ok response a status code equal to that given
func (o *ReadListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read list Ok response
func (o *ReadListOK) Code() int {
	return 200
}

func (o *ReadListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/list/{id}][%d] readListOk %s", 200, payload)
}

func (o *ReadListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/list/{id}][%d] readListOk %s", 200, payload)
}

func (o *ReadListOK) GetPayload() *models.EngineList {
	return o.Payload
}

func (o *ReadListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadListDefault creates a ReadListDefault with default headers values
func NewReadListDefault(code int) *ReadListDefault {
	return &ReadListDefault{
		_statusCode: code,
	}
}

/*
ReadListDefault describes a response with status code -1, with default header values.

Server error
*/
type ReadListDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this read list default response has a 2xx status code
func (o *ReadListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read list default response has a 3xx status code
func (o *ReadListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read list default response has a 4xx status code
func (o *ReadListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read list default response has a 5xx status code
func (o *ReadListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read list default response a status code equal to that given
func (o *ReadListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read list default response
func (o *ReadListDefault) Code() int {
	return o._statusCode
}

func (o *ReadListDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/list/{id}][%d] ReadList default %s", o._statusCode, payload)
}

func (o *ReadListDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/list/{id}][%d] ReadList default %s", o._statusCode, payload)
}

func (o *ReadListDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *ReadListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
