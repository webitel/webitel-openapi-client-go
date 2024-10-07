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

// DevicesUpdateDevice2ParamsBody Device profile
//
// swagger:model devicesUpdateDevice2ParamsBody
type DevicesUpdateDevice2ParamsBody struct {

	// credentials
	//
	// username
	Account string `json:"account,omitempty"`

	// provision
	//
	// vendor brand name
	Brand string `json:"brand,omitempty"`

	// unix
	CreatedAt string `json:"created_at,omitempty"`

	// user
	CreatedBy *APIUserID `json:"created_by,omitempty"`

	// unix
	DeletedAt string `json:"deleted_at,omitempty"`

	// user
	DeletedBy *APIUserID `json:"deleted_by,omitempty"`

	// Hotdesk: this workstation aliases
	//
	// act as a hotdesk
	Hotdesk bool `json:"hotdesk,omitempty"`

	// associated aliases
	Hotdesks []string `json:"hotdesks"`

	// ip
	IP string `json:"ip,omitempty"`

	// caller profile
	//  webitel.adt.caller caller = 4;
	LoggedIn string `json:"logged_in,omitempty"`

	// static
	Mac string `json:"mac,omitempty"`

	// model
	Model string `json:"model,omitempty"`

	// kind of
	// DeviceType type = 19;
	//
	// display name
	Name string `json:"name,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// provision extra data
	Provision map[string]string `json:"provision,omitempty"`

	// TODO: something like PresenceStatus instead of Registration
	//  repeated Registration contacts = 16;
	Reged bool `json:"reged,omitempty"`

	// unix
	UpdatedAt string `json:"updated_at,omitempty"`

	// user
	UpdatedBy *APIUserID `json:"updated_by,omitempty"`

	// [optional] The user that 'owns' the device;
	User *APIUserID `json:"user,omitempty"`
}

// Validate validates this devices update device2 params body
func (m *DevicesUpdateDevice2ParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeletedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DevicesUpdateDevice2ParamsBody) validateCreatedBy(formats strfmt.Registry) error {
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

func (m *DevicesUpdateDevice2ParamsBody) validateDeletedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.DeletedBy) { // not required
		return nil
	}

	if m.DeletedBy != nil {
		if err := m.DeletedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleted_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleted_by")
			}
			return err
		}
	}

	return nil
}

func (m *DevicesUpdateDevice2ParamsBody) validateUpdatedBy(formats strfmt.Registry) error {
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

func (m *DevicesUpdateDevice2ParamsBody) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this devices update device2 params body based on the context it is used
func (m *DevicesUpdateDevice2ParamsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeletedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DevicesUpdateDevice2ParamsBody) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *DevicesUpdateDevice2ParamsBody) contextValidateDeletedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.DeletedBy != nil {

		if swag.IsZero(m.DeletedBy) { // not required
			return nil
		}

		if err := m.DeletedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleted_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleted_by")
			}
			return err
		}
	}

	return nil
}

func (m *DevicesUpdateDevice2ParamsBody) contextValidateUpdatedBy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *DevicesUpdateDevice2ParamsBody) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {

		if swag.IsZero(m.User) { // not required
			return nil
		}

		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DevicesUpdateDevice2ParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DevicesUpdateDevice2ParamsBody) UnmarshalBinary(b []byte) error {
	var res DevicesUpdateDevice2ParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
