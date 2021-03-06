// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ScanPatternSimulation scan pattern simulation
// swagger:model ScanPatternSimulation
type ScanPatternSimulation struct {
	Simulation

	PartBasedSimulationParameters

	ScanPatternSimulationParameters
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ScanPatternSimulation) UnmarshalJSON(raw []byte) error {

	var aO0 Simulation
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Simulation = aO0

	var aO1 PartBasedSimulationParameters
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.PartBasedSimulationParameters = aO1

	var aO2 ScanPatternSimulationParameters
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.ScanPatternSimulationParameters = aO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ScanPatternSimulation) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.Simulation)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.PartBasedSimulationParameters)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.ScanPatternSimulationParameters)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this scan pattern simulation
func (m *ScanPatternSimulation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.Simulation.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.PartBasedSimulationParameters.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.ScanPatternSimulationParameters.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ScanPatternSimulation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScanPatternSimulation) UnmarshalBinary(b []byte) error {
	var res ScanPatternSimulation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
