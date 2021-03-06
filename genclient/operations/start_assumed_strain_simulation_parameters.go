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
)

// NewStartAssumedStrainSimulationParams creates a new StartAssumedStrainSimulationParams object
// with the default values initialized.
func NewStartAssumedStrainSimulationParams() *StartAssumedStrainSimulationParams {
	var ()
	return &StartAssumedStrainSimulationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStartAssumedStrainSimulationParamsWithTimeout creates a new StartAssumedStrainSimulationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStartAssumedStrainSimulationParamsWithTimeout(timeout time.Duration) *StartAssumedStrainSimulationParams {
	var ()
	return &StartAssumedStrainSimulationParams{

		timeout: timeout,
	}
}

// NewStartAssumedStrainSimulationParamsWithContext creates a new StartAssumedStrainSimulationParams object
// with the default values initialized, and the ability to set a context for a request
func NewStartAssumedStrainSimulationParamsWithContext(ctx context.Context) *StartAssumedStrainSimulationParams {
	var ()
	return &StartAssumedStrainSimulationParams{

		Context: ctx,
	}
}

// NewStartAssumedStrainSimulationParamsWithHTTPClient creates a new StartAssumedStrainSimulationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStartAssumedStrainSimulationParamsWithHTTPClient(client *http.Client) *StartAssumedStrainSimulationParams {
	var ()
	return &StartAssumedStrainSimulationParams{
		HTTPClient: client,
	}
}

/*StartAssumedStrainSimulationParams contains all the parameters to send to the API endpoint
for the start assumed strain simulation operation typically these are written to a http.Request
*/
type StartAssumedStrainSimulationParams struct {

	/*ID
	  simulation identifier

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the start assumed strain simulation params
func (o *StartAssumedStrainSimulationParams) WithTimeout(timeout time.Duration) *StartAssumedStrainSimulationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start assumed strain simulation params
func (o *StartAssumedStrainSimulationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start assumed strain simulation params
func (o *StartAssumedStrainSimulationParams) WithContext(ctx context.Context) *StartAssumedStrainSimulationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start assumed strain simulation params
func (o *StartAssumedStrainSimulationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start assumed strain simulation params
func (o *StartAssumedStrainSimulationParams) WithHTTPClient(client *http.Client) *StartAssumedStrainSimulationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start assumed strain simulation params
func (o *StartAssumedStrainSimulationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the start assumed strain simulation params
func (o *StartAssumedStrainSimulationParams) WithID(id int32) *StartAssumedStrainSimulationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the start assumed strain simulation params
func (o *StartAssumedStrainSimulationParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *StartAssumedStrainSimulationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
