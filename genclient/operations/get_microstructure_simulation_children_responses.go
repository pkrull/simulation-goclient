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

// GetMicrostructureSimulationChildrenReader is a Reader for the GetMicrostructureSimulationChildren structure.
type GetMicrostructureSimulationChildrenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMicrostructureSimulationChildrenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMicrostructureSimulationChildrenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetMicrostructureSimulationChildrenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetMicrostructureSimulationChildrenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetMicrostructureSimulationChildrenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMicrostructureSimulationChildrenOK creates a GetMicrostructureSimulationChildrenOK with default headers values
func NewGetMicrostructureSimulationChildrenOK() *GetMicrostructureSimulationChildrenOK {
	return &GetMicrostructureSimulationChildrenOK{}
}

/*GetMicrostructureSimulationChildrenOK handles this case with default header values.

Successfully found the list of items
*/
type GetMicrostructureSimulationChildrenOK struct {
	/*Contains paging information in json format - totalCount, totalPages
	 */
	XPagination string

	Payload []*models.MicrostructureSimulation
}

func (o *GetMicrostructureSimulationChildrenOK) Error() string {
	return fmt.Sprintf("[GET /microstructuresimulations/{id}/children][%d] getMicrostructureSimulationChildrenOK  %+v", 200, o.Payload)
}

func (o *GetMicrostructureSimulationChildrenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Pagination
	o.XPagination = response.GetHeader("X-Pagination")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMicrostructureSimulationChildrenUnauthorized creates a GetMicrostructureSimulationChildrenUnauthorized with default headers values
func NewGetMicrostructureSimulationChildrenUnauthorized() *GetMicrostructureSimulationChildrenUnauthorized {
	return &GetMicrostructureSimulationChildrenUnauthorized{}
}

/*GetMicrostructureSimulationChildrenUnauthorized handles this case with default header values.

Not authorized
*/
type GetMicrostructureSimulationChildrenUnauthorized struct {
	Payload *models.Error
}

func (o *GetMicrostructureSimulationChildrenUnauthorized) Error() string {
	return fmt.Sprintf("[GET /microstructuresimulations/{id}/children][%d] getMicrostructureSimulationChildrenUnauthorized  %+v", 401, o.Payload)
}

func (o *GetMicrostructureSimulationChildrenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMicrostructureSimulationChildrenForbidden creates a GetMicrostructureSimulationChildrenForbidden with default headers values
func NewGetMicrostructureSimulationChildrenForbidden() *GetMicrostructureSimulationChildrenForbidden {
	return &GetMicrostructureSimulationChildrenForbidden{}
}

/*GetMicrostructureSimulationChildrenForbidden handles this case with default header values.

Forbidden
*/
type GetMicrostructureSimulationChildrenForbidden struct {
	Payload *models.Error
}

func (o *GetMicrostructureSimulationChildrenForbidden) Error() string {
	return fmt.Sprintf("[GET /microstructuresimulations/{id}/children][%d] getMicrostructureSimulationChildrenForbidden  %+v", 403, o.Payload)
}

func (o *GetMicrostructureSimulationChildrenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMicrostructureSimulationChildrenDefault creates a GetMicrostructureSimulationChildrenDefault with default headers values
func NewGetMicrostructureSimulationChildrenDefault(code int) *GetMicrostructureSimulationChildrenDefault {
	return &GetMicrostructureSimulationChildrenDefault{
		_statusCode: code,
	}
}

/*GetMicrostructureSimulationChildrenDefault handles this case with default header values.

unexpected error
*/
type GetMicrostructureSimulationChildrenDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get microstructure simulation children default response
func (o *GetMicrostructureSimulationChildrenDefault) Code() int {
	return o._statusCode
}

func (o *GetMicrostructureSimulationChildrenDefault) Error() string {
	return fmt.Sprintf("[GET /microstructuresimulations/{id}/children][%d] getMicrostructureSimulationChildren default  %+v", o._statusCode, o.Payload)
}

func (o *GetMicrostructureSimulationChildrenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}