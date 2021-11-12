// Code generated by go-swagger; DO NOT EDIT.

package email

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SetEmailConfigurationReader is a Reader for the SetEmailConfiguration structure.
type SetEmailConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetEmailConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSetEmailConfigurationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetEmailConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSetEmailConfigurationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetEmailConfigurationNoContent creates a SetEmailConfigurationNoContent with default headers values
func NewSetEmailConfigurationNoContent() *SetEmailConfigurationNoContent {
	return &SetEmailConfigurationNoContent{}
}

/* SetEmailConfigurationNoContent describes a response with status code 204, with default header values.

Email configuration was successfully updated
*/
type SetEmailConfigurationNoContent struct {
}

func (o *SetEmailConfigurationNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/email][%d] setEmailConfigurationNoContent ", 204)
}

func (o *SetEmailConfigurationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetEmailConfigurationBadRequest creates a SetEmailConfigurationBadRequest with default headers values
func NewSetEmailConfigurationBadRequest() *SetEmailConfigurationBadRequest {
	return &SetEmailConfigurationBadRequest{}
}

/* SetEmailConfigurationBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type SetEmailConfigurationBadRequest struct {
}

func (o *SetEmailConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/email][%d] setEmailConfigurationBadRequest ", 400)
}

func (o *SetEmailConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetEmailConfigurationForbidden creates a SetEmailConfigurationForbidden with default headers values
func NewSetEmailConfigurationForbidden() *SetEmailConfigurationForbidden {
	return &SetEmailConfigurationForbidden{}
}

/* SetEmailConfigurationForbidden describes a response with status code 403, with default header values.

Insufficient permissions to update the email configuration
*/
type SetEmailConfigurationForbidden struct {
}

func (o *SetEmailConfigurationForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/email][%d] setEmailConfigurationForbidden ", 403)
}

func (o *SetEmailConfigurationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
