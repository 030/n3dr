// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateRepository18Reader is a Reader for the UpdateRepository18 structure.
type UpdateRepository18Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepository18Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRepository18NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRepository18Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRepository18Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRepository18NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepository18NoContent creates a UpdateRepository18NoContent with default headers values
func NewUpdateRepository18NoContent() *UpdateRepository18NoContent {
	return &UpdateRepository18NoContent{}
}

/*
UpdateRepository18NoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateRepository18NoContent struct {
}

// IsSuccess returns true when this update repository18 no content response has a 2xx status code
func (o *UpdateRepository18NoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update repository18 no content response has a 3xx status code
func (o *UpdateRepository18NoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository18 no content response has a 4xx status code
func (o *UpdateRepository18NoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update repository18 no content response has a 5xx status code
func (o *UpdateRepository18NoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository18 no content response a status code equal to that given
func (o *UpdateRepository18NoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update repository18 no content response
func (o *UpdateRepository18NoContent) Code() int {
	return 204
}

func (o *UpdateRepository18NoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/hosted/{repositoryName}][%d] updateRepository18NoContent ", 204)
}

func (o *UpdateRepository18NoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/hosted/{repositoryName}][%d] updateRepository18NoContent ", 204)
}

func (o *UpdateRepository18NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository18Unauthorized creates a UpdateRepository18Unauthorized with default headers values
func NewUpdateRepository18Unauthorized() *UpdateRepository18Unauthorized {
	return &UpdateRepository18Unauthorized{}
}

/*
UpdateRepository18Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateRepository18Unauthorized struct {
}

// IsSuccess returns true when this update repository18 unauthorized response has a 2xx status code
func (o *UpdateRepository18Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository18 unauthorized response has a 3xx status code
func (o *UpdateRepository18Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository18 unauthorized response has a 4xx status code
func (o *UpdateRepository18Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository18 unauthorized response has a 5xx status code
func (o *UpdateRepository18Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository18 unauthorized response a status code equal to that given
func (o *UpdateRepository18Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update repository18 unauthorized response
func (o *UpdateRepository18Unauthorized) Code() int {
	return 401
}

func (o *UpdateRepository18Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/hosted/{repositoryName}][%d] updateRepository18Unauthorized ", 401)
}

func (o *UpdateRepository18Unauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/hosted/{repositoryName}][%d] updateRepository18Unauthorized ", 401)
}

func (o *UpdateRepository18Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository18Forbidden creates a UpdateRepository18Forbidden with default headers values
func NewUpdateRepository18Forbidden() *UpdateRepository18Forbidden {
	return &UpdateRepository18Forbidden{}
}

/*
UpdateRepository18Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateRepository18Forbidden struct {
}

// IsSuccess returns true when this update repository18 forbidden response has a 2xx status code
func (o *UpdateRepository18Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository18 forbidden response has a 3xx status code
func (o *UpdateRepository18Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository18 forbidden response has a 4xx status code
func (o *UpdateRepository18Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository18 forbidden response has a 5xx status code
func (o *UpdateRepository18Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository18 forbidden response a status code equal to that given
func (o *UpdateRepository18Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update repository18 forbidden response
func (o *UpdateRepository18Forbidden) Code() int {
	return 403
}

func (o *UpdateRepository18Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/hosted/{repositoryName}][%d] updateRepository18Forbidden ", 403)
}

func (o *UpdateRepository18Forbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/hosted/{repositoryName}][%d] updateRepository18Forbidden ", 403)
}

func (o *UpdateRepository18Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository18NotFound creates a UpdateRepository18NotFound with default headers values
func NewUpdateRepository18NotFound() *UpdateRepository18NotFound {
	return &UpdateRepository18NotFound{}
}

/*
UpdateRepository18NotFound describes a response with status code 404, with default header values.

Repository not found
*/
type UpdateRepository18NotFound struct {
}

// IsSuccess returns true when this update repository18 not found response has a 2xx status code
func (o *UpdateRepository18NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository18 not found response has a 3xx status code
func (o *UpdateRepository18NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository18 not found response has a 4xx status code
func (o *UpdateRepository18NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository18 not found response has a 5xx status code
func (o *UpdateRepository18NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository18 not found response a status code equal to that given
func (o *UpdateRepository18NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update repository18 not found response
func (o *UpdateRepository18NotFound) Code() int {
	return 404
}

func (o *UpdateRepository18NotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/hosted/{repositoryName}][%d] updateRepository18NotFound ", 404)
}

func (o *UpdateRepository18NotFound) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/hosted/{repositoryName}][%d] updateRepository18NotFound ", 404)
}

func (o *UpdateRepository18NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
