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

// PhonesResetPhonesReader is a Reader for the PhonesResetPhones structure.
type PhonesResetPhonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PhonesResetPhonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPhonesResetPhonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /contacts/{contact_id}/phones] Phones_ResetPhones", response, response.Code())
	}
}

// NewPhonesResetPhonesOK creates a PhonesResetPhonesOK with default headers values
func NewPhonesResetPhonesOK() *PhonesResetPhonesOK {
	return &PhonesResetPhonesOK{}
}

/*
PhonesResetPhonesOK describes a response with status code 200, with default header values.

A successful response.
*/
type PhonesResetPhonesOK struct {
	Payload *models.WebitelContactsPhoneList
}

// IsSuccess returns true when this phones reset phones Ok response has a 2xx status code
func (o *PhonesResetPhonesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this phones reset phones Ok response has a 3xx status code
func (o *PhonesResetPhonesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this phones reset phones Ok response has a 4xx status code
func (o *PhonesResetPhonesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this phones reset phones Ok response has a 5xx status code
func (o *PhonesResetPhonesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this phones reset phones Ok response a status code equal to that given
func (o *PhonesResetPhonesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the phones reset phones Ok response
func (o *PhonesResetPhonesOK) Code() int {
	return 200
}

func (o *PhonesResetPhonesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /contacts/{contact_id}/phones][%d] phonesResetPhonesOk %s", 200, payload)
}

func (o *PhonesResetPhonesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /contacts/{contact_id}/phones][%d] phonesResetPhonesOk %s", 200, payload)
}

func (o *PhonesResetPhonesOK) GetPayload() *models.WebitelContactsPhoneList {
	return o.Payload
}

func (o *PhonesResetPhonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebitelContactsPhoneList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
