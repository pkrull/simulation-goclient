// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MicrostructureSimulationParameters microstructure simulation parameters
// swagger:model MicrostructureSimulationParameters
type MicrostructureSimulationParameters struct {

	// Height of part geometry, 0.001 to 0.01 meters
	// Required: true
	// Maximum: 0.01
	// Minimum: 0.001
	GeometryHeight *float64 `json:"geometryHeight"`

	// Length of part geometry, 0.001 to 0.01 meters
	// Required: true
	// Maximum: 0.01
	// Minimum: 0.001
	GeometryLength *float64 `json:"geometryLength"`

	// Width of part geometry, 0.001 to 0.01 meters
	// Required: true
	// Maximum: 0.01
	// Minimum: 0.001
	GeometryWidth *float64 `json:"geometryWidth"`

	// Array of hatch spacing values to simulate across. Each value must be 0.00001 to 0.001 meters
	// Required: true
	HatchSpacingValues []float64 `json:"hatchSpacingValues"`

	// Array of printer heater temperature values in degrees kelvin. Each value must be 293 and 474.
	HeaterTemperatureValues []float64 `json:"heaterTemperatureValues"`

	// Array of laser power values to simulate across. Each value must be 10 to 1000 watts
	LaserWattageValues []float64 `json:"laserWattageValues"`

	// Array of layer rotation angles to simulate across. Each value must be 0 to 180 degrees
	// Required: true
	LayerRotationAngleValues []*float64 `json:"layerRotationAngleValues"`

	// Array of powder deposit layer thickness values to simulate across. Each value must be 0.00001 to 0.0001 meters
	// Required: true
	LayerThicknessValues []float64 `json:"layerThicknessValues"`

	// List of sensor locations within the part where the microstructure will be evaluated
	// Required: true
	MicrostructureSensors []*MicrostructureSensor `json:"microstructureSensors"`

	// Seed value used for random number generation. If not provided it will be auto-generated.
	RandomSeed int32 `json:"randomSeed,omitempty"`

	// Array of scan speed values to simulate across. Each value must be 0.01 to 10 meters/second
	// Required: true
	ScanSpeedValues []float64 `json:"scanSpeedValues"`

	// Array of slicing strip width values to simulate across. Each value must be 0 to 0.1 meters. A value of 0 indicates no striping.
	// Required: true
	SlicingStripeWidthValues []*float64 `json:"slicingStripeWidthValues"`

	// Array of starting layer angles to simulate across. Each value must be 0 to 180 degrees
	// Required: true
	StartingLayerAngleValues []*float64 `json:"startingLayerAngleValues"`

	// True if user is providing cooling rate, thermal gradient and melt pool dimensions for each sensor. False if thermal solver should calculate these values. See MicrostructureSensor.
	// Required: true
	UserProvidedThermalData *bool `json:"userProvidedThermalData"`
}

