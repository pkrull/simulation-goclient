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

// PutAssumedStrainSimulationReader is a Reader for the PutAssumedStrainSimulation structure.
type PutAssumedStrainSimulationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAssumedStrainSimulationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutAssumedStrainSimulationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPutAssumedStrainSimulationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPutAssumedStrainSimulationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutAssumedStrainSimulationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutAssumedStrainSimulationOK creates a PutAssumedStrainSimulationOK with default headers values
func NewPutAssumedStrainSimulationOK() *PutAssumedStrainSimulationOK {
	return &PutAssumedStrainSimulationOK{}
}

/*PutAssumedStrainSimulationOK handles this case with default header values.

Successfully updated a simulation
*/
type PutAssumedStrainSimulationOK struct {
	Payload *models.AssumedStrainSimulation
}

func (o *PutAssumedStrainSimulationOK) Error() string {
	return fmt.Sprintf("[PUT /assumedstrainsimulations/{id}][%d] putAssumedStrainSimulationOK  %+v", 200, o.Payload)
}

func (o *PutAssumedStrainSimulationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssumedStrainSimulation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAssumedStrainSimulationUnauthorized creates a PutAssumedStrainSimulationUnauthorized with default headers values
func NewPutAssumedStrainSimulationUnauthorized() *PutAssumedStrainSimulationUnauthorized {
	return &PutAssumedStrainSimulationUnauthorized{}
}

/*PutAssumedStrainSimulationUnauthorized handles this case with default header values.

Not authorized
*/
type PutAssumedStrainSimulationUnauthorized struct {
	Payload *models.Error
}

func (o *PutAssumedStrainSimulationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /assumedstrainsimulations/{id}][%d] putAssumedStrainSimulationUnauthorized  %+v", 401, o.Payload)
}

func (o *PutAssumedStrainSimulationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAssumedStrainSimulationForbidden creates a PutAssumedStrainSimulationForbidden with default headers values
func NewPutAssumedStrainSimulationForbidden() *PutAssumedStrainSimulationForbidden {
	return &PutAssumedStrainSimulationForbidden{}
}

/*PutAssumedStrainSimulationForbidden handles this case with default header values.

Forbidden
*/
type PutAssumedStrainSimulationForbidden struct {
	Payload *models.Error
}

func (o *PutAssumedStrainSimulationForbidden) Error() string {
	return fmt.Sprintf("[PUT /assumedstrainsimulations/{id}][%d] putAssumedStrainSimulationForbidden  %+v", 403, o.Payload)
}

func (o *PutAssumedStrainSimulationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAssumedStrainSimulationDefault creates a PutAssumedStrainSimulationDefault with default headers values
func NewPutAssumedStrainSimulationDefault(code int) *PutAssumedStrainSimulationDefault {
	return &PutAssumedStrainSimulationDefault{
		_statusCode: code,
	}
}

/*PutAssumedStrainSimulationDefault handles this case with default header values.

unexpected error
*/
type PutAssumedStrainSimulationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put assumed strain simulation default response
func (o *PutAssumedStrainSimulationDefault) Code() int {
	return o._statusCode
}

func (o *PutAssumedStrainSimulationDefault) Error() string {
	return fmt.Sprintf("[PUT /assumedstrainsimulations/{id}][%d] putAssumedStrainSimulation default  %+v", o._statusCode, o.Payload)
}

func (o *PutAssumedStrainSimulationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
