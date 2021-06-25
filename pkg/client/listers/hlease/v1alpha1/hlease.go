/*
Copyright The Kubernetes Authors.

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
	v1alpha1 "harmonycloud.cn/common/hlease/pkg/apis/hlease/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HLeaseLister helps list HLeases.
type HLeaseLister interface {
	// List lists all HLeases in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.HLease, err error)
	// HLeases returns an object that can list and get HLeases.
	HLeases(namespace string) HLeaseNamespaceLister
	HLeaseListerExpansion
}

// hLeaseLister implements the HLeaseLister interface.
type hLeaseLister struct {
	indexer cache.Indexer
}

// NewHLeaseLister returns a new HLeaseLister.
func NewHLeaseLister(indexer cache.Indexer) HLeaseLister {
	return &hLeaseLister{indexer: indexer}
}

// List lists all HLeases in the indexer.
func (s *hLeaseLister) List(selector labels.Selector) (ret []*v1alpha1.HLease, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HLease))
	})
	return ret, err
}

// HLeases returns an object that can list and get HLeases.
func (s *hLeaseLister) HLeases(namespace string) HLeaseNamespaceLister {
	return hLeaseNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HLeaseNamespaceLister helps list and get HLeases.
type HLeaseNamespaceLister interface {
	// List lists all HLeases in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.HLease, err error)
	// Get retrieves the HLease from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.HLease, error)
	HLeaseNamespaceListerExpansion
}

// hLeaseNamespaceLister implements the HLeaseNamespaceLister
// interface.
type hLeaseNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HLeases in the indexer for a given namespace.
func (s hLeaseNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.HLease, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HLease))
	})
	return ret, err
}

// Get retrieves the HLease from the indexer for a given namespace and name.
func (s hLeaseNamespaceLister) Get(name string) (*v1alpha1.HLease, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("hlease"), name)
	}
	return obj.(*v1alpha1.HLease), nil
}