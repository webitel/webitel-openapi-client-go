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

// RolesDeleteRoleReader is a Reader for the RolesDeleteRole structure.
type RolesDeleteRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RolesDeleteRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRolesDeleteRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /roles/{id}] Roles_DeleteRole", response, response.Code())
	}
}

// NewRolesDeleteRoleOK creates a RolesDeleteRoleOK with default headers values
func NewRolesDeleteRoleOK() *RolesDeleteRoleOK {
	return &RolesDeleteRoleOK{}
}

/*
RolesDeleteRoleOK describes a response with status code 200, with default header values.

A successful response.
*/
type RolesDeleteRoleOK struct {
	Payload *models.APIDeleteRoleResponse
}

// IsSuccess returns true when this roles delete role Ok response has a 2xx status code
func (o *RolesDeleteRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this roles delete role Ok response has a 3xx status code
func (o *RolesDeleteRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this roles delete role Ok response has a 4xx status code
func (o *RolesDeleteRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this roles delete role Ok response has a 5xx status code
func (o *RolesDeleteRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this roles delete role Ok response a status code equal to that given
func (o *RolesDeleteRoleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the roles delete role Ok response
func (o *RolesDeleteRoleOK) Code() int {
	return 200
}

func (o *RolesDeleteRoleOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /roles/{id}][%d] rolesDeleteRoleOk %s", 200, payload)
}

func (o *RolesDeleteRoleOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /roles/{id}][%d] rolesDeleteRoleOk %s", 200, payload)
}

func (o *RolesDeleteRoleOK) GetPayload() *models.APIDeleteRoleResponse {
	return o.Payload
}

func (o *RolesDeleteRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIDeleteRoleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
