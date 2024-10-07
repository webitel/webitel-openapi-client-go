// Code generated by go-swagger; DO NOT EDIT.

package routing_chat_plan_service

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

// SearchChatPlanReader is a Reader for the SearchChatPlan structure.
type SearchChatPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchChatPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchChatPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSearchChatPlanDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchChatPlanOK creates a SearchChatPlanOK with default headers values
func NewSearchChatPlanOK() *SearchChatPlanOK {
	return &SearchChatPlanOK{}
}

/*
SearchChatPlanOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchChatPlanOK struct {
	Payload *models.EngineListChatPlan
}

// IsSuccess returns true when this search chat plan Ok response has a 2xx status code
func (o *SearchChatPlanOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search chat plan Ok response has a 3xx status code
func (o *SearchChatPlanOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search chat plan Ok response has a 4xx status code
func (o *SearchChatPlanOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search chat plan Ok response has a 5xx status code
func (o *SearchChatPlanOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search chat plan Ok response a status code equal to that given
func (o *SearchChatPlanOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search chat plan Ok response
func (o *SearchChatPlanOK) Code() int {
	return 200
}

func (o *SearchChatPlanOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /routing/outbound/chat][%d] searchChatPlanOk %s", 200, payload)
}

func (o *SearchChatPlanOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /routing/outbound/chat][%d] searchChatPlanOk %s", 200, payload)
}

func (o *SearchChatPlanOK) GetPayload() *models.EngineListChatPlan {
	return o.Payload
}

func (o *SearchChatPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineListChatPlan)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchChatPlanDefault creates a SearchChatPlanDefault with default headers values
func NewSearchChatPlanDefault(code int) *SearchChatPlanDefault {
	return &SearchChatPlanDefault{
		_statusCode: code,
	}
}

/*
SearchChatPlanDefault describes a response with status code -1, with default header values.

Server error
*/
type SearchChatPlanDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this search chat plan default response has a 2xx status code
func (o *SearchChatPlanDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this search chat plan default response has a 3xx status code
func (o *SearchChatPlanDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this search chat plan default response has a 4xx status code
func (o *SearchChatPlanDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this search chat plan default response has a 5xx status code
func (o *SearchChatPlanDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this search chat plan default response a status code equal to that given
func (o *SearchChatPlanDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the search chat plan default response
func (o *SearchChatPlanDefault) Code() int {
	return o._statusCode
}

func (o *SearchChatPlanDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /routing/outbound/chat][%d] SearchChatPlan default %s", o._statusCode, payload)
}

func (o *SearchChatPlanDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /routing/outbound/chat][%d] SearchChatPlan default %s", o._statusCode, payload)
}

func (o *SearchChatPlanDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *SearchChatPlanDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
