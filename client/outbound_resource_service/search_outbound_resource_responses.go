// Code generated by go-swagger; DO NOT EDIT.

package outbound_resource_service

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

// SearchOutboundResourceReader is a Reader for the SearchOutboundResource structure.
type SearchOutboundResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchOutboundResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchOutboundResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSearchOutboundResourceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchOutboundResourceOK creates a SearchOutboundResourceOK with default headers values
func NewSearchOutboundResourceOK() *SearchOutboundResourceOK {
	return &SearchOutboundResourceOK{}
}

/*
SearchOutboundResourceOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchOutboundResourceOK struct {
	Payload *models.EngineListOutboundResource
}

// IsSuccess returns true when this search outbound resource Ok response has a 2xx status code
func (o *SearchOutboundResourceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search outbound resource Ok response has a 3xx status code
func (o *SearchOutboundResourceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search outbound resource Ok response has a 4xx status code
func (o *SearchOutboundResourceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search outbound resource Ok response has a 5xx status code
func (o *SearchOutboundResourceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search outbound resource Ok response a status code equal to that given
func (o *SearchOutboundResourceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search outbound resource Ok response
func (o *SearchOutboundResourceOK) Code() int {
	return 200
}

func (o *SearchOutboundResourceOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/resources][%d] searchOutboundResourceOk %s", 200, payload)
}

func (o *SearchOutboundResourceOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/resources][%d] searchOutboundResourceOk %s", 200, payload)
}

func (o *SearchOutboundResourceOK) GetPayload() *models.EngineListOutboundResource {
	return o.Payload
}

func (o *SearchOutboundResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineListOutboundResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchOutboundResourceDefault creates a SearchOutboundResourceDefault with default headers values
func NewSearchOutboundResourceDefault(code int) *SearchOutboundResourceDefault {
	return &SearchOutboundResourceDefault{
		_statusCode: code,
	}
}

/*
SearchOutboundResourceDefault describes a response with status code -1, with default header values.

Server error
*/
type SearchOutboundResourceDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this search outbound resource default response has a 2xx status code
func (o *SearchOutboundResourceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this search outbound resource default response has a 3xx status code
func (o *SearchOutboundResourceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this search outbound resource default response has a 4xx status code
func (o *SearchOutboundResourceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this search outbound resource default response has a 5xx status code
func (o *SearchOutboundResourceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this search outbound resource default response a status code equal to that given
func (o *SearchOutboundResourceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the search outbound resource default response
func (o *SearchOutboundResourceDefault) Code() int {
	return o._statusCode
}

func (o *SearchOutboundResourceDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/resources][%d] SearchOutboundResource default %s", o._statusCode, payload)
}

func (o *SearchOutboundResourceDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/resources][%d] SearchOutboundResource default %s", o._statusCode, payload)
}

func (o *SearchOutboundResourceDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *SearchOutboundResourceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
