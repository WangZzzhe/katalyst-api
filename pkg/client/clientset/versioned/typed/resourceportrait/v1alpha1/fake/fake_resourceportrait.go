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

	v1alpha1 "github.com/kubewharf/katalyst-api/pkg/apis/resourceportrait/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeResourcePortraits implements ResourcePortraitInterface
type FakeResourcePortraits struct {
	Fake *FakeResourceportraitV1alpha1
	ns   string
}

var resourceportraitsResource = schema.GroupVersionResource{Group: "resourceportrait.katalyst.kubewharf.io", Version: "v1alpha1", Resource: "resourceportraits"}

var resourceportraitsKind = schema.GroupVersionKind{Group: "resourceportrait.katalyst.kubewharf.io", Version: "v1alpha1", Kind: "ResourcePortrait"}

// Get takes name of the resourcePortrait, and returns the corresponding resourcePortrait object, and an error if there is any.
func (c *FakeResourcePortraits) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ResourcePortrait, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(resourceportraitsResource, c.ns, name), &v1alpha1.ResourcePortrait{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourcePortrait), err
}

// List takes label and field selectors, and returns the list of ResourcePortraits that match those selectors.
func (c *FakeResourcePortraits) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ResourcePortraitList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(resourceportraitsResource, resourceportraitsKind, c.ns, opts), &v1alpha1.ResourcePortraitList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ResourcePortraitList{ListMeta: obj.(*v1alpha1.ResourcePortraitList).ListMeta}
	for _, item := range obj.(*v1alpha1.ResourcePortraitList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resourcePortraits.
func (c *FakeResourcePortraits) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(resourceportraitsResource, c.ns, opts))

}

// Create takes the representation of a resourcePortrait and creates it.  Returns the server's representation of the resourcePortrait, and an error, if there is any.
func (c *FakeResourcePortraits) Create(ctx context.Context, resourcePortrait *v1alpha1.ResourcePortrait, opts v1.CreateOptions) (result *v1alpha1.ResourcePortrait, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(resourceportraitsResource, c.ns, resourcePortrait), &v1alpha1.ResourcePortrait{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourcePortrait), err
}

// Update takes the representation of a resourcePortrait and updates it. Returns the server's representation of the resourcePortrait, and an error, if there is any.
func (c *FakeResourcePortraits) Update(ctx context.Context, resourcePortrait *v1alpha1.ResourcePortrait, opts v1.UpdateOptions) (result *v1alpha1.ResourcePortrait, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(resourceportraitsResource, c.ns, resourcePortrait), &v1alpha1.ResourcePortrait{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourcePortrait), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeResourcePortraits) UpdateStatus(ctx context.Context, resourcePortrait *v1alpha1.ResourcePortrait, opts v1.UpdateOptions) (*v1alpha1.ResourcePortrait, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(resourceportraitsResource, "status", c.ns, resourcePortrait), &v1alpha1.ResourcePortrait{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourcePortrait), err
}

// Delete takes name of the resourcePortrait and deletes it. Returns an error if one occurs.
func (c *FakeResourcePortraits) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(resourceportraitsResource, c.ns, name, opts), &v1alpha1.ResourcePortrait{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResourcePortraits) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(resourceportraitsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ResourcePortraitList{})
	return err
}

// Patch applies the patch and returns the patched resourcePortrait.
func (c *FakeResourcePortraits) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ResourcePortrait, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resourceportraitsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ResourcePortrait{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourcePortrait), err
}
