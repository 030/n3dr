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

// NewUpdateRepository33Params creates a new UpdateRepository33Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRepository33Params() *UpdateRepository33Params {
	return &UpdateRepository33Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRepository33ParamsWithTimeout creates a new UpdateRepository33Params object
// with the ability to set a timeout on a request.
func NewUpdateRepository33ParamsWithTimeout(timeout time.Duration) *UpdateRepository33Params {
	return &UpdateRepository33Params{
		timeout: timeout,
	}
}

// NewUpdateRepository33ParamsWithContext creates a new UpdateRepository33Params object
// with the ability to set a context for a request.
func NewUpdateRepository33ParamsWithContext(ctx context.Context) *UpdateRepository33Params {
	return &UpdateRepository33Params{
		Context: ctx,
	}
}

// NewUpdateRepository33ParamsWithHTTPClient creates a new UpdateRepository33Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRepository33ParamsWithHTTPClient(client *http.Client) *UpdateRepository33Params {
	return &UpdateRepository33Params{
		HTTPClient: client,
	}
}

/* UpdateRepository33Params contains all the parameters to send to the API endpoint
   for the update repository 33 operation.

   Typically these are written to a http.Request.
*/
type UpdateRepository33Params struct {

	// Body.
	Body *models.RProxyRepositoryAPIRequest

	/* RepositoryName.

	   Name of the repository to update
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update repository 33 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRepository33Params) WithDefaults() *UpdateRepository33Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update repository 33 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRepository33Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update repository 33 params
func (o *UpdateRepository33Params) WithTimeout(timeout time.Duration) *UpdateRepository33Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update repository 33 params
func (o *UpdateRepository33Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update repository 33 params
func (o *UpdateRepository33Params) WithContext(ctx context.Context) *UpdateRepository33Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update repository 33 params
func (o *UpdateRepository33Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update repository 33 params
func (o *UpdateRepository33Params) WithHTTPClient(client *http.Client) *UpdateRepository33Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update repository 33 params
func (o *UpdateRepository33Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update repository 33 params
func (o *UpdateRepository33Params) WithBody(body *models.RProxyRepositoryAPIRequest) *UpdateRepository33Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update repository 33 params
func (o *UpdateRepository33Params) SetBody(body *models.RProxyRepositoryAPIRequest) {
	o.Body = body
}

// WithRepositoryName adds the repositoryName to the update repository 33 params
func (o *UpdateRepository33Params) WithRepositoryName(repositoryName string) *UpdateRepository33Params {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the update repository 33 params
func (o *UpdateRepository33Params) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRepository33Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param repositoryName
	if err := r.SetPathParam("repositoryName", o.RepositoryName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
