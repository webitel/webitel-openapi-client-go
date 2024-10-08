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

// LDAPUpdateLDAPCatalog2Reader is a Reader for the LDAPUpdateLDAPCatalog2 structure.
type LDAPUpdateLDAPCatalog2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LDAPUpdateLDAPCatalog2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLDAPUpdateLDAPCatalog2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /ldap/{catalog.id}] LDAP_UpdateLDAPCatalog2", response, response.Code())
	}
}

// NewLDAPUpdateLDAPCatalog2OK creates a LDAPUpdateLDAPCatalog2OK with default headers values
func NewLDAPUpdateLDAPCatalog2OK() *LDAPUpdateLDAPCatalog2OK {
	return &LDAPUpdateLDAPCatalog2OK{}
}

/*
LDAPUpdateLDAPCatalog2OK describes a response with status code 200, with default header values.

A successful response.
*/
type LDAPUpdateLDAPCatalog2OK struct {
	Payload *models.APILDAPCatalog
}

// IsSuccess returns true when this ldap update Ldap catalog2 Ok response has a 2xx status code
func (o *LDAPUpdateLDAPCatalog2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ldap update Ldap catalog2 Ok response has a 3xx status code
func (o *LDAPUpdateLDAPCatalog2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ldap update Ldap catalog2 Ok response has a 4xx status code
func (o *LDAPUpdateLDAPCatalog2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ldap update Ldap catalog2 Ok response has a 5xx status code
func (o *LDAPUpdateLDAPCatalog2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this ldap update Ldap catalog2 Ok response a status code equal to that given
func (o *LDAPUpdateLDAPCatalog2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ldap update Ldap catalog2 Ok response
func (o *LDAPUpdateLDAPCatalog2OK) Code() int {
	return 200
}

func (o *LDAPUpdateLDAPCatalog2OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ldap/{catalog.id}][%d] ldapUpdateLdapCatalog2Ok %s", 200, payload)
}

func (o *LDAPUpdateLDAPCatalog2OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ldap/{catalog.id}][%d] ldapUpdateLdapCatalog2Ok %s", 200, payload)
}

func (o *LDAPUpdateLDAPCatalog2OK) GetPayload() *models.APILDAPCatalog {
	return o.Payload
}

func (o *LDAPUpdateLDAPCatalog2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APILDAPCatalog)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
