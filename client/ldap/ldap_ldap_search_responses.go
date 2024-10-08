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

// LDAPLDAPSearchReader is a Reader for the LDAPLDAPSearch structure.
type LDAPLDAPSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LDAPLDAPSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLDAPLDAPSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /ldap/{catalog_id}/search] LDAP_LDAPSearch", response, response.Code())
	}
}

// NewLDAPLDAPSearchOK creates a LDAPLDAPSearchOK with default headers values
func NewLDAPLDAPSearchOK() *LDAPLDAPSearchOK {
	return &LDAPLDAPSearchOK{}
}

/*
LDAPLDAPSearchOK describes a response with status code 200, with default header values.

A successful response.
*/
type LDAPLDAPSearchOK struct {
	Payload *models.APILDAPSearchResponse
}

// IsSuccess returns true when this ldap Ldap search Ok response has a 2xx status code
func (o *LDAPLDAPSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ldap Ldap search Ok response has a 3xx status code
func (o *LDAPLDAPSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ldap Ldap search Ok response has a 4xx status code
func (o *LDAPLDAPSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ldap Ldap search Ok response has a 5xx status code
func (o *LDAPLDAPSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ldap Ldap search Ok response a status code equal to that given
func (o *LDAPLDAPSearchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ldap Ldap search Ok response
func (o *LDAPLDAPSearchOK) Code() int {
	return 200
}

func (o *LDAPLDAPSearchOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ldap/{catalog_id}/search][%d] ldapLdapSearchOk %s", 200, payload)
}

func (o *LDAPLDAPSearchOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ldap/{catalog_id}/search][%d] ldapLdapSearchOk %s", 200, payload)
}

func (o *LDAPLDAPSearchOK) GetPayload() *models.APILDAPSearchResponse {
	return o.Payload
}

func (o *LDAPLDAPSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APILDAPSearchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
