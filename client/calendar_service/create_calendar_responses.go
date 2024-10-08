// Code generated by go-swagger; DO NOT EDIT.

package calendar_service

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

// CreateCalendarReader is a Reader for the CreateCalendar structure.
type CreateCalendarReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCalendarReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateCalendarOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateCalendarDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateCalendarOK creates a CreateCalendarOK with default headers values
func NewCreateCalendarOK() *CreateCalendarOK {
	return &CreateCalendarOK{}
}

/*
CreateCalendarOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateCalendarOK struct {
	Payload *models.EngineCalendar
}

// IsSuccess returns true when this create calendar Ok response has a 2xx status code
func (o *CreateCalendarOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create calendar Ok response has a 3xx status code
func (o *CreateCalendarOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create calendar Ok response has a 4xx status code
func (o *CreateCalendarOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create calendar Ok response has a 5xx status code
func (o *CreateCalendarOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create calendar Ok response a status code equal to that given
func (o *CreateCalendarOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create calendar Ok response
func (o *CreateCalendarOK) Code() int {
	return 200
}

func (o *CreateCalendarOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /calendars][%d] createCalendarOk %s", 200, payload)
}

func (o *CreateCalendarOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /calendars][%d] createCalendarOk %s", 200, payload)
}

func (o *CreateCalendarOK) GetPayload() *models.EngineCalendar {
	return o.Payload
}

func (o *CreateCalendarOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineCalendar)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCalendarDefault creates a CreateCalendarDefault with default headers values
func NewCreateCalendarDefault(code int) *CreateCalendarDefault {
	return &CreateCalendarDefault{
		_statusCode: code,
	}
}

/*
CreateCalendarDefault describes a response with status code -1, with default header values.

Server error
*/
type CreateCalendarDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this create calendar default response has a 2xx status code
func (o *CreateCalendarDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create calendar default response has a 3xx status code
func (o *CreateCalendarDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create calendar default response has a 4xx status code
func (o *CreateCalendarDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create calendar default response has a 5xx status code
func (o *CreateCalendarDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create calendar default response a status code equal to that given
func (o *CreateCalendarDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create calendar default response
func (o *CreateCalendarDefault) Code() int {
	return o._statusCode
}

func (o *CreateCalendarDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /calendars][%d] CreateCalendar default %s", o._statusCode, payload)
}

func (o *CreateCalendarDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /calendars][%d] CreateCalendar default %s", o._statusCode, payload)
}

func (o *CreateCalendarDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *CreateCalendarDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
