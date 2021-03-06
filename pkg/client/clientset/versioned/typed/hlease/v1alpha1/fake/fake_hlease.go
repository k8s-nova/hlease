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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/k8s-nova/hlease/pkg/apis/hlease/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHLeases implements HLeaseInterface
type FakeHLeases struct {
	Fake *FakeIsolateV1alpha1
	ns   string
}

var hleasesResource = schema.GroupVersionResource{Group: "isolate.harmonycloud.cn", Version: "v1alpha1", Resource: "hleases"}

var hleasesKind = schema.GroupVersionKind{Group: "isolate.harmonycloud.cn", Version: "v1alpha1", Kind: "HLease"}

// Get takes name of the hLease, and returns the corresponding hLease object, and an error if there is any.
func (c *FakeHLeases) Get(name string, options v1.GetOptions) (result *v1alpha1.HLease, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hleasesResource, c.ns, name), &v1alpha1.HLease{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HLease), err
}

// List takes label and field selectors, and returns the list of HLeases that match those selectors.
func (c *FakeHLeases) List(opts v1.ListOptions) (result *v1alpha1.HLeaseList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hleasesResource, hleasesKind, c.ns, opts), &v1alpha1.HLeaseList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HLeaseList{ListMeta: obj.(*v1alpha1.HLeaseList).ListMeta}
	for _, item := range obj.(*v1alpha1.HLeaseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hLeases.
func (c *FakeHLeases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hleasesResource, c.ns, opts))

}

// Create takes the representation of a hLease and creates it.  Returns the server's representation of the hLease, and an error, if there is any.
func (c *FakeHLeases) Create(hLease *v1alpha1.HLease) (result *v1alpha1.HLease, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hleasesResource, c.ns, hLease), &v1alpha1.HLease{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HLease), err
}

// Update takes the representation of a hLease and updates it. Returns the server's representation of the hLease, and an error, if there is any.
func (c *FakeHLeases) Update(hLease *v1alpha1.HLease) (result *v1alpha1.HLease, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hleasesResource, c.ns, hLease), &v1alpha1.HLease{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HLease), err
}

// Delete takes name of the hLease and deletes it. Returns an error if one occurs.
func (c *FakeHLeases) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(hleasesResource, c.ns, name), &v1alpha1.HLease{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHLeases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hleasesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.HLeaseList{})
	return err
}

// Patch applies the patch and returns the patched hLease.
func (c *FakeHLeases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.HLease, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hleasesResource, c.ns, name, pt, data, subresources...), &v1alpha1.HLease{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HLease), err
}
