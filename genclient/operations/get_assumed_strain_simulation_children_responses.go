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

// GetAssumedStrainSimulationChildrenReader is a Reader for the GetAssumedStrainSimulationChildren structure.
type GetAssumedStrainSimulationChildrenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAssumedStrainSimulationChildrenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAssumedStrainSimulationChildrenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetAssumedStrainSimulationChildrenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAssumedStrainSimulationChildrenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetAssumedStrainSimulationChildrenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAssumedStrainSimulationChildrenOK creates a GetAssumedStrainSimulationChildrenOK with default headers values
func NewGetAssumedStrainSimulationChildrenOK() *GetAssumedStrainSimulationChildrenOK {
	return &GetAssumedStrainSimulationChildrenOK{}
}

/*GetAssumedStrainSimulationChildrenOK handles this case with default header values.

Successfully found the list of items
*/
type GetAssumedStrainSimulationChildrenOK struct {
	/*Contains paging information in json format - totalCount, totalPages
	 */
	XPagination string

	Payload []*models.AssumedStrainSimulation
}

func (o *GetAssumedStrainSimulationChildrenOK) Error() string {
	return fmt.Sprintf("[GET /assumedstrainsimulations/{id}/children][%d] getAssumedStrainSimulationChildrenOK  %+v", 200, o.Payload)
}

func (o *GetAssumedStrainSimulationChildrenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Pagination
	o.XPagination = response.GetHeader("X-Pagination")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssumedStrainSimulationChildrenUnauthorized creates a GetAssumedStrainSimulationChildrenUnauthorized with default headers values
func NewGetAssumedStrainSimulationChildrenUnauthorized() *GetAssumedStrainSimulationChildrenUnauthorized {
	return &GetAssumedStrainSimulationChildrenUnauthorized{}
}

/*GetAssumedStrainSimulationChildrenUnauthorized handles this case with default header values.

Not authorized
*/
type GetAssumedStrainSimulationChildrenUnauthorized struct {
	Payload *models.Error
}

func (o *GetAssumedStrainSimulationChildrenUnauthorized) Error() string {
	return fmt.Sprintf("[GET /assumedstrainsimulations/{id}/children][%d] getAssumedStrainSimulationChildrenUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAssumedStrainSimulationChildrenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssumedStrainSimulationChildrenForbidden creates a GetAssumedStrainSimulationChildrenForbidden with default headers values
func NewGetAssumedStrainSimulationChildrenForbidden() *GetAssumedStrainSimulationChildrenForbidden {
	return &GetAssumedStrainSimulationChildrenForbidden{}
}

/*GetAssumedStrainSimulationChildrenForbidden handles this case with default header values.

Forbidden
*/
type GetAssumedStrainSimulationChildrenForbidden struct {
	Payload *models.Error
}

func (o *GetAssumedStrainSimulationChildrenForbidden) Error() string {
	return fmt.Sprintf("[GET /assumedstrainsimulations/{id}/children][%d] getAssumedStrainSimulationChildrenForbidden  %+v", 403, o.Payload)
}

func (o *GetAssumedStrainSimulationChildrenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssumedStrainSimulationChildrenDefault creates a GetAssumedStrainSimulationChildrenDefault with default headers values
func NewGetAssumedStrainSimulationChildrenDefault(code int) *GetAssumedStrainSimulationChildrenDefault {
	return &GetAssumedStrainSimulationChildrenDefault{
		_statusCode: code,
	}
}

/*GetAssumedStrainSimulationChildrenDefault handles this case with default header values.

unexpected error
*/
type GetAssumedStrainSimulationChildrenDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get assumed strain simulation children default response
func (o *GetAssumedStrainSimulationChildrenDefault) Code() int {
	return o._statusCode
}

func (o *GetAssumedStrainSimulationChildrenDefault) Error() string {
	return fmt.Sprintf("[GET /assumedstrainsimulations/{id}/children][%d] getAssumedStrainSimulationChildren default  %+v", o._statusCode, o.Payload)
}

func (o *GetAssumedStrainSimulationChildrenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
