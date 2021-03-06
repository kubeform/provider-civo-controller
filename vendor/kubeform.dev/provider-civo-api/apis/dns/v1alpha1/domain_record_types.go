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

type DomainRecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainRecordSpec   `json:"spec,omitempty"`
	Status            DomainRecordStatus `json:"status,omitempty"`
}

type DomainRecordSpec struct {
	State *DomainRecordSpecResource `json:"state,omitempty" tf:"-"`

	Resource DomainRecordSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type DomainRecordSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The account ID of this resource
	// +optional
	AccountID *string `json:"accountID,omitempty" tf:"account_id"`
	// Timestamp when this resource was created
	// +optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at"`
	// ID from domain name
	DomainID *string `json:"domainID" tf:"domain_id"`
	// The portion before the domain name (e.g. www) or an @ for the apex/root domain (you cannot use an A record with an amex/root domain)
	Name *string `json:"name" tf:"name"`
	// Useful for MX records only, the priority mail should be attempted it (defaults to 10)
	// +optional
	Priority *int64 `json:"priority,omitempty" tf:"priority"`
	// How long caching DNS servers should cache this record for, in seconds (the minimum is 600 and the default if unspecified is 600)
	Ttl *int64 `json:"ttl" tf:"ttl"`
	// The choice of RR type from a, cname, mx or txt
	Type *string `json:"type" tf:"type"`
	// Timestamp when this resource was updated
	// +optional
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at"`
	// The IP address (A or MX), hostname (CNAME or MX) or text value (TXT) to serve for this record
	Value *string `json:"value" tf:"value"`
}

type DomainRecordStatus struct {
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

// DomainRecordList is a list of DomainRecords
type DomainRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DomainRecord CRD objects
	Items []DomainRecord `json:"items,omitempty"`
}
