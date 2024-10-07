// Code generated by go-swagger; DO NOT EDIT.

package customers

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

// CustomersLicenseUsageReader is a Reader for the CustomersLicenseUsage structure.
type CustomersLicenseUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomersLicenseUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomersLicenseUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /customer/{customer_id}/license] Customers_LicenseUsage", response, response.Code())
	}
}

// NewCustomersLicenseUsageOK creates a CustomersLicenseUsageOK with default headers values
func NewCustomersLicenseUsageOK() *CustomersLicenseUsageOK {
	return &CustomersLicenseUsageOK{}
}

/*
CustomersLicenseUsageOK describes a response with status code 200, with default header values.

A successful response.
*/
type CustomersLicenseUsageOK struct {
	Payload *models.APILicenseUsageResponse
}

// IsSuccess returns true when this customers license usage Ok response has a 2xx status code
func (o *CustomersLicenseUsageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customers license usage Ok response has a 3xx status code
func (o *CustomersLicenseUsageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customers license usage Ok response has a 4xx status code
func (o *CustomersLicenseUsageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customers license usage Ok response has a 5xx status code
func (o *CustomersLicenseUsageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customers license usage Ok response a status code equal to that given
func (o *CustomersLicenseUsageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customers license usage Ok response
func (o *CustomersLicenseUsageOK) Code() int {
	return 200
}

func (o *CustomersLicenseUsageOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /customer/{customer_id}/license][%d] customersLicenseUsageOk %s", 200, payload)
}

func (o *CustomersLicenseUsageOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /customer/{customer_id}/license][%d] customersLicenseUsageOk %s", 200, payload)
}

func (o *CustomersLicenseUsageOK) GetPayload() *models.APILicenseUsageResponse {
	return o.Payload
}

func (o *CustomersLicenseUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APILicenseUsageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
