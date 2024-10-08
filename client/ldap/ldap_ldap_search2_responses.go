// Code generated by go-swagger; DO NOT EDIT.

package ldap

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

// LDAPLDAPSearch2Reader is a Reader for the LDAPLDAPSearch2 structure.
type LDAPLDAPSearch2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LDAPLDAPSearch2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLDAPLDAPSearch2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ldap/{catalog_id}/search] LDAP_LDAPSearch2", response, response.Code())
	}
}

// NewLDAPLDAPSearch2OK creates a LDAPLDAPSearch2OK with default headers values
func NewLDAPLDAPSearch2OK() *LDAPLDAPSearch2OK {
	return &LDAPLDAPSearch2OK{}
}

/*
LDAPLDAPSearch2OK describes a response with status code 200, with default header values.

A successful response.
*/
type LDAPLDAPSearch2OK struct {
	Payload *models.APILDAPSearchResponse
}

// IsSuccess returns true when this ldap Ldap search2 Ok response has a 2xx status code
func (o *LDAPLDAPSearch2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ldap Ldap search2 Ok response has a 3xx status code
func (o *LDAPLDAPSearch2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ldap Ldap search2 Ok response has a 4xx status code
func (o *LDAPLDAPSearch2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ldap Ldap search2 Ok response has a 5xx status code
func (o *LDAPLDAPSearch2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this ldap Ldap search2 Ok response a status code equal to that given
func (o *LDAPLDAPSearch2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ldap Ldap search2 Ok response
func (o *LDAPLDAPSearch2OK) Code() int {
	return 200
}

func (o *LDAPLDAPSearch2OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ldap/{catalog_id}/search][%d] ldapLdapSearch2Ok %s", 200, payload)
}

func (o *LDAPLDAPSearch2OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ldap/{catalog_id}/search][%d] ldapLdapSearch2Ok %s", 200, payload)
}

func (o *LDAPLDAPSearch2OK) GetPayload() *models.APILDAPSearchResponse {
	return o.Payload
}

func (o *LDAPLDAPSearch2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APILDAPSearchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
