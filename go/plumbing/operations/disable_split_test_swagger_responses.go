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

// DisableSplitTestReader is a Reader for the DisableSplitTest structure.
type DisableSplitTestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DisableSplitTestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDisableSplitTestNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDisableSplitTestDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDisableSplitTestNoContent creates a DisableSplitTestNoContent with default headers values
func NewDisableSplitTestNoContent() *DisableSplitTestNoContent {
	return &DisableSplitTestNoContent{}
}

/*DisableSplitTestNoContent handles this case with default header values.

disabled
*/
type DisableSplitTestNoContent struct {
}

func (o *DisableSplitTestNoContent) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/traffic_splits/{split_test_id}/unpublish][%d] disableSplitTestNoContent ", 204)
}

func (o *DisableSplitTestNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisableSplitTestDefault creates a DisableSplitTestDefault with default headers values
func NewDisableSplitTestDefault(code int) *DisableSplitTestDefault {
	return &DisableSplitTestDefault{
		_statusCode: code,
	}
}

/*DisableSplitTestDefault handles this case with default header values.

error
*/
type DisableSplitTestDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the disable split test default response
func (o *DisableSplitTestDefault) Code() int {
	return o._statusCode
}

func (o *DisableSplitTestDefault) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/traffic_splits/{split_test_id}/unpublish][%d] disableSplitTest default  %+v", o._statusCode, o.Payload)
}

func (o *DisableSplitTestDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DisableSplitTestDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
