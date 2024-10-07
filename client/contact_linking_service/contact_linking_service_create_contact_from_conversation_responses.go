// Code generated by go-swagger; DO NOT EDIT.

package contact_linking_service

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

// ContactLinkingServiceCreateContactFromConversationReader is a Reader for the ContactLinkingServiceCreateContactFromConversation structure.
type ContactLinkingServiceCreateContactFromConversationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContactLinkingServiceCreateContactFromConversationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContactLinkingServiceCreateContactFromConversationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /chat/{conversation_id}/contact] ContactLinkingService_CreateContactFromConversation", response, response.Code())
	}
}

// NewContactLinkingServiceCreateContactFromConversationOK creates a ContactLinkingServiceCreateContactFromConversationOK with default headers values
func NewContactLinkingServiceCreateContactFromConversationOK() *ContactLinkingServiceCreateContactFromConversationOK {
	return &ContactLinkingServiceCreateContactFromConversationOK{}
}

/*
ContactLinkingServiceCreateContactFromConversationOK describes a response with status code 200, with default header values.

A successful response.
*/
type ContactLinkingServiceCreateContactFromConversationOK struct {
	Payload *models.WebitelChatLookup
}

// IsSuccess returns true when this contact linking service create contact from conversation Ok response has a 2xx status code
func (o *ContactLinkingServiceCreateContactFromConversationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this contact linking service create contact from conversation Ok response has a 3xx status code
func (o *ContactLinkingServiceCreateContactFromConversationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this contact linking service create contact from conversation Ok response has a 4xx status code
func (o *ContactLinkingServiceCreateContactFromConversationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this contact linking service create contact from conversation Ok response has a 5xx status code
func (o *ContactLinkingServiceCreateContactFromConversationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this contact linking service create contact from conversation Ok response a status code equal to that given
func (o *ContactLinkingServiceCreateContactFromConversationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the contact linking service create contact from conversation Ok response
func (o *ContactLinkingServiceCreateContactFromConversationOK) Code() int {
	return 200
}

func (o *ContactLinkingServiceCreateContactFromConversationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /chat/{conversation_id}/contact][%d] contactLinkingServiceCreateContactFromConversationOk %s", 200, payload)
}

func (o *ContactLinkingServiceCreateContactFromConversationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /chat/{conversation_id}/contact][%d] contactLinkingServiceCreateContactFromConversationOk %s", 200, payload)
}

func (o *ContactLinkingServiceCreateContactFromConversationOK) GetPayload() *models.WebitelChatLookup {
	return o.Payload
}

func (o *ContactLinkingServiceCreateContactFromConversationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebitelChatLookup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
