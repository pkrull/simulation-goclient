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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3dsim/simulation-goclient/models"
)

// NewPostMicrostructureSimulationParams creates a new PostMicrostructureSimulationParams object
// with the default values initialized.
func NewPostMicrostructureSimulationParams() *PostMicrostructureSimulationParams {
	var ()
	return &PostMicrostructureSimulationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMicrostructureSimulationParamsWithTimeout creates a new PostMicrostructureSimulationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMicrostructureSimulationParamsWithTimeout(timeout time.Duration) *PostMicrostructureSimulationParams {
	var ()
	return &PostMicrostructureSimulationParams{

		timeout: timeout,
	}
}

// NewPostMicrostructureSimulationParamsWithContext creates a new PostMicrostructureSimulationParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMicrostructureSimulationParamsWithContext(ctx context.Context) *PostMicrostructureSimulationParams {
	var ()
	return &PostMicrostructureSimulationParams{

		Context: ctx,
	}
}

// NewPostMicrostructureSimulationParamsWithHTTPClient creates a new PostMicrostructureSimulationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMicrostructureSimulationParamsWithHTTPClient(client *http.Client) *PostMicrostructureSimulationParams {
	var ()
	return &PostMicrostructureSimulationParams{
		HTTPClient: client,
	}
}

/*PostMicrostructureSimulationParams contains all the parameters to send to the API endpoint
for the post microstructure simulation operation typically these are written to a http.Request
*/
type PostMicrostructureSimulationParams struct {

	/*MicrostructureSimulation
	  MicrostructureSimulation fields required to add the simulation

	*/
	MicrostructureSimulation *models.MicrostructureSimulation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post microstructure simulation params
func (o *PostMicrostructureSimulationParams) WithTimeout(timeout time.Duration) *PostMicrostructureSimulationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post microstructure simulation params
func (o *PostMicrostructureSimulationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post microstructure simulation params
func (o *PostMicrostructureSimulationParams) WithContext(ctx context.Context) *PostMicrostructureSimulationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post microstructure simulation params
func (o *PostMicrostructureSimulationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post microstructure simulation params
func (o *PostMicrostructureSimulationParams) WithHTTPClient(client *http.Client) *PostMicrostructureSimulationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post microstructure simulation params
func (o *PostMicrostructureSimulationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMicrostructureSimulation adds the microstructureSimulation to the post microstructure simulation params
func (o *PostMicrostructureSimulationParams) WithMicrostructureSimulation(microstructureSimulation *models.MicrostructureSimulation) *PostMicrostructureSimulationParams {
	o.SetMicrostructureSimulation(microstructureSimulation)
	return o
}

// SetMicrostructureSimulation adds the microstructureSimulation to the post microstructure simulation params
func (o *PostMicrostructureSimulationParams) SetMicrostructureSimulation(microstructureSimulation *models.MicrostructureSimulation) {
	o.MicrostructureSimulation = microstructureSimulation
}

// WriteToRequest writes these params to a swagger request
func (o *PostMicrostructureSimulationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.MicrostructureSimulation == nil {
		o.MicrostructureSimulation = new(models.MicrostructureSimulation)
	}

	if err := r.SetBodyParam(o.MicrostructureSimulation); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
