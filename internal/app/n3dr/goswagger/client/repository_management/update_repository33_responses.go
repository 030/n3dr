// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateRepository33Reader is a Reader for the UpdateRepository33 structure.
type UpdateRepository33Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepository33Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRepository33NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRepository33Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRepository33Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRepository33NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepository33NoContent creates a UpdateRepository33NoContent with default headers values
func NewUpdateRepository33NoContent() *UpdateRepository33NoContent {
	return &UpdateRepository33NoContent{}
}

/*
UpdateRepository33NoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateRepository33NoContent struct {
}

// IsSuccess returns true when this update repository33 no content response has a 2xx status code
func (o *UpdateRepository33NoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update repository33 no content response has a 3xx status code
func (o *UpdateRepository33NoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository33 no content response has a 4xx status code
func (o *UpdateRepository33NoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update repository33 no content response has a 5xx status code
func (o *UpdateRepository33NoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository33 no content response a status code equal to that given
func (o *UpdateRepository33NoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update repository33 no content response
func (o *UpdateRepository33NoContent) Code() int {
	return 204
}

func (o *UpdateRepository33NoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/r/proxy/{repositoryName}][%d] updateRepository33NoContent ", 204)
}

func (o *UpdateRepository33NoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/r/proxy/{repositoryName}][%d] updateRepository33NoContent ", 204)
}

func (o *UpdateRepository33NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository33Unauthorized creates a UpdateRepository33Unauthorized with default headers values
func NewUpdateRepository33Unauthorized() *UpdateRepository33Unauthorized {
	return &UpdateRepository33Unauthorized{}
}

/*
UpdateRepository33Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateRepository33Unauthorized struct {
}

// IsSuccess returns true when this update repository33 unauthorized response has a 2xx status code
func (o *UpdateRepository33Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository33 unauthorized response has a 3xx status code
func (o *UpdateRepository33Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository33 unauthorized response has a 4xx status code
func (o *UpdateRepository33Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository33 unauthorized response has a 5xx status code
func (o *UpdateRepository33Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository33 unauthorized response a status code equal to that given
func (o *UpdateRepository33Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update repository33 unauthorized response
func (o *UpdateRepository33Unauthorized) Code() int {
	return 401
}

func (o *UpdateRepository33Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/r/proxy/{repositoryName}][%d] updateRepository33Unauthorized ", 401)
}

func (o *UpdateRepository33Unauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/r/proxy/{repositoryName}][%d] updateRepository33Unauthorized ", 401)
}

func (o *UpdateRepository33Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository33Forbidden creates a UpdateRepository33Forbidden with default headers values
func NewUpdateRepository33Forbidden() *UpdateRepository33Forbidden {
	return &UpdateRepository33Forbidden{}
}

/*
UpdateRepository33Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateRepository33Forbidden struct {
}

// IsSuccess returns true when this update repository33 forbidden response has a 2xx status code
func (o *UpdateRepository33Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository33 forbidden response has a 3xx status code
func (o *UpdateRepository33Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository33 forbidden response has a 4xx status code
func (o *UpdateRepository33Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository33 forbidden response has a 5xx status code
func (o *UpdateRepository33Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository33 forbidden response a status code equal to that given
func (o *UpdateRepository33Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update repository33 forbidden response
func (o *UpdateRepository33Forbidden) Code() int {
	return 403
}

func (o *UpdateRepository33Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/r/proxy/{repositoryName}][%d] updateRepository33Forbidden ", 403)
}

func (o *UpdateRepository33Forbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/r/proxy/{repositoryName}][%d] updateRepository33Forbidden ", 403)
}

func (o *UpdateRepository33Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository33NotFound creates a UpdateRepository33NotFound with default headers values
func NewUpdateRepository33NotFound() *UpdateRepository33NotFound {
	return &UpdateRepository33NotFound{}
}

/*
UpdateRepository33NotFound describes a response with status code 404, with default header values.

Repository not found
*/
type UpdateRepository33NotFound struct {
}

// IsSuccess returns true when this update repository33 not found response has a 2xx status code
func (o *UpdateRepository33NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository33 not found response has a 3xx status code
func (o *UpdateRepository33NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository33 not found response has a 4xx status code
func (o *UpdateRepository33NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository33 not found response has a 5xx status code
func (o *UpdateRepository33NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository33 not found response a status code equal to that given
func (o *UpdateRepository33NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update repository33 not found response
func (o *UpdateRepository33NotFound) Code() int {
	return 404
}

func (o *UpdateRepository33NotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/r/proxy/{repositoryName}][%d] updateRepository33NotFound ", 404)
}

func (o *UpdateRepository33NotFound) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/r/proxy/{repositoryName}][%d] updateRepository33NotFound ", 404)
}

func (o *UpdateRepository33NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
