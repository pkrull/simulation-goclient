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

// PorositySimulationParameters porosity simulation parameters
// swagger:model PorositySimulationParameters
type PorositySimulationParameters struct {

	// Must be between 0.001 to 0.01 meters
	// Required: true
	// Maximum: 0.01
	// Minimum: 0.001
	GeometryHeight *float64 `json:"geometryHeight"`

	// Must be between 0.001 to 0.01 meters
	// Required: true
	// Maximum: 0.01
	// Minimum: 0.001
	GeometryLength *float64 `json:"geometryLength"`

	// Must be between 0.001 to 0.01 meters
	// Required: true
	// Maximum: 0.01
	// Minimum: 0.001
	GeometryWidth *float64 `json:"geometryWidth"`

	// Array of hatch spacing values to simulate across, Each value must be between 0.00001 to 0.001 meters
	// Required: true
	HatchSpacingValues []float64 `json:"hatchSpacingValues"`

	// heater temperature in degrees kelvin
	// Maximum: 474
	// Minimum: 293
	HeaterTemperature float64 `json:"heaterTemperature,omitempty"`

	// Array of laser power values to simulate across, Each value must be between 10 to 1000 watts
	// Required: true
	LaserWattageValues []float64 `json:"laserWattageValues"`

	// Must be between 0 to 180 degrees
	// Required: true
	// Maximum: 180
	// Minimum: 0
	LayerRotationAngle *float64 `json:"layerRotationAngle"`

	// Array of layer thickness values to simulate across, Each value must be between 0.00001 to 0.0001 meters
	// Required: true
	LayerThicknessValues []float64 `json:"layerThicknessValues"`

	// Number of mesh layers used to simulate a single deposit layer
	// Maximum: 100
	// Minimum: 1
	MeshLayersPerLayer int32 `json:"meshLayersPerLayer,omitempty"`

	// Array of origin x values to simulate across, Each value must be between 0 to 0.2 meters
	// Required: true
	OriginXValues []*float64 `json:"originXValues"`

	// Array of origin y values to simulate across, Each value must be between 0 to 0.2 meters
	// Required: true
	OriginYValues []*float64 `json:"originYValues"`

	// Array of scan speed values to simulate across, Each value must be between 0.01 to 10 meters/second
	// Required: true
	ScanSpeedValues []float64 `json:"scanSpeedValues"`

	// Array of slicing strip width values to simulate across, Each value must be between 0.001 to 0.1 meters
	// Required: true
	SlicingStripeWidthValues []float64 `json:"slicingStripeWidthValues"`

	// Array of solidus temperature values to simulate across
	// Required: true
	SolidusTemperatureValues []float64 `json:"solidusTemperatureValues"`

	// Must be between 0 to 180 degrees
	// Required: true
	// Maximum: 180
	// Minimum: 0
	StartingLayerAngle *float64 `json:"startingLayerAngle"`
}

