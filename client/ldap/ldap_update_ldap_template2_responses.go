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

// LDAPUpdateLDAPTemplate2Reader is a Reader for the LDAPUpdateLDAPTemplate2 structure.
type LDAPUpdateLDAPTemplate2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LDAPUpdateLDAPTemplate2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLDAPUpdateLDAPTemplate2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /ldap/{template.catalog.id}/templates/{template.id}] LDAP_UpdateLDAPTemplate2", response, response.Code())
	}
}

// NewLDAPUpdateLDAPTemplate2OK creates a LDAPUpdateLDAPTemplate2OK with default headers values
func NewLDAPUpdateLDAPTemplate2OK() *LDAPUpdateLDAPTemplate2OK {
	return &LDAPUpdateLDAPTemplate2OK{}
}

/*
LDAPUpdateLDAPTemplate2OK describes a response with status code 200, with default header values.

A successful response.
*/
type LDAPUpdateLDAPTemplate2OK struct {
	Payload *models.APILDAPTemplate
}

// IsSuccess returns true when this ldap update Ldap template2 Ok response has a 2xx status code
func (o *LDAPUpdateLDAPTemplate2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ldap update Ldap template2 Ok response has a 3xx status code
func (o *LDAPUpdateLDAPTemplate2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ldap update Ldap template2 Ok response has a 4xx status code
func (o *LDAPUpdateLDAPTemplate2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ldap update Ldap template2 Ok response has a 5xx status code
func (o *LDAPUpdateLDAPTemplate2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this ldap update Ldap template2 Ok response a status code equal to that given
func (o *LDAPUpdateLDAPTemplate2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ldap update Ldap template2 Ok response
func (o *LDAPUpdateLDAPTemplate2OK) Code() int {
	return 200
}

func (o *LDAPUpdateLDAPTemplate2OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ldap/{template.catalog.id}/templates/{template.id}][%d] ldapUpdateLdapTemplate2Ok %s", 200, payload)
}

func (o *LDAPUpdateLDAPTemplate2OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ldap/{template.catalog.id}/templates/{template.id}][%d] ldapUpdateLdapTemplate2Ok %s", 200, payload)
}

func (o *LDAPUpdateLDAPTemplate2OK) GetPayload() *models.APILDAPTemplate {
	return o.Payload
}

func (o *LDAPUpdateLDAPTemplate2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APILDAPTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
