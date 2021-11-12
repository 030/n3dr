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

	"github.com/030/n3dr/internal/go-swagger/models"
)

// NewCreateRepository12Params creates a new CreateRepository12Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRepository12Params() *CreateRepository12Params {
	return &CreateRepository12Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRepository12ParamsWithTimeout creates a new CreateRepository12Params object
// with the ability to set a timeout on a request.
func NewCreateRepository12ParamsWithTimeout(timeout time.Duration) *CreateRepository12Params {
	return &CreateRepository12Params{
		timeout: timeout,
	}
}

// NewCreateRepository12ParamsWithContext creates a new CreateRepository12Params object
// with the ability to set a context for a request.
func NewCreateRepository12ParamsWithContext(ctx context.Context) *CreateRepository12Params {
	return &CreateRepository12Params{
		Context: ctx,
	}
}

// NewCreateRepository12ParamsWithHTTPClient creates a new CreateRepository12Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRepository12ParamsWithHTTPClient(client *http.Client) *CreateRepository12Params {
	return &CreateRepository12Params{
		HTTPClient: client,
	}
}

/* CreateRepository12Params contains all the parameters to send to the API endpoint
   for the create repository 12 operation.

   Typically these are written to a http.Request.
*/
type CreateRepository12Params struct {

	// Body.
	Body *models.NugetHostedRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create repository 12 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository12Params) WithDefaults() *CreateRepository12Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create repository 12 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository12Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create repository 12 params
func (o *CreateRepository12Params) WithTimeout(timeout time.Duration) *CreateRepository12Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create repository 12 params
func (o *CreateRepository12Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create repository 12 params
func (o *CreateRepository12Params) WithContext(ctx context.Context) *CreateRepository12Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create repository 12 params
func (o *CreateRepository12Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create repository 12 params
func (o *CreateRepository12Params) WithHTTPClient(client *http.Client) *CreateRepository12Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create repository 12 params
func (o *CreateRepository12Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create repository 12 params
func (o *CreateRepository12Params) WithBody(body *models.NugetHostedRepositoryAPIRequest) *CreateRepository12Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create repository 12 params
func (o *CreateRepository12Params) SetBody(body *models.NugetHostedRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRepository12Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
