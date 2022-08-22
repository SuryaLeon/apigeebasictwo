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

type HostObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type HostParameters struct {

	// +kubebuilder:validation:Optional
	BaseURL *string `json:"baseUrl,omitempty" tf:"base_url,omitempty"`

	// +kubebuilder:validation:Required
	EnvironmentName *string `json:"environmentName" tf:"environment_name,omitempty"`

	// +kubebuilder:validation:Required
	HostAliases []*string `json:"hostAliases" tf:"host_aliases,omitempty"`

	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	SSLClientAuthEnabled *bool `json:"sslClientAuthEnabled,omitempty" tf:"ssl_client_auth_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	SSLEnabled *bool `json:"sslEnabled,omitempty" tf:"ssl_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	SSLIgnoreValidationErrors *bool `json:"sslIgnoreValidationErrors,omitempty" tf:"ssl_ignore_validation_errors,omitempty"`

	// +kubebuilder:validation:Optional
	SSLKeyalias *string `json:"sslKeyalias,omitempty" tf:"ssl_keyalias,omitempty"`

	// +kubebuilder:validation:Optional
	SSLKeystore *string `json:"sslKeystore,omitempty" tf:"ssl_keystore,omitempty"`

	// +kubebuilder:validation:Optional
	SSLTruststore *string `json:"sslTruststore,omitempty" tf:"ssl_truststore,omitempty"`
}

// HostSpec defines the desired state of Host
type HostSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HostParameters `json:"forProvider"`
}

// HostStatus defines the observed state of Host.
type HostStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HostObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Host is the Schema for the Hosts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,apigeebasictwojet}
type Host struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HostSpec   `json:"spec"`
	Status            HostStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostList contains a list of Hosts
type HostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Host `json:"items"`
}

// Repository type metadata.
var (
	Host_Kind             = "Host"
	Host_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Host_Kind}.String()
	Host_KindAPIVersion   = Host_Kind + "." + CRDGroupVersion.String()
	Host_GroupVersionKind = CRDGroupVersion.WithKind(Host_Kind)
)

func init() {
	SchemeBuilder.Register(&Host{}, &HostList{})
}
