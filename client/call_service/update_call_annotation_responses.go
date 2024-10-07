// Code generated by go-swagger; DO NOT EDIT.

package call_service

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

// UpdateCallAnnotationReader is a Reader for the UpdateCallAnnotation structure.
type UpdateCallAnnotationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCallAnnotationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCallAnnotationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateCallAnnotationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateCallAnnotationOK creates a UpdateCallAnnotationOK with default headers values
func NewUpdateCallAnnotationOK() *UpdateCallAnnotationOK {
	return &UpdateCallAnnotationOK{}
}

/*
UpdateCallAnnotationOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateCallAnnotationOK struct {
	Payload *models.EngineCallAnnotation
}

// IsSuccess returns true when this update call annotation Ok response has a 2xx status code
func (o *UpdateCallAnnotationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update call annotation Ok response has a 3xx status code
func (o *UpdateCallAnnotationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update call annotation Ok response has a 4xx status code
func (o *UpdateCallAnnotationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update call annotation Ok response has a 5xx status code
func (o *UpdateCallAnnotationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update call annotation Ok response a status code equal to that given
func (o *UpdateCallAnnotationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update call annotation Ok response
func (o *UpdateCallAnnotationOK) Code() int {
	return 200
}

func (o *UpdateCallAnnotationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /calls/history/{call_id}/annotation/{id}][%d] updateCallAnnotationOk %s", 200, payload)
}

func (o *UpdateCallAnnotationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /calls/history/{call_id}/annotation/{id}][%d] updateCallAnnotationOk %s", 200, payload)
}

func (o *UpdateCallAnnotationOK) GetPayload() *models.EngineCallAnnotation {
	return o.Payload
}

func (o *UpdateCallAnnotationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineCallAnnotation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCallAnnotationDefault creates a UpdateCallAnnotationDefault with default headers values
func NewUpdateCallAnnotationDefault(code int) *UpdateCallAnnotationDefault {
	return &UpdateCallAnnotationDefault{
		_statusCode: code,
	}
}

/*
UpdateCallAnnotationDefault describes a response with status code -1, with default header values.

Server error
*/
type UpdateCallAnnotationDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this update call annotation default response has a 2xx status code
func (o *UpdateCallAnnotationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update call annotation default response has a 3xx status code
func (o *UpdateCallAnnotationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update call annotation default response has a 4xx status code
func (o *UpdateCallAnnotationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update call annotation default response has a 5xx status code
func (o *UpdateCallAnnotationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update call annotation default response a status code equal to that given
func (o *UpdateCallAnnotationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update call annotation default response
func (o *UpdateCallAnnotationDefault) Code() int {
	return o._statusCode
}

func (o *UpdateCallAnnotationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /calls/history/{call_id}/annotation/{id}][%d] UpdateCallAnnotation default %s", o._statusCode, payload)
}

func (o *UpdateCallAnnotationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /calls/history/{call_id}/annotation/{id}][%d] UpdateCallAnnotation default %s", o._statusCode, payload)
}

func (o *UpdateCallAnnotationDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *UpdateCallAnnotationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
