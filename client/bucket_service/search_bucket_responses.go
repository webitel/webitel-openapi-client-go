// Code generated by go-swagger; DO NOT EDIT.

package bucket_service

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

// SearchBucketReader is a Reader for the SearchBucket structure.
type SearchBucketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchBucketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchBucketOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSearchBucketDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchBucketOK creates a SearchBucketOK with default headers values
func NewSearchBucketOK() *SearchBucketOK {
	return &SearchBucketOK{}
}

/*
SearchBucketOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchBucketOK struct {
	Payload *models.EngineListBucket
}

// IsSuccess returns true when this search bucket Ok response has a 2xx status code
func (o *SearchBucketOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search bucket Ok response has a 3xx status code
func (o *SearchBucketOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search bucket Ok response has a 4xx status code
func (o *SearchBucketOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search bucket Ok response has a 5xx status code
func (o *SearchBucketOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search bucket Ok response a status code equal to that given
func (o *SearchBucketOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search bucket Ok response
func (o *SearchBucketOK) Code() int {
	return 200
}

func (o *SearchBucketOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/buckets][%d] searchBucketOk %s", 200, payload)
}

func (o *SearchBucketOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/buckets][%d] searchBucketOk %s", 200, payload)
}

func (o *SearchBucketOK) GetPayload() *models.EngineListBucket {
	return o.Payload
}

func (o *SearchBucketOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineListBucket)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchBucketDefault creates a SearchBucketDefault with default headers values
func NewSearchBucketDefault(code int) *SearchBucketDefault {
	return &SearchBucketDefault{
		_statusCode: code,
	}
}

/*
SearchBucketDefault describes a response with status code -1, with default header values.

Server error
*/
type SearchBucketDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this search bucket default response has a 2xx status code
func (o *SearchBucketDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this search bucket default response has a 3xx status code
func (o *SearchBucketDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this search bucket default response has a 4xx status code
func (o *SearchBucketDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this search bucket default response has a 5xx status code
func (o *SearchBucketDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this search bucket default response a status code equal to that given
func (o *SearchBucketDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the search bucket default response
func (o *SearchBucketDefault) Code() int {
	return o._statusCode
}

func (o *SearchBucketDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/buckets][%d] SearchBucket default %s", o._statusCode, payload)
}

func (o *SearchBucketDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/buckets][%d] SearchBucket default %s", o._statusCode, payload)
}

func (o *SearchBucketDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *SearchBucketDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
