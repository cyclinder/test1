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

// PostIpamIPHandlerFunc turns a function with the right signature into a post ipam IP handler
type PostIpamIPHandlerFunc func(PostIpamIPParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostIpamIPHandlerFunc) Handle(params PostIpamIPParams) middleware.Responder {
	return fn(params)
}

// PostIpamIPHandler interface for that can handle valid post ipam IP params
type PostIpamIPHandler interface {
	Handle(PostIpamIPParams) middleware.Responder
}

// NewPostIpamIP creates a new http.Handler for the post ipam IP operation
func NewPostIpamIP(ctx *middleware.Context, handler PostIpamIPHandler) *PostIpamIP {
	return &PostIpamIP{Context: ctx, Handler: handler}
}

/*
	PostIpamIP swagger:route POST /ipam/ip daemonset postIpamIp

# Get ip from spiderpool daemon

Send a request to daemonset to ask for an ip assignment
*/
type PostIpamIP struct {
	Context *middleware.Context
	Handler PostIpamIPHandler
}

func (o *PostIpamIP) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostIpamIPParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
