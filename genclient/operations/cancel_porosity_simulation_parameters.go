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

// NewCancelPorositySimulationParams creates a new CancelPorositySimulationParams object
// with the default values initialized.
func NewCancelPorositySimulationParams() *CancelPorositySimulationParams {
	var ()
	return &CancelPorositySimulationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCancelPorositySimulationParamsWithTimeout creates a new CancelPorositySimulationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCancelPorositySimulationParamsWithTimeout(timeout time.Duration) *CancelPorositySimulationParams {
	var ()
	return &CancelPorositySimulationParams{

		timeout: timeout,
	}
}

// NewCancelPorositySimulationParamsWithContext creates a new CancelPorositySimulationParams object
// with the default values initialized, and the ability to set a context for a request
func NewCancelPorositySimulationParamsWithContext(ctx context.Context) *CancelPorositySimulationParams {
	var ()
	return &CancelPorositySimulationParams{

		Context: ctx,
	}
}

// NewCancelPorositySimulationParamsWithHTTPClient creates a new CancelPorositySimulationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCancelPorositySimulationParamsWithHTTPClient(client *http.Client) *CancelPorositySimulationParams {
	var ()
	return &CancelPorositySimulationParams{
		HTTPClient: client,
	}
}

/*CancelPorositySimulationParams contains all the parameters to send to the API endpoint
for the cancel porosity simulation operation typically these are written to a http.Request
*/
type CancelPorositySimulationParams struct {

	/*ID
	  simulation identifier

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cancel porosity simulation params
func (o *CancelPorositySimulationParams) WithTimeout(timeout time.Duration) *CancelPorositySimulationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel porosity simulation params
func (o *CancelPorositySimulationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel porosity simulation params
func (o *CancelPorositySimulationParams) WithContext(ctx context.Context) *CancelPorositySimulationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel porosity simulation params
func (o *CancelPorositySimulationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel porosity simulation params
func (o *CancelPorositySimulationParams) WithHTTPClient(client *http.Client) *CancelPorositySimulationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel porosity simulation params
func (o *CancelPorositySimulationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the cancel porosity simulation params
func (o *CancelPorositySimulationParams) WithID(id int32) *CancelPorositySimulationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the cancel porosity simulation params
func (o *CancelPorositySimulationParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CancelPorositySimulationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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