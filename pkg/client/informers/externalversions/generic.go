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

package informers

import (
	"fmt"

	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v2"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/cache"

	apiresourcev1alpha1 "github.com/kcp-dev/kcp/pkg/apis/apiresource/v1alpha1"
	apisv1alpha1 "github.com/kcp-dev/kcp/pkg/apis/apis/v1alpha1"
	schedulingv1alpha1 "github.com/kcp-dev/kcp/pkg/apis/scheduling/v1alpha1"
	tenancyv1alpha1 "github.com/kcp-dev/kcp/pkg/apis/tenancy/v1alpha1"
	tenancyv1beta1 "github.com/kcp-dev/kcp/pkg/apis/tenancy/v1beta1"
	topologyv1alpha1 "github.com/kcp-dev/kcp/pkg/apis/topology/v1alpha1"
	workloadv1alpha1 "github.com/kcp-dev/kcp/pkg/apis/workload/v1alpha1"
)

type GenericClusterInformer interface {
	Cluster(logicalcluster.Name) GenericInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() kcpcache.GenericClusterLister
}

type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericClusterInformer struct {
	informer kcpcache.ScopeableSharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.informer
}

// Lister returns the GenericClusterLister.
func (f *genericClusterInformer) Lister() kcpcache.GenericClusterLister {
	return kcpcache.NewGenericClusterLister(f.Informer().GetIndexer(), f.resource)
}

