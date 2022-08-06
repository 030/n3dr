// Code generated by go-swagger; DO NOT EDIT.

package security_management_privileges

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new security management privileges API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for security management privileges API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreatePrivilege(params *CreatePrivilegeParams, opts ...ClientOption) error

	CreatePrivilege1(params *CreatePrivilege1Params, opts ...ClientOption) error

	CreatePrivilege2(params *CreatePrivilege2Params, opts ...ClientOption) error

	CreatePrivilege3(params *CreatePrivilege3Params, opts ...ClientOption) error

	CreatePrivilege4(params *CreatePrivilege4Params, opts ...ClientOption) error

	CreatePrivilege5(params *CreatePrivilege5Params, opts ...ClientOption) error

	DeletePrivilege(params *DeletePrivilegeParams, opts ...ClientOption) error

	GetPrivilege(params *GetPrivilegeParams, opts ...ClientOption) (*GetPrivilegeOK, error)

	GetPrivileges(params *GetPrivilegesParams, opts ...ClientOption) (*GetPrivilegesOK, error)

	UpdatePrivilege(params *UpdatePrivilegeParams, opts ...ClientOption) error

	UpdatePrivilege1(params *UpdatePrivilege1Params, opts ...ClientOption) error

	UpdatePrivilege2(params *UpdatePrivilege2Params, opts ...ClientOption) error

	UpdatePrivilege3(params *UpdatePrivilege3Params, opts ...ClientOption) error

	UpdatePrivilege4(params *UpdatePrivilege4Params, opts ...ClientOption) error

	UpdatePrivilege5(params *UpdatePrivilege5Params, opts ...ClientOption) error

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreatePrivilege creates a wildcard type privilege
*/
func (a *Client) CreatePrivilege(params *CreatePrivilegeParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePrivilegeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createPrivilege",
		Method:             "POST",
		PathPattern:        "/v1/security/privileges/wildcard",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreatePrivilegeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
  CreatePrivilege1 creates an application type privilege
*/
func (a *Client) CreatePrivilege1(params *CreatePrivilege1Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePrivilege1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "createPrivilege_1",
		Method:             "POST",
		PathPattern:        "/v1/security/privileges/application",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreatePrivilege1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
  CreatePrivilege2 creates a repository content selector type privilege
*/
func (a *Client) CreatePrivilege2(params *CreatePrivilege2Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePrivilege2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "createPrivilege_2",
		Method:             "POST",
		PathPattern:        "/v1/security/privileges/repository-content-selector",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreatePrivilege2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
  CreatePrivilege3 creates a repository admin type privilege
*/
func (a *Client) CreatePrivilege3(params *CreatePrivilege3Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePrivilege3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "createPrivilege_3",
		Method:             "POST",
		PathPattern:        "/v1/security/privileges/repository-admin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreatePrivilege3Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
  CreatePrivilege4 creates a repository view type privilege
*/
func (a *Client) CreatePrivilege4(params *CreatePrivilege4Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePrivilege4Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "createPrivilege_4",
		Method:             "POST",
		PathPattern:        "/v1/security/privileges/repository-view",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreatePrivilege4Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
  CreatePrivilege5 creates a script type privilege
*/
func (a *Client) CreatePrivilege5(params *CreatePrivilege5Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePrivilege5Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "createPrivilege_5",
		Method:             "POST",
		PathPattern:        "/v1/security/privileges/script",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreatePrivilege5Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
  DeletePrivilege deletes a privilege by name
*/
func (a *Client) DeletePrivilege(params *DeletePrivilegeParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePrivilegeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deletePrivilege",
		Method:             "DELETE",
		PathPattern:        "/v1/security/privileges/{privilegeName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeletePrivilegeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
  GetPrivilege retrieves a privilege by name
*/
func (a *Client) GetPrivilege(params *GetPrivilegeParams, opts ...ClientOption) (*GetPrivilegeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivilegeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPrivilege",
		Method:             "GET",
		PathPattern:        "/v1/security/privileges/{privilegeName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPrivilegeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPrivilegeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPrivilege: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPrivileges retrieves a list of privileges
*/
func (a *Client) GetPrivileges(params *GetPrivilegesParams, opts ...ClientOption) (*GetPrivilegesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivilegesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPrivileges",
		Method:             "GET",
		PathPattern:        "/v1/security/privileges",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPrivilegesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPrivilegesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPrivileges: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdatePrivilege updates a wildcard type privilege
*/
func (a *Client) UpdatePrivilege(params *UpdatePrivilegeParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePrivilegeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updatePrivilege",
		Method:             "PUT",
		PathPattern:        "/v1/security/privileges/wildcard/{privilegeName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdatePrivilegeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
  UpdatePrivilege1 updates an application type privilege
*/
func (a *Client) UpdatePrivilege1(params *UpdatePrivilege1Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePrivilege1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "updatePrivilege_1",
		Method:             "PUT",
		PathPattern:        "/v1/security/privileges/application/{privilegeName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdatePrivilege1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
  UpdatePrivilege2 updates a repository view type privilege
*/
func (a *Client) UpdatePrivilege2(params *UpdatePrivilege2Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePrivilege2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "updatePrivilege_2",
		Method:             "PUT",
		PathPattern:        "/v1/security/privileges/repository-view/{privilegeName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdatePrivilege2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
  UpdatePrivilege3 updates a repository content selector type privilege
*/
func (a *Client) UpdatePrivilege3(params *UpdatePrivilege3Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePrivilege3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "updatePrivilege_3",
		Method:             "PUT",
		PathPattern:        "/v1/security/privileges/repository-content-selector/{privilegeName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdatePrivilege3Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
  UpdatePrivilege4 updates a repository admin type privilege
*/
func (a *Client) UpdatePrivilege4(params *UpdatePrivilege4Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePrivilege4Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "updatePrivilege_4",
		Method:             "PUT",
		PathPattern:        "/v1/security/privileges/repository-admin/{privilegeName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdatePrivilege4Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
  UpdatePrivilege5 updates a script type privilege
*/
func (a *Client) UpdatePrivilege5(params *UpdatePrivilege5Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePrivilege5Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "updatePrivilege_5",
		Method:             "PUT",
		PathPattern:        "/v1/security/privileges/script/{privilegeName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdatePrivilege5Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
