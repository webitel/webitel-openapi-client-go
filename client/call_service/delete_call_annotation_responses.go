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

// DeleteCallAnnotationReader is a Reader for the DeleteCallAnnotation structure.
type DeleteCallAnnotationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCallAnnotationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCallAnnotationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteCallAnnotationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCallAnnotationOK creates a DeleteCallAnnotationOK with default headers values
func NewDeleteCallAnnotationOK() *DeleteCallAnnotationOK {
	return &DeleteCallAnnotationOK{}
}

/*
DeleteCallAnnotationOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteCallAnnotationOK struct {
	Payload *models.EngineCallAnnotation
}

// IsSuccess returns true when this delete call annotation Ok response has a 2xx status code
func (o *DeleteCallAnnotationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete call annotation Ok response has a 3xx status code
func (o *DeleteCallAnnotationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete call annotation Ok response has a 4xx status code
func (o *DeleteCallAnnotationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete call annotation Ok response has a 5xx status code
func (o *DeleteCallAnnotationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete call annotation Ok response a status code equal to that given
func (o *DeleteCallAnnotationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete call annotation Ok response
func (o *DeleteCallAnnotationOK) Code() int {
	return 200
}

func (o *DeleteCallAnnotationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /calls/history/{call_id}/annotation/{id}][%d] deleteCallAnnotationOk %s", 200, payload)
}

func (o *DeleteCallAnnotationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /calls/history/{call_id}/annotation/{id}][%d] deleteCallAnnotationOk %s", 200, payload)
}

func (o *DeleteCallAnnotationOK) GetPayload() *models.EngineCallAnnotation {
	return o.Payload
}

func (o *DeleteCallAnnotationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineCallAnnotation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCallAnnotationDefault creates a DeleteCallAnnotationDefault with default headers values
func NewDeleteCallAnnotationDefault(code int) *DeleteCallAnnotationDefault {
	return &DeleteCallAnnotationDefault{
		_statusCode: code,
	}
}

/*
DeleteCallAnnotationDefault describes a response with status code -1, with default header values.

Server error
*/
type DeleteCallAnnotationDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this delete call annotation default response has a 2xx status code
func (o *DeleteCallAnnotationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete call annotation default response has a 3xx status code
func (o *DeleteCallAnnotationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete call annotation default response has a 4xx status code
func (o *DeleteCallAnnotationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete call annotation default response has a 5xx status code
func (o *DeleteCallAnnotationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete call annotation default response a status code equal to that given
func (o *DeleteCallAnnotationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete call annotation default response
func (o *DeleteCallAnnotationDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCallAnnotationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /calls/history/{call_id}/annotation/{id}][%d] DeleteCallAnnotation default %s", o._statusCode, payload)
}

func (o *DeleteCallAnnotationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /calls/history/{call_id}/annotation/{id}][%d] DeleteCallAnnotation default %s", o._statusCode, payload)
}

func (o *DeleteCallAnnotationDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *DeleteCallAnnotationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
