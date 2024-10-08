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

// APILDAPUpdateLDAPCatalogBodyCatalog LDAP Catalog changes.
//
// LDAP Catalog changes.
//
// swagger:model apiLdapUpdateLdapCatalogBodyCatalog
type APILDAPUpdateLDAPCatalogBodyCatalog struct {

	// base_dn, aka domain e.g.: 'dc=example,dc=org'
	BaseDn string `json:"base_dn,omitempty"`

	// ----- BIND: Authorization -----
	//
	// authorization method e.g.: SIMPLE, SAML, NTLM, etc.
	Bind string `json:"bind,omitempty"`

	// basic: last operation details
	//
	// unix
	CreatedAt string `json:"created_at,omitempty"`

	// user
	CreatedBy *APIUserID `json:"created_by,omitempty"`

	// ----- RSYNC Processing -----
	// int32  priority = 4;
	//
	// [CRON]OS SPEC ! github.com/robfig/cron/v3
	Cron string `json:"cron,omitempty"`

	// last modification timestamp attribute type
	EntryModify string `json:"entry_modify,omitempty"`

	// ----- DOMAIN: Entry model -----
	//
	// entry unique id (except DN which may be updated)
	EntryUUID string `json:"entry_uuid,omitempty"`

	// Meaningfull title name
	Name string `json:"name,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// Last OR Current SYNC process details
	Process *APILDAProcess `json:"process,omitempty"`

	// sync
	Sync *APILDAProcessOptions `json:"sync,omitempty"`

	// // // ----- MAP: user(s) settings -----
	// LDAPUsersMap users = 26; // DEPRECATED: Use Template(s) instead
	// SYNC Object's search and mappings
	Templates []*APILDAPTemplate `json:"templates"`

	// tls
	TLS *APILDAPCatalogTLSConfig `json:"tls,omitempty"`

	// unix
	UpdatedAt string `json:"updated_at,omitempty"`

	// user
	UpdatedBy *APIUserID `json:"updated_by,omitempty"`

	// ----- Connection Options -----
	//
	// URL e.g.: [(ldap|ldapi|ldaps)://]host[:port]
	URL string `json:"url,omitempty"`

	// bind_dn
	Username string `json:"username,omitempty"`
}

// Validate validates this api Ldap update Ldap catalog body catalog
func (m *APILDAPUpdateLDAPCatalogBodyCatalog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSync(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTLS(formats); err != nil {
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

func (m *APILDAPUpdateLDAPCatalogBodyCatalog) validateCreatedBy(formats strfmt.Registry) error {
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

func (m *APILDAPUpdateLDAPCatalogBodyCatalog) validateProcess(formats strfmt.Registry) error {
	if swag.IsZero(m.Process) { // not required
		return nil
	}

	if m.Process != nil {
		if err := m.Process.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("process")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("process")
			}
			return err
		}
	}

	return nil
}

func (m *APILDAPUpdateLDAPCatalogBodyCatalog) validateSync(formats strfmt.Registry) error {
	if swag.IsZero(m.Sync) { // not required
		return nil
	}

	if m.Sync != nil {
		if err := m.Sync.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sync")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sync")
			}
			return err
		}
	}

	return nil
}

func (m *APILDAPUpdateLDAPCatalogBodyCatalog) validateTemplates(formats strfmt.Registry) error {
	if swag.IsZero(m.Templates) { // not required
		return nil
	}

	for i := 0; i < len(m.Templates); i++ {
		if swag.IsZero(m.Templates[i]) { // not required
			continue
		}

		if m.Templates[i] != nil {
			if err := m.Templates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("templates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("templates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APILDAPUpdateLDAPCatalogBodyCatalog) validateTLS(formats strfmt.Registry) error {
	if swag.IsZero(m.TLS) { // not required
		return nil
	}

	if m.TLS != nil {
		if err := m.TLS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tls")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tls")
			}
			return err
		}
	}

	return nil
}

func (m *APILDAPUpdateLDAPCatalogBodyCatalog) validateUpdatedBy(formats strfmt.Registry) error {
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

// ContextValidate validate this api Ldap update Ldap catalog body catalog based on the context it is used
func (m *APILDAPUpdateLDAPCatalogBodyCatalog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProcess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSync(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTemplates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTLS(ctx, formats); err != nil {
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

func (m *APILDAPUpdateLDAPCatalogBodyCatalog) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *APILDAPUpdateLDAPCatalogBodyCatalog) contextValidateProcess(ctx context.Context, formats strfmt.Registry) error {

	if m.Process != nil {

		if swag.IsZero(m.Process) { // not required
			return nil
		}

		if err := m.Process.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("process")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("process")
			}
			return err
		}
	}

	return nil
}

func (m *APILDAPUpdateLDAPCatalogBodyCatalog) contextValidateSync(ctx context.Context, formats strfmt.Registry) error {

	if m.Sync != nil {

		if swag.IsZero(m.Sync) { // not required
			return nil
		}

		if err := m.Sync.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sync")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sync")
			}
			return err
		}
	}

	return nil
}

func (m *APILDAPUpdateLDAPCatalogBodyCatalog) contextValidateTemplates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Templates); i++ {

		if m.Templates[i] != nil {

			if swag.IsZero(m.Templates[i]) { // not required
				return nil
			}

			if err := m.Templates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("templates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("templates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APILDAPUpdateLDAPCatalogBodyCatalog) contextValidateTLS(ctx context.Context, formats strfmt.Registry) error {

	if m.TLS != nil {

		if swag.IsZero(m.TLS) { // not required
			return nil
		}

		if err := m.TLS.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tls")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tls")
			}
			return err
		}
	}

	return nil
}

func (m *APILDAPUpdateLDAPCatalogBodyCatalog) contextValidateUpdatedBy(ctx context.Context, formats strfmt.Registry) error {

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
func (m *APILDAPUpdateLDAPCatalogBodyCatalog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APILDAPUpdateLDAPCatalogBodyCatalog) UnmarshalBinary(b []byte) error {
	var res APILDAPUpdateLDAPCatalogBodyCatalog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
