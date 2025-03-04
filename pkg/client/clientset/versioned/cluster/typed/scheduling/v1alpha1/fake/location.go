//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1alpha1

import (
	"context"

	"github.com/kcp-dev/logicalcluster/v2"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/testing"

	schedulingv1alpha1 "github.com/kcp-dev/kcp/pkg/apis/scheduling/v1alpha1"
	schedulingv1alpha1client "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/scheduling/v1alpha1"
)

var locationsResource = schema.GroupVersionResource{Group: "scheduling.kcp.dev", Version: "v1alpha1", Resource: "locations"}
var locationsKind = schema.GroupVersionKind{Group: "scheduling.kcp.dev", Version: "v1alpha1", Kind: "Location"}

type locationsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *locationsClusterClient) Cluster(cluster logicalcluster.Name) schedulingv1alpha1client.LocationInterface {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &locationsClient{Fake: c.Fake, Cluster: cluster}
}

// List takes label and field selectors, and returns the list of Locations that match those selectors across all clusters.
func (c *locationsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*schedulingv1alpha1.LocationList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(locationsResource, locationsKind, logicalcluster.Wildcard, opts), &schedulingv1alpha1.LocationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &schedulingv1alpha1.LocationList{ListMeta: obj.(*schedulingv1alpha1.LocationList).ListMeta}
	for _, item := range obj.(*schedulingv1alpha1.LocationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested Locations across all clusters.
func (c *locationsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(locationsResource, logicalcluster.Wildcard, opts))
}

type locationsClient struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
}

func (c *locationsClient) Create(ctx context.Context, location *schedulingv1alpha1.Location, opts metav1.CreateOptions) (*schedulingv1alpha1.Location, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(locationsResource, c.Cluster, location), &schedulingv1alpha1.Location{})
	if obj == nil {
		return nil, err
	}
	return obj.(*schedulingv1alpha1.Location), err
}

func (c *locationsClient) Update(ctx context.Context, location *schedulingv1alpha1.Location, opts metav1.UpdateOptions) (*schedulingv1alpha1.Location, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(locationsResource, c.Cluster, location), &schedulingv1alpha1.Location{})
	if obj == nil {
		return nil, err
	}
	return obj.(*schedulingv1alpha1.Location), err
}

func (c *locationsClient) UpdateStatus(ctx context.Context, location *schedulingv1alpha1.Location, opts metav1.UpdateOptions) (*schedulingv1alpha1.Location, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(locationsResource, c.Cluster, "status", location), &schedulingv1alpha1.Location{})
	if obj == nil {
		return nil, err
	}
	return obj.(*schedulingv1alpha1.Location), err
}

func (c *locationsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(locationsResource, c.Cluster, name, opts), &schedulingv1alpha1.Location{})
	return err
}

func (c *locationsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(locationsResource, c.Cluster, listOpts)

	_, err := c.Fake.Invokes(action, &schedulingv1alpha1.LocationList{})
	return err
}

func (c *locationsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*schedulingv1alpha1.Location, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(locationsResource, c.Cluster, name), &schedulingv1alpha1.Location{})
	if obj == nil {
		return nil, err
	}
	return obj.(*schedulingv1alpha1.Location), err
}

// List takes label and field selectors, and returns the list of Locations that match those selectors.
func (c *locationsClient) List(ctx context.Context, opts metav1.ListOptions) (*schedulingv1alpha1.LocationList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(locationsResource, locationsKind, c.Cluster, opts), &schedulingv1alpha1.LocationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &schedulingv1alpha1.LocationList{ListMeta: obj.(*schedulingv1alpha1.LocationList).ListMeta}
	for _, item := range obj.(*schedulingv1alpha1.LocationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *locationsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(locationsResource, c.Cluster, opts))
}

func (c *locationsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*schedulingv1alpha1.Location, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(locationsResource, c.Cluster, name, pt, data, subresources...), &schedulingv1alpha1.Location{})
	if obj == nil {
		return nil, err
	}
	return obj.(*schedulingv1alpha1.Location), err
}
