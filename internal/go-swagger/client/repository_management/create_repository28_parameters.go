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

// NewCreateRepository28Params creates a new CreateRepository28Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRepository28Params() *CreateRepository28Params {
	return &CreateRepository28Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRepository28ParamsWithTimeout creates a new CreateRepository28Params object
// with the ability to set a timeout on a request.
func NewCreateRepository28ParamsWithTimeout(timeout time.Duration) *CreateRepository28Params {
	return &CreateRepository28Params{
		timeout: timeout,
	}
}

// NewCreateRepository28ParamsWithContext creates a new CreateRepository28Params object
// with the ability to set a context for a request.
func NewCreateRepository28ParamsWithContext(ctx context.Context) *CreateRepository28Params {
	return &CreateRepository28Params{
		Context: ctx,
	}
}

// NewCreateRepository28ParamsWithHTTPClient creates a new CreateRepository28Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRepository28ParamsWithHTTPClient(client *http.Client) *CreateRepository28Params {
	return &CreateRepository28Params{
		HTTPClient: client,
	}
}

/* CreateRepository28Params contains all the parameters to send to the API endpoint
   for the create repository 28 operation.

   Typically these are written to a http.Request.
*/
type CreateRepository28Params struct {

	// Body.
	Body *models.PypiProxyRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create repository 28 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository28Params) WithDefaults() *CreateRepository28Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create repository 28 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository28Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create repository 28 params
func (o *CreateRepository28Params) WithTimeout(timeout time.Duration) *CreateRepository28Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create repository 28 params
func (o *CreateRepository28Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create repository 28 params
func (o *CreateRepository28Params) WithContext(ctx context.Context) *CreateRepository28Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create repository 28 params
func (o *CreateRepository28Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create repository 28 params
func (o *CreateRepository28Params) WithHTTPClient(client *http.Client) *CreateRepository28Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create repository 28 params
func (o *CreateRepository28Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create repository 28 params
func (o *CreateRepository28Params) WithBody(body *models.PypiProxyRepositoryAPIRequest) *CreateRepository28Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create repository 28 params
func (o *CreateRepository28Params) SetBody(body *models.PypiProxyRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRepository28Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
