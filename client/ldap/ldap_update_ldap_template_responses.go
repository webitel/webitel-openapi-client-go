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

// LDAPUpdateLDAPTemplateReader is a Reader for the LDAPUpdateLDAPTemplate structure.
type LDAPUpdateLDAPTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LDAPUpdateLDAPTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLDAPUpdateLDAPTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /ldap/{template.catalog.id}/templates/{template.id}] LDAP_UpdateLDAPTemplate", response, response.Code())
	}
}

// NewLDAPUpdateLDAPTemplateOK creates a LDAPUpdateLDAPTemplateOK with default headers values
func NewLDAPUpdateLDAPTemplateOK() *LDAPUpdateLDAPTemplateOK {
	return &LDAPUpdateLDAPTemplateOK{}
}

/*
LDAPUpdateLDAPTemplateOK describes a response with status code 200, with default header values.

A successful response.
*/
type LDAPUpdateLDAPTemplateOK struct {
	Payload *models.APILDAPTemplate
}

// IsSuccess returns true when this ldap update Ldap template Ok response has a 2xx status code
func (o *LDAPUpdateLDAPTemplateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ldap update Ldap template Ok response has a 3xx status code
func (o *LDAPUpdateLDAPTemplateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ldap update Ldap template Ok response has a 4xx status code
func (o *LDAPUpdateLDAPTemplateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ldap update Ldap template Ok response has a 5xx status code
func (o *LDAPUpdateLDAPTemplateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ldap update Ldap template Ok response a status code equal to that given
func (o *LDAPUpdateLDAPTemplateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ldap update Ldap template Ok response
func (o *LDAPUpdateLDAPTemplateOK) Code() int {
	return 200
}

func (o *LDAPUpdateLDAPTemplateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ldap/{template.catalog.id}/templates/{template.id}][%d] ldapUpdateLdapTemplateOk %s", 200, payload)
}

func (o *LDAPUpdateLDAPTemplateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ldap/{template.catalog.id}/templates/{template.id}][%d] ldapUpdateLdapTemplateOk %s", 200, payload)
}

func (o *LDAPUpdateLDAPTemplateOK) GetPayload() *models.APILDAPTemplate {
	return o.Payload
}

func (o *LDAPUpdateLDAPTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APILDAPTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
