// Code generated by go-swagger; DO NOT EDIT.

package o_auth2_federation

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

// OAuth2FederationDeleteOAuthServiceReader is a Reader for the OAuth2FederationDeleteOAuthService structure.
type OAuth2FederationDeleteOAuthServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OAuth2FederationDeleteOAuthServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOAuth2FederationDeleteOAuthServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /oauth/apps] OAuth2Federation_DeleteOAuthService", response, response.Code())
	}
}

// NewOAuth2FederationDeleteOAuthServiceOK creates a OAuth2FederationDeleteOAuthServiceOK with default headers values
func NewOAuth2FederationDeleteOAuthServiceOK() *OAuth2FederationDeleteOAuthServiceOK {
	return &OAuth2FederationDeleteOAuthServiceOK{}
}

/*
OAuth2FederationDeleteOAuthServiceOK describes a response with status code 200, with default header values.

A successful response.
*/
type OAuth2FederationDeleteOAuthServiceOK struct {
	Payload models.APIDeleteOAuthServiceResponse
}

// IsSuccess returns true when this o auth2 federation delete o auth service Ok response has a 2xx status code
func (o *OAuth2FederationDeleteOAuthServiceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this o auth2 federation delete o auth service Ok response has a 3xx status code
func (o *OAuth2FederationDeleteOAuthServiceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this o auth2 federation delete o auth service Ok response has a 4xx status code
func (o *OAuth2FederationDeleteOAuthServiceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this o auth2 federation delete o auth service Ok response has a 5xx status code
func (o *OAuth2FederationDeleteOAuthServiceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this o auth2 federation delete o auth service Ok response a status code equal to that given
func (o *OAuth2FederationDeleteOAuthServiceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the o auth2 federation delete o auth service Ok response
func (o *OAuth2FederationDeleteOAuthServiceOK) Code() int {
	return 200
}

func (o *OAuth2FederationDeleteOAuthServiceOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /oauth/apps][%d] oAuth2FederationDeleteOAuthServiceOk %s", 200, payload)
}

func (o *OAuth2FederationDeleteOAuthServiceOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /oauth/apps][%d] oAuth2FederationDeleteOAuthServiceOk %s", 200, payload)
}

func (o *OAuth2FederationDeleteOAuthServiceOK) GetPayload() models.APIDeleteOAuthServiceResponse {
	return o.Payload
}

func (o *OAuth2FederationDeleteOAuthServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
