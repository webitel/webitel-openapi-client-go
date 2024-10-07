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

// PhonesLocatePhoneReader is a Reader for the PhonesLocatePhone structure.
type PhonesLocatePhoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PhonesLocatePhoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPhonesLocatePhoneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /contacts/{contact_id}/phones/{etag}] Phones_LocatePhone", response, response.Code())
	}
}

// NewPhonesLocatePhoneOK creates a PhonesLocatePhoneOK with default headers values
func NewPhonesLocatePhoneOK() *PhonesLocatePhoneOK {
	return &PhonesLocatePhoneOK{}
}

/*
PhonesLocatePhoneOK describes a response with status code 200, with default header values.

A successful response.
*/
type PhonesLocatePhoneOK struct {
	Payload *models.WebitelContactsPhoneNumber
}

// IsSuccess returns true when this phones locate phone Ok response has a 2xx status code
func (o *PhonesLocatePhoneOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this phones locate phone Ok response has a 3xx status code
func (o *PhonesLocatePhoneOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this phones locate phone Ok response has a 4xx status code
func (o *PhonesLocatePhoneOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this phones locate phone Ok response has a 5xx status code
func (o *PhonesLocatePhoneOK) IsServerError() bool {
	return false
}

// IsCode returns true when this phones locate phone Ok response a status code equal to that given
func (o *PhonesLocatePhoneOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the phones locate phone Ok response
func (o *PhonesLocatePhoneOK) Code() int {
	return 200
}

func (o *PhonesLocatePhoneOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /contacts/{contact_id}/phones/{etag}][%d] phonesLocatePhoneOk %s", 200, payload)
}

func (o *PhonesLocatePhoneOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /contacts/{contact_id}/phones/{etag}][%d] phonesLocatePhoneOk %s", 200, payload)
}

func (o *PhonesLocatePhoneOK) GetPayload() *models.WebitelContactsPhoneNumber {
	return o.Payload
}

func (o *PhonesLocatePhoneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebitelContactsPhoneNumber)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
