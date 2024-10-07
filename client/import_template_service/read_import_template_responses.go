// Code generated by go-swagger; DO NOT EDIT.

package import_template_service

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

// ReadImportTemplateReader is a Reader for the ReadImportTemplate structure.
type ReadImportTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadImportTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadImportTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /storage/import_templates/{id}] ReadImportTemplate", response, response.Code())
	}
}

// NewReadImportTemplateOK creates a ReadImportTemplateOK with default headers values
func NewReadImportTemplateOK() *ReadImportTemplateOK {
	return &ReadImportTemplateOK{}
}

/*
ReadImportTemplateOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadImportTemplateOK struct {
	Payload *models.StorageImportTemplate
}

// IsSuccess returns true when this read import template Ok response has a 2xx status code
func (o *ReadImportTemplateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read import template Ok response has a 3xx status code
func (o *ReadImportTemplateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read import template Ok response has a 4xx status code
func (o *ReadImportTemplateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read import template Ok response has a 5xx status code
func (o *ReadImportTemplateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read import template Ok response a status code equal to that given
func (o *ReadImportTemplateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read import template Ok response
func (o *ReadImportTemplateOK) Code() int {
	return 200
}

func (o *ReadImportTemplateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/import_templates/{id}][%d] readImportTemplateOk %s", 200, payload)
}

func (o *ReadImportTemplateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/import_templates/{id}][%d] readImportTemplateOk %s", 200, payload)
}

func (o *ReadImportTemplateOK) GetPayload() *models.StorageImportTemplate {
	return o.Payload
}

func (o *ReadImportTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageImportTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
