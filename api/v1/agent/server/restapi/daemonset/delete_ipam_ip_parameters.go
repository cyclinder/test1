// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package daemonset

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"github.com/spidernet-io/spiderpool/api/v1/agent/models"
)

// NewDeleteIpamIPParams creates a new DeleteIpamIPParams object
//
// There are no default values defined in the spec.
func NewDeleteIpamIPParams() DeleteIpamIPParams {

	return DeleteIpamIPParams{}
}

// DeleteIpamIPParams contains all the bound params for the delete ipam IP operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteIpamIP
type DeleteIpamIPParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	IpamDelArgs *models.IpamDelArgs
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteIpamIPParams() beforehand.
func (o *DeleteIpamIPParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.IpamDelArgs
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("ipamDelArgs", "body", ""))
			} else {
				res = append(res, errors.NewParseError("ipamDelArgs", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.IpamDelArgs = &body
			}
		}
	} else {
		res = append(res, errors.Required("ipamDelArgs", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
