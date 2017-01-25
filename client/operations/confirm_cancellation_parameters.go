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

// NewConfirmCancellationParams creates a new ConfirmCancellationParams object
// with the default values initialized.
func NewConfirmCancellationParams() *ConfirmCancellationParams {
	var ()
	return &ConfirmCancellationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConfirmCancellationParamsWithTimeout creates a new ConfirmCancellationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConfirmCancellationParamsWithTimeout(timeout time.Duration) *ConfirmCancellationParams {
	var ()
	return &ConfirmCancellationParams{

		timeout: timeout,
	}
}

// NewConfirmCancellationParamsWithContext creates a new ConfirmCancellationParams object
// with the default values initialized, and the ability to set a context for a request
func NewConfirmCancellationParamsWithContext(ctx context.Context) *ConfirmCancellationParams {
	var ()
	return &ConfirmCancellationParams{

		Context: ctx,
	}
}

/*ConfirmCancellationParams contains all the parameters to send to the API endpoint
for the confirm cancellation operation typically these are written to a http.Request
*/
type ConfirmCancellationParams struct {

	/*ID
	  ID of part

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the confirm cancellation params
func (o *ConfirmCancellationParams) WithTimeout(timeout time.Duration) *ConfirmCancellationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the confirm cancellation params
func (o *ConfirmCancellationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the confirm cancellation params
func (o *ConfirmCancellationParams) WithContext(ctx context.Context) *ConfirmCancellationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the confirm cancellation params
func (o *ConfirmCancellationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the confirm cancellation params
func (o *ConfirmCancellationParams) WithID(id int32) *ConfirmCancellationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the confirm cancellation params
func (o *ConfirmCancellationParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ConfirmCancellationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
