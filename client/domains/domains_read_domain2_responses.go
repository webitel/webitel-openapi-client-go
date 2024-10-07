// Code generated by go-swagger; DO NOT EDIT.

package domains

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

// DomainsReadDomain2Reader is a Reader for the DomainsReadDomain2 structure.
type DomainsReadDomain2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainsReadDomain2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDomainsReadDomain2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /domains/{dc}] Domains_ReadDomain2", response, response.Code())
	}
}

// NewDomainsReadDomain2OK creates a DomainsReadDomain2OK with default headers values
func NewDomainsReadDomain2OK() *DomainsReadDomain2OK {
	return &DomainsReadDomain2OK{}
}

/*
DomainsReadDomain2OK describes a response with status code 200, with default header values.

A successful response.
*/
type DomainsReadDomain2OK struct {
	Payload *models.APIReadDomainResponse
}

// IsSuccess returns true when this domains read domain2 Ok response has a 2xx status code
func (o *DomainsReadDomain2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this domains read domain2 Ok response has a 3xx status code
func (o *DomainsReadDomain2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domains read domain2 Ok response has a 4xx status code
func (o *DomainsReadDomain2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this domains read domain2 Ok response has a 5xx status code
func (o *DomainsReadDomain2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this domains read domain2 Ok response a status code equal to that given
func (o *DomainsReadDomain2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the domains read domain2 Ok response
func (o *DomainsReadDomain2OK) Code() int {
	return 200
}

func (o *DomainsReadDomain2OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /domains/{dc}][%d] domainsReadDomain2Ok %s", 200, payload)
}

func (o *DomainsReadDomain2OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /domains/{dc}][%d] domainsReadDomain2Ok %s", 200, payload)
}

func (o *DomainsReadDomain2OK) GetPayload() *models.APIReadDomainResponse {
	return o.Payload
}

func (o *DomainsReadDomain2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIReadDomainResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
