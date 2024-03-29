// Code generated by go-swagger; DO NOT EDIT.

package security_management_privileges

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreatePrivilege4Reader is a Reader for the CreatePrivilege4 structure.
type CreatePrivilege4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePrivilege4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewCreatePrivilege4BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreatePrivilege4Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreatePrivilege4BadRequest creates a CreatePrivilege4BadRequest with default headers values
func NewCreatePrivilege4BadRequest() *CreatePrivilege4BadRequest {
	return &CreatePrivilege4BadRequest{}
}

/*
CreatePrivilege4BadRequest describes a response with status code 400, with default header values.

Privilege object not configured properly.
*/
type CreatePrivilege4BadRequest struct {
}

// IsSuccess returns true when this create privilege4 bad request response has a 2xx status code
func (o *CreatePrivilege4BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create privilege4 bad request response has a 3xx status code
func (o *CreatePrivilege4BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create privilege4 bad request response has a 4xx status code
func (o *CreatePrivilege4BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create privilege4 bad request response has a 5xx status code
func (o *CreatePrivilege4BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create privilege4 bad request response a status code equal to that given
func (o *CreatePrivilege4BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create privilege4 bad request response
func (o *CreatePrivilege4BadRequest) Code() int {
	return 400
}

func (o *CreatePrivilege4BadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/security/privileges/repository-view][%d] createPrivilege4BadRequest ", 400)
}

func (o *CreatePrivilege4BadRequest) String() string {
	return fmt.Sprintf("[POST /v1/security/privileges/repository-view][%d] createPrivilege4BadRequest ", 400)
}

func (o *CreatePrivilege4BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreatePrivilege4Forbidden creates a CreatePrivilege4Forbidden with default headers values
func NewCreatePrivilege4Forbidden() *CreatePrivilege4Forbidden {
	return &CreatePrivilege4Forbidden{}
}

/*
CreatePrivilege4Forbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the operation.
*/
type CreatePrivilege4Forbidden struct {
}

// IsSuccess returns true when this create privilege4 forbidden response has a 2xx status code
func (o *CreatePrivilege4Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create privilege4 forbidden response has a 3xx status code
func (o *CreatePrivilege4Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create privilege4 forbidden response has a 4xx status code
func (o *CreatePrivilege4Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create privilege4 forbidden response has a 5xx status code
func (o *CreatePrivilege4Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create privilege4 forbidden response a status code equal to that given
func (o *CreatePrivilege4Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create privilege4 forbidden response
func (o *CreatePrivilege4Forbidden) Code() int {
	return 403
}

func (o *CreatePrivilege4Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/security/privileges/repository-view][%d] createPrivilege4Forbidden ", 403)
}

func (o *CreatePrivilege4Forbidden) String() string {
	return fmt.Sprintf("[POST /v1/security/privileges/repository-view][%d] createPrivilege4Forbidden ", 403)
}

func (o *CreatePrivilege4Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
