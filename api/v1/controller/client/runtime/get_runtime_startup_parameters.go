// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package runtime

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
)

// NewGetRuntimeStartupParams creates a new GetRuntimeStartupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRuntimeStartupParams() *GetRuntimeStartupParams {
	return &GetRuntimeStartupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRuntimeStartupParamsWithTimeout creates a new GetRuntimeStartupParams object
// with the ability to set a timeout on a request.
func NewGetRuntimeStartupParamsWithTimeout(timeout time.Duration) *GetRuntimeStartupParams {
	return &GetRuntimeStartupParams{
		timeout: timeout,
	}
}

// NewGetRuntimeStartupParamsWithContext creates a new GetRuntimeStartupParams object
// with the ability to set a context for a request.
func NewGetRuntimeStartupParamsWithContext(ctx context.Context) *GetRuntimeStartupParams {
	return &GetRuntimeStartupParams{
		Context: ctx,
	}
}

// NewGetRuntimeStartupParamsWithHTTPClient creates a new GetRuntimeStartupParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRuntimeStartupParamsWithHTTPClient(client *http.Client) *GetRuntimeStartupParams {
	return &GetRuntimeStartupParams{
		HTTPClient: client,
	}
}

/*
GetRuntimeStartupParams contains all the parameters to send to the API endpoint

	for the get runtime startup operation.

	Typically these are written to a http.Request.
*/
type GetRuntimeStartupParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get runtime startup params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRuntimeStartupParams) WithDefaults() *GetRuntimeStartupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get runtime startup params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRuntimeStartupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get runtime startup params
func (o *GetRuntimeStartupParams) WithTimeout(timeout time.Duration) *GetRuntimeStartupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get runtime startup params
func (o *GetRuntimeStartupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get runtime startup params
func (o *GetRuntimeStartupParams) WithContext(ctx context.Context) *GetRuntimeStartupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get runtime startup params
func (o *GetRuntimeStartupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get runtime startup params
func (o *GetRuntimeStartupParams) WithHTTPClient(client *http.Client) *GetRuntimeStartupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get runtime startup params
func (o *GetRuntimeStartupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetRuntimeStartupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
