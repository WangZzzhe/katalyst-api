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
	time "time"

	nodev1alpha1 "github.com/kubewharf/katalyst-api/pkg/apis/node/v1alpha1"
	versioned "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kubewharf/katalyst-api/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kubewharf/katalyst-api/pkg/client/listers/node/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CustomNodeResourceInformer provides access to a shared informer and lister for
// CustomNodeResources.
type CustomNodeResourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CustomNodeResourceLister
}

type customNodeResourceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCustomNodeResourceInformer constructs a new informer for CustomNodeResource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCustomNodeResourceInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCustomNodeResourceInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredCustomNodeResourceInformer constructs a new informer for CustomNodeResource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCustomNodeResourceInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NodeV1alpha1().CustomNodeResources().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NodeV1alpha1().CustomNodeResources().Watch(options)
			},
		},
		&nodev1alpha1.CustomNodeResource{},
		resyncPeriod,
		indexers,
	)
}

func (f *customNodeResourceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCustomNodeResourceInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *customNodeResourceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&nodev1alpha1.CustomNodeResource{}, f.defaultInformer)
}

func (f *customNodeResourceInformer) Lister() v1alpha1.CustomNodeResourceLister {
	return v1alpha1.NewCustomNodeResourceLister(f.Informer().GetIndexer())
}
