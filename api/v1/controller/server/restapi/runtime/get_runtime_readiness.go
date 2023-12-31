// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package runtime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetRuntimeReadinessHandlerFunc turns a function with the right signature into a get runtime readiness handler
type GetRuntimeReadinessHandlerFunc func(GetRuntimeReadinessParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRuntimeReadinessHandlerFunc) Handle(params GetRuntimeReadinessParams) middleware.Responder {
	return fn(params)
}

// GetRuntimeReadinessHandler interface for that can handle valid get runtime readiness params
type GetRuntimeReadinessHandler interface {
	Handle(GetRuntimeReadinessParams) middleware.Responder
}

// NewGetRuntimeReadiness creates a new http.Handler for the get runtime readiness operation
func NewGetRuntimeReadiness(ctx *middleware.Context, handler GetRuntimeReadinessHandler) *GetRuntimeReadiness {
	return &GetRuntimeReadiness{Context: ctx, Handler: handler}
}

/*
	GetRuntimeReadiness swagger:route GET /runtime/readiness runtime getRuntimeReadiness

# Readiness probe

Check pod readiness probe
*/
type GetRuntimeReadiness struct {
	Context *middleware.Context
	Handler GetRuntimeReadinessHandler
}

func (o *GetRuntimeReadiness) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetRuntimeReadinessParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
