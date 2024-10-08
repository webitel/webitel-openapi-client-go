// Code generated by go-swagger; DO NOT EDIT.

package labels

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

// LabelsResetLabelsReader is a Reader for the LabelsResetLabels structure.
type LabelsResetLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LabelsResetLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLabelsResetLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /contacts/{contact_id}/labels] Labels_ResetLabels", response, response.Code())
	}
}

// NewLabelsResetLabelsOK creates a LabelsResetLabelsOK with default headers values
func NewLabelsResetLabelsOK() *LabelsResetLabelsOK {
	return &LabelsResetLabelsOK{}
}

/*
LabelsResetLabelsOK describes a response with status code 200, with default header values.

A successful response.
*/
type LabelsResetLabelsOK struct {
	Payload *models.WebitelContactsLabelList
}

// IsSuccess returns true when this labels reset labels Ok response has a 2xx status code
func (o *LabelsResetLabelsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this labels reset labels Ok response has a 3xx status code
func (o *LabelsResetLabelsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this labels reset labels Ok response has a 4xx status code
func (o *LabelsResetLabelsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this labels reset labels Ok response has a 5xx status code
func (o *LabelsResetLabelsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this labels reset labels Ok response a status code equal to that given
func (o *LabelsResetLabelsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the labels reset labels Ok response
func (o *LabelsResetLabelsOK) Code() int {
	return 200
}

func (o *LabelsResetLabelsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /contacts/{contact_id}/labels][%d] labelsResetLabelsOk %s", 200, payload)
}

func (o *LabelsResetLabelsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /contacts/{contact_id}/labels][%d] labelsResetLabelsOk %s", 200, payload)
}

func (o *LabelsResetLabelsOK) GetPayload() *models.WebitelContactsLabelList {
	return o.Payload
}

func (o *LabelsResetLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebitelContactsLabelList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
