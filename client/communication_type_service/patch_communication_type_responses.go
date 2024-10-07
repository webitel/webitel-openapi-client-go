// Code generated by go-swagger; DO NOT EDIT.

package communication_type_service

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

// PatchCommunicationTypeReader is a Reader for the PatchCommunicationType structure.
type PatchCommunicationTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchCommunicationTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchCommunicationTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchCommunicationTypeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchCommunicationTypeOK creates a PatchCommunicationTypeOK with default headers values
func NewPatchCommunicationTypeOK() *PatchCommunicationTypeOK {
	return &PatchCommunicationTypeOK{}
}

/*
PatchCommunicationTypeOK describes a response with status code 200, with default header values.

A successful response.
*/
type PatchCommunicationTypeOK struct {
	Payload *models.EngineCommunicationType
}

// IsSuccess returns true when this patch communication type Ok response has a 2xx status code
func (o *PatchCommunicationTypeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch communication type Ok response has a 3xx status code
func (o *PatchCommunicationTypeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch communication type Ok response has a 4xx status code
func (o *PatchCommunicationTypeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch communication type Ok response has a 5xx status code
func (o *PatchCommunicationTypeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch communication type Ok response a status code equal to that given
func (o *PatchCommunicationTypeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch communication type Ok response
func (o *PatchCommunicationTypeOK) Code() int {
	return 200
}

func (o *PatchCommunicationTypeOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /call_center/communication_type/{id}][%d] patchCommunicationTypeOk %s", 200, payload)
}

func (o *PatchCommunicationTypeOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /call_center/communication_type/{id}][%d] patchCommunicationTypeOk %s", 200, payload)
}

func (o *PatchCommunicationTypeOK) GetPayload() *models.EngineCommunicationType {
	return o.Payload
}

func (o *PatchCommunicationTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineCommunicationType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCommunicationTypeDefault creates a PatchCommunicationTypeDefault with default headers values
func NewPatchCommunicationTypeDefault(code int) *PatchCommunicationTypeDefault {
	return &PatchCommunicationTypeDefault{
		_statusCode: code,
	}
}

/*
PatchCommunicationTypeDefault describes a response with status code -1, with default header values.

Server error
*/
type PatchCommunicationTypeDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this patch communication type default response has a 2xx status code
func (o *PatchCommunicationTypeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch communication type default response has a 3xx status code
func (o *PatchCommunicationTypeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch communication type default response has a 4xx status code
func (o *PatchCommunicationTypeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch communication type default response has a 5xx status code
func (o *PatchCommunicationTypeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch communication type default response a status code equal to that given
func (o *PatchCommunicationTypeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the patch communication type default response
func (o *PatchCommunicationTypeDefault) Code() int {
	return o._statusCode
}

func (o *PatchCommunicationTypeDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /call_center/communication_type/{id}][%d] PatchCommunicationType default %s", o._statusCode, payload)
}

func (o *PatchCommunicationTypeDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /call_center/communication_type/{id}][%d] PatchCommunicationType default %s", o._statusCode, payload)
}

func (o *PatchCommunicationTypeDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *PatchCommunicationTypeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
