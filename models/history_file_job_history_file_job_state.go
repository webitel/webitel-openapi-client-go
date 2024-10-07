// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// HistoryFileJobHistoryFileJobState history file job history file job state
//
// swagger:model HistoryFileJobHistoryFileJobState
type HistoryFileJobHistoryFileJobState string

func NewHistoryFileJobHistoryFileJobState(value HistoryFileJobHistoryFileJobState) *HistoryFileJobHistoryFileJobState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HistoryFileJobHistoryFileJobState.
func (m HistoryFileJobHistoryFileJobState) Pointer() *HistoryFileJobHistoryFileJobState {
	return &m
}

const (

	// HistoryFileJobHistoryFileJobStateIdle captures enum value "idle"
	HistoryFileJobHistoryFileJobStateIdle HistoryFileJobHistoryFileJobState = "idle"

	// HistoryFileJobHistoryFileJobStateActive captures enum value "active"
	HistoryFileJobHistoryFileJobStateActive HistoryFileJobHistoryFileJobState = "active"

	// HistoryFileJobHistoryFileJobStateFinished captures enum value "finished"
	HistoryFileJobHistoryFileJobStateFinished HistoryFileJobHistoryFileJobState = "finished"

	// HistoryFileJobHistoryFileJobStateError captures enum value "error"
	HistoryFileJobHistoryFileJobStateError HistoryFileJobHistoryFileJobState = "error"
)

// for schema
var historyFileJobHistoryFileJobStateEnum []interface{}

func init() {
	var res []HistoryFileJobHistoryFileJobState
	if err := json.Unmarshal([]byte(`["idle","active","finished","error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		historyFileJobHistoryFileJobStateEnum = append(historyFileJobHistoryFileJobStateEnum, v)
	}
}

func (m HistoryFileJobHistoryFileJobState) validateHistoryFileJobHistoryFileJobStateEnum(path, location string, value HistoryFileJobHistoryFileJobState) error {
	if err := validate.EnumCase(path, location, value, historyFileJobHistoryFileJobStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this history file job history file job state
func (m HistoryFileJobHistoryFileJobState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHistoryFileJobHistoryFileJobStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this history file job history file job state based on context it is used
func (m HistoryFileJobHistoryFileJobState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
