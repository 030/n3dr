// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/030/n3dr/internal/goswagger/models"
)

// NewCreateRepository21Params creates a new CreateRepository21Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRepository21Params() *CreateRepository21Params {
	return &CreateRepository21Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRepository21ParamsWithTimeout creates a new CreateRepository21Params object
// with the ability to set a timeout on a request.
func NewCreateRepository21ParamsWithTimeout(timeout time.Duration) *CreateRepository21Params {
	return &CreateRepository21Params{
		timeout: timeout,
	}
}

// NewCreateRepository21ParamsWithContext creates a new CreateRepository21Params object
// with the ability to set a context for a request.
func NewCreateRepository21ParamsWithContext(ctx context.Context) *CreateRepository21Params {
	return &CreateRepository21Params{
		Context: ctx,
	}
}

// NewCreateRepository21ParamsWithHTTPClient creates a new CreateRepository21Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRepository21ParamsWithHTTPClient(client *http.Client) *CreateRepository21Params {
	return &CreateRepository21Params{
		HTTPClient: client,
	}
}

/* CreateRepository21Params contains all the parameters to send to the API endpoint
   for the create repository 21 operation.

   Typically these are written to a http.Request.
*/
type CreateRepository21Params struct {

	// Body.
	Body *models.YumHostedRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create repository 21 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository21Params) WithDefaults() *CreateRepository21Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create repository 21 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository21Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create repository 21 params
func (o *CreateRepository21Params) WithTimeout(timeout time.Duration) *CreateRepository21Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create repository 21 params
func (o *CreateRepository21Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create repository 21 params
func (o *CreateRepository21Params) WithContext(ctx context.Context) *CreateRepository21Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create repository 21 params
func (o *CreateRepository21Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create repository 21 params
func (o *CreateRepository21Params) WithHTTPClient(client *http.Client) *CreateRepository21Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create repository 21 params
func (o *CreateRepository21Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create repository 21 params
func (o *CreateRepository21Params) WithBody(body *models.YumHostedRepositoryAPIRequest) *CreateRepository21Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create repository 21 params
func (o *CreateRepository21Params) SetBody(body *models.YumHostedRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRepository21Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}