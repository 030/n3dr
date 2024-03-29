// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DisableRepositoryHealthCheckReader is a Reader for the DisableRepositoryHealthCheck structure.
type DisableRepositoryHealthCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DisableRepositoryHealthCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDisableRepositoryHealthCheckNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDisableRepositoryHealthCheckUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDisableRepositoryHealthCheckForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDisableRepositoryHealthCheckNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDisableRepositoryHealthCheckNoContent creates a DisableRepositoryHealthCheckNoContent with default headers values
func NewDisableRepositoryHealthCheckNoContent() *DisableRepositoryHealthCheckNoContent {
	return &DisableRepositoryHealthCheckNoContent{}
}

/*
DisableRepositoryHealthCheckNoContent describes a response with status code 204, with default header values.

Repository Health Check disabled
*/
type DisableRepositoryHealthCheckNoContent struct {
}

// IsSuccess returns true when this disable repository health check no content response has a 2xx status code
func (o *DisableRepositoryHealthCheckNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this disable repository health check no content response has a 3xx status code
func (o *DisableRepositoryHealthCheckNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable repository health check no content response has a 4xx status code
func (o *DisableRepositoryHealthCheckNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this disable repository health check no content response has a 5xx status code
func (o *DisableRepositoryHealthCheckNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this disable repository health check no content response a status code equal to that given
func (o *DisableRepositoryHealthCheckNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the disable repository health check no content response
func (o *DisableRepositoryHealthCheckNoContent) Code() int {
	return 204
}

func (o *DisableRepositoryHealthCheckNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/repositories/{repositoryName}/health-check][%d] disableRepositoryHealthCheckNoContent ", 204)
}

func (o *DisableRepositoryHealthCheckNoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/repositories/{repositoryName}/health-check][%d] disableRepositoryHealthCheckNoContent ", 204)
}

func (o *DisableRepositoryHealthCheckNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisableRepositoryHealthCheckUnauthorized creates a DisableRepositoryHealthCheckUnauthorized with default headers values
func NewDisableRepositoryHealthCheckUnauthorized() *DisableRepositoryHealthCheckUnauthorized {
	return &DisableRepositoryHealthCheckUnauthorized{}
}

/*
DisableRepositoryHealthCheckUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type DisableRepositoryHealthCheckUnauthorized struct {
}

// IsSuccess returns true when this disable repository health check unauthorized response has a 2xx status code
func (o *DisableRepositoryHealthCheckUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this disable repository health check unauthorized response has a 3xx status code
func (o *DisableRepositoryHealthCheckUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable repository health check unauthorized response has a 4xx status code
func (o *DisableRepositoryHealthCheckUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this disable repository health check unauthorized response has a 5xx status code
func (o *DisableRepositoryHealthCheckUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this disable repository health check unauthorized response a status code equal to that given
func (o *DisableRepositoryHealthCheckUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the disable repository health check unauthorized response
func (o *DisableRepositoryHealthCheckUnauthorized) Code() int {
	return 401
}

func (o *DisableRepositoryHealthCheckUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/repositories/{repositoryName}/health-check][%d] disableRepositoryHealthCheckUnauthorized ", 401)
}

func (o *DisableRepositoryHealthCheckUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/repositories/{repositoryName}/health-check][%d] disableRepositoryHealthCheckUnauthorized ", 401)
}

func (o *DisableRepositoryHealthCheckUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisableRepositoryHealthCheckForbidden creates a DisableRepositoryHealthCheckForbidden with default headers values
func NewDisableRepositoryHealthCheckForbidden() *DisableRepositoryHealthCheckForbidden {
	return &DisableRepositoryHealthCheckForbidden{}
}

/*
DisableRepositoryHealthCheckForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type DisableRepositoryHealthCheckForbidden struct {
}

// IsSuccess returns true when this disable repository health check forbidden response has a 2xx status code
func (o *DisableRepositoryHealthCheckForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this disable repository health check forbidden response has a 3xx status code
func (o *DisableRepositoryHealthCheckForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable repository health check forbidden response has a 4xx status code
func (o *DisableRepositoryHealthCheckForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this disable repository health check forbidden response has a 5xx status code
func (o *DisableRepositoryHealthCheckForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this disable repository health check forbidden response a status code equal to that given
func (o *DisableRepositoryHealthCheckForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the disable repository health check forbidden response
func (o *DisableRepositoryHealthCheckForbidden) Code() int {
	return 403
}

func (o *DisableRepositoryHealthCheckForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/repositories/{repositoryName}/health-check][%d] disableRepositoryHealthCheckForbidden ", 403)
}

func (o *DisableRepositoryHealthCheckForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/repositories/{repositoryName}/health-check][%d] disableRepositoryHealthCheckForbidden ", 403)
}

func (o *DisableRepositoryHealthCheckForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisableRepositoryHealthCheckNotFound creates a DisableRepositoryHealthCheckNotFound with default headers values
func NewDisableRepositoryHealthCheckNotFound() *DisableRepositoryHealthCheckNotFound {
	return &DisableRepositoryHealthCheckNotFound{}
}

/*
DisableRepositoryHealthCheckNotFound describes a response with status code 404, with default header values.

Repository not found
*/
type DisableRepositoryHealthCheckNotFound struct {
}

// IsSuccess returns true when this disable repository health check not found response has a 2xx status code
func (o *DisableRepositoryHealthCheckNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this disable repository health check not found response has a 3xx status code
func (o *DisableRepositoryHealthCheckNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable repository health check not found response has a 4xx status code
func (o *DisableRepositoryHealthCheckNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this disable repository health check not found response has a 5xx status code
func (o *DisableRepositoryHealthCheckNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this disable repository health check not found response a status code equal to that given
func (o *DisableRepositoryHealthCheckNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the disable repository health check not found response
func (o *DisableRepositoryHealthCheckNotFound) Code() int {
	return 404
}

func (o *DisableRepositoryHealthCheckNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/repositories/{repositoryName}/health-check][%d] disableRepositoryHealthCheckNotFound ", 404)
}

func (o *DisableRepositoryHealthCheckNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/repositories/{repositoryName}/health-check][%d] disableRepositoryHealthCheckNotFound ", 404)
}

func (o *DisableRepositoryHealthCheckNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
