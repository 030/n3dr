// Code generated by go-swagger; DO NOT EDIT.

package security_management_anonymous_access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/030/n3dr/internal/app/n3dr/goswagger/models"
)

// UpdateReader is a Reader for the Update structure.
type UpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateOK creates a UpdateOK with default headers values
func NewUpdateOK() *UpdateOK {
	return &UpdateOK{}
}

/* UpdateOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateOK struct {
	Payload *models.AnonymousAccessSettingsXO
}

func (o *UpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1/security/anonymous][%d] updateOK  %+v", 200, o.Payload)
}
func (o *UpdateOK) GetPayload() *models.AnonymousAccessSettingsXO {
	return o.Payload
}

func (o *UpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AnonymousAccessSettingsXO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateForbidden creates a UpdateForbidden with default headers values
func NewUpdateForbidden() *UpdateForbidden {
	return &UpdateForbidden{}
}

/* UpdateForbidden describes a response with status code 403, with default header values.

Insufficient permissions to update settings
*/
type UpdateForbidden struct {
}

func (o *UpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/security/anonymous][%d] updateForbidden ", 403)
}

func (o *UpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}