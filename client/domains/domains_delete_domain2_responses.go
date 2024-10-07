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

// DomainsDeleteDomain2Reader is a Reader for the DomainsDeleteDomain2 structure.
type DomainsDeleteDomain2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainsDeleteDomain2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDomainsDeleteDomain2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /domains/{dc}] Domains_DeleteDomain2", response, response.Code())
	}
}

// NewDomainsDeleteDomain2OK creates a DomainsDeleteDomain2OK with default headers values
func NewDomainsDeleteDomain2OK() *DomainsDeleteDomain2OK {
	return &DomainsDeleteDomain2OK{}
}

/*
DomainsDeleteDomain2OK describes a response with status code 200, with default header values.

A successful response.
*/
type DomainsDeleteDomain2OK struct {
	Payload models.APIDeleteDomainResponse
}

// IsSuccess returns true when this domains delete domain2 Ok response has a 2xx status code
func (o *DomainsDeleteDomain2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this domains delete domain2 Ok response has a 3xx status code
func (o *DomainsDeleteDomain2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domains delete domain2 Ok response has a 4xx status code
func (o *DomainsDeleteDomain2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this domains delete domain2 Ok response has a 5xx status code
func (o *DomainsDeleteDomain2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this domains delete domain2 Ok response a status code equal to that given
func (o *DomainsDeleteDomain2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the domains delete domain2 Ok response
func (o *DomainsDeleteDomain2OK) Code() int {
	return 200
}

func (o *DomainsDeleteDomain2OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /domains/{dc}][%d] domainsDeleteDomain2Ok %s", 200, payload)
}

func (o *DomainsDeleteDomain2OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /domains/{dc}][%d] domainsDeleteDomain2Ok %s", 200, payload)
}

func (o *DomainsDeleteDomain2OK) GetPayload() models.APIDeleteDomainResponse {
	return o.Payload
}

func (o *DomainsDeleteDomain2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
