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

// SearchTeamHookReader is a Reader for the SearchTeamHook structure.
type SearchTeamHookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchTeamHookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchTeamHookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSearchTeamHookDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchTeamHookOK creates a SearchTeamHookOK with default headers values
func NewSearchTeamHookOK() *SearchTeamHookOK {
	return &SearchTeamHookOK{}
}

/*
SearchTeamHookOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchTeamHookOK struct {
	Payload *models.EngineListTeamHook
}

// IsSuccess returns true when this search team hook Ok response has a 2xx status code
func (o *SearchTeamHookOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search team hook Ok response has a 3xx status code
func (o *SearchTeamHookOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search team hook Ok response has a 4xx status code
func (o *SearchTeamHookOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search team hook Ok response has a 5xx status code
func (o *SearchTeamHookOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search team hook Ok response a status code equal to that given
func (o *SearchTeamHookOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search team hook Ok response
func (o *SearchTeamHookOK) Code() int {
	return 200
}

func (o *SearchTeamHookOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/teams/{team_id}/hooks][%d] searchTeamHookOk %s", 200, payload)
}

func (o *SearchTeamHookOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/teams/{team_id}/hooks][%d] searchTeamHookOk %s", 200, payload)
}

func (o *SearchTeamHookOK) GetPayload() *models.EngineListTeamHook {
	return o.Payload
}

func (o *SearchTeamHookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineListTeamHook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchTeamHookDefault creates a SearchTeamHookDefault with default headers values
func NewSearchTeamHookDefault(code int) *SearchTeamHookDefault {
	return &SearchTeamHookDefault{
		_statusCode: code,
	}
}

/*
SearchTeamHookDefault describes a response with status code -1, with default header values.

Server error
*/
type SearchTeamHookDefault struct {
	_statusCode int

	Payload *models.EngineAPIError
}

// IsSuccess returns true when this search team hook default response has a 2xx status code
func (o *SearchTeamHookDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this search team hook default response has a 3xx status code
func (o *SearchTeamHookDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this search team hook default response has a 4xx status code
func (o *SearchTeamHookDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this search team hook default response has a 5xx status code
func (o *SearchTeamHookDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this search team hook default response a status code equal to that given
func (o *SearchTeamHookDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the search team hook default response
func (o *SearchTeamHookDefault) Code() int {
	return o._statusCode
}

func (o *SearchTeamHookDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/teams/{team_id}/hooks][%d] SearchTeamHook default %s", o._statusCode, payload)
}

func (o *SearchTeamHookDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /call_center/teams/{team_id}/hooks][%d] SearchTeamHook default %s", o._statusCode, payload)
}

func (o *SearchTeamHookDefault) GetPayload() *models.EngineAPIError {
	return o.Payload
}

func (o *SearchTeamHookDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EngineAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
