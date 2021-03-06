/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AutoscalerObservation struct {
}

type AutoscalerParameters struct {

	// The maximum number of nodes to autoscale to.
	// +kubebuilder:validation:Required
	Max *int64 `json:"max" tf:"max,omitempty"`

	// The minimum number of nodes to autoscale to.
	// +kubebuilder:validation:Required
	Min *int64 `json:"min" tf:"min,omitempty"`
}

type ClusterObservation struct {
	APIEndpoints []*string `json:"apiEndpoints,omitempty" tf:"api_endpoints,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ClusterParameters struct {

	// Defines settings for the Kubernetes Control Plane.
	// +kubebuilder:validation:Optional
	ControlPlane []ControlPlaneParameters `json:"controlPlane,omitempty" tf:"control_plane,omitempty"`

	// The desired Kubernetes version for this Kubernetes cluster in the format of <major>.<minor>. The latest supported patch version will be deployed.
	// +kubebuilder:validation:Required
	K8SVersion *string `json:"k8sVersion" tf:"k8s_version,omitempty"`

	// The unique label for the cluster.
	// +kubebuilder:validation:Required
	Label *string `json:"label" tf:"label,omitempty"`

	// A node pool in the cluster.
	// +kubebuilder:validation:Required
	Pool []PoolParameters `json:"pool" tf:"pool,omitempty"`

	// This cluster's location.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// An array of tags applied to this object. Tags are for organizational purposes only.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ControlPlaneObservation struct {
}

type ControlPlaneParameters struct {

	// Defines whether High Availability is enabled for the Control Plane Components of the cluster.
	// +kubebuilder:validation:Optional
	HighAvailability *bool `json:"highAvailability,omitempty" tf:"high_availability,omitempty"`
}

type NodesObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InstanceID *int64 `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type NodesParameters struct {
}

type PoolObservation struct {
	ID *int64 `json:"id,omitempty" tf:"id,omitempty"`

	Nodes []NodesObservation `json:"nodes,omitempty" tf:"nodes,omitempty"`
}

type PoolParameters struct {

	// When specified, the number of nodes autoscales within the defined minimum and maximum values.
	// +kubebuilder:validation:Optional
	Autoscaler []AutoscalerParameters `json:"autoscaler,omitempty" tf:"autoscaler,omitempty"`

	// The number of nodes in the Node Pool.
	// +kubebuilder:validation:Required
	Count *int64 `json:"count" tf:"count,omitempty"`

	// A Linode Type for all of the nodes in the Node Pool.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,linodejet}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
