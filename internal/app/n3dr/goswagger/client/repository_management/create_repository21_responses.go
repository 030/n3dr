// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateRepository21Reader is a Reader for the CreateRepository21 structure.
type CreateRepository21Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepository21Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRepository21Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRepository21Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRepository21Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRepository21Created creates a CreateRepository21Created with default headers values
func NewCreateRepository21Created() *CreateRepository21Created {
	return &CreateRepository21Created{}
}

/*
CreateRepository21Created describes a response with status code 201, with default header values.

Repository created
*/
type CreateRepository21Created struct {
}

// IsSuccess returns true when this create repository21 created response has a 2xx status code
func (o *CreateRepository21Created) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create repository21 created response has a 3xx status code
func (o *CreateRepository21Created) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository21 created response has a 4xx status code
func (o *CreateRepository21Created) IsClientError() bool {
	return false
}

// IsServerError returns true when this create repository21 created response has a 5xx status code
func (o *CreateRepository21Created) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository21 created response a status code equal to that given
func (o *CreateRepository21Created) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create repository21 created response
func (o *CreateRepository21Created) Code() int {
	return 201
}

func (o *CreateRepository21Created) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/yum/hosted][%d] createRepository21Created ", 201)
}

func (o *CreateRepository21Created) String() string {
	return fmt.Sprintf("[POST /v1/repositories/yum/hosted][%d] createRepository21Created ", 201)
}

func (o *CreateRepository21Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository21Unauthorized creates a CreateRepository21Unauthorized with default headers values
func NewCreateRepository21Unauthorized() *CreateRepository21Unauthorized {
	return &CreateRepository21Unauthorized{}
}

/*
CreateRepository21Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateRepository21Unauthorized struct {
}

// IsSuccess returns true when this create repository21 unauthorized response has a 2xx status code
func (o *CreateRepository21Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository21 unauthorized response has a 3xx status code
func (o *CreateRepository21Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository21 unauthorized response has a 4xx status code
func (o *CreateRepository21Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository21 unauthorized response has a 5xx status code
func (o *CreateRepository21Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository21 unauthorized response a status code equal to that given
func (o *CreateRepository21Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create repository21 unauthorized response
func (o *CreateRepository21Unauthorized) Code() int {
	return 401
}

func (o *CreateRepository21Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/yum/hosted][%d] createRepository21Unauthorized ", 401)
}

func (o *CreateRepository21Unauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/yum/hosted][%d] createRepository21Unauthorized ", 401)
}

func (o *CreateRepository21Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository21Forbidden creates a CreateRepository21Forbidden with default headers values
func NewCreateRepository21Forbidden() *CreateRepository21Forbidden {
	return &CreateRepository21Forbidden{}
}

/*
CreateRepository21Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateRepository21Forbidden struct {
}

// IsSuccess returns true when this create repository21 forbidden response has a 2xx status code
func (o *CreateRepository21Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository21 forbidden response has a 3xx status code
func (o *CreateRepository21Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository21 forbidden response has a 4xx status code
func (o *CreateRepository21Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository21 forbidden response has a 5xx status code
func (o *CreateRepository21Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository21 forbidden response a status code equal to that given
func (o *CreateRepository21Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create repository21 forbidden response
func (o *CreateRepository21Forbidden) Code() int {
	return 403
}

func (o *CreateRepository21Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/yum/hosted][%d] createRepository21Forbidden ", 403)
}

func (o *CreateRepository21Forbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/yum/hosted][%d] createRepository21Forbidden ", 403)
}

func (o *CreateRepository21Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
