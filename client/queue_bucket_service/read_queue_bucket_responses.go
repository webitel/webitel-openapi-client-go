// Code generated by go-swagger; DO NOT EDIT.

package queue_bucket_service

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

// ReadQueueBucketReader is a Reader for the ReadQueueBucket structure.
type ReadQueueBucketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadQueueBucketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadQueueBucketOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewReadQueueBucketDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadQueueBucketOK creates a ReadQueueBucketOK with default headers values
func NewReadQueueBucketOK() *ReadQueueBucketOK {
	return &ReadQueueBucketOK{}
}

/*
ReadQueueBucketOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadQueueBucketOK struct {
	Payload *models.EngineQueueBucket
}

// IsSuccess returns true when this read queue bucket Ok response has a 2xx status code
func (o *ReadQueueBucketOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read queue bucket Ok response has a 3xx status code
func (o *ReadQueueBucketOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read queue bucket Ok response has a 4xx status code
func (o *ReadQueueBucketOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read queue bucket Ok response has a 5xx status code
func (o *ReadQueueBucketOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read queue bucket Ok response a status code equal to that given
func (o *ReadQueueBucketOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read queue bucket Ok response
func (o *ReadQueueBucketOK) Code() int {
	return 200
}

func (o *ReadQueueBucketOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/queues/{queue_id}/buckets/{id}][%d] readQueueBucketOk %s", 200, payload)
}

func (o *ReadQueueBucketOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/queues/{queue_id}/buckets/{id}][%d] readQueueBucketOk %s", 200, payload)
}

func (o *ReadQueueBucketOK) GetPayload() *models.EngineQueueBucket {
	return o.Payload
}

func (o *ReadQueueBucketOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineQueueBucket)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadQueueBucketDefault creates a ReadQueueBucketDefault with default headers values
func NewReadQueueBucketDefault(code int) *ReadQueueBucketDefault {
	return &ReadQueueBucketDefault{
		_statusCode: code,
	}
}

/*
ReadQueueBucketDefault describes a response with status code -1, with default header values.

Server error
*/
type ReadQueueBucketDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this read queue bucket default response has a 2xx status code
func (o *ReadQueueBucketDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read queue bucket default response has a 3xx status code
func (o *ReadQueueBucketDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read queue bucket default response has a 4xx status code
func (o *ReadQueueBucketDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read queue bucket default response has a 5xx status code
func (o *ReadQueueBucketDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read queue bucket default response a status code equal to that given
func (o *ReadQueueBucketDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read queue bucket default response
func (o *ReadQueueBucketDefault) Code() int {
	return o._statusCode
}

func (o *ReadQueueBucketDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/queues/{queue_id}/buckets/{id}][%d] ReadQueueBucket default %s", o._statusCode, payload)
}

func (o *ReadQueueBucketDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/queues/{queue_id}/buckets/{id}][%d] ReadQueueBucket default %s", o._statusCode, payload)
}

func (o *ReadQueueBucketDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *ReadQueueBucketDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
