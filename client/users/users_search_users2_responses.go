// Code generated by go-swagger; DO NOT EDIT.

package users

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

// UsersSearchUsers2Reader is a Reader for the UsersSearchUsers2 structure.
type UsersSearchUsers2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersSearchUsers2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersSearchUsers2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /users/search] Users_SearchUsers2", response, response.Code())
	}
}

// NewUsersSearchUsers2OK creates a UsersSearchUsers2OK with default headers values
func NewUsersSearchUsers2OK() *UsersSearchUsers2OK {
	return &UsersSearchUsers2OK{}
}

/*
UsersSearchUsers2OK describes a response with status code 200, with default header values.

A successful response.
*/
type UsersSearchUsers2OK struct {
	Payload *models.APISearchUsersResponse
}

// IsSuccess returns true when this users search users2 Ok response has a 2xx status code
func (o *UsersSearchUsers2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users search users2 Ok response has a 3xx status code
func (o *UsersSearchUsers2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users search users2 Ok response has a 4xx status code
func (o *UsersSearchUsers2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users search users2 Ok response has a 5xx status code
func (o *UsersSearchUsers2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this users search users2 Ok response a status code equal to that given
func (o *UsersSearchUsers2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the users search users2 Ok response
func (o *UsersSearchUsers2OK) Code() int {
	return 200
}

func (o *UsersSearchUsers2OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users/search][%d] usersSearchUsers2Ok %s", 200, payload)
}

func (o *UsersSearchUsers2OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users/search][%d] usersSearchUsers2Ok %s", 200, payload)
}

func (o *UsersSearchUsers2OK) GetPayload() *models.APISearchUsersResponse {
	return o.Payload
}

func (o *UsersSearchUsers2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APISearchUsersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
