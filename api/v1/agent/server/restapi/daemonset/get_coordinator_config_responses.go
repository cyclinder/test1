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

// GetCoordinatorConfigOKCode is the HTTP code returned for type GetCoordinatorConfigOK
const GetCoordinatorConfigOKCode int = 200

/*
GetCoordinatorConfigOK Success

swagger:response getCoordinatorConfigOK
*/
type GetCoordinatorConfigOK struct {

	/*
	  In: Body
	*/
	Payload *models.CoordinatorConfig `json:"body,omitempty"`
}

// NewGetCoordinatorConfigOK creates GetCoordinatorConfigOK with default headers values
func NewGetCoordinatorConfigOK() *GetCoordinatorConfigOK {

	return &GetCoordinatorConfigOK{}
}

// WithPayload adds the payload to the get coordinator config o k response
func (o *GetCoordinatorConfigOK) WithPayload(payload *models.CoordinatorConfig) *GetCoordinatorConfigOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get coordinator config o k response
func (o *GetCoordinatorConfigOK) SetPayload(payload *models.CoordinatorConfig) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCoordinatorConfigOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCoordinatorConfigFailureCode is the HTTP code returned for type GetCoordinatorConfigFailure
const GetCoordinatorConfigFailureCode int = 500

/*
GetCoordinatorConfigFailure Config not initialized successfully or others

swagger:response getCoordinatorConfigFailure
*/
type GetCoordinatorConfigFailure struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewGetCoordinatorConfigFailure creates GetCoordinatorConfigFailure with default headers values
func NewGetCoordinatorConfigFailure() *GetCoordinatorConfigFailure {

	return &GetCoordinatorConfigFailure{}
}

// WithPayload adds the payload to the get coordinator config failure response
func (o *GetCoordinatorConfigFailure) WithPayload(payload models.Error) *GetCoordinatorConfigFailure {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get coordinator config failure response
func (o *GetCoordinatorConfigFailure) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCoordinatorConfigFailure) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
