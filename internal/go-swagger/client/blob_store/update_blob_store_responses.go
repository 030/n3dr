// Code generated by go-swagger; DO NOT EDIT.

package blob_store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateBlobStoreReader is a Reader for the UpdateBlobStore structure.
type UpdateBlobStoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBlobStoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateBlobStoreNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateBlobStoreBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateBlobStoreUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateBlobStoreForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateBlobStoreNoContent creates a UpdateBlobStoreNoContent with default headers values
func NewUpdateBlobStoreNoContent() *UpdateBlobStoreNoContent {
	return &UpdateBlobStoreNoContent{}
}

/* UpdateBlobStoreNoContent describes a response with status code 204, with default header values.

S3 blob store updated
*/
type UpdateBlobStoreNoContent struct {
}

func (o *UpdateBlobStoreNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/blobstores/s3/{name}][%d] updateBlobStoreNoContent ", 204)
}

func (o *UpdateBlobStoreNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateBlobStoreBadRequest creates a UpdateBlobStoreBadRequest with default headers values
func NewUpdateBlobStoreBadRequest() *UpdateBlobStoreBadRequest {
	return &UpdateBlobStoreBadRequest{}
}

/* UpdateBlobStoreBadRequest describes a response with status code 400, with default header values.

Specified S3 blob store doesn't exist
*/
type UpdateBlobStoreBadRequest struct {
}

func (o *UpdateBlobStoreBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/blobstores/s3/{name}][%d] updateBlobStoreBadRequest ", 400)
}

func (o *UpdateBlobStoreBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateBlobStoreUnauthorized creates a UpdateBlobStoreUnauthorized with default headers values
func NewUpdateBlobStoreUnauthorized() *UpdateBlobStoreUnauthorized {
	return &UpdateBlobStoreUnauthorized{}
}

/* UpdateBlobStoreUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateBlobStoreUnauthorized struct {
}

func (o *UpdateBlobStoreUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/blobstores/s3/{name}][%d] updateBlobStoreUnauthorized ", 401)
}

func (o *UpdateBlobStoreUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateBlobStoreForbidden creates a UpdateBlobStoreForbidden with default headers values
func NewUpdateBlobStoreForbidden() *UpdateBlobStoreForbidden {
	return &UpdateBlobStoreForbidden{}
}

/* UpdateBlobStoreForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateBlobStoreForbidden struct {
}

func (o *UpdateBlobStoreForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/blobstores/s3/{name}][%d] updateBlobStoreForbidden ", 403)
}

func (o *UpdateBlobStoreForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
