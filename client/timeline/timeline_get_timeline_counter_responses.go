// Code generated by go-swagger; DO NOT EDIT.

package timeline

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

// TimelineGetTimelineCounterReader is a Reader for the TimelineGetTimelineCounter structure.
type TimelineGetTimelineCounterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TimelineGetTimelineCounterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTimelineGetTimelineCounterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /contacts/{contact_id}/timeline/counter] Timeline_GetTimelineCounter", response, response.Code())
	}
}

// NewTimelineGetTimelineCounterOK creates a TimelineGetTimelineCounterOK with default headers values
func NewTimelineGetTimelineCounterOK() *TimelineGetTimelineCounterOK {
	return &TimelineGetTimelineCounterOK{}
}

/*
TimelineGetTimelineCounterOK describes a response with status code 200, with default header values.

A successful response.
*/
type TimelineGetTimelineCounterOK struct {
	Payload *models.WebitelContactsGetTimelineCounterResponse
}

// IsSuccess returns true when this timeline get timeline counter Ok response has a 2xx status code
func (o *TimelineGetTimelineCounterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this timeline get timeline counter Ok response has a 3xx status code
func (o *TimelineGetTimelineCounterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this timeline get timeline counter Ok response has a 4xx status code
func (o *TimelineGetTimelineCounterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this timeline get timeline counter Ok response has a 5xx status code
func (o *TimelineGetTimelineCounterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this timeline get timeline counter Ok response a status code equal to that given
func (o *TimelineGetTimelineCounterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the timeline get timeline counter Ok response
func (o *TimelineGetTimelineCounterOK) Code() int {
	return 200
}

func (o *TimelineGetTimelineCounterOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /contacts/{contact_id}/timeline/counter][%d] timelineGetTimelineCounterOk %s", 200, payload)
}

func (o *TimelineGetTimelineCounterOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /contacts/{contact_id}/timeline/counter][%d] timelineGetTimelineCounterOk %s", 200, payload)
}

func (o *TimelineGetTimelineCounterOK) GetPayload() *models.WebitelContactsGetTimelineCounterResponse {
	return o.Payload
}

func (o *TimelineGetTimelineCounterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebitelContactsGetTimelineCounterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
