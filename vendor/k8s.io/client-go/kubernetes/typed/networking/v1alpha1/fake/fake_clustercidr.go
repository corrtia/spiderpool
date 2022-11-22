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
	"context"
	json "encoding/json"
	"fmt"

	v1alpha1 "k8s.io/api/networking/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	networkingv1alpha1 "k8s.io/client-go/applyconfigurations/networking/v1alpha1"
	testing "k8s.io/client-go/testing"
)

// FakeClusterCIDRs implements ClusterCIDRInterface
type FakeClusterCIDRs struct {
	Fake *FakeNetworkingV1alpha1
}

var clustercidrsResource = schema.GroupVersionResource{Group: "networking.k8s.io", Version: "v1alpha1", Resource: "clustercidrs"}

var clustercidrsKind = schema.GroupVersionKind{Group: "networking.k8s.io", Version: "v1alpha1", Kind: "ClusterCIDR"}

// Get takes name of the clusterCIDR, and returns the corresponding clusterCIDR object, and an error if there is any.
func (c *FakeClusterCIDRs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterCIDR, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clustercidrsResource, name), &v1alpha1.ClusterCIDR{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterCIDR), err
}

// List takes label and field selectors, and returns the list of ClusterCIDRs that match those selectors.
func (c *FakeClusterCIDRs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterCIDRList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clustercidrsResource, clustercidrsKind, opts), &v1alpha1.ClusterCIDRList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterCIDRList{ListMeta: obj.(*v1alpha1.ClusterCIDRList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterCIDRList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterCIDRs.
func (c *FakeClusterCIDRs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clustercidrsResource, opts))
}

// Create takes the representation of a clusterCIDR and creates it.  Returns the server's representation of the clusterCIDR, and an error, if there is any.
func (c *FakeClusterCIDRs) Create(ctx context.Context, clusterCIDR *v1alpha1.ClusterCIDR, opts v1.CreateOptions) (result *v1alpha1.ClusterCIDR, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clustercidrsResource, clusterCIDR), &v1alpha1.ClusterCIDR{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterCIDR), err
}

// Update takes the representation of a clusterCIDR and updates it. Returns the server's representation of the clusterCIDR, and an error, if there is any.
func (c *FakeClusterCIDRs) Update(ctx context.Context, clusterCIDR *v1alpha1.ClusterCIDR, opts v1.UpdateOptions) (result *v1alpha1.ClusterCIDR, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clustercidrsResource, clusterCIDR), &v1alpha1.ClusterCIDR{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterCIDR), err
}

// Delete takes name of the clusterCIDR and deletes it. Returns an error if one occurs.
func (c *FakeClusterCIDRs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clustercidrsResource, name, opts), &v1alpha1.ClusterCIDR{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterCIDRs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clustercidrsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterCIDRList{})
	return err
}

// Patch applies the patch and returns the patched clusterCIDR.
func (c *FakeClusterCIDRs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterCIDR, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clustercidrsResource, name, pt, data, subresources...), &v1alpha1.ClusterCIDR{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterCIDR), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied clusterCIDR.
func (c *FakeClusterCIDRs) Apply(ctx context.Context, clusterCIDR *networkingv1alpha1.ClusterCIDRApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ClusterCIDR, err error) {
	if clusterCIDR == nil {
		return nil, fmt.Errorf("clusterCIDR provided to Apply must not be nil")
	}
	data, err := json.Marshal(clusterCIDR)
	if err != nil {
		return nil, err
	}
	name := clusterCIDR.Name
	if name == nil {
		return nil, fmt.Errorf("clusterCIDR.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clustercidrsResource, *name, types.ApplyPatchType, data), &v1alpha1.ClusterCIDR{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterCIDR), err
}
