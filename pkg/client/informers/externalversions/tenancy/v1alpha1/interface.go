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
	"github.com/kcp-dev/kcp/pkg/client/informers/externalversions/internalinterfaces"
)

type ClusterInterface interface {
	// ClusterWorkspaces returns a ClusterWorkspaceClusterInformer
	ClusterWorkspaces() ClusterWorkspaceClusterInformer
	// ClusterWorkspaceTypes returns a ClusterWorkspaceTypeClusterInformer
	ClusterWorkspaceTypes() ClusterWorkspaceTypeClusterInformer
	// ClusterWorkspaceShards returns a ClusterWorkspaceShardClusterInformer
	ClusterWorkspaceShards() ClusterWorkspaceShardClusterInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new ClusterInterface.
func New(f internalinterfaces.SharedInformerFactory, tweakListOptions internalinterfaces.TweakListOptionsFunc) ClusterInterface {
	return &version{factory: f, tweakListOptions: tweakListOptions}
}

// ClusterWorkspaces returns a ClusterWorkspaceClusterInformer
func (v *version) ClusterWorkspaces() ClusterWorkspaceClusterInformer {
	return &clusterWorkspaceClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterWorkspaceTypes returns a ClusterWorkspaceTypeClusterInformer
func (v *version) ClusterWorkspaceTypes() ClusterWorkspaceTypeClusterInformer {
	return &clusterWorkspaceTypeClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterWorkspaceShards returns a ClusterWorkspaceShardClusterInformer
func (v *version) ClusterWorkspaceShards() ClusterWorkspaceShardClusterInformer {
	return &clusterWorkspaceShardClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

type Interface interface {
	// ClusterWorkspaces returns a ClusterWorkspaceInformer
	ClusterWorkspaces() ClusterWorkspaceInformer
	// ClusterWorkspaceTypes returns a ClusterWorkspaceTypeInformer
	ClusterWorkspaceTypes() ClusterWorkspaceTypeInformer
	// ClusterWorkspaceShards returns a ClusterWorkspaceShardInformer
	ClusterWorkspaceShards() ClusterWorkspaceShardInformer
}

type scopedVersion struct {
	factory          internalinterfaces.SharedScopedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// New returns a new ClusterInterface.
func NewScoped(f internalinterfaces.SharedScopedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &scopedVersion{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ClusterWorkspaces returns a ClusterWorkspaceInformer
func (v *scopedVersion) ClusterWorkspaces() ClusterWorkspaceInformer {
	return &clusterWorkspaceScopedInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterWorkspaceTypes returns a ClusterWorkspaceTypeInformer
func (v *scopedVersion) ClusterWorkspaceTypes() ClusterWorkspaceTypeInformer {
	return &clusterWorkspaceTypeScopedInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterWorkspaceShards returns a ClusterWorkspaceShardInformer
func (v *scopedVersion) ClusterWorkspaceShards() ClusterWorkspaceShardInformer {
	return &clusterWorkspaceShardScopedInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
