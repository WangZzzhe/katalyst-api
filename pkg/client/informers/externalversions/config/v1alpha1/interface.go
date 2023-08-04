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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/kubewharf/katalyst-api/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// CustomNodeConfigs returns a CustomNodeConfigInformer.
	CustomNodeConfigs() CustomNodeConfigInformer
	// KatalystCustomConfigs returns a KatalystCustomConfigInformer.
	KatalystCustomConfigs() KatalystCustomConfigInformer
	// NodeOvercommitConfigs returns a NodeOvercommitConfigInformer.
	NodeOvercommitConfigs() NodeOvercommitConfigInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// CustomNodeConfigs returns a CustomNodeConfigInformer.
func (v *version) CustomNodeConfigs() CustomNodeConfigInformer {
	return &customNodeConfigInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// KatalystCustomConfigs returns a KatalystCustomConfigInformer.
func (v *version) KatalystCustomConfigs() KatalystCustomConfigInformer {
	return &katalystCustomConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NodeOvercommitConfigs returns a NodeOvercommitConfigInformer.
func (v *version) NodeOvercommitConfigs() NodeOvercommitConfigInformer {
	return &nodeOvercommitConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
