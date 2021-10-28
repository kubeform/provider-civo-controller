/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Cluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec,omitempty"`
	Status            ClusterStatus `json:"status,omitempty"`
}

type ClusterSpecInstalledApplications struct {
	// Name of application
	// +optional
	Application *string `json:"application,omitempty" tf:"application"`
	// Category of the application
	// +optional
	Category *string `json:"category,omitempty" tf:"category"`
	// Application installation status (`true` if installed)
	// +optional
	Installed *bool `json:"installed,omitempty" tf:"installed"`
	// Version of application
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
}

type ClusterSpecInstances struct {
	// Instance's CPU cores
	// +optional
	CpuCores *int64 `json:"cpuCores,omitempty" tf:"cpu_cores"`
	// Instance's disk (GB)
	// +optional
	DiskGb *int64 `json:"diskGb,omitempty" tf:"disk_gb"`
	// Instance's hostname
	// +optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname"`
	// Instance's RAM (MB)
	// +optional
	RamMb *int64 `json:"ramMb,omitempty" tf:"ram_mb"`
	// Instance's size
	// +optional
	Size *string `json:"size,omitempty" tf:"size"`
	// Instance's status
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// Instance's tags
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
}

type ClusterSpecPoolsInstances struct {
	// Instance's CPU cores
	// +optional
	CpuCores *int64 `json:"cpuCores,omitempty" tf:"cpu_cores"`
	// Instance's disk (GB)
	// +optional
	DiskGb *int64 `json:"diskGb,omitempty" tf:"disk_gb"`
	// Instance's hostname
	// +optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname"`
	// Instance's RAM (MB)
	// +optional
	RamMb *int64 `json:"ramMb,omitempty" tf:"ram_mb"`
	// Instance's size
	// +optional
	Size *string `json:"size,omitempty" tf:"size"`
	// Instance's status
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// Instance's tags
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
}

type ClusterSpecPools struct {
	// Number of nodes in the nodepool
	// +optional
	Count *int64 `json:"count,omitempty" tf:"count"`
	// Instance names in the nodepool
	// +optional
	InstanceNames []string `json:"instanceNames,omitempty" tf:"instance_names"`
	// +optional
	Instances []ClusterSpecPoolsInstances `json:"instances,omitempty" tf:"instances"`
	// Size of the nodes in the nodepool
	// +optional
	Size *string `json:"size,omitempty" tf:"size"`
}

type ClusterSpec struct {
	State *ClusterSpecResource `json:"state,omitempty" tf:"-"`

	Resource ClusterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ClusterSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The API server endpoint of the cluster
	// +optional
	ApiEndpoint *string `json:"apiEndpoint,omitempty" tf:"api_endpoint"`
	// Comma separated list of applications to install. Spaces within application names are fine, but shouldn't be either side of the comma. Application names are case-sensitive; the available applications can be listed with the Civo CLI: 'civo kubernetes applications ls'. If you want to remove a default installed application, prefix it with a '-', e.g. -Traefik. For application that supports plans, you can use 'app_name:app_plan' format e.g. 'Linkerd:Linkerd & Jaeger' or 'MariaDB:5GB'.
	// +optional
	Applications *string `json:"applications,omitempty" tf:"applications"`
	// The timestamp when the cluster was created
	// +optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at"`
	// The DNS name of the cluster
	// +optional
	DnsEntry *string `json:"dnsEntry,omitempty" tf:"dns_entry"`
	// The existing firewall ID to use for this cluster
	FirewallID *string `json:"firewallID" tf:"firewall_id"`
	// +optional
	InstalledApplications []ClusterSpecInstalledApplications `json:"installedApplications,omitempty" tf:"installed_applications"`
	// +optional
	Instances []ClusterSpecInstances `json:"instances,omitempty" tf:"instances"`
	// The kubeconfig of the cluster
	// +optional
	Kubeconfig *string `json:"-" sensitive:"true" tf:"kubeconfig"`
	// The version of k3s to install (optional, the default is currently the latest available)
	// +optional
	KubernetesVersion *string `json:"kubernetesVersion,omitempty" tf:"kubernetes_version"`
	// The IP address of the master node
	// +optional
	MasterIP *string `json:"masterIP,omitempty" tf:"master_ip"`
	// Name for your cluster, must be unique within your account
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// The network for the cluster, if not declare we use the default one
	// +optional
	NetworkID *string `json:"networkID,omitempty" tf:"network_id"`
	// The number of instances to create (optional, the default at the time of writing is 3)
	// +optional
	NumTargetNodes *int64 `json:"numTargetNodes,omitempty" tf:"num_target_nodes"`
	// +optional
	Pools []ClusterSpecPools `json:"pools,omitempty" tf:"pools"`
	// When cluster is ready, this will return `true`
	// +optional
	Ready *bool `json:"ready,omitempty" tf:"ready"`
	// The region for the cluster, if not declare we use the region in declared in the provider
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
	// Status of the cluster
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// Space separated list of tags, to be used freely as required
	// +optional
	Tags *string `json:"tags,omitempty" tf:"tags"`
	// The size of each node (optional, the default is currently g3.k3s.medium)
	// +optional
	TargetNodesSize *string `json:"targetNodesSize,omitempty" tf:"target_nodes_size"`
}

type ClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ClusterList is a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Cluster CRD objects
	Items []Cluster `json:"items,omitempty"`
}
