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

// NewGetRuntimeLivenessParams creates a new GetRuntimeLivenessParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRuntimeLivenessParams() *GetRuntimeLivenessParams {
	return &GetRuntimeLivenessParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRuntimeLivenessParamsWithTimeout creates a new GetRuntimeLivenessParams object
// with the ability to set a timeout on a request.
func NewGetRuntimeLivenessParamsWithTimeout(timeout time.Duration) *GetRuntimeLivenessParams {
	return &GetRuntimeLivenessParams{
		timeout: timeout,
	}
}

// NewGetRuntimeLivenessParamsWithContext creates a new GetRuntimeLivenessParams object
// with the ability to set a context for a request.
func NewGetRuntimeLivenessParamsWithContext(ctx context.Context) *GetRuntimeLivenessParams {
	return &GetRuntimeLivenessParams{
		Context: ctx,
	}
}

// NewGetRuntimeLivenessParamsWithHTTPClient creates a new GetRuntimeLivenessParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRuntimeLivenessParamsWithHTTPClient(client *http.Client) *GetRuntimeLivenessParams {
	return &GetRuntimeLivenessParams{
		HTTPClient: client,
	}
}

/*
GetRuntimeLivenessParams contains all the parameters to send to the API endpoint

	for the get runtime liveness operation.

	Typically these are written to a http.Request.
*/
type GetRuntimeLivenessParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get runtime liveness params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRuntimeLivenessParams) WithDefaults() *GetRuntimeLivenessParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get runtime liveness params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRuntimeLivenessParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get runtime liveness params
func (o *GetRuntimeLivenessParams) WithTimeout(timeout time.Duration) *GetRuntimeLivenessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get runtime liveness params
func (o *GetRuntimeLivenessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get runtime liveness params
func (o *GetRuntimeLivenessParams) WithContext(ctx context.Context) *GetRuntimeLivenessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get runtime liveness params
func (o *GetRuntimeLivenessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get runtime liveness params
func (o *GetRuntimeLivenessParams) WithHTTPClient(client *http.Client) *GetRuntimeLivenessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get runtime liveness params
func (o *GetRuntimeLivenessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetRuntimeLivenessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
