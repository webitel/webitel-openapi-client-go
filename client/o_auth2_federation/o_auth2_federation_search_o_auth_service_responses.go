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

// OAuth2FederationSearchOAuthServiceReader is a Reader for the OAuth2FederationSearchOAuthService structure.
type OAuth2FederationSearchOAuthServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OAuth2FederationSearchOAuthServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOAuth2FederationSearchOAuthServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /oauth/apps] OAuth2Federation_SearchOAuthService", response, response.Code())
	}
}

// NewOAuth2FederationSearchOAuthServiceOK creates a OAuth2FederationSearchOAuthServiceOK with default headers values
func NewOAuth2FederationSearchOAuthServiceOK() *OAuth2FederationSearchOAuthServiceOK {
	return &OAuth2FederationSearchOAuthServiceOK{}
}

/*
OAuth2FederationSearchOAuthServiceOK describes a response with status code 200, with default header values.

A successful response.
*/
type OAuth2FederationSearchOAuthServiceOK struct {
	Payload *models.APISearchOAuthServiceResponse
}

// IsSuccess returns true when this o auth2 federation search o auth service Ok response has a 2xx status code
func (o *OAuth2FederationSearchOAuthServiceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this o auth2 federation search o auth service Ok response has a 3xx status code
func (o *OAuth2FederationSearchOAuthServiceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this o auth2 federation search o auth service Ok response has a 4xx status code
func (o *OAuth2FederationSearchOAuthServiceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this o auth2 federation search o auth service Ok response has a 5xx status code
func (o *OAuth2FederationSearchOAuthServiceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this o auth2 federation search o auth service Ok response a status code equal to that given
func (o *OAuth2FederationSearchOAuthServiceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the o auth2 federation search o auth service Ok response
func (o *OAuth2FederationSearchOAuthServiceOK) Code() int {
	return 200
}

func (o *OAuth2FederationSearchOAuthServiceOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /oauth/apps][%d] oAuth2FederationSearchOAuthServiceOk %s", 200, payload)
}

func (o *OAuth2FederationSearchOAuthServiceOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /oauth/apps][%d] oAuth2FederationSearchOAuthServiceOk %s", 200, payload)
}

func (o *OAuth2FederationSearchOAuthServiceOK) GetPayload() *models.APISearchOAuthServiceResponse {
	return o.Payload
}

func (o *OAuth2FederationSearchOAuthServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APISearchOAuthServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
