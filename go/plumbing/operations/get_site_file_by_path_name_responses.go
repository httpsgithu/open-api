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

// GetSiteFileByPathNameReader is a Reader for the GetSiteFileByPathName structure.
type GetSiteFileByPathNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSiteFileByPathNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSiteFileByPathNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSiteFileByPathNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSiteFileByPathNameOK creates a GetSiteFileByPathNameOK with default headers values
func NewGetSiteFileByPathNameOK() *GetSiteFileByPathNameOK {
	return &GetSiteFileByPathNameOK{}
}

/*GetSiteFileByPathNameOK handles this case with default header values.

OK
*/
type GetSiteFileByPathNameOK struct {
	Payload *models.File
}

func (o *GetSiteFileByPathNameOK) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/files/{file_path}][%d] getSiteFileByPathNameOK  %+v", 200, o.Payload)
}

func (o *GetSiteFileByPathNameOK) GetPayload() *models.File {
	return o.Payload
}

func (o *GetSiteFileByPathNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.File)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSiteFileByPathNameDefault creates a GetSiteFileByPathNameDefault with default headers values
func NewGetSiteFileByPathNameDefault(code int) *GetSiteFileByPathNameDefault {
	return &GetSiteFileByPathNameDefault{
		_statusCode: code,
	}
}

/*GetSiteFileByPathNameDefault handles this case with default header values.

error
*/
type GetSiteFileByPathNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get site file by path name default response
func (o *GetSiteFileByPathNameDefault) Code() int {
	return o._statusCode
}

func (o *GetSiteFileByPathNameDefault) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/files/{file_path}][%d] getSiteFileByPathName default  %+v", o._statusCode, o.Payload)
}

func (o *GetSiteFileByPathNameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSiteFileByPathNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
