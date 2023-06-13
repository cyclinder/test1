// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSpiderIPPools implements SpiderIPPoolInterface
type FakeSpiderIPPools struct {
	Fake *FakeSpiderpoolV2beta1
}

var spiderippoolsResource = schema.GroupVersionResource{Group: "spiderpool.spidernet.io", Version: "v2beta1", Resource: "spiderippools"}

var spiderippoolsKind = schema.GroupVersionKind{Group: "spiderpool.spidernet.io", Version: "v2beta1", Kind: "SpiderIPPool"}

// Get takes name of the spiderIPPool, and returns the corresponding spiderIPPool object, and an error if there is any.
func (c *FakeSpiderIPPools) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2beta1.SpiderIPPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(spiderippoolsResource, name), &v2beta1.SpiderIPPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2beta1.SpiderIPPool), err
}

// List takes label and field selectors, and returns the list of SpiderIPPools that match those selectors.
func (c *FakeSpiderIPPools) List(ctx context.Context, opts v1.ListOptions) (result *v2beta1.SpiderIPPoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(spiderippoolsResource, spiderippoolsKind, opts), &v2beta1.SpiderIPPoolList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2beta1.SpiderIPPoolList{ListMeta: obj.(*v2beta1.SpiderIPPoolList).ListMeta}
	for _, item := range obj.(*v2beta1.SpiderIPPoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested spiderIPPools.
func (c *FakeSpiderIPPools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(spiderippoolsResource, opts))
}

// Create takes the representation of a spiderIPPool and creates it.  Returns the server's representation of the spiderIPPool, and an error, if there is any.
func (c *FakeSpiderIPPools) Create(ctx context.Context, spiderIPPool *v2beta1.SpiderIPPool, opts v1.CreateOptions) (result *v2beta1.SpiderIPPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(spiderippoolsResource, spiderIPPool), &v2beta1.SpiderIPPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2beta1.SpiderIPPool), err
}

// Update takes the representation of a spiderIPPool and updates it. Returns the server's representation of the spiderIPPool, and an error, if there is any.
func (c *FakeSpiderIPPools) Update(ctx context.Context, spiderIPPool *v2beta1.SpiderIPPool, opts v1.UpdateOptions) (result *v2beta1.SpiderIPPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(spiderippoolsResource, spiderIPPool), &v2beta1.SpiderIPPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2beta1.SpiderIPPool), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSpiderIPPools) UpdateStatus(ctx context.Context, spiderIPPool *v2beta1.SpiderIPPool, opts v1.UpdateOptions) (*v2beta1.SpiderIPPool, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(spiderippoolsResource, "status", spiderIPPool), &v2beta1.SpiderIPPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2beta1.SpiderIPPool), err
}

// Delete takes name of the spiderIPPool and deletes it. Returns an error if one occurs.
func (c *FakeSpiderIPPools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(spiderippoolsResource, name, opts), &v2beta1.SpiderIPPool{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSpiderIPPools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(spiderippoolsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v2beta1.SpiderIPPoolList{})
	return err
}

// Patch applies the patch and returns the patched spiderIPPool.
func (c *FakeSpiderIPPools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2beta1.SpiderIPPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(spiderippoolsResource, name, pt, data, subresources...), &v2beta1.SpiderIPPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2beta1.SpiderIPPool), err
}
