// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateRepository21Reader is a Reader for the UpdateRepository21 structure.
type UpdateRepository21Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepository21Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRepository21NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRepository21Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRepository21Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepository21NoContent creates a UpdateRepository21NoContent with default headers values
func NewUpdateRepository21NoContent() *UpdateRepository21NoContent {
	return &UpdateRepository21NoContent{}
}

/*
UpdateRepository21NoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateRepository21NoContent struct {
}

// IsSuccess returns true when this update repository21 no content response has a 2xx status code
func (o *UpdateRepository21NoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update repository21 no content response has a 3xx status code
func (o *UpdateRepository21NoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository21 no content response has a 4xx status code
func (o *UpdateRepository21NoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update repository21 no content response has a 5xx status code
func (o *UpdateRepository21NoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository21 no content response a status code equal to that given
func (o *UpdateRepository21NoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update repository21 no content response
func (o *UpdateRepository21NoContent) Code() int {
	return 204
}

func (o *UpdateRepository21NoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/hosted/{repositoryName}][%d] updateRepository21NoContent ", 204)
}

func (o *UpdateRepository21NoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/hosted/{repositoryName}][%d] updateRepository21NoContent ", 204)
}

func (o *UpdateRepository21NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository21Unauthorized creates a UpdateRepository21Unauthorized with default headers values
func NewUpdateRepository21Unauthorized() *UpdateRepository21Unauthorized {
	return &UpdateRepository21Unauthorized{}
}

/*
UpdateRepository21Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateRepository21Unauthorized struct {
}

// IsSuccess returns true when this update repository21 unauthorized response has a 2xx status code
func (o *UpdateRepository21Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository21 unauthorized response has a 3xx status code
func (o *UpdateRepository21Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository21 unauthorized response has a 4xx status code
func (o *UpdateRepository21Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository21 unauthorized response has a 5xx status code
func (o *UpdateRepository21Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository21 unauthorized response a status code equal to that given
func (o *UpdateRepository21Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update repository21 unauthorized response
func (o *UpdateRepository21Unauthorized) Code() int {
	return 401
}

func (o *UpdateRepository21Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/hosted/{repositoryName}][%d] updateRepository21Unauthorized ", 401)
}

func (o *UpdateRepository21Unauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/hosted/{repositoryName}][%d] updateRepository21Unauthorized ", 401)
}

func (o *UpdateRepository21Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository21Forbidden creates a UpdateRepository21Forbidden with default headers values
func NewUpdateRepository21Forbidden() *UpdateRepository21Forbidden {
	return &UpdateRepository21Forbidden{}
}

/*
UpdateRepository21Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateRepository21Forbidden struct {
}

// IsSuccess returns true when this update repository21 forbidden response has a 2xx status code
func (o *UpdateRepository21Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository21 forbidden response has a 3xx status code
func (o *UpdateRepository21Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository21 forbidden response has a 4xx status code
func (o *UpdateRepository21Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository21 forbidden response has a 5xx status code
func (o *UpdateRepository21Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository21 forbidden response a status code equal to that given
func (o *UpdateRepository21Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update repository21 forbidden response
func (o *UpdateRepository21Forbidden) Code() int {
	return 403
}

func (o *UpdateRepository21Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/hosted/{repositoryName}][%d] updateRepository21Forbidden ", 403)
}

func (o *UpdateRepository21Forbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/hosted/{repositoryName}][%d] updateRepository21Forbidden ", 403)
}

func (o *UpdateRepository21Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
