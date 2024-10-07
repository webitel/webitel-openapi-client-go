// Code generated by go-swagger; DO NOT EDIT.

package catalog

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

// CatalogGetCustomersReader is a Reader for the CatalogGetCustomers structure.
type CatalogGetCustomersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogGetCustomersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCatalogGetCustomersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /chat/customers] Catalog_GetCustomers", response, response.Code())
	}
}

// NewCatalogGetCustomersOK creates a CatalogGetCustomersOK with default headers values
func NewCatalogGetCustomersOK() *CatalogGetCustomersOK {
	return &CatalogGetCustomersOK{}
}

/*
CatalogGetCustomersOK describes a response with status code 200, with default header values.

A successful response.
*/
type CatalogGetCustomersOK struct {
	Payload *models.WebitelChatChatCustomers
}

// IsSuccess returns true when this catalog get customers Ok response has a 2xx status code
func (o *CatalogGetCustomersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this catalog get customers Ok response has a 3xx status code
func (o *CatalogGetCustomersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog get customers Ok response has a 4xx status code
func (o *CatalogGetCustomersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this catalog get customers Ok response has a 5xx status code
func (o *CatalogGetCustomersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog get customers Ok response a status code equal to that given
func (o *CatalogGetCustomersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the catalog get customers Ok response
func (o *CatalogGetCustomersOK) Code() int {
	return 200
}

func (o *CatalogGetCustomersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /chat/customers][%d] catalogGetCustomersOk %s", 200, payload)
}

func (o *CatalogGetCustomersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /chat/customers][%d] catalogGetCustomersOk %s", 200, payload)
}

func (o *CatalogGetCustomersOK) GetPayload() *models.WebitelChatChatCustomers {
	return o.Payload
}

func (o *CatalogGetCustomersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebitelChatChatCustomers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
