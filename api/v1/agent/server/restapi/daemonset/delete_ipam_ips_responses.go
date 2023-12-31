// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package daemonset

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/spidernet-io/spiderpool/api/v1/agent/models"
)

// DeleteIpamIpsOKCode is the HTTP code returned for type DeleteIpamIpsOK
const DeleteIpamIpsOKCode int = 200

/*
DeleteIpamIpsOK Success

swagger:response deleteIpamIpsOK
*/
type DeleteIpamIpsOK struct {
}

// NewDeleteIpamIpsOK creates DeleteIpamIpsOK with default headers values
func NewDeleteIpamIpsOK() *DeleteIpamIpsOK {

	return &DeleteIpamIpsOK{}
}

// WriteResponse to the client
func (o *DeleteIpamIpsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteIpamIpsFailureCode is the HTTP code returned for type DeleteIpamIpsFailure
const DeleteIpamIpsFailureCode int = 500

/*
DeleteIpamIpsFailure Addresses release failure

swagger:response deleteIpamIpsFailure
*/
type DeleteIpamIpsFailure struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewDeleteIpamIpsFailure creates DeleteIpamIpsFailure with default headers values
func NewDeleteIpamIpsFailure() *DeleteIpamIpsFailure {

	return &DeleteIpamIpsFailure{}
}

// WithPayload adds the payload to the delete ipam ips failure response
func (o *DeleteIpamIpsFailure) WithPayload(payload models.Error) *DeleteIpamIpsFailure {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete ipam ips failure response
func (o *DeleteIpamIpsFailure) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteIpamIpsFailure) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
