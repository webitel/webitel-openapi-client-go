// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EngineSchemaVariable engine schema variable
//
// swagger:model engineSchemaVariable
type EngineSchemaVariable struct {

	// encrypt
	Encrypt bool `json:"encrypt,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// value
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this engine schema variable
func (m *EngineSchemaVariable) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this engine schema variable based on context it is used
func (m *EngineSchemaVariable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EngineSchemaVariable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineSchemaVariable) UnmarshalBinary(b []byte) error {
	var res EngineSchemaVariable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
