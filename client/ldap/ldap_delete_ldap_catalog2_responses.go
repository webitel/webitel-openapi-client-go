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

// LDAPDeleteLDAPCatalog2Reader is a Reader for the LDAPDeleteLDAPCatalog2 structure.
type LDAPDeleteLDAPCatalog2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LDAPDeleteLDAPCatalog2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLDAPDeleteLDAPCatalog2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ldap] LDAP_DeleteLDAPCatalog2", response, response.Code())
	}
}

// NewLDAPDeleteLDAPCatalog2OK creates a LDAPDeleteLDAPCatalog2OK with default headers values
func NewLDAPDeleteLDAPCatalog2OK() *LDAPDeleteLDAPCatalog2OK {
	return &LDAPDeleteLDAPCatalog2OK{}
}

/*
LDAPDeleteLDAPCatalog2OK describes a response with status code 200, with default header values.

A successful response.
*/
type LDAPDeleteLDAPCatalog2OK struct {
	Payload *models.APILDAPCatalog
}

// IsSuccess returns true when this ldap delete Ldap catalog2 Ok response has a 2xx status code
func (o *LDAPDeleteLDAPCatalog2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ldap delete Ldap catalog2 Ok response has a 3xx status code
func (o *LDAPDeleteLDAPCatalog2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ldap delete Ldap catalog2 Ok response has a 4xx status code
func (o *LDAPDeleteLDAPCatalog2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ldap delete Ldap catalog2 Ok response has a 5xx status code
func (o *LDAPDeleteLDAPCatalog2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this ldap delete Ldap catalog2 Ok response a status code equal to that given
func (o *LDAPDeleteLDAPCatalog2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ldap delete Ldap catalog2 Ok response
func (o *LDAPDeleteLDAPCatalog2OK) Code() int {
	return 200
}

func (o *LDAPDeleteLDAPCatalog2OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ldap][%d] ldapDeleteLdapCatalog2Ok %s", 200, payload)
}

func (o *LDAPDeleteLDAPCatalog2OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ldap][%d] ldapDeleteLdapCatalog2Ok %s", 200, payload)
}

func (o *LDAPDeleteLDAPCatalog2OK) GetPayload() *models.APILDAPCatalog {
	return o.Payload
}

func (o *LDAPDeleteLDAPCatalog2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APILDAPCatalog)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
