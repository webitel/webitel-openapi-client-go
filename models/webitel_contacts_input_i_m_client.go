// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WebitelContactsInputIMClient Input of the contact IM client.
//
// swagger:model webitel.contacts.InputIMClient
type WebitelContactsInputIMClient struct {

	// Id of Agent created this IM client.
	CreatedBy string `json:"created_by,omitempty"`

	// External user id
	ExternalUser string `json:"external_user,omitempty"`

	// App (Text-Gateway) used to connect the IM client.
	GatewayID string `json:"gateway_id,omitempty"`

	// protocol
	Protocol string `json:"protocol,omitempty"`

	// [Via] App(-specific) peer(-id) to connect[from] the IM client.
	Via string `json:"via,omitempty"`
}

// Validate validates this webitel contacts input i m client
func (m *WebitelContactsInputIMClient) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this webitel contacts input i m client based on context it is used
func (m *WebitelContactsInputIMClient) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebitelContactsInputIMClient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebitelContactsInputIMClient) UnmarshalBinary(b []byte) error {
	var res WebitelContactsInputIMClient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
