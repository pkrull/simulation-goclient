// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3dsim/simulation-goclient/models"
)

// CancelSingleBeadSimulationReader is a Reader for the CancelSingleBeadSimulation structure.
type CancelSingleBeadSimulationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelSingleBeadSimulationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCancelSingleBeadSimulationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCancelSingleBeadSimulationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCancelSingleBeadSimulationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCancelSingleBeadSimulationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCancelSingleBeadSimulationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCancelSingleBeadSimulationOK creates a CancelSingleBeadSimulationOK with default headers values
func NewCancelSingleBeadSimulationOK() *CancelSingleBeadSimulationOK {
	return &CancelSingleBeadSimulationOK{}
}

/*CancelSingleBeadSimulationOK handles this case with default header values.

Simulation was successfully cancelled.
*/
type CancelSingleBeadSimulationOK struct {
}

func (o *CancelSingleBeadSimulationOK) Error() string {
	return fmt.Sprintf("[PUT /singlebeadsimulations/{id}/cancel][%d] cancelSingleBeadSimulationOK ", 200)
}

func (o *CancelSingleBeadSimulationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCancelSingleBeadSimulationUnauthorized creates a CancelSingleBeadSimulationUnauthorized with default headers values
func NewCancelSingleBeadSimulationUnauthorized() *CancelSingleBeadSimulationUnauthorized {
	return &CancelSingleBeadSimulationUnauthorized{}
}

/*CancelSingleBeadSimulationUnauthorized handles this case with default header values.

Not authorized
*/
type CancelSingleBeadSimulationUnauthorized struct {
	Payload *models.Error
}

func (o *CancelSingleBeadSimulationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /singlebeadsimulations/{id}/cancel][%d] cancelSingleBeadSimulationUnauthorized  %+v", 401, o.Payload)
}

func (o *CancelSingleBeadSimulationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelSingleBeadSimulationForbidden creates a CancelSingleBeadSimulationForbidden with default headers values
func NewCancelSingleBeadSimulationForbidden() *CancelSingleBeadSimulationForbidden {
	return &CancelSingleBeadSimulationForbidden{}
}

/*CancelSingleBeadSimulationForbidden handles this case with default header values.

Forbidden
*/
type CancelSingleBeadSimulationForbidden struct {
	Payload *models.Error
}

func (o *CancelSingleBeadSimulationForbidden) Error() string {
	return fmt.Sprintf("[PUT /singlebeadsimulations/{id}/cancel][%d] cancelSingleBeadSimulationForbidden  %+v", 403, o.Payload)
}

func (o *CancelSingleBeadSimulationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelSingleBeadSimulationNotFound creates a CancelSingleBeadSimulationNotFound with default headers values
func NewCancelSingleBeadSimulationNotFound() *CancelSingleBeadSimulationNotFound {
	return &CancelSingleBeadSimulationNotFound{}
}

/*CancelSingleBeadSimulationNotFound handles this case with default header values.

Simulation not found (id invalid)
*/
type CancelSingleBeadSimulationNotFound struct {
	Payload *models.Error
}

func (o *CancelSingleBeadSimulationNotFound) Error() string {
	return fmt.Sprintf("[PUT /singlebeadsimulations/{id}/cancel][%d] cancelSingleBeadSimulationNotFound  %+v", 404, o.Payload)
}

func (o *CancelSingleBeadSimulationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelSingleBeadSimulationDefault creates a CancelSingleBeadSimulationDefault with default headers values
func NewCancelSingleBeadSimulationDefault(code int) *CancelSingleBeadSimulationDefault {
	return &CancelSingleBeadSimulationDefault{
		_statusCode: code,
	}
}

/*CancelSingleBeadSimulationDefault handles this case with default header values.

unexpected error
*/
type CancelSingleBeadSimulationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the cancel single bead simulation default response
func (o *CancelSingleBeadSimulationDefault) Code() int {
	return o._statusCode
}

func (o *CancelSingleBeadSimulationDefault) Error() string {
	return fmt.Sprintf("[PUT /singlebeadsimulations/{id}/cancel][%d] cancelSingleBeadSimulation default  %+v", o._statusCode, o.Payload)
}

func (o *CancelSingleBeadSimulationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
