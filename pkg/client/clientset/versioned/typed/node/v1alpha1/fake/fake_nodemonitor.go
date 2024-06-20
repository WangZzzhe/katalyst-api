/*
Copyright 2022 The Katalyst Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/kubewharf/katalyst-api/pkg/apis/node/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNodeMonitors implements NodeMonitorInterface
type FakeNodeMonitors struct {
	Fake *FakeNodeV1alpha1
}

var nodemonitorsResource = schema.GroupVersionResource{Group: "node.katalyst.kubewharf.io", Version: "v1alpha1", Resource: "nodemonitors"}

var nodemonitorsKind = schema.GroupVersionKind{Group: "node.katalyst.kubewharf.io", Version: "v1alpha1", Kind: "NodeMonitor"}

// Get takes name of the nodeMonitor, and returns the corresponding nodeMonitor object, and an error if there is any.
func (c *FakeNodeMonitors) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NodeMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(nodemonitorsResource, name), &v1alpha1.NodeMonitor{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeMonitor), err
}

// List takes label and field selectors, and returns the list of NodeMonitors that match those selectors.
func (c *FakeNodeMonitors) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NodeMonitorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(nodemonitorsResource, nodemonitorsKind, opts), &v1alpha1.NodeMonitorList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NodeMonitorList{ListMeta: obj.(*v1alpha1.NodeMonitorList).ListMeta}
	for _, item := range obj.(*v1alpha1.NodeMonitorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodeMonitors.
func (c *FakeNodeMonitors) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(nodemonitorsResource, opts))
}

// Create takes the representation of a nodeMonitor and creates it.  Returns the server's representation of the nodeMonitor, and an error, if there is any.
func (c *FakeNodeMonitors) Create(ctx context.Context, nodeMonitor *v1alpha1.NodeMonitor, opts v1.CreateOptions) (result *v1alpha1.NodeMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(nodemonitorsResource, nodeMonitor), &v1alpha1.NodeMonitor{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeMonitor), err
}

// Update takes the representation of a nodeMonitor and updates it. Returns the server's representation of the nodeMonitor, and an error, if there is any.
func (c *FakeNodeMonitors) Update(ctx context.Context, nodeMonitor *v1alpha1.NodeMonitor, opts v1.UpdateOptions) (result *v1alpha1.NodeMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(nodemonitorsResource, nodeMonitor), &v1alpha1.NodeMonitor{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeMonitor), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNodeMonitors) UpdateStatus(ctx context.Context, nodeMonitor *v1alpha1.NodeMonitor, opts v1.UpdateOptions) (*v1alpha1.NodeMonitor, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(nodemonitorsResource, "status", nodeMonitor), &v1alpha1.NodeMonitor{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeMonitor), err
}

// Delete takes name of the nodeMonitor and deletes it. Returns an error if one occurs.
func (c *FakeNodeMonitors) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(nodemonitorsResource, name), &v1alpha1.NodeMonitor{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodeMonitors) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(nodemonitorsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NodeMonitorList{})
	return err
}

// Patch applies the patch and returns the patched nodeMonitor.
func (c *FakeNodeMonitors) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NodeMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(nodemonitorsResource, name, pt, data, subresources...), &v1alpha1.NodeMonitor{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeMonitor), err
}
