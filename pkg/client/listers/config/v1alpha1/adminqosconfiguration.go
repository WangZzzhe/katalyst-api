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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kubewharf/katalyst-api/pkg/apis/config/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AdminQoSConfigurationLister helps list AdminQoSConfigurations.
// All objects returned here must be treated as read-only.
type AdminQoSConfigurationLister interface {
	// List lists all AdminQoSConfigurations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AdminQoSConfiguration, err error)
	// AdminQoSConfigurations returns an object that can list and get AdminQoSConfigurations.
	AdminQoSConfigurations(namespace string) AdminQoSConfigurationNamespaceLister
	AdminQoSConfigurationListerExpansion
}

// adminQoSConfigurationLister implements the AdminQoSConfigurationLister interface.
type adminQoSConfigurationLister struct {
	indexer cache.Indexer
}

// NewAdminQoSConfigurationLister returns a new AdminQoSConfigurationLister.
func NewAdminQoSConfigurationLister(indexer cache.Indexer) AdminQoSConfigurationLister {
	return &adminQoSConfigurationLister{indexer: indexer}
}

// List lists all AdminQoSConfigurations in the indexer.
func (s *adminQoSConfigurationLister) List(selector labels.Selector) (ret []*v1alpha1.AdminQoSConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AdminQoSConfiguration))
	})
	return ret, err
}

// AdminQoSConfigurations returns an object that can list and get AdminQoSConfigurations.
func (s *adminQoSConfigurationLister) AdminQoSConfigurations(namespace string) AdminQoSConfigurationNamespaceLister {
	return adminQoSConfigurationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AdminQoSConfigurationNamespaceLister helps list and get AdminQoSConfigurations.
// All objects returned here must be treated as read-only.
type AdminQoSConfigurationNamespaceLister interface {
	// List lists all AdminQoSConfigurations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AdminQoSConfiguration, err error)
	// Get retrieves the AdminQoSConfiguration from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AdminQoSConfiguration, error)
	AdminQoSConfigurationNamespaceListerExpansion
}

// adminQoSConfigurationNamespaceLister implements the AdminQoSConfigurationNamespaceLister
// interface.
type adminQoSConfigurationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AdminQoSConfigurations in the indexer for a given namespace.
func (s adminQoSConfigurationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AdminQoSConfiguration, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AdminQoSConfiguration))
	})
	return ret, err
}

// Get retrieves the AdminQoSConfiguration from the indexer for a given namespace and name.
func (s adminQoSConfigurationNamespaceLister) Get(name string) (*v1alpha1.AdminQoSConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("adminqosconfiguration"), name)
	}
	return obj.(*v1alpha1.AdminQoSConfiguration), nil
}
