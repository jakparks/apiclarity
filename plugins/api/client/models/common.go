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

// Common common
//
// swagger:model Common
type Common struct {

	// truncated body
	TruncatedBody bool `json:"TruncatedBody,omitempty"`

	// body
	// Format: byte
	Body strfmt.Base64 `json:"body,omitempty"`

	// headers
	Headers []*Header `json:"headers"`

	// time
	Time int64 `json:"time,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this common
func (m *Common) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeaders(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Common) validateHeaders(formats strfmt.Registry) error {
	if swag.IsZero(m.Headers) { // not required
		return nil
	}

	for i := 0; i < len(m.Headers); i++ {
		if swag.IsZero(m.Headers[i]) { // not required
			continue
		}

		if m.Headers[i] != nil {
			if err := m.Headers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("headers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this common based on the context it is used
func (m *Common) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHeaders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Common) contextValidateHeaders(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Headers); i++ {

		if m.Headers[i] != nil {
			if err := m.Headers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("headers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Common) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Common) UnmarshalBinary(b []byte) error {
	var res Common
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
