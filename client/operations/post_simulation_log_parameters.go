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

// NewPostSimulationLogParams creates a new PostSimulationLogParams object
// with the default values initialized.
func NewPostSimulationLogParams() *PostSimulationLogParams {
	var ()
	return &PostSimulationLogParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSimulationLogParamsWithTimeout creates a new PostSimulationLogParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSimulationLogParamsWithTimeout(timeout time.Duration) *PostSimulationLogParams {
	var ()
	return &PostSimulationLogParams{

		timeout: timeout,
	}
}

// NewPostSimulationLogParamsWithContext creates a new PostSimulationLogParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSimulationLogParamsWithContext(ctx context.Context) *PostSimulationLogParams {
	var ()
	return &PostSimulationLogParams{

		Context: ctx,
	}
}

/*PostSimulationLogParams contains all the parameters to send to the API endpoint
for the post simulation log operation typically these are written to a http.Request
*/
type PostSimulationLogParams struct {

	/*SimulationLog
	  A log for a simulation

	*/
	SimulationLog *models.SimulationLog
	/*ID
	  simulation identifier

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post simulation log params
func (o *PostSimulationLogParams) WithTimeout(timeout time.Duration) *PostSimulationLogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post simulation log params
func (o *PostSimulationLogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post simulation log params
func (o *PostSimulationLogParams) WithContext(ctx context.Context) *PostSimulationLogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post simulation log params
func (o *PostSimulationLogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithSimulationLog adds the simulationLog to the post simulation log params
func (o *PostSimulationLogParams) WithSimulationLog(simulationLog *models.SimulationLog) *PostSimulationLogParams {
	o.SetSimulationLog(simulationLog)
	return o
}

// SetSimulationLog adds the simulationLog to the post simulation log params
func (o *PostSimulationLogParams) SetSimulationLog(simulationLog *models.SimulationLog) {
	o.SimulationLog = simulationLog
}

// WithID adds the id to the post simulation log params
func (o *PostSimulationLogParams) WithID(id int32) *PostSimulationLogParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post simulation log params
func (o *PostSimulationLogParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostSimulationLogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.SimulationLog == nil {
		o.SimulationLog = new(models.SimulationLog)
	}

	if err := r.SetBodyParam(o.SimulationLog); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
