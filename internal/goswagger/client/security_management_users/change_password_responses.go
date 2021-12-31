// Code generated by go-swagger; DO NOT EDIT.

package security_management_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ChangePasswordReader is a Reader for the ChangePassword structure.
type ChangePasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangePasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewChangePasswordBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewChangePasswordForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewChangePasswordNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewChangePasswordBadRequest creates a ChangePasswordBadRequest with default headers values
func NewChangePasswordBadRequest() *ChangePasswordBadRequest {
	return &ChangePasswordBadRequest{}
}

/* ChangePasswordBadRequest describes a response with status code 400, with default header values.

Password was not supplied in the body of the request
*/
type ChangePasswordBadRequest struct {
}

func (o *ChangePasswordBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/security/users/{userId}/change-password][%d] changePasswordBadRequest ", 400)
}

func (o *ChangePasswordBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangePasswordForbidden creates a ChangePasswordForbidden with default headers values
func NewChangePasswordForbidden() *ChangePasswordForbidden {
	return &ChangePasswordForbidden{}
}

/* ChangePasswordForbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the operation.
*/
type ChangePasswordForbidden struct {
}

func (o *ChangePasswordForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/security/users/{userId}/change-password][%d] changePasswordForbidden ", 403)
}

func (o *ChangePasswordForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangePasswordNotFound creates a ChangePasswordNotFound with default headers values
func NewChangePasswordNotFound() *ChangePasswordNotFound {
	return &ChangePasswordNotFound{}
}

/* ChangePasswordNotFound describes a response with status code 404, with default header values.

User not found in the system.
*/
type ChangePasswordNotFound struct {
}

func (o *ChangePasswordNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/security/users/{userId}/change-password][%d] changePasswordNotFound ", 404)
}

func (o *ChangePasswordNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}