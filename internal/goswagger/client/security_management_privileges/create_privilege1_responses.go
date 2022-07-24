// Code generated by go-swagger; DO NOT EDIT.

package security_management_privileges

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreatePrivilege1Reader is a Reader for the CreatePrivilege1 structure.
type CreatePrivilege1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePrivilege1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewCreatePrivilege1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreatePrivilege1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreatePrivilege1BadRequest creates a CreatePrivilege1BadRequest with default headers values
func NewCreatePrivilege1BadRequest() *CreatePrivilege1BadRequest {
	return &CreatePrivilege1BadRequest{}
}

/* CreatePrivilege1BadRequest describes a response with status code 400, with default header values.

Privilege object not configured properly.
*/
type CreatePrivilege1BadRequest struct {
}

func (o *CreatePrivilege1BadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/security/privileges/wildcard][%d] createPrivilege1BadRequest ", 400)
}

func (o *CreatePrivilege1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreatePrivilege1Forbidden creates a CreatePrivilege1Forbidden with default headers values
func NewCreatePrivilege1Forbidden() *CreatePrivilege1Forbidden {
	return &CreatePrivilege1Forbidden{}
}

/* CreatePrivilege1Forbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the operation.
*/
type CreatePrivilege1Forbidden struct {
}

func (o *CreatePrivilege1Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/security/privileges/wildcard][%d] createPrivilege1Forbidden ", 403)
}

func (o *CreatePrivilege1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
