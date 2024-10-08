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

// UsersSearchUsersReader is a Reader for the UsersSearchUsers structure.
type UsersSearchUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersSearchUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersSearchUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /users] Users_SearchUsers", response, response.Code())
	}
}

// NewUsersSearchUsersOK creates a UsersSearchUsersOK with default headers values
func NewUsersSearchUsersOK() *UsersSearchUsersOK {
	return &UsersSearchUsersOK{}
}

/*
UsersSearchUsersOK describes a response with status code 200, with default header values.

A successful response.
*/
type UsersSearchUsersOK struct {
	Payload *models.APISearchUsersResponse
}

// IsSuccess returns true when this users search users Ok response has a 2xx status code
func (o *UsersSearchUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users search users Ok response has a 3xx status code
func (o *UsersSearchUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users search users Ok response has a 4xx status code
func (o *UsersSearchUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users search users Ok response has a 5xx status code
func (o *UsersSearchUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users search users Ok response a status code equal to that given
func (o *UsersSearchUsersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the users search users Ok response
func (o *UsersSearchUsersOK) Code() int {
	return 200
}

func (o *UsersSearchUsersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users][%d] usersSearchUsersOk %s", 200, payload)
}

func (o *UsersSearchUsersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users][%d] usersSearchUsersOk %s", 200, payload)
}

func (o *UsersSearchUsersOK) GetPayload() *models.APISearchUsersResponse {
	return o.Payload
}

func (o *UsersSearchUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APISearchUsersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
