// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WfmLookupEntity wfm lookup entity
//
// swagger:model wfmLookupEntity
type WfmLookupEntity struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this wfm lookup entity
func (m *WfmLookupEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this wfm lookup entity based on context it is used
func (m *WfmLookupEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WfmLookupEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WfmLookupEntity) UnmarshalBinary(b []byte) error {
	var res WfmLookupEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
