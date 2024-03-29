// Code generated by go-swagger; DO NOT EDIT.

package security_management_privileges

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdatePrivilege2Reader is a Reader for the UpdatePrivilege2 structure.
type UpdatePrivilege2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePrivilege2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewUpdatePrivilege2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdatePrivilege2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdatePrivilege2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdatePrivilege2BadRequest creates a UpdatePrivilege2BadRequest with default headers values
func NewUpdatePrivilege2BadRequest() *UpdatePrivilege2BadRequest {
	return &UpdatePrivilege2BadRequest{}
}

/*
UpdatePrivilege2BadRequest describes a response with status code 400, with default header values.

Privilege object not configured properly.
*/
type UpdatePrivilege2BadRequest struct {
}

// IsSuccess returns true when this update privilege2 bad request response has a 2xx status code
func (o *UpdatePrivilege2BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update privilege2 bad request response has a 3xx status code
func (o *UpdatePrivilege2BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update privilege2 bad request response has a 4xx status code
func (o *UpdatePrivilege2BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update privilege2 bad request response has a 5xx status code
func (o *UpdatePrivilege2BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update privilege2 bad request response a status code equal to that given
func (o *UpdatePrivilege2BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update privilege2 bad request response
func (o *UpdatePrivilege2BadRequest) Code() int {
	return 400
}

func (o *UpdatePrivilege2BadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/security/privileges/repository-view/{privilegeName}][%d] updatePrivilege2BadRequest ", 400)
}

func (o *UpdatePrivilege2BadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/security/privileges/repository-view/{privilegeName}][%d] updatePrivilege2BadRequest ", 400)
}

func (o *UpdatePrivilege2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePrivilege2Forbidden creates a UpdatePrivilege2Forbidden with default headers values
func NewUpdatePrivilege2Forbidden() *UpdatePrivilege2Forbidden {
	return &UpdatePrivilege2Forbidden{}
}

/*
UpdatePrivilege2Forbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the operation.
*/
type UpdatePrivilege2Forbidden struct {
}

// IsSuccess returns true when this update privilege2 forbidden response has a 2xx status code
func (o *UpdatePrivilege2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update privilege2 forbidden response has a 3xx status code
func (o *UpdatePrivilege2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update privilege2 forbidden response has a 4xx status code
func (o *UpdatePrivilege2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update privilege2 forbidden response has a 5xx status code
func (o *UpdatePrivilege2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update privilege2 forbidden response a status code equal to that given
func (o *UpdatePrivilege2Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update privilege2 forbidden response
func (o *UpdatePrivilege2Forbidden) Code() int {
	return 403
}

func (o *UpdatePrivilege2Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/security/privileges/repository-view/{privilegeName}][%d] updatePrivilege2Forbidden ", 403)
}

func (o *UpdatePrivilege2Forbidden) String() string {
	return fmt.Sprintf("[PUT /v1/security/privileges/repository-view/{privilegeName}][%d] updatePrivilege2Forbidden ", 403)
}

func (o *UpdatePrivilege2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePrivilege2NotFound creates a UpdatePrivilege2NotFound with default headers values
func NewUpdatePrivilege2NotFound() *UpdatePrivilege2NotFound {
	return &UpdatePrivilege2NotFound{}
}

/*
UpdatePrivilege2NotFound describes a response with status code 404, with default header values.

Privilege not found in the system.
*/
type UpdatePrivilege2NotFound struct {
}

// IsSuccess returns true when this update privilege2 not found response has a 2xx status code
func (o *UpdatePrivilege2NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update privilege2 not found response has a 3xx status code
func (o *UpdatePrivilege2NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update privilege2 not found response has a 4xx status code
func (o *UpdatePrivilege2NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update privilege2 not found response has a 5xx status code
func (o *UpdatePrivilege2NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update privilege2 not found response a status code equal to that given
func (o *UpdatePrivilege2NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update privilege2 not found response
func (o *UpdatePrivilege2NotFound) Code() int {
	return 404
}

func (o *UpdatePrivilege2NotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/security/privileges/repository-view/{privilegeName}][%d] updatePrivilege2NotFound ", 404)
}

func (o *UpdatePrivilege2NotFound) String() string {
	return fmt.Sprintf("[PUT /v1/security/privileges/repository-view/{privilegeName}][%d] updatePrivilege2NotFound ", 404)
}

func (o *UpdatePrivilege2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
