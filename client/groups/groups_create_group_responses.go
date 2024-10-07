// Code generated by go-swagger; DO NOT EDIT.

package groups

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

// GroupsCreateGroupReader is a Reader for the GroupsCreateGroup structure.
type GroupsCreateGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupsCreateGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupsCreateGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /contacts/groups] Groups_CreateGroup", response, response.Code())
	}
}

// NewGroupsCreateGroupOK creates a GroupsCreateGroupOK with default headers values
func NewGroupsCreateGroupOK() *GroupsCreateGroupOK {
	return &GroupsCreateGroupOK{}
}

/*
GroupsCreateGroupOK describes a response with status code 200, with default header values.

A successful response.
*/
type GroupsCreateGroupOK struct {
	Payload *models.WebitelContactsGroup
}

// IsSuccess returns true when this groups create group Ok response has a 2xx status code
func (o *GroupsCreateGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this groups create group Ok response has a 3xx status code
func (o *GroupsCreateGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this groups create group Ok response has a 4xx status code
func (o *GroupsCreateGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this groups create group Ok response has a 5xx status code
func (o *GroupsCreateGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this groups create group Ok response a status code equal to that given
func (o *GroupsCreateGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the groups create group Ok response
func (o *GroupsCreateGroupOK) Code() int {
	return 200
}

func (o *GroupsCreateGroupOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /contacts/groups][%d] groupsCreateGroupOk %s", 200, payload)
}

func (o *GroupsCreateGroupOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /contacts/groups][%d] groupsCreateGroupOk %s", 200, payload)
}

func (o *GroupsCreateGroupOK) GetPayload() *models.WebitelContactsGroup {
	return o.Payload
}

func (o *GroupsCreateGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebitelContactsGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
