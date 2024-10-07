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

// WebitelContactsCommentFormat webitel contacts comment format
//
// swagger:model webitel.contacts.Comment.Format
type WebitelContactsCommentFormat struct {

	// bold
	Bold WebitelContactsCommentFormatBold `json:"bold,omitempty"`

	// codeblock
	Codeblock *WebitelContactsCommentFormatCodeblock `json:"codeblock,omitempty"`

	// italic
	Italic WebitelContactsCommentFormatItalic `json:"italic,omitempty"`

	// Length text runes count.
	Length int32 `json:"length,omitempty"`

	// link
	Link *WebitelContactsCommentFormatLink `json:"link,omitempty"`

	// monospace
	Monospace WebitelContactsCommentFormatMonospace `json:"monospace,omitempty"`

	// Offset text runes count.
	Offset int32 `json:"offset,omitempty"`

	// strikethrough
	Strikethrough WebitelContactsCommentFormatStrikethrough `json:"strikethrough,omitempty"`

	// underline
	Underline WebitelContactsCommentFormatUnderline `json:"underline,omitempty"`
}

// Validate validates this webitel contacts comment format
func (m *WebitelContactsCommentFormat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCodeblock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLink(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebitelContactsCommentFormat) validateCodeblock(formats strfmt.Registry) error {
	if swag.IsZero(m.Codeblock) { // not required
		return nil
	}

	if m.Codeblock != nil {
		if err := m.Codeblock.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("codeblock")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("codeblock")
			}
			return err
		}
	}

	return nil
}

func (m *WebitelContactsCommentFormat) validateLink(formats strfmt.Registry) error {
	if swag.IsZero(m.Link) { // not required
		return nil
	}

	if m.Link != nil {
		if err := m.Link.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("link")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("link")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this webitel contacts comment format based on the context it is used
func (m *WebitelContactsCommentFormat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCodeblock(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLink(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebitelContactsCommentFormat) contextValidateCodeblock(ctx context.Context, formats strfmt.Registry) error {

	if m.Codeblock != nil {

		if swag.IsZero(m.Codeblock) { // not required
			return nil
		}

		if err := m.Codeblock.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("codeblock")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("codeblock")
			}
			return err
		}
	}

	return nil
}

func (m *WebitelContactsCommentFormat) contextValidateLink(ctx context.Context, formats strfmt.Registry) error {

	if m.Link != nil {

		if swag.IsZero(m.Link) { // not required
			return nil
		}

		if err := m.Link.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("link")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("link")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebitelContactsCommentFormat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebitelContactsCommentFormat) UnmarshalBinary(b []byte) error {
	var res WebitelContactsCommentFormat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
