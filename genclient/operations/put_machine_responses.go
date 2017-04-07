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

// PutMachineReader is a Reader for the PutMachine structure.
type PutMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutMachineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPutMachineUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPutMachineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutMachineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutMachineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutMachineOK creates a PutMachineOK with default headers values
func NewPutMachineOK() *PutMachineOK {
	return &PutMachineOK{}
}

/*PutMachineOK handles this case with default header values.

Machine was successfully updated.
*/
type PutMachineOK struct {
	Payload *models.Machine
}

func (o *PutMachineOK) Error() string {
	return fmt.Sprintf("[PUT /machines/{id}][%d] putMachineOK  %+v", 200, o.Payload)
}

func (o *PutMachineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Machine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutMachineUnauthorized creates a PutMachineUnauthorized with default headers values
func NewPutMachineUnauthorized() *PutMachineUnauthorized {
	return &PutMachineUnauthorized{}
}

/*PutMachineUnauthorized handles this case with default header values.

Not authorized
*/
type PutMachineUnauthorized struct {
	Payload *models.Error
}

func (o *PutMachineUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /machines/{id}][%d] putMachineUnauthorized  %+v", 401, o.Payload)
}

func (o *PutMachineUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutMachineForbidden creates a PutMachineForbidden with default headers values
func NewPutMachineForbidden() *PutMachineForbidden {
	return &PutMachineForbidden{}
}

/*PutMachineForbidden handles this case with default header values.

Forbidden
*/
type PutMachineForbidden struct {
	Payload *models.Error
}

func (o *PutMachineForbidden) Error() string {
	return fmt.Sprintf("[PUT /machines/{id}][%d] putMachineForbidden  %+v", 403, o.Payload)
}

func (o *PutMachineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutMachineNotFound creates a PutMachineNotFound with default headers values
func NewPutMachineNotFound() *PutMachineNotFound {
	return &PutMachineNotFound{}
}

/*PutMachineNotFound handles this case with default header values.

Machine not found (id invalid)
*/
type PutMachineNotFound struct {
}

func (o *PutMachineNotFound) Error() string {
	return fmt.Sprintf("[PUT /machines/{id}][%d] putMachineNotFound ", 404)
}

func (o *PutMachineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutMachineDefault creates a PutMachineDefault with default headers values
func NewPutMachineDefault(code int) *PutMachineDefault {
	return &PutMachineDefault{
		_statusCode: code,
	}
}

/*PutMachineDefault handles this case with default header values.

unexpected error
*/
type PutMachineDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put machine default response
func (o *PutMachineDefault) Code() int {
	return o._statusCode
}

func (o *PutMachineDefault) Error() string {
	return fmt.Sprintf("[PUT /machines/{id}][%d] putMachine default  %+v", o._statusCode, o.Payload)
}

func (o *PutMachineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}