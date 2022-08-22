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

type KeystoreObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type KeystoreParameters struct {

	// +kubebuilder:validation:Required
	EnvironmentName *string `json:"environmentName" tf:"environment_name,omitempty"`
}

// KeystoreSpec defines the desired state of Keystore
type KeystoreSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KeystoreParameters `json:"forProvider"`
}

// KeystoreStatus defines the observed state of Keystore.
type KeystoreStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KeystoreObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Keystore is the Schema for the Keystores API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,apigeebasictwojet}
type Keystore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeystoreSpec   `json:"spec"`
	Status            KeystoreStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeystoreList contains a list of Keystores
type KeystoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Keystore `json:"items"`
}

// Repository type metadata.
var (
	Keystore_Kind             = "Keystore"
	Keystore_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Keystore_Kind}.String()
	Keystore_KindAPIVersion   = Keystore_Kind + "." + CRDGroupVersion.String()
	Keystore_GroupVersionKind = CRDGroupVersion.WithKind(Keystore_Kind)
)

func init() {
	SchemeBuilder.Register(&Keystore{}, &KeystoreList{})
}
