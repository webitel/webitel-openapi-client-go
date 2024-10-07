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

// WebitelContactsPhoto A contact's photo.
// A picture shown next to the contact's name
// to help others recognize the contact.
//
// swagger:model webitel.contacts.Photo
type WebitelContactsPhoto struct {

	// The user who created this Field.
	CreatedAt string `json:"created_at,omitempty"`

	// Timestamp(milli) of the Field creation.
	CreatedBy *WebitelContactsLookup `json:"created_by,omitempty"`

	// Unique ID of the latest version of the update.
	// This ID changes after any update to the underlying value(s).
	Etag string `json:"etag,omitempty"`

	// The unique ID of the association. Never changes.
	ID string `json:"id,omitempty"`

	// File unique ID within 'storage' service
	PhotoID string `json:"photo_id,omitempty"`

	// The URL of the photo.
	// You can change the desired size by appending
	// a query parameter sz={size} at the end of the url,
	// where {size} is the size in pixels.
	// Example: https://lh3.googleusercontent.com/-T_wVWLlmg7w/AAAAAAAAAAI/AAAAAAAABa8/00gzXvDBYqw/s100/photo.jpg?sz=50
	PhotoURL string `json:"photo_url,omitempty"`

	// True if the photo is a default photo; false if the photo is a user-provided photo.
	Primary bool `json:"primary,omitempty"`

	// Timestamp(milli) of the last Field update.
	// Take part in Etag generation.
	UpdatedAt string `json:"updated_at,omitempty"`

	// The user who performed last Update.
	UpdatedBy *WebitelContactsLookup `json:"updated_by,omitempty"`

	// Version of the latest update. Numeric sequence.
	Ver int32 `json:"ver,omitempty"`
}

// Validate validates this webitel contacts photo
func (m *WebitelContactsPhoto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedBy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebitelContactsPhoto) validateCreatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created_by")
			}
			return err
		}
	}

	return nil
}

func (m *WebitelContactsPhoto) validateUpdatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedBy) { // not required
		return nil
	}

	if m.UpdatedBy != nil {
		if err := m.UpdatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updated_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updated_by")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this webitel contacts photo based on the context it is used
func (m *WebitelContactsPhoto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebitelContactsPhoto) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.CreatedBy != nil {

		if swag.IsZero(m.CreatedBy) { // not required
			return nil
		}

		if err := m.CreatedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created_by")
			}
			return err
		}
	}

	return nil
}

func (m *WebitelContactsPhoto) contextValidateUpdatedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.UpdatedBy != nil {

		if swag.IsZero(m.UpdatedBy) { // not required
			return nil
		}

		if err := m.UpdatedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updated_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updated_by")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebitelContactsPhoto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebitelContactsPhoto) UnmarshalBinary(b []byte) error {
	var res WebitelContactsPhoto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
