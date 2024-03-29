// Code generated by go-swagger; DO NOT EDIT.

package lifecycle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// BounceReader is a Reader for the Bounce structure.
type BounceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BounceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewBounceDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewBounceDefault creates a BounceDefault with default headers values
func NewBounceDefault(code int) *BounceDefault {
	return &BounceDefault{
		_statusCode: code,
	}
}

/*
BounceDefault describes a response with status code -1, with default header values.

successful operation
*/
type BounceDefault struct {
	_statusCode int
}

// IsSuccess returns true when this bounce default response has a 2xx status code
func (o *BounceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this bounce default response has a 3xx status code
func (o *BounceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this bounce default response has a 4xx status code
func (o *BounceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this bounce default response has a 5xx status code
func (o *BounceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this bounce default response a status code equal to that given
func (o *BounceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the bounce default response
func (o *BounceDefault) Code() int {
	return o._statusCode
}

func (o *BounceDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/lifecycle/bounce][%d] bounce default ", o._statusCode)
}

func (o *BounceDefault) String() string {
	return fmt.Sprintf("[PUT /v1/lifecycle/bounce][%d] bounce default ", o._statusCode)
}

func (o *BounceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