// Validate validates this microstructure simulation parameters
func (m *MicrostructureSimulationParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGeometryHeight(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGeometryLength(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGeometryWidth(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHatchSpacingValues(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHeaterTemperatureValues(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLaserWattageValues(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLayerRotationAngleValues(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLayerThicknessValues(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMicrostructureSensors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateScanSpeedValues(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSlicingStripeWidthValues(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartingLayerAngleValues(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUserProvidedThermalData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MicrostructureSimulationParameters) validateGeometryHeight(formats strfmt.Registry) error {

	if err := validate.Required("geometryHeight", "body", m.GeometryHeight); err != nil {
		return err
	}

	if err := validate.Minimum("geometryHeight", "body", float64(*m.GeometryHeight), 0.001, false); err != nil {
		return err
	}

	if err := validate.Maximum("geometryHeight", "body", float64(*m.GeometryHeight), 0.01, false); err != nil {
		return err
	}

	return nil
}

func (m *MicrostructureSimulationParameters) validateGeometryLength(formats strfmt.Registry) error {

	if err := validate.Required("geometryLength", "body", m.GeometryLength); err != nil {
		return err
	}

	if err := validate.Minimum("geometryLength", "body", float64(*m.GeometryLength), 0.001, false); err != nil {
		return err
	}

	if err := validate.Maximum("geometryLength", "body", float64(*m.GeometryLength), 0.01, false); err != nil {
		return err
	}

	return nil
}

func (m *MicrostructureSimulationParameters) validateGeometryWidth(formats strfmt.Registry) error {

	if err := validate.Required("geometryWidth", "body", m.GeometryWidth); err != nil {
		return err
	}

	if err := validate.Minimum("geometryWidth", "body", float64(*m.GeometryWidth), 0.001, false); err != nil {
		return err
	}

	if err := validate.Maximum("geometryWidth", "body", float64(*m.GeometryWidth), 0.01, false); err != nil {
		return err
	}

	return nil
}

func (m *MicrostructureSimulationParameters) validateHatchSpacingValues(formats strfmt.Registry) error {

	if err := validate.Required("hatchSpacingValues", "body", m.HatchSpacingValues); err != nil {
		return err
	}

	for i := 0; i < len(m.HatchSpacingValues); i++ {

		if err := validate.Minimum("hatchSpacingValues"+"."+strconv.Itoa(i), "body", float64(m.HatchSpacingValues[i]), 1e-05, false); err != nil {
			return err
		}

		if err := validate.Maximum("hatchSpacingValues"+"."+strconv.Itoa(i), "body", float64(m.HatchSpacingValues[i]), 0.001, false); err != nil {
			return err
		}

	}

	return nil
}

func (m *MicrostructureSimulationParameters) validateHeaterTemperatureValues(formats strfmt.Registry) error {

	if swag.IsZero(m.HeaterTemperatureValues) { // not required
		return nil
	}

	for i := 0; i < len(m.HeaterTemperatureValues); i++ {

		if err := validate.Minimum("heaterTemperatureValues"+"."+strconv.Itoa(i), "body", float64(m.HeaterTemperatureValues[i]), 293, false); err != nil {
			return err
		}

		if err := validate.Maximum("heaterTemperatureValues"+"."+strconv.Itoa(i), "body", float64(m.HeaterTemperatureValues[i]), 474, false); err != nil {
			return err
		}

	}

	return nil
}

func (m *MicrostructureSimulationParameters) validateLaserWattageValues(formats strfmt.Registry) error {

	if swag.IsZero(m.LaserWattageValues) { // not required
		return nil
	}

	for i := 0; i < len(m.LaserWattageValues); i++ {

		if err := validate.Minimum("laserWattageValues"+"."+strconv.Itoa(i), "body", float64(m.LaserWattageValues[i]), 10, false); err != nil {
			return err
		}

		if err := validate.Maximum("laserWattageValues"+"."+strconv.Itoa(i), "body", float64(m.LaserWattageValues[i]), 1000, false); err != nil {
			return err
		}

	}

	return nil
}

func (m *MicrostructureSimulationParameters) validateLayerRotationAngleValues(formats strfmt.Registry) error {

	if err := validate.Required("layerRotationAngleValues", "body", m.LayerRotationAngleValues); err != nil {
		return err
	}

	for i := 0; i < len(m.LayerRotationAngleValues); i++ {

		if swag.IsZero(m.LayerRotationAngleValues[i]) { // not required
			continue
		}

		if err := validate.Minimum("layerRotationAngleValues"+"."+strconv.Itoa(i), "body", float64(*m.LayerRotationAngleValues[i]), 0, false); err != nil {
			return err
		}

		if err := validate.Maximum("layerRotationAngleValues"+"."+strconv.Itoa(i), "body", float64(*m.LayerRotationAngleValues[i]), 180, false); err != nil {
			return err
		}

	}

	return nil
}

func (m *MicrostructureSimulationParameters) validateLayerThicknessValues(formats strfmt.Registry) error {

	if err := validate.Required("layerThicknessValues", "body", m.LayerThicknessValues); err != nil {
		return err
	}

	for i := 0; i < len(m.LayerThicknessValues); i++ {

		if err := validate.Minimum("layerThicknessValues"+"."+strconv.Itoa(i), "body", float64(m.LayerThicknessValues[i]), 1e-05, false); err != nil {
			return err
		}

		if err := validate.Maximum("layerThicknessValues"+"."+strconv.Itoa(i), "body", float64(m.LayerThicknessValues[i]), 0.0001, false); err != nil {
			return err
		}

	}

	return nil
}

func (m *MicrostructureSimulationParameters) validateMicrostructureSensors(formats strfmt.Registry) error {

	if err := validate.Required("microstructureSensors", "body", m.MicrostructureSensors); err != nil {
		return err
	}

	for i := 0; i < len(m.MicrostructureSensors); i++ {

		if swag.IsZero(m.MicrostructureSensors[i]) { // not required
			continue
		}

		if m.MicrostructureSensors[i] != nil {

			if err := m.MicrostructureSensors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("microstructureSensors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MicrostructureSimulationParameters) validateScanSpeedValues(formats strfmt.Registry) error {

	if err := validate.Required("scanSpeedValues", "body", m.ScanSpeedValues); err != nil {
		return err
	}

	for i := 0; i < len(m.ScanSpeedValues); i++ {

		if err := validate.Minimum("scanSpeedValues"+"."+strconv.Itoa(i), "body", float64(m.ScanSpeedValues[i]), 0.01, false); err != nil {
			return err
		}

		if err := validate.Maximum("scanSpeedValues"+"."+strconv.Itoa(i), "body", float64(m.ScanSpeedValues[i]), 10, false); err != nil {
			return err
		}

	}

	return nil
}

func (m *MicrostructureSimulationParameters) validateSlicingStripeWidthValues(formats strfmt.Registry) error {

	if err := validate.Required("slicingStripeWidthValues", "body", m.SlicingStripeWidthValues); err != nil {
		return err
	}

	for i := 0; i < len(m.SlicingStripeWidthValues); i++ {

		if swag.IsZero(m.SlicingStripeWidthValues[i]) { // not required
			continue
		}

		if err := validate.Minimum("slicingStripeWidthValues"+"."+strconv.Itoa(i), "body", float64(*m.SlicingStripeWidthValues[i]), 0, false); err != nil {
			return err
		}

		if err := validate.Maximum("slicingStripeWidthValues"+"."+strconv.Itoa(i), "body", float64(*m.SlicingStripeWidthValues[i]), 0.1, false); err != nil {
			return err
		}

	}

	return nil
}

func (m *MicrostructureSimulationParameters) validateStartingLayerAngleValues(formats strfmt.Registry) error {

	if err := validate.Required("startingLayerAngleValues", "body", m.StartingLayerAngleValues); err != nil {
		return err
	}

	for i := 0; i < len(m.StartingLayerAngleValues); i++ {

		if swag.IsZero(m.StartingLayerAngleValues[i]) { // not required
			continue
		}

		if err := validate.Minimum("startingLayerAngleValues"+"."+strconv.Itoa(i), "body", float64(*m.StartingLayerAngleValues[i]), 0, false); err != nil {
			return err
		}

		if err := validate.Maximum("startingLayerAngleValues"+"."+strconv.Itoa(i), "body", float64(*m.StartingLayerAngleValues[i]), 180, false); err != nil {
			return err
		}

	}

	return nil
}

func (m *MicrostructureSimulationParameters) validateUserProvidedThermalData(formats strfmt.Registry) error {

	if err := validate.Required("userProvidedThermalData", "body", m.UserProvidedThermalData); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MicrostructureSimulationParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MicrostructureSimulationParameters) UnmarshalBinary(b []byte) error {
	var res MicrostructureSimulationParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
