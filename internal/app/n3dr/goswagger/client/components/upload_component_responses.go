// Code generated by go-swagger; DO NOT EDIT.

package components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UploadComponentReader is a Reader for the UploadComponent structure.
type UploadComponentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadComponentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 403:
		result := NewUploadComponentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUploadComponentUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUploadComponentForbidden creates a UploadComponentForbidden with default headers values
func NewUploadComponentForbidden() *UploadComponentForbidden {
	return &UploadComponentForbidden{}
}

/*
UploadComponentForbidden describes a response with status code 403, with default header values.

Insufficient permissions to upload a component
*/
type UploadComponentForbidden struct {
}

// IsSuccess returns true when this upload component forbidden response has a 2xx status code
func (o *UploadComponentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload component forbidden response has a 3xx status code
func (o *UploadComponentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload component forbidden response has a 4xx status code
func (o *UploadComponentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload component forbidden response has a 5xx status code
func (o *UploadComponentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this upload component forbidden response a status code equal to that given
func (o *UploadComponentForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the upload component forbidden response
func (o *UploadComponentForbidden) Code() int {
	return 403
}

func (o *UploadComponentForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/components][%d] uploadComponentForbidden ", 403)
}

func (o *UploadComponentForbidden) String() string {
	return fmt.Sprintf("[POST /v1/components][%d] uploadComponentForbidden ", 403)
}

func (o *UploadComponentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUploadComponentUnprocessableEntity creates a UploadComponentUnprocessableEntity with default headers values
func NewUploadComponentUnprocessableEntity() *UploadComponentUnprocessableEntity {
	return &UploadComponentUnprocessableEntity{}
}

/*
UploadComponentUnprocessableEntity describes a response with status code 422, with default header values.

Parameter 'repository' is required
*/
type UploadComponentUnprocessableEntity struct {
}

// IsSuccess returns true when this upload component unprocessable entity response has a 2xx status code
func (o *UploadComponentUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload component unprocessable entity response has a 3xx status code
func (o *UploadComponentUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload component unprocessable entity response has a 4xx status code
func (o *UploadComponentUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload component unprocessable entity response has a 5xx status code
func (o *UploadComponentUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this upload component unprocessable entity response a status code equal to that given
func (o *UploadComponentUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the upload component unprocessable entity response
func (o *UploadComponentUnprocessableEntity) Code() int {
	return 422
}

func (o *UploadComponentUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /v1/components][%d] uploadComponentUnprocessableEntity ", 422)
}

func (o *UploadComponentUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /v1/components][%d] uploadComponentUnprocessableEntity ", 422)
}

func (o *UploadComponentUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
