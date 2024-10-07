// Code generated by go-swagger; DO NOT EDIT.

package team_hook_service

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

// UpdateTeamHookReader is a Reader for the UpdateTeamHook structure.
type UpdateTeamHookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTeamHookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTeamHookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateTeamHookDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateTeamHookOK creates a UpdateTeamHookOK with default headers values
func NewUpdateTeamHookOK() *UpdateTeamHookOK {
	return &UpdateTeamHookOK{}
}

/*
UpdateTeamHookOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateTeamHookOK struct {
	Payload *models.EngineTeamHook
}

// IsSuccess returns true when this update team hook Ok response has a 2xx status code
func (o *UpdateTeamHookOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update team hook Ok response has a 3xx status code
func (o *UpdateTeamHookOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update team hook Ok response has a 4xx status code
func (o *UpdateTeamHookOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update team hook Ok response has a 5xx status code
func (o *UpdateTeamHookOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update team hook Ok response a status code equal to that given
func (o *UpdateTeamHookOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update team hook Ok response
func (o *UpdateTeamHookOK) Code() int {
	return 200
}

func (o *UpdateTeamHookOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/teams/{team_id}/hooks/{id}][%d] updateTeamHookOk %s", 200, payload)
}

func (o *UpdateTeamHookOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/teams/{team_id}/hooks/{id}][%d] updateTeamHookOk %s", 200, payload)
}

func (o *UpdateTeamHookOK) GetPayload() *models.EngineTeamHook {
	return o.Payload
}

func (o *UpdateTeamHookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineTeamHook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamHookDefault creates a UpdateTeamHookDefault with default headers values
func NewUpdateTeamHookDefault(code int) *UpdateTeamHookDefault {
	return &UpdateTeamHookDefault{
		_statusCode: code,
	}
}

/*
UpdateTeamHookDefault describes a response with status code -1, with default header values.

Server error
*/
type UpdateTeamHookDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this update team hook default response has a 2xx status code
func (o *UpdateTeamHookDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update team hook default response has a 3xx status code
func (o *UpdateTeamHookDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update team hook default response has a 4xx status code
func (o *UpdateTeamHookDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update team hook default response has a 5xx status code
func (o *UpdateTeamHookDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update team hook default response a status code equal to that given
func (o *UpdateTeamHookDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update team hook default response
func (o *UpdateTeamHookDefault) Code() int {
	return o._statusCode
}

func (o *UpdateTeamHookDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/teams/{team_id}/hooks/{id}][%d] UpdateTeamHook default %s", o._statusCode, payload)
}

func (o *UpdateTeamHookDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /call_center/teams/{team_id}/hooks/{id}][%d] UpdateTeamHook default %s", o._statusCode, payload)
}

func (o *UpdateTeamHookDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *UpdateTeamHookDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
