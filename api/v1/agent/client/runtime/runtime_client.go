// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package runtime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new runtime API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for runtime API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetRuntimeLiveness(params *GetRuntimeLivenessParams, opts ...ClientOption) (*GetRuntimeLivenessOK, error)

	GetRuntimeReadiness(params *GetRuntimeReadinessParams, opts ...ClientOption) (*GetRuntimeReadinessOK, error)

	GetRuntimeStartup(params *GetRuntimeStartupParams, opts ...ClientOption) (*GetRuntimeStartupOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetRuntimeLiveness livenesses probe

Check pod liveness probe
*/
func (a *Client) GetRuntimeLiveness(params *GetRuntimeLivenessParams, opts ...ClientOption) (*GetRuntimeLivenessOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRuntimeLivenessParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetRuntimeLiveness",
		Method:             "GET",
		PathPattern:        "/runtime/liveness",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRuntimeLivenessReader{formats: a.formats},
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
	success, ok := result.(*GetRuntimeLivenessOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetRuntimeLiveness: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetRuntimeReadiness readinesses probe

Check pod readiness probe
*/
func (a *Client) GetRuntimeReadiness(params *GetRuntimeReadinessParams, opts ...ClientOption) (*GetRuntimeReadinessOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRuntimeReadinessParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetRuntimeReadiness",
		Method:             "GET",
		PathPattern:        "/runtime/readiness",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRuntimeReadinessReader{formats: a.formats},
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
	success, ok := result.(*GetRuntimeReadinessOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetRuntimeReadiness: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetRuntimeStartup startups probe

Check pod startup probe
*/
func (a *Client) GetRuntimeStartup(params *GetRuntimeStartupParams, opts ...ClientOption) (*GetRuntimeStartupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRuntimeStartupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetRuntimeStartup",
		Method:             "GET",
		PathPattern:        "/runtime/startup",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRuntimeStartupReader{formats: a.formats},
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
	success, ok := result.(*GetRuntimeStartupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetRuntimeStartup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
