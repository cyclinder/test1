// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package daemonset

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetWorkloadendpointHandlerFunc turns a function with the right signature into a get workloadendpoint handler
type GetWorkloadendpointHandlerFunc func(GetWorkloadendpointParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetWorkloadendpointHandlerFunc) Handle(params GetWorkloadendpointParams) middleware.Responder {
	return fn(params)
}

// GetWorkloadendpointHandler interface for that can handle valid get workloadendpoint params
type GetWorkloadendpointHandler interface {
	Handle(GetWorkloadendpointParams) middleware.Responder
}

// NewGetWorkloadendpoint creates a new http.Handler for the get workloadendpoint operation
func NewGetWorkloadendpoint(ctx *middleware.Context, handler GetWorkloadendpointHandler) *GetWorkloadendpoint {
	return &GetWorkloadendpoint{Context: ctx, Handler: handler}
}

/*
	GetWorkloadendpoint swagger:route GET /workloadendpoint daemonset getWorkloadendpoint

# Get workloadendpoint status

Get workloadendpoint details for spiderflat use
*/
type GetWorkloadendpoint struct {
	Context *middleware.Context
	Handler GetWorkloadendpointHandler
}

func (o *GetWorkloadendpoint) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetWorkloadendpointParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
