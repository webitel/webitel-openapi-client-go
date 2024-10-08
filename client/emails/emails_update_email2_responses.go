// Code generated by go-swagger; DO NOT EDIT.

package emails

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

// EmailsUpdateEmail2Reader is a Reader for the EmailsUpdateEmail2 structure.
type EmailsUpdateEmail2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmailsUpdateEmail2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmailsUpdateEmail2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /contacts/{contact_id}/emails/{etag}] Emails_UpdateEmail2", response, response.Code())
	}
}

// NewEmailsUpdateEmail2OK creates a EmailsUpdateEmail2OK with default headers values
func NewEmailsUpdateEmail2OK() *EmailsUpdateEmail2OK {
	return &EmailsUpdateEmail2OK{}
}

/*
EmailsUpdateEmail2OK describes a response with status code 200, with default header values.

A successful response.
*/
type EmailsUpdateEmail2OK struct {
	Payload *models.WebitelContactsEmailList
}

// IsSuccess returns true when this emails update email2 Ok response has a 2xx status code
func (o *EmailsUpdateEmail2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this emails update email2 Ok response has a 3xx status code
func (o *EmailsUpdateEmail2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this emails update email2 Ok response has a 4xx status code
func (o *EmailsUpdateEmail2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this emails update email2 Ok response has a 5xx status code
func (o *EmailsUpdateEmail2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this emails update email2 Ok response a status code equal to that given
func (o *EmailsUpdateEmail2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the emails update email2 Ok response
func (o *EmailsUpdateEmail2OK) Code() int {
	return 200
}

func (o *EmailsUpdateEmail2OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /contacts/{contact_id}/emails/{etag}][%d] emailsUpdateEmail2Ok %s", 200, payload)
}

func (o *EmailsUpdateEmail2OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /contacts/{contact_id}/emails/{etag}][%d] emailsUpdateEmail2Ok %s", 200, payload)
}

func (o *EmailsUpdateEmail2OK) GetPayload() *models.WebitelContactsEmailList {
	return o.Payload
}

func (o *EmailsUpdateEmail2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebitelContactsEmailList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
