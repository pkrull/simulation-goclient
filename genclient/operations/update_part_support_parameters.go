// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3dsim/simulation-goclient/models"
)

// NewUpdatePartSupportParams creates a new UpdatePartSupportParams object
// with the default values initialized.
func NewUpdatePartSupportParams() *UpdatePartSupportParams {
	var ()
	return &UpdatePartSupportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePartSupportParamsWithTimeout creates a new UpdatePartSupportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePartSupportParamsWithTimeout(timeout time.Duration) *UpdatePartSupportParams {
	var ()
	return &UpdatePartSupportParams{

		timeout: timeout,
	}
}

// NewUpdatePartSupportParamsWithContext creates a new UpdatePartSupportParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePartSupportParamsWithContext(ctx context.Context) *UpdatePartSupportParams {
	var ()
	return &UpdatePartSupportParams{

		Context: ctx,
	}
}

// NewUpdatePartSupportParamsWithHTTPClient creates a new UpdatePartSupportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePartSupportParamsWithHTTPClient(client *http.Client) *UpdatePartSupportParams {
	var ()
	return &UpdatePartSupportParams{
		HTTPClient: client,
	}
}

/*UpdatePartSupportParams contains all the parameters to send to the API endpoint
for the update part support operation typically these are written to a http.Request
*/
type UpdatePartSupportParams struct {

	/*PartSupport
	  PartSupport to update.

	*/
	PartSupport *models.PartSupport
	/*PartID
	  ID of part

	*/
	PartID int32
	/*SupportID
	  ID of support to get

	*/
	SupportID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update part support params
func (o *UpdatePartSupportParams) WithTimeout(timeout time.Duration) *UpdatePartSupportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update part support params
func (o *UpdatePartSupportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update part support params
func (o *UpdatePartSupportParams) WithContext(ctx context.Context) *UpdatePartSupportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update part support params
func (o *UpdatePartSupportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update part support params
func (o *UpdatePartSupportParams) WithHTTPClient(client *http.Client) *UpdatePartSupportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update part support params
func (o *UpdatePartSupportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPartSupport adds the partSupport to the update part support params
func (o *UpdatePartSupportParams) WithPartSupport(partSupport *models.PartSupport) *UpdatePartSupportParams {
	o.SetPartSupport(partSupport)
	return o
}

// SetPartSupport adds the partSupport to the update part support params
func (o *UpdatePartSupportParams) SetPartSupport(partSupport *models.PartSupport) {
	o.PartSupport = partSupport
}

// WithPartID adds the partID to the update part support params
func (o *UpdatePartSupportParams) WithPartID(partID int32) *UpdatePartSupportParams {
	o.SetPartID(partID)
	return o
}

// SetPartID adds the partId to the update part support params
func (o *UpdatePartSupportParams) SetPartID(partID int32) {
	o.PartID = partID
}

// WithSupportID adds the supportID to the update part support params
func (o *UpdatePartSupportParams) WithSupportID(supportID int32) *UpdatePartSupportParams {
	o.SetSupportID(supportID)
	return o
}

// SetSupportID adds the supportId to the update part support params
func (o *UpdatePartSupportParams) SetSupportID(supportID int32) {
	o.SupportID = supportID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePartSupportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PartSupport == nil {
		o.PartSupport = new(models.PartSupport)
	}

	if err := r.SetBodyParam(o.PartSupport); err != nil {
		return err
	}

	// path param partId
	if err := r.SetPathParam("partId", swag.FormatInt32(o.PartID)); err != nil {
		return err
	}

	// path param supportId
	if err := r.SetPathParam("supportId", swag.FormatInt32(o.SupportID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