// Validate validates this porosity simulation parameters
func (m *PorositySimulationParameters) Validate(formats strfmt.Registry) error {
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

	if err := m.validateHeaterTemperature(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLaserWattageValues(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLayerRotationAngle(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLayerThicknessValues(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMeshLayersPerLayer(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOriginXValues(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOriginYValues(formats); err != nil {
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

	if err := m.validateSolidusTemperatureValues(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartingLayerAngle(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PorositySimulationParameters) validateGeometryHeight(formats strfmt.Registry) error {

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

func (m *PorositySimulationParameters) validateGeometryLength(formats strfmt.Registry) error {

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

func (m *PorositySimulationParameters) validateGeometryWidth(formats strfmt.Registry) error {

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

func (m *PorositySimulationParameters) validateHatchSpacingValues(formats strfmt.Registry) error {

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

func (m *PorositySimulationParameters) validateHeaterTemperature(formats strfmt.Registry) error {

	if swag.IsZero(m.HeaterTemperature) { // not required
		return nil
	}

	if err := validate.Minimum("heaterTemperature", "body", float64(m.HeaterTemperature), 293, false); err != nil {
		return err
	}

	if err := validate.Maximum("heaterTemperature", "body", float64(m.HeaterTemperature), 474, false); err != nil {
		return err
	}

	return nil
}

func (m *PorositySimulationParameters) validateLaserWattageValues(formats strfmt.Registry) error {

	if err := validate.Required("laserWattageValues", "body", m.LaserWattageValues); err != nil {
		return err
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

func (m *PorositySimulationParameters) validateLayerRotationAngle(formats strfmt.Registry) error {

	if err := validate.Required("layerRotationAngle", "body", m.LayerRotationAngle); err != nil {
		return err
	}

	if err := validate.Minimum("layerRotationAngle", "body", float64(*m.LayerRotationAngle), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("layerRotationAngle", "body", float64(*m.LayerRotationAngle), 180, false); err != nil {
		return err
	}

	return nil
}

func (m *PorositySimulationParameters) validateLayerThicknessValues(formats strfmt.Registry) error {

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

func (m *PorositySimulationParameters) validateMeshLayersPerLayer(formats strfmt.Registry) error {

	if swag.IsZero(m.MeshLayersPerLayer) { // not required
		return nil
	}

	if err := validate.MinimumInt("meshLayersPerLayer", "body", int64(m.MeshLayersPerLayer), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("meshLayersPerLayer", "body", int64(m.MeshLayersPerLayer), 100, false); err != nil {
		return err
	}

	return nil
}

func (m *PorositySimulationParameters) validateOriginXValues(formats strfmt.Registry) error {

	if err := validate.Required("originXValues", "body", m.OriginXValues); err != nil {
		return err
	}

	for i := 0; i < len(m.OriginXValues); i++ {

		if swag.IsZero(m.OriginXValues[i]) { // not required
			continue
		}

		if err := validate.Minimum("originXValues"+"."+strconv.Itoa(i), "body", float64(*m.OriginXValues[i]), 0, false); err != nil {
			return err
		}

		if err := validate.Maximum("originXValues"+"."+strconv.Itoa(i), "body", float64(*m.OriginXValues[i]), 0.2, false); err != nil {
			return err
		}

	}

	return nil
}

func (m *PorositySimulationParameters) validateOriginYValues(formats strfmt.Registry) error {

	if err := validate.Required("originYValues", "body", m.OriginYValues); err != nil {
		return err
	}

	for i := 0; i < len(m.OriginYValues); i++ {

		if swag.IsZero(m.OriginYValues[i]) { // not required
			continue
		}

		if err := validate.Minimum("originYValues"+"."+strconv.Itoa(i), "body", float64(*m.OriginYValues[i]), 0, false); err != nil {
			return err
		}

		if err := validate.Maximum("originYValues"+"."+strconv.Itoa(i), "body", float64(*m.OriginYValues[i]), 0.2, false); err != nil {
			return err
		}

	}

	return nil
}

func (m *PorositySimulationParameters) validateScanSpeedValues(formats strfmt.Registry) error {

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

func (m *PorositySimulationParameters) validateSlicingStripeWidthValues(formats strfmt.Registry) error {

	if err := validate.Required("slicingStripeWidthValues", "body", m.SlicingStripeWidthValues); err != nil {
		return err
	}

	for i := 0; i < len(m.SlicingStripeWidthValues); i++ {

		if err := validate.Minimum("slicingStripeWidthValues"+"."+strconv.Itoa(i), "body", float64(m.SlicingStripeWidthValues[i]), 0.001, false); err != nil {
			return err
		}

		if err := validate.Maximum("slicingStripeWidthValues"+"."+strconv.Itoa(i), "body", float64(m.SlicingStripeWidthValues[i]), 0.1, false); err != nil {
			return err
		}

	}

	return nil
}

func (m *PorositySimulationParameters) validateSolidusTemperatureValues(formats strfmt.Registry) error {

	if err := validate.Required("solidusTemperatureValues", "body", m.SolidusTemperatureValues); err != nil {
		return err
	}

	return nil
}

func (m *PorositySimulationParameters) validateStartingLayerAngle(formats strfmt.Registry) error {

	if err := validate.Required("startingLayerAngle", "body", m.StartingLayerAngle); err != nil {
		return err
	}

	if err := validate.Minimum("startingLayerAngle", "body", float64(*m.StartingLayerAngle), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("startingLayerAngle", "body", float64(*m.StartingLayerAngle), 180, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PorositySimulationParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PorositySimulationParameters) UnmarshalBinary(b []byte) error {
	var res PorositySimulationParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
