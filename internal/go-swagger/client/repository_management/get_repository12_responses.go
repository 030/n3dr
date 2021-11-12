// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/030/n3dr/internal/go-swagger/models"
)

// GetRepository12Reader is a Reader for the GetRepository12 structure.
type GetRepository12Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepository12Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepository12OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepository12OK creates a GetRepository12OK with default headers values
func NewGetRepository12OK() *GetRepository12OK {
	return &GetRepository12OK{}
}

/* GetRepository12OK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepository12OK struct {
	Payload *models.SimpleAPIHostedRepository
}

func (o *GetRepository12OK) Error() string {
	return fmt.Sprintf("[GET /v1/repositories/nuget/hosted/{repositoryName}][%d] getRepository12OK  %+v", 200, o.Payload)
}
func (o *GetRepository12OK) GetPayload() *models.SimpleAPIHostedRepository {
	return o.Payload
}

func (o *GetRepository12OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleAPIHostedRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
