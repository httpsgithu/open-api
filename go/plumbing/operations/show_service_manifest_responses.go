// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netlify/open-api/v2/go/models"
)

// ShowServiceManifestReader is a Reader for the ShowServiceManifest structure.
type ShowServiceManifestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowServiceManifestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewShowServiceManifestCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewShowServiceManifestDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewShowServiceManifestCreated creates a ShowServiceManifestCreated with default headers values
func NewShowServiceManifestCreated() *ShowServiceManifestCreated {
	return &ShowServiceManifestCreated{}
}

/*ShowServiceManifestCreated handles this case with default header values.

retrieving from provider
*/
type ShowServiceManifestCreated struct {
	Payload interface{}
}

func (o *ShowServiceManifestCreated) Error() string {
	return fmt.Sprintf("[GET /services/{addonName}/manifest][%d] showServiceManifestCreated  %+v", 201, o.Payload)
}

func (o *ShowServiceManifestCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *ShowServiceManifestCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowServiceManifestDefault creates a ShowServiceManifestDefault with default headers values
func NewShowServiceManifestDefault(code int) *ShowServiceManifestDefault {
	return &ShowServiceManifestDefault{
		_statusCode: code,
	}
}

/*ShowServiceManifestDefault handles this case with default header values.

error
*/
type ShowServiceManifestDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the show service manifest default response
func (o *ShowServiceManifestDefault) Code() int {
	return o._statusCode
}

func (o *ShowServiceManifestDefault) Error() string {
	return fmt.Sprintf("[GET /services/{addonName}/manifest][%d] showServiceManifest default  %+v", o._statusCode, o.Payload)
}

func (o *ShowServiceManifestDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ShowServiceManifestDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
