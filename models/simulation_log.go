// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SimulationLog simulation log
// swagger:model SimulationLog
type SimulationLog struct {

	// the Amazon SWF activity id
	ActivityID string `json:"activityId,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// level of the log
	Level string `json:"level,omitempty"`

	// time stamp of log
	// Required: true
	LoggedAt *strfmt.DateTime `json:"loggedAt"`

	// message
	// Required: true
	Message *string `json:"message"`

	// id of associated simulation, set server-side, read-only
	SimulationID int32 `json:"simulationId,omitempty"`
}

// Validate validates this simulation log
func (m *SimulationLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLevel(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLoggedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var simulationLogTypeLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Info","Trace","Error","Warn"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		simulationLogTypeLevelPropEnum = append(simulationLogTypeLevelPropEnum, v)
	}
}

const (
	// SimulationLogLevelInfo captures enum value "Info"
	SimulationLogLevelInfo string = "Info"
	// SimulationLogLevelTrace captures enum value "Trace"
	SimulationLogLevelTrace string = "Trace"
	// SimulationLogLevelError captures enum value "Error"
	SimulationLogLevelError string = "Error"
	// SimulationLogLevelWarn captures enum value "Warn"
	SimulationLogLevelWarn string = "Warn"
)

// prop value enum
func (m *SimulationLog) validateLevelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, simulationLogTypeLevelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SimulationLog) validateLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.Level) { // not required
		return nil
	}

	// value enum
	if err := m.validateLevelEnum("level", "body", m.Level); err != nil {
		return err
	}

	return nil
}

func (m *SimulationLog) validateLoggedAt(formats strfmt.Registry) error {

	if err := validate.Required("loggedAt", "body", m.LoggedAt); err != nil {
		return err
	}

	return nil
}

func (m *SimulationLog) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SimulationLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SimulationLog) UnmarshalBinary(b []byte) error {
	var res SimulationLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
