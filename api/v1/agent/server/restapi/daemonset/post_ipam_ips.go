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

// PostIpamIpsHandlerFunc turns a function with the right signature into a post ipam ips handler
type PostIpamIpsHandlerFunc func(PostIpamIpsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostIpamIpsHandlerFunc) Handle(params PostIpamIpsParams) middleware.Responder {
	return fn(params)
}

// PostIpamIpsHandler interface for that can handle valid post ipam ips params
type PostIpamIpsHandler interface {
	Handle(PostIpamIpsParams) middleware.Responder
}

// NewPostIpamIps creates a new http.Handler for the post ipam ips operation
func NewPostIpamIps(ctx *middleware.Context, handler PostIpamIpsHandler) *PostIpamIps {
	return &PostIpamIps{Context: ctx, Handler: handler}
}

/*
	PostIpamIps swagger:route POST /ipam/ips daemonset postIpamIps

# Assign multiple ip as a batch

Assign multiple ip for a pod, case for spiderflat compent
*/
type PostIpamIps struct {
	Context *middleware.Context
	Handler PostIpamIpsHandler
}

func (o *PostIpamIps) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostIpamIpsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
