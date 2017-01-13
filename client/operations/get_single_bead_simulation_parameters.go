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

// NewGetSingleBeadSimulationParams creates a new GetSingleBeadSimulationParams object
// with the default values initialized.
func NewGetSingleBeadSimulationParams() *GetSingleBeadSimulationParams {
	var ()
	return &GetSingleBeadSimulationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSingleBeadSimulationParamsWithTimeout creates a new GetSingleBeadSimulationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSingleBeadSimulationParamsWithTimeout(timeout time.Duration) *GetSingleBeadSimulationParams {
	var ()
	return &GetSingleBeadSimulationParams{

		timeout: timeout,
	}
}

// NewGetSingleBeadSimulationParamsWithContext creates a new GetSingleBeadSimulationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSingleBeadSimulationParamsWithContext(ctx context.Context) *GetSingleBeadSimulationParams {
	var ()
	return &GetSingleBeadSimulationParams{

		Context: ctx,
	}
}

/*GetSingleBeadSimulationParams contains all the parameters to send to the API endpoint
for the get single bead simulation operation typically these are written to a http.Request
*/
type GetSingleBeadSimulationParams struct {

	/*ID
	  simulation identifier

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get single bead simulation params
func (o *GetSingleBeadSimulationParams) WithTimeout(timeout time.Duration) *GetSingleBeadSimulationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get single bead simulation params
func (o *GetSingleBeadSimulationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get single bead simulation params
func (o *GetSingleBeadSimulationParams) WithContext(ctx context.Context) *GetSingleBeadSimulationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get single bead simulation params
func (o *GetSingleBeadSimulationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the get single bead simulation params
func (o *GetSingleBeadSimulationParams) WithID(id int32) *GetSingleBeadSimulationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get single bead simulation params
func (o *GetSingleBeadSimulationParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSingleBeadSimulationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
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