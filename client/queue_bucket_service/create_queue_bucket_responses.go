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

// CreateQueueBucketReader is a Reader for the CreateQueueBucket structure.
type CreateQueueBucketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateQueueBucketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateQueueBucketOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateQueueBucketDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateQueueBucketOK creates a CreateQueueBucketOK with default headers values
func NewCreateQueueBucketOK() *CreateQueueBucketOK {
	return &CreateQueueBucketOK{}
}

/*
CreateQueueBucketOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateQueueBucketOK struct {
	Payload *models.EngineQueueBucket
}

// IsSuccess returns true when this create queue bucket Ok response has a 2xx status code
func (o *CreateQueueBucketOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create queue bucket Ok response has a 3xx status code
func (o *CreateQueueBucketOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create queue bucket Ok response has a 4xx status code
func (o *CreateQueueBucketOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create queue bucket Ok response has a 5xx status code
func (o *CreateQueueBucketOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create queue bucket Ok response a status code equal to that given
func (o *CreateQueueBucketOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create queue bucket Ok response
func (o *CreateQueueBucketOK) Code() int {
	return 200
}

func (o *CreateQueueBucketOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/queues/{queue_id}/buckets][%d] createQueueBucketOk %s", 200, payload)
}

func (o *CreateQueueBucketOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/queues/{queue_id}/buckets][%d] createQueueBucketOk %s", 200, payload)
}

func (o *CreateQueueBucketOK) GetPayload() *models.EngineQueueBucket {
	return o.Payload
}

func (o *CreateQueueBucketOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineQueueBucket)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQueueBucketDefault creates a CreateQueueBucketDefault with default headers values
func NewCreateQueueBucketDefault(code int) *CreateQueueBucketDefault {
	return &CreateQueueBucketDefault{
		_statusCode: code,
	}
}

/*
CreateQueueBucketDefault describes a response with status code -1, with default header values.

Server error
*/
type CreateQueueBucketDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this create queue bucket default response has a 2xx status code
func (o *CreateQueueBucketDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create queue bucket default response has a 3xx status code
func (o *CreateQueueBucketDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create queue bucket default response has a 4xx status code
func (o *CreateQueueBucketDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create queue bucket default response has a 5xx status code
func (o *CreateQueueBucketDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create queue bucket default response a status code equal to that given
func (o *CreateQueueBucketDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create queue bucket default response
func (o *CreateQueueBucketDefault) Code() int {
	return o._statusCode
}

func (o *CreateQueueBucketDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/queues/{queue_id}/buckets][%d] CreateQueueBucket default %s", o._statusCode, payload)
}

func (o *CreateQueueBucketDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /call_center/queues/{queue_id}/buckets][%d] CreateQueueBucket default %s", o._statusCode, payload)
}

func (o *CreateQueueBucketDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *CreateQueueBucketDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
