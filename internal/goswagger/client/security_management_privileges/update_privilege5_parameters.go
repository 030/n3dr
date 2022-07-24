// Code generated by go-swagger; DO NOT EDIT.

package security_management_privileges

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

// NewUpdatePrivilege5Params creates a new UpdatePrivilege5Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdatePrivilege5Params() *UpdatePrivilege5Params {
	return &UpdatePrivilege5Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePrivilege5ParamsWithTimeout creates a new UpdatePrivilege5Params object
// with the ability to set a timeout on a request.
func NewUpdatePrivilege5ParamsWithTimeout(timeout time.Duration) *UpdatePrivilege5Params {
	return &UpdatePrivilege5Params{
		timeout: timeout,
	}
}

// NewUpdatePrivilege5ParamsWithContext creates a new UpdatePrivilege5Params object
// with the ability to set a context for a request.
func NewUpdatePrivilege5ParamsWithContext(ctx context.Context) *UpdatePrivilege5Params {
	return &UpdatePrivilege5Params{
		Context: ctx,
	}
}

// NewUpdatePrivilege5ParamsWithHTTPClient creates a new UpdatePrivilege5Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdatePrivilege5ParamsWithHTTPClient(client *http.Client) *UpdatePrivilege5Params {
	return &UpdatePrivilege5Params{
		HTTPClient: client,
	}
}

/* UpdatePrivilege5Params contains all the parameters to send to the API endpoint
   for the update privilege 5 operation.

   Typically these are written to a http.Request.
*/
type UpdatePrivilege5Params struct {

	/* Body.

	   The privilege to update.
	*/
	Body *models.APIPrivilegeScriptRequest

	/* PrivilegeName.

	   The name of the privilege to update.
	*/
	PrivilegeName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update privilege 5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePrivilege5Params) WithDefaults() *UpdatePrivilege5Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update privilege 5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePrivilege5Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update privilege 5 params
func (o *UpdatePrivilege5Params) WithTimeout(timeout time.Duration) *UpdatePrivilege5Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update privilege 5 params
func (o *UpdatePrivilege5Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update privilege 5 params
func (o *UpdatePrivilege5Params) WithContext(ctx context.Context) *UpdatePrivilege5Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update privilege 5 params
func (o *UpdatePrivilege5Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update privilege 5 params
func (o *UpdatePrivilege5Params) WithHTTPClient(client *http.Client) *UpdatePrivilege5Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update privilege 5 params
func (o *UpdatePrivilege5Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update privilege 5 params
func (o *UpdatePrivilege5Params) WithBody(body *models.APIPrivilegeScriptRequest) *UpdatePrivilege5Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update privilege 5 params
func (o *UpdatePrivilege5Params) SetBody(body *models.APIPrivilegeScriptRequest) {
	o.Body = body
}

// WithPrivilegeName adds the privilegeName to the update privilege 5 params
func (o *UpdatePrivilege5Params) WithPrivilegeName(privilegeName string) *UpdatePrivilege5Params {
	o.SetPrivilegeName(privilegeName)
	return o
}

// SetPrivilegeName adds the privilegeName to the update privilege 5 params
func (o *UpdatePrivilege5Params) SetPrivilegeName(privilegeName string) {
	o.PrivilegeName = privilegeName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePrivilege5Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param privilegeName
	if err := r.SetPathParam("privilegeName", o.PrivilegeName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
