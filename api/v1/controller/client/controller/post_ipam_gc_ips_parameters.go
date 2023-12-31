// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package controller

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

// NewPostIpamGcIpsParams creates a new PostIpamGcIpsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostIpamGcIpsParams() *PostIpamGcIpsParams {
	return &PostIpamGcIpsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostIpamGcIpsParamsWithTimeout creates a new PostIpamGcIpsParams object
// with the ability to set a timeout on a request.
func NewPostIpamGcIpsParamsWithTimeout(timeout time.Duration) *PostIpamGcIpsParams {
	return &PostIpamGcIpsParams{
		timeout: timeout,
	}
}

// NewPostIpamGcIpsParamsWithContext creates a new PostIpamGcIpsParams object
// with the ability to set a context for a request.
func NewPostIpamGcIpsParamsWithContext(ctx context.Context) *PostIpamGcIpsParams {
	return &PostIpamGcIpsParams{
		Context: ctx,
	}
}

// NewPostIpamGcIpsParamsWithHTTPClient creates a new PostIpamGcIpsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostIpamGcIpsParamsWithHTTPClient(client *http.Client) *PostIpamGcIpsParams {
	return &PostIpamGcIpsParams{
		HTTPClient: client,
	}
}

/*
PostIpamGcIpsParams contains all the parameters to send to the API endpoint

	for the post ipam gc ips operation.

	Typically these are written to a http.Request.
*/
type PostIpamGcIpsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post ipam gc ips params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostIpamGcIpsParams) WithDefaults() *PostIpamGcIpsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post ipam gc ips params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostIpamGcIpsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post ipam gc ips params
func (o *PostIpamGcIpsParams) WithTimeout(timeout time.Duration) *PostIpamGcIpsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post ipam gc ips params
func (o *PostIpamGcIpsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post ipam gc ips params
func (o *PostIpamGcIpsParams) WithContext(ctx context.Context) *PostIpamGcIpsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post ipam gc ips params
func (o *PostIpamGcIpsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post ipam gc ips params
func (o *PostIpamGcIpsParams) WithHTTPClient(client *http.Client) *PostIpamGcIpsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post ipam gc ips params
func (o *PostIpamGcIpsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostIpamGcIpsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
