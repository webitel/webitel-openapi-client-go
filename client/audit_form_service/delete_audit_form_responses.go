// Code generated by go-swagger; DO NOT EDIT.

package audit_form_service

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

// DeleteAuditFormReader is a Reader for the DeleteAuditForm structure.
type DeleteAuditFormReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAuditFormReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAuditFormOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteAuditFormDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAuditFormOK creates a DeleteAuditFormOK with default headers values
func NewDeleteAuditFormOK() *DeleteAuditFormOK {
	return &DeleteAuditFormOK{}
}

/*
DeleteAuditFormOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteAuditFormOK struct {
	Payload *models.EngineAuditForm
}

// IsSuccess returns true when this delete audit form Ok response has a 2xx status code
func (o *DeleteAuditFormOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete audit form Ok response has a 3xx status code
func (o *DeleteAuditFormOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete audit form Ok response has a 4xx status code
func (o *DeleteAuditFormOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete audit form Ok response has a 5xx status code
func (o *DeleteAuditFormOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete audit form Ok response a status code equal to that given
func (o *DeleteAuditFormOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete audit form Ok response
func (o *DeleteAuditFormOK) Code() int {
	return 200
}

func (o *DeleteAuditFormOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/audit/forms/{id}][%d] deleteAuditFormOk %s", 200, payload)
}

func (o *DeleteAuditFormOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/audit/forms/{id}][%d] deleteAuditFormOk %s", 200, payload)
}

func (o *DeleteAuditFormOK) GetPayload() *models.EngineAuditForm {
	return o.Payload
}

func (o *DeleteAuditFormOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAuditForm)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuditFormDefault creates a DeleteAuditFormDefault with default headers values
func NewDeleteAuditFormDefault(code int) *DeleteAuditFormDefault {
	return &DeleteAuditFormDefault{
		_statusCode: code,
	}
}

/*
DeleteAuditFormDefault describes a response with status code -1, with default header values.

Server error
*/
type DeleteAuditFormDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this delete audit form default response has a 2xx status code
func (o *DeleteAuditFormDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete audit form default response has a 3xx status code
func (o *DeleteAuditFormDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete audit form default response has a 4xx status code
func (o *DeleteAuditFormDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete audit form default response has a 5xx status code
func (o *DeleteAuditFormDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete audit form default response a status code equal to that given
func (o *DeleteAuditFormDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete audit form default response
func (o *DeleteAuditFormDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAuditFormDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/audit/forms/{id}][%d] DeleteAuditForm default %s", o._statusCode, payload)
}

func (o *DeleteAuditFormDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /call_center/audit/forms/{id}][%d] DeleteAuditForm default %s", o._statusCode, payload)
}

func (o *DeleteAuditFormDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *DeleteAuditFormDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
