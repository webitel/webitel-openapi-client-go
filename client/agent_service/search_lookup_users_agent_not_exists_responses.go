// Code generated by go-swagger; DO NOT EDIT.

package agent_service

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

// SearchLookupUsersAgentNotExistsReader is a Reader for the SearchLookupUsersAgentNotExists structure.
type SearchLookupUsersAgentNotExistsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchLookupUsersAgentNotExistsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchLookupUsersAgentNotExistsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSearchLookupUsersAgentNotExistsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchLookupUsersAgentNotExistsOK creates a SearchLookupUsersAgentNotExistsOK with default headers values
func NewSearchLookupUsersAgentNotExistsOK() *SearchLookupUsersAgentNotExistsOK {
	return &SearchLookupUsersAgentNotExistsOK{}
}

/*
SearchLookupUsersAgentNotExistsOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchLookupUsersAgentNotExistsOK struct {
	Payload *models.EngineListAgentUser
}

// IsSuccess returns true when this search lookup users agent not exists Ok response has a 2xx status code
func (o *SearchLookupUsersAgentNotExistsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search lookup users agent not exists Ok response has a 3xx status code
func (o *SearchLookupUsersAgentNotExistsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search lookup users agent not exists Ok response has a 4xx status code
func (o *SearchLookupUsersAgentNotExistsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search lookup users agent not exists Ok response has a 5xx status code
func (o *SearchLookupUsersAgentNotExistsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search lookup users agent not exists Ok response a status code equal to that given
func (o *SearchLookupUsersAgentNotExistsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search lookup users agent not exists Ok response
func (o *SearchLookupUsersAgentNotExistsOK) Code() int {
	return 200
}

func (o *SearchLookupUsersAgentNotExistsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/lookups/agents/users][%d] searchLookupUsersAgentNotExistsOk %s", 200, payload)
}

func (o *SearchLookupUsersAgentNotExistsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/lookups/agents/users][%d] searchLookupUsersAgentNotExistsOk %s", 200, payload)
}

func (o *SearchLookupUsersAgentNotExistsOK) GetPayload() *models.EngineListAgentUser {
	return o.Payload
}

func (o *SearchLookupUsersAgentNotExistsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineListAgentUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchLookupUsersAgentNotExistsDefault creates a SearchLookupUsersAgentNotExistsDefault with default headers values
func NewSearchLookupUsersAgentNotExistsDefault(code int) *SearchLookupUsersAgentNotExistsDefault {
	return &SearchLookupUsersAgentNotExistsDefault{
		_statusCode: code,
	}
}

/*
SearchLookupUsersAgentNotExistsDefault describes a response with status code -1, with default header values.

Server error
*/
type SearchLookupUsersAgentNotExistsDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this search lookup users agent not exists default response has a 2xx status code
func (o *SearchLookupUsersAgentNotExistsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this search lookup users agent not exists default response has a 3xx status code
func (o *SearchLookupUsersAgentNotExistsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this search lookup users agent not exists default response has a 4xx status code
func (o *SearchLookupUsersAgentNotExistsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this search lookup users agent not exists default response has a 5xx status code
func (o *SearchLookupUsersAgentNotExistsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this search lookup users agent not exists default response a status code equal to that given
func (o *SearchLookupUsersAgentNotExistsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the search lookup users agent not exists default response
func (o *SearchLookupUsersAgentNotExistsDefault) Code() int {
	return o._statusCode
}

func (o *SearchLookupUsersAgentNotExistsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/lookups/agents/users][%d] SearchLookupUsersAgentNotExists default %s", o._statusCode, payload)
}

func (o *SearchLookupUsersAgentNotExistsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/lookups/agents/users][%d] SearchLookupUsersAgentNotExists default %s", o._statusCode, payload)
}

func (o *SearchLookupUsersAgentNotExistsDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *SearchLookupUsersAgentNotExistsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
