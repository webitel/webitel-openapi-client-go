// Code generated by go-swagger; DO NOT EDIT.

package variables

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

// VariablesDeleteVariableReader is a Reader for the VariablesDeleteVariable structure.
type VariablesDeleteVariableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VariablesDeleteVariableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVariablesDeleteVariableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /contacts/{contact_id}/variables/{etag}] Variables_DeleteVariable", response, response.Code())
	}
}

// NewVariablesDeleteVariableOK creates a VariablesDeleteVariableOK with default headers values
func NewVariablesDeleteVariableOK() *VariablesDeleteVariableOK {
	return &VariablesDeleteVariableOK{}
}

/*
VariablesDeleteVariableOK describes a response with status code 200, with default header values.

A successful response.
*/
type VariablesDeleteVariableOK struct {
	Payload *models.WebitelContactsVariable
}

// IsSuccess returns true when this variables delete variable Ok response has a 2xx status code
func (o *VariablesDeleteVariableOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this variables delete variable Ok response has a 3xx status code
func (o *VariablesDeleteVariableOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this variables delete variable Ok response has a 4xx status code
func (o *VariablesDeleteVariableOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this variables delete variable Ok response has a 5xx status code
func (o *VariablesDeleteVariableOK) IsServerError() bool {
	return false
}

// IsCode returns true when this variables delete variable Ok response a status code equal to that given
func (o *VariablesDeleteVariableOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the variables delete variable Ok response
func (o *VariablesDeleteVariableOK) Code() int {
	return 200
}

func (o *VariablesDeleteVariableOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /contacts/{contact_id}/variables/{etag}][%d] variablesDeleteVariableOk %s", 200, payload)
}

func (o *VariablesDeleteVariableOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /contacts/{contact_id}/variables/{etag}][%d] variablesDeleteVariableOk %s", 200, payload)
}

func (o *VariablesDeleteVariableOK) GetPayload() *models.WebitelContactsVariable {
	return o.Payload
}

func (o *VariablesDeleteVariableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebitelContactsVariable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
