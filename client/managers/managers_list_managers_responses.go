// Code generated by go-swagger; DO NOT EDIT.

package managers

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

// ManagersListManagersReader is a Reader for the ManagersListManagers structure.
type ManagersListManagersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ManagersListManagersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewManagersListManagersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /contacts/{contact_id}/managers] Managers_ListManagers", response, response.Code())
	}
}

// NewManagersListManagersOK creates a ManagersListManagersOK with default headers values
func NewManagersListManagersOK() *ManagersListManagersOK {
	return &ManagersListManagersOK{}
}

/*
ManagersListManagersOK describes a response with status code 200, with default header values.

A successful response.
*/
type ManagersListManagersOK struct {
	Payload *models.WebitelContactsManagerList
}

// IsSuccess returns true when this managers list managers Ok response has a 2xx status code
func (o *ManagersListManagersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this managers list managers Ok response has a 3xx status code
func (o *ManagersListManagersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this managers list managers Ok response has a 4xx status code
func (o *ManagersListManagersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this managers list managers Ok response has a 5xx status code
func (o *ManagersListManagersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this managers list managers Ok response a status code equal to that given
func (o *ManagersListManagersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the managers list managers Ok response
func (o *ManagersListManagersOK) Code() int {
	return 200
}

func (o *ManagersListManagersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /contacts/{contact_id}/managers][%d] managersListManagersOk %s", 200, payload)
}

func (o *ManagersListManagersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /contacts/{contact_id}/managers][%d] managersListManagersOk %s", 200, payload)
}

func (o *ManagersListManagersOK) GetPayload() *models.WebitelContactsManagerList {
	return o.Payload
}

func (o *ManagersListManagersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebitelContactsManagerList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
