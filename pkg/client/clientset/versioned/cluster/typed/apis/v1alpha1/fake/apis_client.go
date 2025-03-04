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
	"github.com/kcp-dev/logicalcluster/v2"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	"k8s.io/client-go/rest"

	kcpapisv1alpha1 "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/cluster/typed/apis/v1alpha1"
	apisv1alpha1 "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/apis/v1alpha1"
)

var _ kcpapisv1alpha1.ApisV1alpha1ClusterInterface = (*ApisV1alpha1ClusterClient)(nil)

type ApisV1alpha1ClusterClient struct {
	*kcptesting.Fake
}

func (c *ApisV1alpha1ClusterClient) Cluster(cluster logicalcluster.Name) apisv1alpha1.ApisV1alpha1Interface {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &ApisV1alpha1Client{Fake: c.Fake, Cluster: cluster}
}

func (c *ApisV1alpha1ClusterClient) APIBindings() kcpapisv1alpha1.APIBindingClusterInterface {
	return &aPIBindingsClusterClient{Fake: c.Fake}
}

func (c *ApisV1alpha1ClusterClient) APIExports() kcpapisv1alpha1.APIExportClusterInterface {
	return &aPIExportsClusterClient{Fake: c.Fake}
}

func (c *ApisV1alpha1ClusterClient) APIExportEndpointSlices() kcpapisv1alpha1.APIExportEndpointSliceClusterInterface {
	return &aPIExportEndpointSlicesClusterClient{Fake: c.Fake}
}

func (c *ApisV1alpha1ClusterClient) APIResourceSchemas() kcpapisv1alpha1.APIResourceSchemaClusterInterface {
	return &aPIResourceSchemasClusterClient{Fake: c.Fake}
}

var _ apisv1alpha1.ApisV1alpha1Interface = (*ApisV1alpha1Client)(nil)

type ApisV1alpha1Client struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
}

func (c *ApisV1alpha1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *ApisV1alpha1Client) APIBindings() apisv1alpha1.APIBindingInterface {
	return &aPIBindingsClient{Fake: c.Fake, Cluster: c.Cluster}
}

func (c *ApisV1alpha1Client) APIExports() apisv1alpha1.APIExportInterface {
	return &aPIExportsClient{Fake: c.Fake, Cluster: c.Cluster}
}

func (c *ApisV1alpha1Client) APIExportEndpointSlices() apisv1alpha1.APIExportEndpointSliceInterface {
	return &aPIExportEndpointSlicesClient{Fake: c.Fake, Cluster: c.Cluster}
}

func (c *ApisV1alpha1Client) APIResourceSchemas() apisv1alpha1.APIResourceSchemaInterface {
	return &aPIResourceSchemasClient{Fake: c.Fake, Cluster: c.Cluster}
}
