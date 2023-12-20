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
	v1alpha1 "github.com/kubewharf/katalyst-api/pkg/apis/autoscaling/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVerticalPodAutoscalerRecommendations implements VerticalPodAutoscalerRecommendationInterface
type FakeVerticalPodAutoscalerRecommendations struct {
	Fake *FakeAutoscalingV1alpha1
	ns   string
}

var verticalpodautoscalerrecommendationsResource = schema.GroupVersionResource{Group: "autoscaling.katalyst.kubewharf.io", Version: "v1alpha1", Resource: "verticalpodautoscalerrecommendations"}

var verticalpodautoscalerrecommendationsKind = schema.GroupVersionKind{Group: "autoscaling.katalyst.kubewharf.io", Version: "v1alpha1", Kind: "VerticalPodAutoscalerRecommendation"}

// Get takes name of the verticalPodAutoscalerRecommendation, and returns the corresponding verticalPodAutoscalerRecommendation object, and an error if there is any.
func (c *FakeVerticalPodAutoscalerRecommendations) Get(name string, options v1.GetOptions) (result *v1alpha1.VerticalPodAutoscalerRecommendation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(verticalpodautoscalerrecommendationsResource, c.ns, name), &v1alpha1.VerticalPodAutoscalerRecommendation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VerticalPodAutoscalerRecommendation), err
}

// List takes label and field selectors, and returns the list of VerticalPodAutoscalerRecommendations that match those selectors.
func (c *FakeVerticalPodAutoscalerRecommendations) List(opts v1.ListOptions) (result *v1alpha1.VerticalPodAutoscalerRecommendationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(verticalpodautoscalerrecommendationsResource, verticalpodautoscalerrecommendationsKind, c.ns, opts), &v1alpha1.VerticalPodAutoscalerRecommendationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VerticalPodAutoscalerRecommendationList{ListMeta: obj.(*v1alpha1.VerticalPodAutoscalerRecommendationList).ListMeta}
	for _, item := range obj.(*v1alpha1.VerticalPodAutoscalerRecommendationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested verticalPodAutoscalerRecommendations.
func (c *FakeVerticalPodAutoscalerRecommendations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(verticalpodautoscalerrecommendationsResource, c.ns, opts))

}

// Create takes the representation of a verticalPodAutoscalerRecommendation and creates it.  Returns the server's representation of the verticalPodAutoscalerRecommendation, and an error, if there is any.
func (c *FakeVerticalPodAutoscalerRecommendations) Create(verticalPodAutoscalerRecommendation *v1alpha1.VerticalPodAutoscalerRecommendation) (result *v1alpha1.VerticalPodAutoscalerRecommendation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(verticalpodautoscalerrecommendationsResource, c.ns, verticalPodAutoscalerRecommendation), &v1alpha1.VerticalPodAutoscalerRecommendation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VerticalPodAutoscalerRecommendation), err
}

// Update takes the representation of a verticalPodAutoscalerRecommendation and updates it. Returns the server's representation of the verticalPodAutoscalerRecommendation, and an error, if there is any.
func (c *FakeVerticalPodAutoscalerRecommendations) Update(verticalPodAutoscalerRecommendation *v1alpha1.VerticalPodAutoscalerRecommendation) (result *v1alpha1.VerticalPodAutoscalerRecommendation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(verticalpodautoscalerrecommendationsResource, c.ns, verticalPodAutoscalerRecommendation), &v1alpha1.VerticalPodAutoscalerRecommendation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VerticalPodAutoscalerRecommendation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVerticalPodAutoscalerRecommendations) UpdateStatus(verticalPodAutoscalerRecommendation *v1alpha1.VerticalPodAutoscalerRecommendation) (*v1alpha1.VerticalPodAutoscalerRecommendation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(verticalpodautoscalerrecommendationsResource, "status", c.ns, verticalPodAutoscalerRecommendation), &v1alpha1.VerticalPodAutoscalerRecommendation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VerticalPodAutoscalerRecommendation), err
}

// Delete takes name of the verticalPodAutoscalerRecommendation and deletes it. Returns an error if one occurs.
func (c *FakeVerticalPodAutoscalerRecommendations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(verticalpodautoscalerrecommendationsResource, c.ns, name), &v1alpha1.VerticalPodAutoscalerRecommendation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVerticalPodAutoscalerRecommendations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(verticalpodautoscalerrecommendationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.VerticalPodAutoscalerRecommendationList{})
	return err
}

// Patch applies the patch and returns the patched verticalPodAutoscalerRecommendation.
func (c *FakeVerticalPodAutoscalerRecommendations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VerticalPodAutoscalerRecommendation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(verticalpodautoscalerrecommendationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.VerticalPodAutoscalerRecommendation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VerticalPodAutoscalerRecommendation), err
}
