// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateRepository10Reader is a Reader for the UpdateRepository10 structure.
type UpdateRepository10Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepository10Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRepository10NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRepository10Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRepository10Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepository10NoContent creates a UpdateRepository10NoContent with default headers values
func NewUpdateRepository10NoContent() *UpdateRepository10NoContent {
	return &UpdateRepository10NoContent{}
}

/*
UpdateRepository10NoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateRepository10NoContent struct {
}

// IsSuccess returns true when this update repository10 no content response has a 2xx status code
func (o *UpdateRepository10NoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update repository10 no content response has a 3xx status code
func (o *UpdateRepository10NoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository10 no content response has a 4xx status code
func (o *UpdateRepository10NoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update repository10 no content response has a 5xx status code
func (o *UpdateRepository10NoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository10 no content response a status code equal to that given
func (o *UpdateRepository10NoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update repository10 no content response
func (o *UpdateRepository10NoContent) Code() int {
	return 204
}

func (o *UpdateRepository10NoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/npm/proxy/{repositoryName}][%d] updateRepository10NoContent ", 204)
}

func (o *UpdateRepository10NoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/npm/proxy/{repositoryName}][%d] updateRepository10NoContent ", 204)
}

func (o *UpdateRepository10NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository10Unauthorized creates a UpdateRepository10Unauthorized with default headers values
func NewUpdateRepository10Unauthorized() *UpdateRepository10Unauthorized {
	return &UpdateRepository10Unauthorized{}
}

/*
UpdateRepository10Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateRepository10Unauthorized struct {
}

// IsSuccess returns true when this update repository10 unauthorized response has a 2xx status code
func (o *UpdateRepository10Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository10 unauthorized response has a 3xx status code
func (o *UpdateRepository10Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository10 unauthorized response has a 4xx status code
func (o *UpdateRepository10Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository10 unauthorized response has a 5xx status code
func (o *UpdateRepository10Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository10 unauthorized response a status code equal to that given
func (o *UpdateRepository10Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update repository10 unauthorized response
func (o *UpdateRepository10Unauthorized) Code() int {
	return 401
}

func (o *UpdateRepository10Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/npm/proxy/{repositoryName}][%d] updateRepository10Unauthorized ", 401)
}

func (o *UpdateRepository10Unauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/npm/proxy/{repositoryName}][%d] updateRepository10Unauthorized ", 401)
}

func (o *UpdateRepository10Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository10Forbidden creates a UpdateRepository10Forbidden with default headers values
func NewUpdateRepository10Forbidden() *UpdateRepository10Forbidden {
	return &UpdateRepository10Forbidden{}
}

/*
UpdateRepository10Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateRepository10Forbidden struct {
}

// IsSuccess returns true when this update repository10 forbidden response has a 2xx status code
func (o *UpdateRepository10Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository10 forbidden response has a 3xx status code
func (o *UpdateRepository10Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository10 forbidden response has a 4xx status code
func (o *UpdateRepository10Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository10 forbidden response has a 5xx status code
func (o *UpdateRepository10Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository10 forbidden response a status code equal to that given
func (o *UpdateRepository10Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update repository10 forbidden response
func (o *UpdateRepository10Forbidden) Code() int {
	return 403
}

func (o *UpdateRepository10Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/npm/proxy/{repositoryName}][%d] updateRepository10Forbidden ", 403)
}

func (o *UpdateRepository10Forbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/npm/proxy/{repositoryName}][%d] updateRepository10Forbidden ", 403)
}

func (o *UpdateRepository10Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
