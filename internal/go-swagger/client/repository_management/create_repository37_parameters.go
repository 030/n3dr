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

// NewCreateRepository37Params creates a new CreateRepository37Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRepository37Params() *CreateRepository37Params {
	return &CreateRepository37Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRepository37ParamsWithTimeout creates a new CreateRepository37Params object
// with the ability to set a timeout on a request.
func NewCreateRepository37ParamsWithTimeout(timeout time.Duration) *CreateRepository37Params {
	return &CreateRepository37Params{
		timeout: timeout,
	}
}

// NewCreateRepository37ParamsWithContext creates a new CreateRepository37Params object
// with the ability to set a context for a request.
func NewCreateRepository37ParamsWithContext(ctx context.Context) *CreateRepository37Params {
	return &CreateRepository37Params{
		Context: ctx,
	}
}

// NewCreateRepository37ParamsWithHTTPClient creates a new CreateRepository37Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRepository37ParamsWithHTTPClient(client *http.Client) *CreateRepository37Params {
	return &CreateRepository37Params{
		HTTPClient: client,
	}
}

/* CreateRepository37Params contains all the parameters to send to the API endpoint
   for the create repository 37 operation.

   Typically these are written to a http.Request.
*/
type CreateRepository37Params struct {

	// Body.
	Body *models.P2ProxyRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create repository 37 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository37Params) WithDefaults() *CreateRepository37Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create repository 37 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository37Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create repository 37 params
func (o *CreateRepository37Params) WithTimeout(timeout time.Duration) *CreateRepository37Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create repository 37 params
func (o *CreateRepository37Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create repository 37 params
func (o *CreateRepository37Params) WithContext(ctx context.Context) *CreateRepository37Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create repository 37 params
func (o *CreateRepository37Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create repository 37 params
func (o *CreateRepository37Params) WithHTTPClient(client *http.Client) *CreateRepository37Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create repository 37 params
func (o *CreateRepository37Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create repository 37 params
func (o *CreateRepository37Params) WithBody(body *models.P2ProxyRepositoryAPIRequest) *CreateRepository37Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create repository 37 params
func (o *CreateRepository37Params) SetBody(body *models.P2ProxyRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRepository37Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
