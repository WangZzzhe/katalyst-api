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

	recommendationv1alpha1 "github.com/kubewharf/katalyst-api/pkg/apis/recommendation/v1alpha1"
	versioned "github.com/kubewharf/katalyst-api/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kubewharf/katalyst-api/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kubewharf/katalyst-api/pkg/client/listers/recommendation/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ResourceRecommendInformer provides access to a shared informer and lister for
// ResourceRecommends.
type ResourceRecommendInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ResourceRecommendLister
}

type resourceRecommendInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewResourceRecommendInformer constructs a new informer for ResourceRecommend type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewResourceRecommendInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredResourceRecommendInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredResourceRecommendInformer constructs a new informer for ResourceRecommend type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredResourceRecommendInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RecommendationV1alpha1().ResourceRecommends(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RecommendationV1alpha1().ResourceRecommends(namespace).Watch(options)
			},
		},
		&recommendationv1alpha1.ResourceRecommend{},
		resyncPeriod,
		indexers,
	)
}

func (f *resourceRecommendInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredResourceRecommendInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *resourceRecommendInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&recommendationv1alpha1.ResourceRecommend{}, f.defaultInformer)
}

func (f *resourceRecommendInformer) Lister() v1alpha1.ResourceRecommendLister {
	return v1alpha1.NewResourceRecommendLister(f.Informer().GetIndexer())
}
