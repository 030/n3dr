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

// NewCreateRepository31Params creates a new CreateRepository31Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRepository31Params() *CreateRepository31Params {
	return &CreateRepository31Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRepository31ParamsWithTimeout creates a new CreateRepository31Params object
// with the ability to set a timeout on a request.
func NewCreateRepository31ParamsWithTimeout(timeout time.Duration) *CreateRepository31Params {
	return &CreateRepository31Params{
		timeout: timeout,
	}
}

// NewCreateRepository31ParamsWithContext creates a new CreateRepository31Params object
// with the ability to set a context for a request.
func NewCreateRepository31ParamsWithContext(ctx context.Context) *CreateRepository31Params {
	return &CreateRepository31Params{
		Context: ctx,
	}
}

// NewCreateRepository31ParamsWithHTTPClient creates a new CreateRepository31Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRepository31ParamsWithHTTPClient(client *http.Client) *CreateRepository31Params {
	return &CreateRepository31Params{
		HTTPClient: client,
	}
}

/* CreateRepository31Params contains all the parameters to send to the API endpoint
   for the create repository 31 operation.

   Typically these are written to a http.Request.
*/
type CreateRepository31Params struct {

	// Body.
	Body *models.RGroupRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create repository 31 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository31Params) WithDefaults() *CreateRepository31Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create repository 31 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository31Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create repository 31 params
func (o *CreateRepository31Params) WithTimeout(timeout time.Duration) *CreateRepository31Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create repository 31 params
func (o *CreateRepository31Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create repository 31 params
func (o *CreateRepository31Params) WithContext(ctx context.Context) *CreateRepository31Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create repository 31 params
func (o *CreateRepository31Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create repository 31 params
func (o *CreateRepository31Params) WithHTTPClient(client *http.Client) *CreateRepository31Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create repository 31 params
func (o *CreateRepository31Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create repository 31 params
func (o *CreateRepository31Params) WithBody(body *models.RGroupRepositoryAPIRequest) *CreateRepository31Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create repository 31 params
func (o *CreateRepository31Params) SetBody(body *models.RGroupRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRepository31Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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