// Code generated by go-swagger; DO NOT EDIT.

package blob_store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateBlobStoreReader is a Reader for the CreateBlobStore structure.
type CreateBlobStoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBlobStoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateBlobStoreCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateBlobStoreUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateBlobStoreForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateBlobStoreCreated creates a CreateBlobStoreCreated with default headers values
func NewCreateBlobStoreCreated() *CreateBlobStoreCreated {
	return &CreateBlobStoreCreated{}
}

/*
CreateBlobStoreCreated describes a response with status code 201, with default header values.

S3 blob store created
*/
type CreateBlobStoreCreated struct {
}

// IsSuccess returns true when this create blob store created response has a 2xx status code
func (o *CreateBlobStoreCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create blob store created response has a 3xx status code
func (o *CreateBlobStoreCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create blob store created response has a 4xx status code
func (o *CreateBlobStoreCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create blob store created response has a 5xx status code
func (o *CreateBlobStoreCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create blob store created response a status code equal to that given
func (o *CreateBlobStoreCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create blob store created response
func (o *CreateBlobStoreCreated) Code() int {
	return 201
}

func (o *CreateBlobStoreCreated) Error() string {
	return fmt.Sprintf("[POST /v1/blobstores/s3][%d] createBlobStoreCreated ", 201)
}

func (o *CreateBlobStoreCreated) String() string {
	return fmt.Sprintf("[POST /v1/blobstores/s3][%d] createBlobStoreCreated ", 201)
}

func (o *CreateBlobStoreCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateBlobStoreUnauthorized creates a CreateBlobStoreUnauthorized with default headers values
func NewCreateBlobStoreUnauthorized() *CreateBlobStoreUnauthorized {
	return &CreateBlobStoreUnauthorized{}
}

/*
CreateBlobStoreUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateBlobStoreUnauthorized struct {
}

// IsSuccess returns true when this create blob store unauthorized response has a 2xx status code
func (o *CreateBlobStoreUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create blob store unauthorized response has a 3xx status code
func (o *CreateBlobStoreUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create blob store unauthorized response has a 4xx status code
func (o *CreateBlobStoreUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create blob store unauthorized response has a 5xx status code
func (o *CreateBlobStoreUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create blob store unauthorized response a status code equal to that given
func (o *CreateBlobStoreUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create blob store unauthorized response
func (o *CreateBlobStoreUnauthorized) Code() int {
	return 401
}

func (o *CreateBlobStoreUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/blobstores/s3][%d] createBlobStoreUnauthorized ", 401)
}

func (o *CreateBlobStoreUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/blobstores/s3][%d] createBlobStoreUnauthorized ", 401)
}

func (o *CreateBlobStoreUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateBlobStoreForbidden creates a CreateBlobStoreForbidden with default headers values
func NewCreateBlobStoreForbidden() *CreateBlobStoreForbidden {
	return &CreateBlobStoreForbidden{}
}

/*
CreateBlobStoreForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateBlobStoreForbidden struct {
}

// IsSuccess returns true when this create blob store forbidden response has a 2xx status code
func (o *CreateBlobStoreForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create blob store forbidden response has a 3xx status code
func (o *CreateBlobStoreForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create blob store forbidden response has a 4xx status code
func (o *CreateBlobStoreForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create blob store forbidden response has a 5xx status code
func (o *CreateBlobStoreForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create blob store forbidden response a status code equal to that given
func (o *CreateBlobStoreForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create blob store forbidden response
func (o *CreateBlobStoreForbidden) Code() int {
	return 403
}

func (o *CreateBlobStoreForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/blobstores/s3][%d] createBlobStoreForbidden ", 403)
}

func (o *CreateBlobStoreForbidden) String() string {
	return fmt.Sprintf("[POST /v1/blobstores/s3][%d] createBlobStoreForbidden ", 403)
}

func (o *CreateBlobStoreForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
