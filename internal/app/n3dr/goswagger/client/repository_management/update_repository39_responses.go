// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateRepository39Reader is a Reader for the UpdateRepository39 structure.
type UpdateRepository39Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepository39Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRepository39NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRepository39Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRepository39Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepository39NoContent creates a UpdateRepository39NoContent with default headers values
func NewUpdateRepository39NoContent() *UpdateRepository39NoContent {
	return &UpdateRepository39NoContent{}
}

/*
UpdateRepository39NoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateRepository39NoContent struct {
}

// IsSuccess returns true when this update repository39 no content response has a 2xx status code
func (o *UpdateRepository39NoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update repository39 no content response has a 3xx status code
func (o *UpdateRepository39NoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository39 no content response has a 4xx status code
func (o *UpdateRepository39NoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update repository39 no content response has a 5xx status code
func (o *UpdateRepository39NoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository39 no content response a status code equal to that given
func (o *UpdateRepository39NoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update repository39 no content response
func (o *UpdateRepository39NoContent) Code() int {
	return 204
}

func (o *UpdateRepository39NoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/hosted/{repositoryName}][%d] updateRepository39NoContent ", 204)
}

func (o *UpdateRepository39NoContent) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/hosted/{repositoryName}][%d] updateRepository39NoContent ", 204)
}

func (o *UpdateRepository39NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository39Unauthorized creates a UpdateRepository39Unauthorized with default headers values
func NewUpdateRepository39Unauthorized() *UpdateRepository39Unauthorized {
	return &UpdateRepository39Unauthorized{}
}

/*
UpdateRepository39Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateRepository39Unauthorized struct {
}

// IsSuccess returns true when this update repository39 unauthorized response has a 2xx status code
func (o *UpdateRepository39Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository39 unauthorized response has a 3xx status code
func (o *UpdateRepository39Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository39 unauthorized response has a 4xx status code
func (o *UpdateRepository39Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository39 unauthorized response has a 5xx status code
func (o *UpdateRepository39Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository39 unauthorized response a status code equal to that given
func (o *UpdateRepository39Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update repository39 unauthorized response
func (o *UpdateRepository39Unauthorized) Code() int {
	return 401
}

func (o *UpdateRepository39Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/hosted/{repositoryName}][%d] updateRepository39Unauthorized ", 401)
}

func (o *UpdateRepository39Unauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/hosted/{repositoryName}][%d] updateRepository39Unauthorized ", 401)
}

func (o *UpdateRepository39Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository39Forbidden creates a UpdateRepository39Forbidden with default headers values
func NewUpdateRepository39Forbidden() *UpdateRepository39Forbidden {
	return &UpdateRepository39Forbidden{}
}

/*
UpdateRepository39Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateRepository39Forbidden struct {
}

// IsSuccess returns true when this update repository39 forbidden response has a 2xx status code
func (o *UpdateRepository39Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update repository39 forbidden response has a 3xx status code
func (o *UpdateRepository39Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update repository39 forbidden response has a 4xx status code
func (o *UpdateRepository39Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update repository39 forbidden response has a 5xx status code
func (o *UpdateRepository39Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update repository39 forbidden response a status code equal to that given
func (o *UpdateRepository39Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update repository39 forbidden response
func (o *UpdateRepository39Forbidden) Code() int {
	return 403
}

func (o *UpdateRepository39Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/hosted/{repositoryName}][%d] updateRepository39Forbidden ", 403)
}

func (o *UpdateRepository39Forbidden) String() string {
	return fmt.Sprintf("[PUT /v1/repositories/bower/hosted/{repositoryName}][%d] updateRepository39Forbidden ", 403)
}

func (o *UpdateRepository39Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
