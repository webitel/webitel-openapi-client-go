// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EngineQuestionAnswer engine question answer
//
// swagger:model engineQuestionAnswer
type EngineQuestionAnswer struct {

	// score
	Score float32 `json:"score,omitempty"`
}

// Validate validates this engine question answer
func (m *EngineQuestionAnswer) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this engine question answer based on context it is used
func (m *EngineQuestionAnswer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EngineQuestionAnswer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineQuestionAnswer) UnmarshalBinary(b []byte) error {
	var res EngineQuestionAnswer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
