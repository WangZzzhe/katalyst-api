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

package versioned

import (
	"fmt"

	autoscalingv1alpha1 "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned/typed/autoscaling/v1alpha1"
	autoscalingv1alpha2 "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned/typed/autoscaling/v1alpha2"
	configv1alpha1 "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned/typed/config/v1alpha1"
	nodev1alpha1 "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned/typed/node/v1alpha1"
	overcommitv1alpha1 "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned/typed/overcommit/v1alpha1"
	recommendationv1alpha1 "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned/typed/recommendation/v1alpha1"
	workloadv1alpha1 "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned/typed/workload/v1alpha1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	AutoscalingV1alpha1() autoscalingv1alpha1.AutoscalingV1alpha1Interface
	AutoscalingV1alpha2() autoscalingv1alpha2.AutoscalingV1alpha2Interface
	ConfigV1alpha1() configv1alpha1.ConfigV1alpha1Interface
	NodeV1alpha1() nodev1alpha1.NodeV1alpha1Interface
	OvercommitV1alpha1() overcommitv1alpha1.OvercommitV1alpha1Interface
	RecommendationV1alpha1() recommendationv1alpha1.RecommendationV1alpha1Interface
	WorkloadV1alpha1() workloadv1alpha1.WorkloadV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	autoscalingV1alpha1    *autoscalingv1alpha1.AutoscalingV1alpha1Client
	autoscalingV1alpha2    *autoscalingv1alpha2.AutoscalingV1alpha2Client
	configV1alpha1         *configv1alpha1.ConfigV1alpha1Client
	nodeV1alpha1           *nodev1alpha1.NodeV1alpha1Client
	overcommitV1alpha1     *overcommitv1alpha1.OvercommitV1alpha1Client
	recommendationV1alpha1 *recommendationv1alpha1.RecommendationV1alpha1Client
	workloadV1alpha1       *workloadv1alpha1.WorkloadV1alpha1Client
}

// AutoscalingV1alpha1 retrieves the AutoscalingV1alpha1Client
func (c *Clientset) AutoscalingV1alpha1() autoscalingv1alpha1.AutoscalingV1alpha1Interface {
	return c.autoscalingV1alpha1
}

// AutoscalingV1alpha2 retrieves the AutoscalingV1alpha2Client
func (c *Clientset) AutoscalingV1alpha2() autoscalingv1alpha2.AutoscalingV1alpha2Interface {
	return c.autoscalingV1alpha2
}

// ConfigV1alpha1 retrieves the ConfigV1alpha1Client
func (c *Clientset) ConfigV1alpha1() configv1alpha1.ConfigV1alpha1Interface {
	return c.configV1alpha1
}

// NodeV1alpha1 retrieves the NodeV1alpha1Client
func (c *Clientset) NodeV1alpha1() nodev1alpha1.NodeV1alpha1Interface {
	return c.nodeV1alpha1
}

// OvercommitV1alpha1 retrieves the OvercommitV1alpha1Client
func (c *Clientset) OvercommitV1alpha1() overcommitv1alpha1.OvercommitV1alpha1Interface {
	return c.overcommitV1alpha1
}

// RecommendationV1alpha1 retrieves the RecommendationV1alpha1Client
func (c *Clientset) RecommendationV1alpha1() recommendationv1alpha1.RecommendationV1alpha1Interface {
	return c.recommendationV1alpha1
}

// WorkloadV1alpha1 retrieves the WorkloadV1alpha1Client
func (c *Clientset) WorkloadV1alpha1() workloadv1alpha1.WorkloadV1alpha1Interface {
	return c.workloadV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("Burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.autoscalingV1alpha1, err = autoscalingv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.autoscalingV1alpha2, err = autoscalingv1alpha2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.configV1alpha1, err = configv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.nodeV1alpha1, err = nodev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.overcommitV1alpha1, err = overcommitv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.recommendationV1alpha1, err = recommendationv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.workloadV1alpha1, err = workloadv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.autoscalingV1alpha1 = autoscalingv1alpha1.NewForConfigOrDie(c)
	cs.autoscalingV1alpha2 = autoscalingv1alpha2.NewForConfigOrDie(c)
	cs.configV1alpha1 = configv1alpha1.NewForConfigOrDie(c)
	cs.nodeV1alpha1 = nodev1alpha1.NewForConfigOrDie(c)
	cs.overcommitV1alpha1 = overcommitv1alpha1.NewForConfigOrDie(c)
	cs.recommendationV1alpha1 = recommendationv1alpha1.NewForConfigOrDie(c)
	cs.workloadV1alpha1 = workloadv1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.autoscalingV1alpha1 = autoscalingv1alpha1.New(c)
	cs.autoscalingV1alpha2 = autoscalingv1alpha2.New(c)
	cs.configV1alpha1 = configv1alpha1.New(c)
	cs.nodeV1alpha1 = nodev1alpha1.New(c)
	cs.overcommitV1alpha1 = overcommitv1alpha1.New(c)
	cs.recommendationV1alpha1 = recommendationv1alpha1.New(c)
	cs.workloadV1alpha1 = workloadv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
