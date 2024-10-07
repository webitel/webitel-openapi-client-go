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

// AggregateHistoryCallReader is a Reader for the AggregateHistoryCall structure.
type AggregateHistoryCallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AggregateHistoryCallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAggregateHistoryCallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAggregateHistoryCallDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAggregateHistoryCallOK creates a AggregateHistoryCallOK with default headers values
func NewAggregateHistoryCallOK() *AggregateHistoryCallOK {
	return &AggregateHistoryCallOK{}
}

/*
AggregateHistoryCallOK describes a response with status code 200, with default header values.

A successful response.
*/
type AggregateHistoryCallOK struct {
	Payload *models.EngineListAggregate
}

// IsSuccess returns true when this aggregate history call Ok response has a 2xx status code
func (o *AggregateHistoryCallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aggregate history call Ok response has a 3xx status code
func (o *AggregateHistoryCallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate history call Ok response has a 4xx status code
func (o *AggregateHistoryCallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregate history call Ok response has a 5xx status code
func (o *AggregateHistoryCallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate history call Ok response a status code equal to that given
func (o *AggregateHistoryCallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the aggregate history call Ok response
func (o *AggregateHistoryCallOK) Code() int {
	return 200
}

func (o *AggregateHistoryCallOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /calls/history/aggregate][%d] aggregateHistoryCallOk %s", 200, payload)
}

func (o *AggregateHistoryCallOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /calls/history/aggregate][%d] aggregateHistoryCallOk %s", 200, payload)
}

func (o *AggregateHistoryCallOK) GetPayload() *models.EngineListAggregate {
	return o.Payload
}

func (o *AggregateHistoryCallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineListAggregate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregateHistoryCallDefault creates a AggregateHistoryCallDefault with default headers values
func NewAggregateHistoryCallDefault(code int) *AggregateHistoryCallDefault {
	return &AggregateHistoryCallDefault{
		_statusCode: code,
	}
}

/*
AggregateHistoryCallDefault describes a response with status code -1, with default header values.

Server error
*/
type AggregateHistoryCallDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this aggregate history call default response has a 2xx status code
func (o *AggregateHistoryCallDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this aggregate history call default response has a 3xx status code
func (o *AggregateHistoryCallDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this aggregate history call default response has a 4xx status code
func (o *AggregateHistoryCallDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this aggregate history call default response has a 5xx status code
func (o *AggregateHistoryCallDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this aggregate history call default response a status code equal to that given
func (o *AggregateHistoryCallDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the aggregate history call default response
func (o *AggregateHistoryCallDefault) Code() int {
	return o._statusCode
}

func (o *AggregateHistoryCallDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /calls/history/aggregate][%d] AggregateHistoryCall default %s", o._statusCode, payload)
}

func (o *AggregateHistoryCallDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /calls/history/aggregate][%d] AggregateHistoryCall default %s", o._statusCode, payload)
}

func (o *AggregateHistoryCallDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *AggregateHistoryCallDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
