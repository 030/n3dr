// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateRepository9Reader is a Reader for the UpdateRepository9 structure.
type UpdateRepository9Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepository9Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRepository9NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRepository9Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRepository9Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepository9NoContent creates a UpdateRepository9NoContent with default headers values
func NewUpdateRepository9NoContent() *UpdateRepository9NoContent {
	return &UpdateRepository9NoContent{}
}

/* UpdateRepository9NoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateRepository9NoContent struct {
}

func (o *UpdateRepository9NoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/npm/hosted/{repositoryName}][%d] updateRepository9NoContent ", 204)
}

func (o *UpdateRepository9NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository9Unauthorized creates a UpdateRepository9Unauthorized with default headers values
func NewUpdateRepository9Unauthorized() *UpdateRepository9Unauthorized {
	return &UpdateRepository9Unauthorized{}
}

/* UpdateRepository9Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateRepository9Unauthorized struct {
}

func (o *UpdateRepository9Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/npm/hosted/{repositoryName}][%d] updateRepository9Unauthorized ", 401)
}

func (o *UpdateRepository9Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository9Forbidden creates a UpdateRepository9Forbidden with default headers values
func NewUpdateRepository9Forbidden() *UpdateRepository9Forbidden {
	return &UpdateRepository9Forbidden{}
}

/* UpdateRepository9Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateRepository9Forbidden struct {
}

func (o *UpdateRepository9Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/npm/hosted/{repositoryName}][%d] updateRepository9Forbidden ", 403)
}

func (o *UpdateRepository9Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}