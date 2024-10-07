// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WebitelChatDialog Chat Dialog. Conversation info.
//
// swagger:model webitel.chat.Dialog
type WebitelChatDialog struct {

	// Timestamp when dialog was closed.
	// Zero value means - connected (online)
	// Otherwise - disconnected (offline)
	Closed string `json:"closed,omitempty"`

	// Close reason if closed from the request sender perspective
	ClosedCause string `json:"closed_cause,omitempty"`

	// Context. Variables. Environment.
	Context map[string]string `json:"context,omitempty"`

	// Timestamp of the latest activity.
	Date string `json:"date,omitempty"`

	// [D]omain[C]omponent primary ID.
	Dc string `json:"dc,omitempty"`

	// [FROM]: Originator.
	// Leg[A]. Contact / User.
	From *WebitelChatPeer `json:"from,omitempty"`

	// The Conversation thread unique ID.
	ID string `json:"id,omitempty"`

	// [TO]: Participants.
	// Leg[A+]. Schema / Agent.
	Members []*WebitelChatChat `json:"members"`

	// The latest (top) message.
	Message *WebitelChatMessage `json:"message,omitempty"`

	// OPTIONAL. Last dialog queue
	Queue *WebitelChatPeer `json:"queue,omitempty"`

	// Timestamp when dialog started.
	Started string `json:"started,omitempty"`

	// Title of the dialog.
	Title string `json:"title,omitempty"`

	// [VIA] Text gateway [FROM] originated thru ...
	Via *WebitelChatPeer `json:"via,omitempty"`
}

// Validate validates this webitel chat dialog
func (m *WebitelChatDialog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMembers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVia(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebitelChatDialog) validateFrom(formats strfmt.Registry) error {
	if swag.IsZero(m.From) { // not required
		return nil
	}

	if m.From != nil {
		if err := m.From.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("from")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("from")
			}
			return err
		}
	}

	return nil
}

func (m *WebitelChatDialog) validateMembers(formats strfmt.Registry) error {
	if swag.IsZero(m.Members) { // not required
		return nil
	}

	for i := 0; i < len(m.Members); i++ {
		if swag.IsZero(m.Members[i]) { // not required
			continue
		}

		if m.Members[i] != nil {
			if err := m.Members[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("members" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WebitelChatDialog) validateMessage(formats strfmt.Registry) error {
	if swag.IsZero(m.Message) { // not required
		return nil
	}

	if m.Message != nil {
		if err := m.Message.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("message")
			}
			return err
		}
	}

	return nil
}

func (m *WebitelChatDialog) validateQueue(formats strfmt.Registry) error {
	if swag.IsZero(m.Queue) { // not required
		return nil
	}

	if m.Queue != nil {
		if err := m.Queue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queue")
			}
			return err
		}
	}

	return nil
}

func (m *WebitelChatDialog) validateVia(formats strfmt.Registry) error {
	if swag.IsZero(m.Via) { // not required
		return nil
	}

	if m.Via != nil {
		if err := m.Via.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("via")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("via")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this webitel chat dialog based on the context it is used
func (m *WebitelChatDialog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFrom(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMembers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQueue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVia(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebitelChatDialog) contextValidateFrom(ctx context.Context, formats strfmt.Registry) error {

	if m.From != nil {

		if swag.IsZero(m.From) { // not required
			return nil
		}

		if err := m.From.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("from")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("from")
			}
			return err
		}
	}

	return nil
}

func (m *WebitelChatDialog) contextValidateMembers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Members); i++ {

		if m.Members[i] != nil {

			if swag.IsZero(m.Members[i]) { // not required
				return nil
			}

			if err := m.Members[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("members" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WebitelChatDialog) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if m.Message != nil {

		if swag.IsZero(m.Message) { // not required
			return nil
		}

		if err := m.Message.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("message")
			}
			return err
		}
	}

	return nil
}

func (m *WebitelChatDialog) contextValidateQueue(ctx context.Context, formats strfmt.Registry) error {

	if m.Queue != nil {

		if swag.IsZero(m.Queue) { // not required
			return nil
		}

		if err := m.Queue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queue")
			}
			return err
		}
	}

	return nil
}

func (m *WebitelChatDialog) contextValidateVia(ctx context.Context, formats strfmt.Registry) error {

	if m.Via != nil {

		if swag.IsZero(m.Via) { // not required
			return nil
		}

		if err := m.Via.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("via")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("via")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebitelChatDialog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebitelChatDialog) UnmarshalBinary(b []byte) error {
	var res WebitelChatDialog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
