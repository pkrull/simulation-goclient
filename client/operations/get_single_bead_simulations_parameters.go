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

// NewGetSingleBeadSimulationsParams creates a new GetSingleBeadSimulationsParams object
// with the default values initialized.
func NewGetSingleBeadSimulationsParams() *GetSingleBeadSimulationsParams {
	var ()
	return &GetSingleBeadSimulationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSingleBeadSimulationsParamsWithTimeout creates a new GetSingleBeadSimulationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSingleBeadSimulationsParamsWithTimeout(timeout time.Duration) *GetSingleBeadSimulationsParams {
	var ()
	return &GetSingleBeadSimulationsParams{

		timeout: timeout,
	}
}

// NewGetSingleBeadSimulationsParamsWithContext creates a new GetSingleBeadSimulationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSingleBeadSimulationsParamsWithContext(ctx context.Context) *GetSingleBeadSimulationsParams {
	var ()
	return &GetSingleBeadSimulationsParams{

		Context: ctx,
	}
}

/*GetSingleBeadSimulationsParams contains all the parameters to send to the API endpoint
for the get single bead simulations operation typically these are written to a http.Request
*/
type GetSingleBeadSimulationsParams struct {

	/*Limit
	  number of items to return within the query

	*/
	Limit *int32
	/*Offset
	  starting paging count; ex. offset of 60 will skip the first 60 items in the list

	*/
	Offset *int32
	/*OrganizationID
	  the organization id to get items for.  Must be provided as API callers only have access to items belonging to their organization.

	*/
	OrganizationID int32
	/*Sort
	  key:direction pairs for one or multiple field sort orders

	*/
	Sort []string
	/*Status
	  simulation status for items retrieved.  If an array of items is sent, they are treated as "OR" operations. e.g. status=InProgress,Requested would yield a list of simulations that are in either state.

	*/
	Status []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get single bead simulations params
func (o *GetSingleBeadSimulationsParams) WithTimeout(timeout time.Duration) *GetSingleBeadSimulationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get single bead simulations params
func (o *GetSingleBeadSimulationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get single bead simulations params
func (o *GetSingleBeadSimulationsParams) WithContext(ctx context.Context) *GetSingleBeadSimulationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get single bead simulations params
func (o *GetSingleBeadSimulationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithLimit adds the limit to the get single bead simulations params
func (o *GetSingleBeadSimulationsParams) WithLimit(limit *int32) *GetSingleBeadSimulationsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get single bead simulations params
func (o *GetSingleBeadSimulationsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get single bead simulations params
func (o *GetSingleBeadSimulationsParams) WithOffset(offset *int32) *GetSingleBeadSimulationsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get single bead simulations params
func (o *GetSingleBeadSimulationsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrganizationID adds the organizationID to the get single bead simulations params
func (o *GetSingleBeadSimulationsParams) WithOrganizationID(organizationID int32) *GetSingleBeadSimulationsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get single bead simulations params
func (o *GetSingleBeadSimulationsParams) SetOrganizationID(organizationID int32) {
	o.OrganizationID = organizationID
}

// WithSort adds the sort to the get single bead simulations params
func (o *GetSingleBeadSimulationsParams) WithSort(sort []string) *GetSingleBeadSimulationsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get single bead simulations params
func (o *GetSingleBeadSimulationsParams) SetSort(sort []string) {
	o.Sort = sort
}

// WithStatus adds the status to the get single bead simulations params
func (o *GetSingleBeadSimulationsParams) WithStatus(status []string) *GetSingleBeadSimulationsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get single bead simulations params
func (o *GetSingleBeadSimulationsParams) SetStatus(status []string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *GetSingleBeadSimulationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	// query param organizationId
	qrOrganizationID := o.OrganizationID
	qOrganizationID := swag.FormatInt32(qrOrganizationID)
	if qOrganizationID != "" {
		if err := r.SetQueryParam("organizationId", qOrganizationID); err != nil {
			return err
		}
	}

	valuesSort := o.Sort

	joinedSort := swag.JoinByFormat(valuesSort, "csv")
	// query array param sort
	if err := r.SetQueryParam("sort", joinedSort...); err != nil {
		return err
	}

	valuesStatus := o.Status

	joinedStatus := swag.JoinByFormat(valuesStatus, "")
	// query array param status
	if err := r.SetQueryParam("status", joinedStatus...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}