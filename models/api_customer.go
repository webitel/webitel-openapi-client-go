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

// APICustomer api customer
//
// swagger:model api.Customer
type APICustomer struct {

	// string version = 2; // version
	// local store details
	//
	// local created ms
	CreatedAt string `json:"created_at,omitempty"`

	// [optional] domains
	Dnsrv []*APIObjectID `json:"dnsrv"`

	// serial number assigned (global::Customer-ID)
	ID string `json:"id,omitempty"`

	// validity boundaries
	//
	// [required] issuer(CA) created at
	IssuedAt string `json:"issued_at,omitempty"`

	// int32 competitive = 13; // zero-based competitive sessions limit ?
	//
	// grants issued
	License []*APILicenseV1 `json:"license"`

	// extensions granted
	//
	// defines map[class]limit usage
	Limit map[string]int32 `json:"limit,omitempty"`

	// [optional]: signature expires; update required after
	NextUpdate string `json:"next_update,omitempty"`

	// [required] valid till
	NotAfter string `json:"not_after,omitempty"`

	// [optional] valid from
	NotBefore string `json:"not_before,omitempty"`

	// registration name
	Organization string `json:"organization,omitempty"`

	// local revoked ms
	RevokedAt string `json:"revoked_at,omitempty"`

	// signature validity
	//
	// [optional]: issuer(CA) signed at
	ThisUpdate string `json:"this_update,omitempty"`

	// last uploaded ms
	UpdatedAt string `json:"updated_at,omitempty"`

	// Verification status
	//
	// [optional] validation results
	Verify *APIVerification `json:"verify,omitempty"`
}

// Validate validates this api customer
func (m *APICustomer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDnsrv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLicense(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerify(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APICustomer) validateDnsrv(formats strfmt.Registry) error {
	if swag.IsZero(m.Dnsrv) { // not required
		return nil
	}

	for i := 0; i < len(m.Dnsrv); i++ {
		if swag.IsZero(m.Dnsrv[i]) { // not required
			continue
		}

		if m.Dnsrv[i] != nil {
			if err := m.Dnsrv[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dnsrv" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dnsrv" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APICustomer) validateLicense(formats strfmt.Registry) error {
	if swag.IsZero(m.License) { // not required
		return nil
	}

	for i := 0; i < len(m.License); i++ {
		if swag.IsZero(m.License[i]) { // not required
			continue
		}

		if m.License[i] != nil {
			if err := m.License[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("license" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("license" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APICustomer) validateVerify(formats strfmt.Registry) error {
	if swag.IsZero(m.Verify) { // not required
		return nil
	}

	if m.Verify != nil {
		if err := m.Verify.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("verify")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("verify")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this api customer based on the context it is used
func (m *APICustomer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDnsrv(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLicense(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVerify(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APICustomer) contextValidateDnsrv(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Dnsrv); i++ {

		if m.Dnsrv[i] != nil {

			if swag.IsZero(m.Dnsrv[i]) { // not required
				return nil
			}

			if err := m.Dnsrv[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dnsrv" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dnsrv" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APICustomer) contextValidateLicense(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.License); i++ {

		if m.License[i] != nil {

			if swag.IsZero(m.License[i]) { // not required
				return nil
			}

			if err := m.License[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("license" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("license" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APICustomer) contextValidateVerify(ctx context.Context, formats strfmt.Registry) error {

	if m.Verify != nil {

		if swag.IsZero(m.Verify) { // not required
			return nil
		}

		if err := m.Verify.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("verify")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("verify")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APICustomer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APICustomer) UnmarshalBinary(b []byte) error {
	var res APICustomer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
