// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateRepository19Reader is a Reader for the CreateRepository19 structure.
type CreateRepository19Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepository19Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRepository19Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRepository19Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRepository19Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRepository19Created creates a CreateRepository19Created with default headers values
func NewCreateRepository19Created() *CreateRepository19Created {
	return &CreateRepository19Created{}
}

/*
CreateRepository19Created describes a response with status code 201, with default header values.

Repository created
*/
type CreateRepository19Created struct {
}

// IsSuccess returns true when this create repository19 created response has a 2xx status code
func (o *CreateRepository19Created) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create repository19 created response has a 3xx status code
func (o *CreateRepository19Created) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository19 created response has a 4xx status code
func (o *CreateRepository19Created) IsClientError() bool {
	return false
}

// IsServerError returns true when this create repository19 created response has a 5xx status code
func (o *CreateRepository19Created) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository19 created response a status code equal to that given
func (o *CreateRepository19Created) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create repository19 created response
func (o *CreateRepository19Created) Code() int {
	return 201
}

func (o *CreateRepository19Created) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/docker/proxy][%d] createRepository19Created ", 201)
}

func (o *CreateRepository19Created) String() string {
	return fmt.Sprintf("[POST /v1/repositories/docker/proxy][%d] createRepository19Created ", 201)
}

func (o *CreateRepository19Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository19Unauthorized creates a CreateRepository19Unauthorized with default headers values
func NewCreateRepository19Unauthorized() *CreateRepository19Unauthorized {
	return &CreateRepository19Unauthorized{}
}

/*
CreateRepository19Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateRepository19Unauthorized struct {
}

// IsSuccess returns true when this create repository19 unauthorized response has a 2xx status code
func (o *CreateRepository19Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository19 unauthorized response has a 3xx status code
func (o *CreateRepository19Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository19 unauthorized response has a 4xx status code
func (o *CreateRepository19Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository19 unauthorized response has a 5xx status code
func (o *CreateRepository19Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository19 unauthorized response a status code equal to that given
func (o *CreateRepository19Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create repository19 unauthorized response
func (o *CreateRepository19Unauthorized) Code() int {
	return 401
}

func (o *CreateRepository19Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/docker/proxy][%d] createRepository19Unauthorized ", 401)
}

func (o *CreateRepository19Unauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/docker/proxy][%d] createRepository19Unauthorized ", 401)
}

func (o *CreateRepository19Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository19Forbidden creates a CreateRepository19Forbidden with default headers values
func NewCreateRepository19Forbidden() *CreateRepository19Forbidden {
	return &CreateRepository19Forbidden{}
}

/*
CreateRepository19Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateRepository19Forbidden struct {
}

// IsSuccess returns true when this create repository19 forbidden response has a 2xx status code
func (o *CreateRepository19Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository19 forbidden response has a 3xx status code
func (o *CreateRepository19Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository19 forbidden response has a 4xx status code
func (o *CreateRepository19Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository19 forbidden response has a 5xx status code
func (o *CreateRepository19Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository19 forbidden response a status code equal to that given
func (o *CreateRepository19Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create repository19 forbidden response
func (o *CreateRepository19Forbidden) Code() int {
	return 403
}

func (o *CreateRepository19Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/docker/proxy][%d] createRepository19Forbidden ", 403)
}

func (o *CreateRepository19Forbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/docker/proxy][%d] createRepository19Forbidden ", 403)
}

func (o *CreateRepository19Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
