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

// SimulationActivity simulation activity
// swagger:model SimulationActivity
type SimulationActivity struct {

	// the Amazon SWF activity id
	// Required: true
	ActivityID *string `json:"activityId"`

	// the name of the Amazon SWF activity
	// Required: true
	ActivityName *string `json:"activityName"`

	// completed time stamp
	CompletedAt *strfmt.DateTime `json:"completedAt,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// number 0 - 100 representing percentage of activity already completed
	PercentComplete int32 `json:"percentComplete,omitempty"`

	// id of associated simulation, set server-side, read-only
	// Required: true
	SimulationID *int32 `json:"simulationId"`

	// started time stamp
	StartedAt strfmt.DateTime `json:"startedAt,omitempty"`

	// status of the activity
	Status string `json:"status,omitempty"`

	// the version of the worker, e.g. 1.9.0
	// Required: true
	WorkerVersion *string `json:"workerVersion"`
}

// Validate validates this simulation activity
func (m *SimulationActivity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivityID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateActivityName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSimulationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWorkerVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SimulationActivity) validateActivityID(formats strfmt.Registry) error {

	if err := validate.Required("activityId", "body", m.ActivityID); err != nil {
		return err
	}

	return nil
}

func (m *SimulationActivity) validateActivityName(formats strfmt.Registry) error {

	if err := validate.Required("activityName", "body", m.ActivityName); err != nil {
		return err
	}

	return nil
}

func (m *SimulationActivity) validateSimulationID(formats strfmt.Registry) error {

	if err := validate.Required("simulationId", "body", m.SimulationID); err != nil {
		return err
	}

	return nil
}

var simulationActivityTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Waiting","Running","Cancelled","Error","Finished"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		simulationActivityTypeStatusPropEnum = append(simulationActivityTypeStatusPropEnum, v)
	}
}

const (
	// SimulationActivityStatusWaiting captures enum value "Waiting"
	SimulationActivityStatusWaiting string = "Waiting"
	// SimulationActivityStatusRunning captures enum value "Running"
	SimulationActivityStatusRunning string = "Running"
	// SimulationActivityStatusCancelled captures enum value "Cancelled"
	SimulationActivityStatusCancelled string = "Cancelled"
	// SimulationActivityStatusError captures enum value "Error"
	SimulationActivityStatusError string = "Error"
	// SimulationActivityStatusFinished captures enum value "Finished"
	SimulationActivityStatusFinished string = "Finished"
)

// prop value enum
func (m *SimulationActivity) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, simulationActivityTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SimulationActivity) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *SimulationActivity) validateWorkerVersion(formats strfmt.Registry) error {

	if err := validate.Required("workerVersion", "body", m.WorkerVersion); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SimulationActivity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SimulationActivity) UnmarshalBinary(b []byte) error {
	var res SimulationActivity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
