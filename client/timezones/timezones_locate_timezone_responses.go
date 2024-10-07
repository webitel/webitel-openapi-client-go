// Code generated by go-swagger; DO NOT EDIT.

package timezones

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

// TimezonesLocateTimezoneReader is a Reader for the TimezonesLocateTimezone structure.
type TimezonesLocateTimezoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TimezonesLocateTimezoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTimezonesLocateTimezoneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /contacts/{contact_id}/timezones/{etag}] Timezones_LocateTimezone", response, response.Code())
	}
}

// NewTimezonesLocateTimezoneOK creates a TimezonesLocateTimezoneOK with default headers values
func NewTimezonesLocateTimezoneOK() *TimezonesLocateTimezoneOK {
	return &TimezonesLocateTimezoneOK{}
}

/*
TimezonesLocateTimezoneOK describes a response with status code 200, with default header values.

A successful response.
*/
type TimezonesLocateTimezoneOK struct {
	Payload *models.WebitelContactsTimezone
}

// IsSuccess returns true when this timezones locate timezone Ok response has a 2xx status code
func (o *TimezonesLocateTimezoneOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this timezones locate timezone Ok response has a 3xx status code
func (o *TimezonesLocateTimezoneOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this timezones locate timezone Ok response has a 4xx status code
func (o *TimezonesLocateTimezoneOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this timezones locate timezone Ok response has a 5xx status code
func (o *TimezonesLocateTimezoneOK) IsServerError() bool {
	return false
}

// IsCode returns true when this timezones locate timezone Ok response a status code equal to that given
func (o *TimezonesLocateTimezoneOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the timezones locate timezone Ok response
func (o *TimezonesLocateTimezoneOK) Code() int {
	return 200
}

func (o *TimezonesLocateTimezoneOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /contacts/{contact_id}/timezones/{etag}][%d] timezonesLocateTimezoneOk %s", 200, payload)
}

func (o *TimezonesLocateTimezoneOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /contacts/{contact_id}/timezones/{etag}][%d] timezonesLocateTimezoneOk %s", 200, payload)
}

func (o *TimezonesLocateTimezoneOK) GetPayload() *models.WebitelContactsTimezone {
	return o.Payload
}

func (o *TimezonesLocateTimezoneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebitelContactsTimezone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
