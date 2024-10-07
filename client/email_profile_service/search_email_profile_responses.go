// Code generated by go-swagger; DO NOT EDIT.

package email_profile_service

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

// SearchEmailProfileReader is a Reader for the SearchEmailProfile structure.
type SearchEmailProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchEmailProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchEmailProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSearchEmailProfileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchEmailProfileOK creates a SearchEmailProfileOK with default headers values
func NewSearchEmailProfileOK() *SearchEmailProfileOK {
	return &SearchEmailProfileOK{}
}

/*
SearchEmailProfileOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchEmailProfileOK struct {
	Payload *models.EngineListEmailProfile
}

// IsSuccess returns true when this search email profile Ok response has a 2xx status code
func (o *SearchEmailProfileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search email profile Ok response has a 3xx status code
func (o *SearchEmailProfileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search email profile Ok response has a 4xx status code
func (o *SearchEmailProfileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search email profile Ok response has a 5xx status code
func (o *SearchEmailProfileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search email profile Ok response a status code equal to that given
func (o *SearchEmailProfileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search email profile Ok response
func (o *SearchEmailProfileOK) Code() int {
	return 200
}

func (o *SearchEmailProfileOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /email/profile][%d] searchEmailProfileOk %s", 200, payload)
}

func (o *SearchEmailProfileOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /email/profile][%d] searchEmailProfileOk %s", 200, payload)
}

func (o *SearchEmailProfileOK) GetPayload() *models.EngineListEmailProfile {
	return o.Payload
}

func (o *SearchEmailProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineListEmailProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchEmailProfileDefault creates a SearchEmailProfileDefault with default headers values
func NewSearchEmailProfileDefault(code int) *SearchEmailProfileDefault {
	return &SearchEmailProfileDefault{
		_statusCode: code,
	}
}

/*
SearchEmailProfileDefault describes a response with status code -1, with default header values.

Server error
*/
type SearchEmailProfileDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this search email profile default response has a 2xx status code
func (o *SearchEmailProfileDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this search email profile default response has a 3xx status code
func (o *SearchEmailProfileDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this search email profile default response has a 4xx status code
func (o *SearchEmailProfileDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this search email profile default response has a 5xx status code
func (o *SearchEmailProfileDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this search email profile default response a status code equal to that given
func (o *SearchEmailProfileDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the search email profile default response
func (o *SearchEmailProfileDefault) Code() int {
	return o._statusCode
}

func (o *SearchEmailProfileDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /email/profile][%d] SearchEmailProfile default %s", o._statusCode, payload)
}

func (o *SearchEmailProfileDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /email/profile][%d] SearchEmailProfile default %s", o._statusCode, payload)
}

func (o *SearchEmailProfileDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *SearchEmailProfileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
