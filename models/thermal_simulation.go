// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ThermalSimulation thermal simulation
// swagger:model ThermalSimulation
type ThermalSimulation struct {
	PartBasedSimulationParameters

	ThermalSimulationParameters
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ThermalSimulation) UnmarshalJSON(raw []byte) error {

	var aO0 PartBasedSimulationParameters
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PartBasedSimulationParameters = aO0

	var aO1 ThermalSimulationParameters
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.ThermalSimulationParameters = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ThermalSimulation) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.PartBasedSimulationParameters)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.ThermalSimulationParameters)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this thermal simulation
func (m *ThermalSimulation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.PartBasedSimulationParameters.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.ThermalSimulationParameters.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ThermalSimulation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThermalSimulation) UnmarshalBinary(b []byte) error {
	var res ThermalSimulation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
