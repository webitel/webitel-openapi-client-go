// Code generated by go-swagger; DO NOT EDIT.

package roles

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

// RolesUpdateRoleReader is a Reader for the RolesUpdateRole structure.
type RolesUpdateRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RolesUpdateRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRolesUpdateRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /roles/{role.id}] Roles_UpdateRole", response, response.Code())
	}
}

// NewRolesUpdateRoleOK creates a RolesUpdateRoleOK with default headers values
func NewRolesUpdateRoleOK() *RolesUpdateRoleOK {
	return &RolesUpdateRoleOK{}
}

/*
RolesUpdateRoleOK describes a response with status code 200, with default header values.

A successful response.
*/
type RolesUpdateRoleOK struct {
	Payload *models.APIUpdateRoleResponse
}

// IsSuccess returns true when this roles update role Ok response has a 2xx status code
func (o *RolesUpdateRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this roles update role Ok response has a 3xx status code
func (o *RolesUpdateRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this roles update role Ok response has a 4xx status code
func (o *RolesUpdateRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this roles update role Ok response has a 5xx status code
func (o *RolesUpdateRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this roles update role Ok response a status code equal to that given
func (o *RolesUpdateRoleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the roles update role Ok response
func (o *RolesUpdateRoleOK) Code() int {
	return 200
}

func (o *RolesUpdateRoleOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /roles/{role.id}][%d] rolesUpdateRoleOk %s", 200, payload)
}

func (o *RolesUpdateRoleOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /roles/{role.id}][%d] rolesUpdateRoleOk %s", 200, payload)
}

func (o *RolesUpdateRoleOK) GetPayload() *models.APIUpdateRoleResponse {
	return o.Payload
}

func (o *RolesUpdateRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIUpdateRoleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
