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
	clientset "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned"
	autoscalingv1alpha1 "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned/typed/autoscaling/v1alpha1"
	fakeautoscalingv1alpha1 "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned/typed/autoscaling/v1alpha1/fake"
	autoscalingv1alpha2 "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned/typed/autoscaling/v1alpha2"
	fakeautoscalingv1alpha2 "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned/typed/autoscaling/v1alpha2/fake"
	configv1alpha1 "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned/typed/config/v1alpha1"
	fakeconfigv1alpha1 "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned/typed/config/v1alpha1/fake"
	nodev1alpha1 "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned/typed/node/v1alpha1"
	fakenodev1alpha1 "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned/typed/node/v1alpha1/fake"
	overcommitv1alpha1 "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned/typed/overcommit/v1alpha1"
	fakeovercommitv1alpha1 "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned/typed/overcommit/v1alpha1/fake"
	recommendationv1alpha1 "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned/typed/recommendation/v1alpha1"
	fakerecommendationv1alpha1 "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned/typed/recommendation/v1alpha1/fake"
	resourceportraitv1alpha1 "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned/typed/resourceportrait/v1alpha1"
	fakeresourceportraitv1alpha1 "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned/typed/resourceportrait/v1alpha1/fake"
	tidev1alpha1 "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned/typed/tide/v1alpha1"
	faketidev1alpha1 "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned/typed/tide/v1alpha1/fake"
	workloadv1alpha1 "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned/typed/workload/v1alpha1"
	fakeworkloadv1alpha1 "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned/typed/workload/v1alpha1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var (
	_ clientset.Interface = &Clientset{}
	_ testing.FakeClient  = &Clientset{}
)

// AutoscalingV1alpha1 retrieves the AutoscalingV1alpha1Client
func (c *Clientset) AutoscalingV1alpha1() autoscalingv1alpha1.AutoscalingV1alpha1Interface {
	return &fakeautoscalingv1alpha1.FakeAutoscalingV1alpha1{Fake: &c.Fake}
}

// AutoscalingV1alpha2 retrieves the AutoscalingV1alpha2Client
func (c *Clientset) AutoscalingV1alpha2() autoscalingv1alpha2.AutoscalingV1alpha2Interface {
	return &fakeautoscalingv1alpha2.FakeAutoscalingV1alpha2{Fake: &c.Fake}
}

// ConfigV1alpha1 retrieves the ConfigV1alpha1Client
func (c *Clientset) ConfigV1alpha1() configv1alpha1.ConfigV1alpha1Interface {
	return &fakeconfigv1alpha1.FakeConfigV1alpha1{Fake: &c.Fake}
}

// NodeV1alpha1 retrieves the NodeV1alpha1Client
func (c *Clientset) NodeV1alpha1() nodev1alpha1.NodeV1alpha1Interface {
	return &fakenodev1alpha1.FakeNodeV1alpha1{Fake: &c.Fake}
}

// OvercommitV1alpha1 retrieves the OvercommitV1alpha1Client
func (c *Clientset) OvercommitV1alpha1() overcommitv1alpha1.OvercommitV1alpha1Interface {
	return &fakeovercommitv1alpha1.FakeOvercommitV1alpha1{Fake: &c.Fake}
}

// RecommendationV1alpha1 retrieves the RecommendationV1alpha1Client
func (c *Clientset) RecommendationV1alpha1() recommendationv1alpha1.RecommendationV1alpha1Interface {
	return &fakerecommendationv1alpha1.FakeRecommendationV1alpha1{Fake: &c.Fake}
}

// ResourceportraitV1alpha1 retrieves the ResourceportraitV1alpha1Client
func (c *Clientset) ResourceportraitV1alpha1() resourceportraitv1alpha1.ResourceportraitV1alpha1Interface {
	return &fakeresourceportraitv1alpha1.FakeResourceportraitV1alpha1{Fake: &c.Fake}
}

// TideV1alpha1 retrieves the TideV1alpha1Client
func (c *Clientset) TideV1alpha1() tidev1alpha1.TideV1alpha1Interface {
	return &faketidev1alpha1.FakeTideV1alpha1{Fake: &c.Fake}
}

// WorkloadV1alpha1 retrieves the WorkloadV1alpha1Client
func (c *Clientset) WorkloadV1alpha1() workloadv1alpha1.WorkloadV1alpha1Interface {
	return &fakeworkloadv1alpha1.FakeWorkloadV1alpha1{Fake: &c.Fake}
}
