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

// PostMicrostructureSimulationReader is a Reader for the PostMicrostructureSimulation structure.
type PostMicrostructureSimulationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMicrostructureSimulationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostMicrostructureSimulationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostMicrostructureSimulationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostMicrostructureSimulationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostMicrostructureSimulationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostMicrostructureSimulationOK creates a PostMicrostructureSimulationOK with default headers values
func NewPostMicrostructureSimulationOK() *PostMicrostructureSimulationOK {
	return &PostMicrostructureSimulationOK{}
}

/*PostMicrostructureSimulationOK handles this case with default header values.

Successfully added a simulation
*/
type PostMicrostructureSimulationOK struct {
	Payload *models.MicrostructureSimulation
}

func (o *PostMicrostructureSimulationOK) Error() string {
	return fmt.Sprintf("[POST /microstructuresimulations][%d] postMicrostructureSimulationOK  %+v", 200, o.Payload)
}

func (o *PostMicrostructureSimulationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MicrostructureSimulation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMicrostructureSimulationUnauthorized creates a PostMicrostructureSimulationUnauthorized with default headers values
func NewPostMicrostructureSimulationUnauthorized() *PostMicrostructureSimulationUnauthorized {
	return &PostMicrostructureSimulationUnauthorized{}
}

/*PostMicrostructureSimulationUnauthorized handles this case with default header values.

Not authorized
*/
type PostMicrostructureSimulationUnauthorized struct {
	Payload *models.Error
}

func (o *PostMicrostructureSimulationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /microstructuresimulations][%d] postMicrostructureSimulationUnauthorized  %+v", 401, o.Payload)
}

func (o *PostMicrostructureSimulationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMicrostructureSimulationForbidden creates a PostMicrostructureSimulationForbidden with default headers values
func NewPostMicrostructureSimulationForbidden() *PostMicrostructureSimulationForbidden {
	return &PostMicrostructureSimulationForbidden{}
}

/*PostMicrostructureSimulationForbidden handles this case with default header values.

Forbidden
*/
type PostMicrostructureSimulationForbidden struct {
	Payload *models.Error
}

func (o *PostMicrostructureSimulationForbidden) Error() string {
	return fmt.Sprintf("[POST /microstructuresimulations][%d] postMicrostructureSimulationForbidden  %+v", 403, o.Payload)
}

func (o *PostMicrostructureSimulationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMicrostructureSimulationDefault creates a PostMicrostructureSimulationDefault with default headers values
func NewPostMicrostructureSimulationDefault(code int) *PostMicrostructureSimulationDefault {
	return &PostMicrostructureSimulationDefault{
		_statusCode: code,
	}
}

/*PostMicrostructureSimulationDefault handles this case with default header values.

unexpected error
*/
type PostMicrostructureSimulationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post microstructure simulation default response
func (o *PostMicrostructureSimulationDefault) Code() int {
	return o._statusCode
}

func (o *PostMicrostructureSimulationDefault) Error() string {
	return fmt.Sprintf("[POST /microstructuresimulations][%d] postMicrostructureSimulation default  %+v", o._statusCode, o.Payload)
}

func (o *PostMicrostructureSimulationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}