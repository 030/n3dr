// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateRepository15Reader is a Reader for the CreateRepository15 structure.
type CreateRepository15Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepository15Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRepository15Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRepository15Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRepository15Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRepository15Created creates a CreateRepository15Created with default headers values
func NewCreateRepository15Created() *CreateRepository15Created {
	return &CreateRepository15Created{}
}

/*
CreateRepository15Created describes a response with status code 201, with default header values.

Repository created
*/
type CreateRepository15Created struct {
}

// IsSuccess returns true when this create repository15 created response has a 2xx status code
func (o *CreateRepository15Created) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create repository15 created response has a 3xx status code
func (o *CreateRepository15Created) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository15 created response has a 4xx status code
func (o *CreateRepository15Created) IsClientError() bool {
	return false
}

// IsServerError returns true when this create repository15 created response has a 5xx status code
func (o *CreateRepository15Created) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository15 created response a status code equal to that given
func (o *CreateRepository15Created) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create repository15 created response
func (o *CreateRepository15Created) Code() int {
	return 201
}

func (o *CreateRepository15Created) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/rubygems/hosted][%d] createRepository15Created ", 201)
}

func (o *CreateRepository15Created) String() string {
	return fmt.Sprintf("[POST /v1/repositories/rubygems/hosted][%d] createRepository15Created ", 201)
}

func (o *CreateRepository15Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository15Unauthorized creates a CreateRepository15Unauthorized with default headers values
func NewCreateRepository15Unauthorized() *CreateRepository15Unauthorized {
	return &CreateRepository15Unauthorized{}
}

/*
CreateRepository15Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateRepository15Unauthorized struct {
}

// IsSuccess returns true when this create repository15 unauthorized response has a 2xx status code
func (o *CreateRepository15Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository15 unauthorized response has a 3xx status code
func (o *CreateRepository15Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository15 unauthorized response has a 4xx status code
func (o *CreateRepository15Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository15 unauthorized response has a 5xx status code
func (o *CreateRepository15Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository15 unauthorized response a status code equal to that given
func (o *CreateRepository15Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create repository15 unauthorized response
func (o *CreateRepository15Unauthorized) Code() int {
	return 401
}

func (o *CreateRepository15Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/rubygems/hosted][%d] createRepository15Unauthorized ", 401)
}

func (o *CreateRepository15Unauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/rubygems/hosted][%d] createRepository15Unauthorized ", 401)
}

func (o *CreateRepository15Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository15Forbidden creates a CreateRepository15Forbidden with default headers values
func NewCreateRepository15Forbidden() *CreateRepository15Forbidden {
	return &CreateRepository15Forbidden{}
}

/*
CreateRepository15Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateRepository15Forbidden struct {
}

// IsSuccess returns true when this create repository15 forbidden response has a 2xx status code
func (o *CreateRepository15Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository15 forbidden response has a 3xx status code
func (o *CreateRepository15Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository15 forbidden response has a 4xx status code
func (o *CreateRepository15Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository15 forbidden response has a 5xx status code
func (o *CreateRepository15Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository15 forbidden response a status code equal to that given
func (o *CreateRepository15Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create repository15 forbidden response
func (o *CreateRepository15Forbidden) Code() int {
	return 403
}

func (o *CreateRepository15Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/rubygems/hosted][%d] createRepository15Forbidden ", 403)
}

func (o *CreateRepository15Forbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/rubygems/hosted][%d] createRepository15Forbidden ", 403)
}

func (o *CreateRepository15Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
