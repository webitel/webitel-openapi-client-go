// Code generated by go-swagger; DO NOT EDIT.

package phones

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

// PhonesMergePhonesReader is a Reader for the PhonesMergePhones structure.
type PhonesMergePhonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PhonesMergePhonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPhonesMergePhonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /contacts/{contact_id}/phones] Phones_MergePhones", response, response.Code())
	}
}

// NewPhonesMergePhonesOK creates a PhonesMergePhonesOK with default headers values
func NewPhonesMergePhonesOK() *PhonesMergePhonesOK {
	return &PhonesMergePhonesOK{}
}

/*
PhonesMergePhonesOK describes a response with status code 200, with default header values.

A successful response.
*/
type PhonesMergePhonesOK struct {
	Payload *models.WebitelContactsPhoneList
}

// IsSuccess returns true when this phones merge phones Ok response has a 2xx status code
func (o *PhonesMergePhonesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this phones merge phones Ok response has a 3xx status code
func (o *PhonesMergePhonesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this phones merge phones Ok response has a 4xx status code
func (o *PhonesMergePhonesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this phones merge phones Ok response has a 5xx status code
func (o *PhonesMergePhonesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this phones merge phones Ok response a status code equal to that given
func (o *PhonesMergePhonesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the phones merge phones Ok response
func (o *PhonesMergePhonesOK) Code() int {
	return 200
}

func (o *PhonesMergePhonesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /contacts/{contact_id}/phones][%d] phonesMergePhonesOk %s", 200, payload)
}

func (o *PhonesMergePhonesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /contacts/{contact_id}/phones][%d] phonesMergePhonesOk %s", 200, payload)
}

func (o *PhonesMergePhonesOK) GetPayload() *models.WebitelContactsPhoneList {
	return o.Payload
}

func (o *PhonesMergePhonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebitelContactsPhoneList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
