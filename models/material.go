package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Material material
// swagger:model Material
type Material struct {

	// false if not archived, true if archived
	// Required: true
	Archived *bool `json:"archived"`

	// configuration
	Configuration *MaterialConfiguration `json:"configuration,omitempty"`

	// configuration history
	ConfigurationHistory []*MaterialConfiguration `json:"configurationHistory"`

	// identifier for the active configuration for this material
	// Required: true
	ConfigurationID *int32 `json:"configurationId"`

	// created time stamp, set server-side, read only field
	// Required: true
	Created *strfmt.DateTime `json:"created"`

	// creating user, set server-side, read only field
	// Required: true
	CreatedBy *string `json:"createdBy"`

	// material description
	// Max Length: 2048
	Description string `json:"description,omitempty"`

	// item identifier
	// Required: true
	ID *int32 `json:"id"`

	// flag whether it is a core material or a custom material
	// Required: true
	IsCore *bool `json:"isCore"`

	// key associated with this material.  e.g. Ti64
	// Required: true
	// Max Length: 16
	Key *string `json:"key"`

	// last modified time stamp, set server-side, read only field
	// Required: true
	LastModified *strfmt.DateTime `json:"lastModified"`

	// modifying user, set server-side, read only field
	// Required: true
	LastModifiedBy *string `json:"lastModifiedBy"`

	// material name
	// Required: true
	// Max Length: 128
	Name *string `json:"name"`

	// organization identifier
	// Required: true
	OrganizationID *int32 `json:"organizationId"`

	// version label
	// Required: true
	// Max Length: 16
	Version *string `json:"version"`
}

// Validate validates this material
func (m *Material) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArchived(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConfiguration(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConfigurationHistory(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConfigurationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIsCore(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLastModified(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLastModifiedBy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Material) validateArchived(formats strfmt.Registry) error {

	if err := validate.Required("archived", "body", m.Archived); err != nil {
		return err
	}

	return nil
}

func (m *Material) validateConfiguration(formats strfmt.Registry) error {

	if swag.IsZero(m.Configuration) { // not required
		return nil
	}

	if m.Configuration != nil {

		if err := m.Configuration.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Material) validateConfigurationHistory(formats strfmt.Registry) error {

	if swag.IsZero(m.ConfigurationHistory) { // not required
		return nil
	}

	for i := 0; i < len(m.ConfigurationHistory); i++ {

		if swag.IsZero(m.ConfigurationHistory[i]) { // not required
			continue
		}

		if m.ConfigurationHistory[i] != nil {

			if err := m.ConfigurationHistory[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Material) validateConfigurationID(formats strfmt.Registry) error {

	if err := validate.Required("configurationId", "body", m.ConfigurationID); err != nil {
		return err
	}

	return nil
}

func (m *Material) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *Material) validateCreatedBy(formats strfmt.Registry) error {

	if err := validate.Required("createdBy", "body", m.CreatedBy); err != nil {
		return err
	}

	return nil
}

func (m *Material) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 2048); err != nil {
		return err
	}

	return nil
}

func (m *Material) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Material) validateIsCore(formats strfmt.Registry) error {

	if err := validate.Required("isCore", "body", m.IsCore); err != nil {
		return err
	}

	return nil
}

func (m *Material) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	if err := validate.MaxLength("key", "body", string(*m.Key), 16); err != nil {
		return err
	}

	return nil
}

func (m *Material) validateLastModified(formats strfmt.Registry) error {

	if err := validate.Required("lastModified", "body", m.LastModified); err != nil {
		return err
	}

	return nil
}

func (m *Material) validateLastModifiedBy(formats strfmt.Registry) error {

	if err := validate.Required("lastModifiedBy", "body", m.LastModifiedBy); err != nil {
		return err
	}

	return nil
}

func (m *Material) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 128); err != nil {
		return err
	}

	return nil
}

func (m *Material) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organizationId", "body", m.OrganizationID); err != nil {
		return err
	}

	return nil
}

func (m *Material) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if err := validate.MaxLength("version", "body", string(*m.Version), 16); err != nil {
		return err
	}

	return nil
}
