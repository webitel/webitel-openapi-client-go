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

// OAuth2FederationUpdateOAuthServiceReader is a Reader for the OAuth2FederationUpdateOAuthService structure.
type OAuth2FederationUpdateOAuthServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OAuth2FederationUpdateOAuthServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOAuth2FederationUpdateOAuthServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /oauth/apps/{changes.id}] OAuth2Federation_UpdateOAuthService", response, response.Code())
	}
}

// NewOAuth2FederationUpdateOAuthServiceOK creates a OAuth2FederationUpdateOAuthServiceOK with default headers values
func NewOAuth2FederationUpdateOAuthServiceOK() *OAuth2FederationUpdateOAuthServiceOK {
	return &OAuth2FederationUpdateOAuthServiceOK{}
}

/*
OAuth2FederationUpdateOAuthServiceOK describes a response with status code 200, with default header values.

A successful response.
*/
type OAuth2FederationUpdateOAuthServiceOK struct {
	Payload *models.APIOAuthService
}

// IsSuccess returns true when this o auth2 federation update o auth service Ok response has a 2xx status code
func (o *OAuth2FederationUpdateOAuthServiceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this o auth2 federation update o auth service Ok response has a 3xx status code
func (o *OAuth2FederationUpdateOAuthServiceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this o auth2 federation update o auth service Ok response has a 4xx status code
func (o *OAuth2FederationUpdateOAuthServiceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this o auth2 federation update o auth service Ok response has a 5xx status code
func (o *OAuth2FederationUpdateOAuthServiceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this o auth2 federation update o auth service Ok response a status code equal to that given
func (o *OAuth2FederationUpdateOAuthServiceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the o auth2 federation update o auth service Ok response
func (o *OAuth2FederationUpdateOAuthServiceOK) Code() int {
	return 200
}

func (o *OAuth2FederationUpdateOAuthServiceOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /oauth/apps/{changes.id}][%d] oAuth2FederationUpdateOAuthServiceOk %s", 200, payload)
}

func (o *OAuth2FederationUpdateOAuthServiceOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /oauth/apps/{changes.id}][%d] oAuth2FederationUpdateOAuthServiceOk %s", 200, payload)
}

func (o *OAuth2FederationUpdateOAuthServiceOK) GetPayload() *models.APIOAuthService {
	return o.Payload
}

func (o *OAuth2FederationUpdateOAuthServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIOAuthService)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
