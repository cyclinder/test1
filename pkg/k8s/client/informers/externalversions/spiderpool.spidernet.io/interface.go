// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

// Code generated by informer-gen. DO NOT EDIT.

package spiderpool

import (
	internalinterfaces "github.com/spidernet-io/spiderpool/pkg/k8s/client/informers/externalversions/internalinterfaces"
	v2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/client/informers/externalversions/spiderpool.spidernet.io/v2beta1"
)

// Interface provides access to each of this group's versions.
type Interface interface {
	// V2beta1 provides access to shared informers for resources in V2beta1.
	V2beta1() v2beta1.Interface
}

type group struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &group{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// V2beta1 returns a new v2beta1.Interface.
func (g *group) V2beta1() v2beta1.Interface {
	return v2beta1.New(g.factory, g.namespace, g.tweakListOptions)
}