// Cluster scopes to a GenericInformer.
func (f *genericClusterInformer) Cluster(cluster logicalcluster.Name) GenericInformer {
	return &genericInformer{
		informer: f.Informer().Cluster(cluster),
		lister:   f.Lister().ByCluster(cluster),
	}
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	lister   cache.GenericLister
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return f.lister
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericClusterInformer, error) {
	switch resource {
	// Group=apiresource.kcp.dev, Version=V1alpha1
	case apiresourcev1alpha1.SchemeGroupVersion.WithResource("apiresourceimports"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Apiresource().V1alpha1().APIResourceImports().Informer()}, nil
	case apiresourcev1alpha1.SchemeGroupVersion.WithResource("negotiatedapiresources"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Apiresource().V1alpha1().NegotiatedAPIResources().Informer()}, nil
	// Group=apis.kcp.dev, Version=V1alpha1
	case apisv1alpha1.SchemeGroupVersion.WithResource("apibindings"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Apis().V1alpha1().APIBindings().Informer()}, nil
	case apisv1alpha1.SchemeGroupVersion.WithResource("apiexports"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Apis().V1alpha1().APIExports().Informer()}, nil
	case apisv1alpha1.SchemeGroupVersion.WithResource("apiexportendpointslices"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Apis().V1alpha1().APIExportEndpointSlices().Informer()}, nil
	case apisv1alpha1.SchemeGroupVersion.WithResource("apiresourceschemas"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Apis().V1alpha1().APIResourceSchemas().Informer()}, nil
	// Group=scheduling.kcp.dev, Version=V1alpha1
	case schedulingv1alpha1.SchemeGroupVersion.WithResource("locations"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Scheduling().V1alpha1().Locations().Informer()}, nil
	case schedulingv1alpha1.SchemeGroupVersion.WithResource("placements"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Scheduling().V1alpha1().Placements().Informer()}, nil
	// Group=tenancy.kcp.dev, Version=V1alpha1
	case tenancyv1alpha1.SchemeGroupVersion.WithResource("clusterworkspaces"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Tenancy().V1alpha1().ClusterWorkspaces().Informer()}, nil
	case tenancyv1alpha1.SchemeGroupVersion.WithResource("clusterworkspacetypes"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Tenancy().V1alpha1().ClusterWorkspaceTypes().Informer()}, nil
	case tenancyv1alpha1.SchemeGroupVersion.WithResource("clusterworkspaceshards"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Tenancy().V1alpha1().ClusterWorkspaceShards().Informer()}, nil
	// Group=tenancy.kcp.dev, Version=V1beta1
	case tenancyv1beta1.SchemeGroupVersion.WithResource("workspaces"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Tenancy().V1beta1().Workspaces().Informer()}, nil
	// Group=topology.kcp.dev, Version=V1alpha1
	case topologyv1alpha1.SchemeGroupVersion.WithResource("partitions"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Topology().V1alpha1().Partitions().Informer()}, nil
	case topologyv1alpha1.SchemeGroupVersion.WithResource("partitionsets"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Topology().V1alpha1().PartitionSets().Informer()}, nil
	// Group=workload.kcp.dev, Version=V1alpha1
	case workloadv1alpha1.SchemeGroupVersion.WithResource("synctargets"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Workload().V1alpha1().SyncTargets().Informer()}, nil
	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedScopedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=apiresource.kcp.dev, Version=V1alpha1
	case apiresourcev1alpha1.SchemeGroupVersion.WithResource("apiresourceimports"):
		informer := f.Apiresource().V1alpha1().APIResourceImports().Informer()
		return &genericInformer{lister: cache.NewGenericLister(informer.GetIndexer(), resource.GroupResource()), informer: informer}, nil
	case apiresourcev1alpha1.SchemeGroupVersion.WithResource("negotiatedapiresources"):
		informer := f.Apiresource().V1alpha1().NegotiatedAPIResources().Informer()
		return &genericInformer{lister: cache.NewGenericLister(informer.GetIndexer(), resource.GroupResource()), informer: informer}, nil
	// Group=apis.kcp.dev, Version=V1alpha1
	case apisv1alpha1.SchemeGroupVersion.WithResource("apibindings"):
		informer := f.Apis().V1alpha1().APIBindings().Informer()
		return &genericInformer{lister: cache.NewGenericLister(informer.GetIndexer(), resource.GroupResource()), informer: informer}, nil
	case apisv1alpha1.SchemeGroupVersion.WithResource("apiexports"):
		informer := f.Apis().V1alpha1().APIExports().Informer()
		return &genericInformer{lister: cache.NewGenericLister(informer.GetIndexer(), resource.GroupResource()), informer: informer}, nil
	case apisv1alpha1.SchemeGroupVersion.WithResource("apiexportendpointslices"):
		informer := f.Apis().V1alpha1().APIExportEndpointSlices().Informer()
		return &genericInformer{lister: cache.NewGenericLister(informer.GetIndexer(), resource.GroupResource()), informer: informer}, nil
	case apisv1alpha1.SchemeGroupVersion.WithResource("apiresourceschemas"):
		informer := f.Apis().V1alpha1().APIResourceSchemas().Informer()
		return &genericInformer{lister: cache.NewGenericLister(informer.GetIndexer(), resource.GroupResource()), informer: informer}, nil
	// Group=scheduling.kcp.dev, Version=V1alpha1
	case schedulingv1alpha1.SchemeGroupVersion.WithResource("locations"):
		informer := f.Scheduling().V1alpha1().Locations().Informer()
		return &genericInformer{lister: cache.NewGenericLister(informer.GetIndexer(), resource.GroupResource()), informer: informer}, nil
	case schedulingv1alpha1.SchemeGroupVersion.WithResource("placements"):
		informer := f.Scheduling().V1alpha1().Placements().Informer()
		return &genericInformer{lister: cache.NewGenericLister(informer.GetIndexer(), resource.GroupResource()), informer: informer}, nil
	// Group=tenancy.kcp.dev, Version=V1alpha1
	case tenancyv1alpha1.SchemeGroupVersion.WithResource("clusterworkspaces"):
		informer := f.Tenancy().V1alpha1().ClusterWorkspaces().Informer()
		return &genericInformer{lister: cache.NewGenericLister(informer.GetIndexer(), resource.GroupResource()), informer: informer}, nil
	case tenancyv1alpha1.SchemeGroupVersion.WithResource("clusterworkspacetypes"):
		informer := f.Tenancy().V1alpha1().ClusterWorkspaceTypes().Informer()
		return &genericInformer{lister: cache.NewGenericLister(informer.GetIndexer(), resource.GroupResource()), informer: informer}, nil
	case tenancyv1alpha1.SchemeGroupVersion.WithResource("clusterworkspaceshards"):
		informer := f.Tenancy().V1alpha1().ClusterWorkspaceShards().Informer()
		return &genericInformer{lister: cache.NewGenericLister(informer.GetIndexer(), resource.GroupResource()), informer: informer}, nil
	// Group=tenancy.kcp.dev, Version=V1beta1
	case tenancyv1beta1.SchemeGroupVersion.WithResource("workspaces"):
		informer := f.Tenancy().V1beta1().Workspaces().Informer()
		return &genericInformer{lister: cache.NewGenericLister(informer.GetIndexer(), resource.GroupResource()), informer: informer}, nil
	// Group=topology.kcp.dev, Version=V1alpha1
	case topologyv1alpha1.SchemeGroupVersion.WithResource("partitions"):
		informer := f.Topology().V1alpha1().Partitions().Informer()
		return &genericInformer{lister: cache.NewGenericLister(informer.GetIndexer(), resource.GroupResource()), informer: informer}, nil
	case topologyv1alpha1.SchemeGroupVersion.WithResource("partitionsets"):
		informer := f.Topology().V1alpha1().PartitionSets().Informer()
		return &genericInformer{lister: cache.NewGenericLister(informer.GetIndexer(), resource.GroupResource()), informer: informer}, nil
	// Group=workload.kcp.dev, Version=V1alpha1
	case workloadv1alpha1.SchemeGroupVersion.WithResource("synctargets"):
		informer := f.Workload().V1alpha1().SyncTargets().Informer()
		return &genericInformer{lister: cache.NewGenericLister(informer.GetIndexer(), resource.GroupResource()), informer: informer}, nil
	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
