// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WfmUpdatePauseTemplateResponse wfm update pause template response
//
// swagger:model wfmUpdatePauseTemplateResponse
type WfmUpdatePauseTemplateResponse struct {

	// item
	Item *WfmPauseTemplate `json:"item,omitempty"`
}

// Validate validates this wfm update pause template response
func (m *WfmUpdatePauseTemplateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WfmUpdatePauseTemplateResponse) validateItem(formats strfmt.Registry) error {
	if swag.IsZero(m.Item) { // not required
		return nil
	}

	if m.Item != nil {
		if err := m.Item.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("item")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("item")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this wfm update pause template response based on the context it is used
func (m *WfmUpdatePauseTemplateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WfmUpdatePauseTemplateResponse) contextValidateItem(ctx context.Context, formats strfmt.Registry) error {

	if m.Item != nil {

		if swag.IsZero(m.Item) { // not required
			return nil
		}

		if err := m.Item.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("item")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("item")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WfmUpdatePauseTemplateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WfmUpdatePauseTemplateResponse) UnmarshalBinary(b []byte) error {
	var res WfmUpdatePauseTemplateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
