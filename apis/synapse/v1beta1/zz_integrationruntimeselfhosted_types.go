// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type IntegrationRuntimeSelfHostedInitParameters struct {

	// Integration runtime description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`
}

type IntegrationRuntimeSelfHostedObservation struct {

	// The primary integration runtime authentication key.
	AuthorizationKeyPrimary *string `json:"authorizationKeyPrimary,omitempty" tf:"authorization_key_primary,omitempty"`

	// The secondary integration runtime authentication key.
	AuthorizationKeySecondary *string `json:"authorizationKeySecondary,omitempty" tf:"authorization_key_secondary,omitempty"`

	// Integration runtime description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Synapse Self-hosted Integration Runtime.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Synapse Workspace ID in which to associate the Integration Runtime with. Changing this forces a new Synapse Self-hosted Integration Runtime to be created.
	SynapseWorkspaceID *string `json:"synapseWorkspaceId,omitempty" tf:"synapse_workspace_id,omitempty"`
}

type IntegrationRuntimeSelfHostedParameters struct {

	// Integration runtime description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Synapse Workspace ID in which to associate the Integration Runtime with. Changing this forces a new Synapse Self-hosted Integration Runtime to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/synapse/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SynapseWorkspaceID *string `json:"synapseWorkspaceId,omitempty" tf:"synapse_workspace_id,omitempty"`

	// Reference to a Workspace in synapse to populate synapseWorkspaceId.
	// +kubebuilder:validation:Optional
	SynapseWorkspaceIDRef *v1.Reference `json:"synapseWorkspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in synapse to populate synapseWorkspaceId.
	// +kubebuilder:validation:Optional
	SynapseWorkspaceIDSelector *v1.Selector `json:"synapseWorkspaceIdSelector,omitempty" tf:"-"`
}

// IntegrationRuntimeSelfHostedSpec defines the desired state of IntegrationRuntimeSelfHosted
type IntegrationRuntimeSelfHostedSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IntegrationRuntimeSelfHostedParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider IntegrationRuntimeSelfHostedInitParameters `json:"initProvider,omitempty"`
}

// IntegrationRuntimeSelfHostedStatus defines the observed state of IntegrationRuntimeSelfHosted.
type IntegrationRuntimeSelfHostedStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IntegrationRuntimeSelfHostedObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// IntegrationRuntimeSelfHosted is the Schema for the IntegrationRuntimeSelfHosteds API. Manages a Synapse Self-hosted Integration Runtime.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IntegrationRuntimeSelfHosted struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IntegrationRuntimeSelfHostedSpec   `json:"spec"`
	Status            IntegrationRuntimeSelfHostedStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IntegrationRuntimeSelfHostedList contains a list of IntegrationRuntimeSelfHosteds
type IntegrationRuntimeSelfHostedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IntegrationRuntimeSelfHosted `json:"items"`
}

// Repository type metadata.
var (
	IntegrationRuntimeSelfHosted_Kind             = "IntegrationRuntimeSelfHosted"
	IntegrationRuntimeSelfHosted_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IntegrationRuntimeSelfHosted_Kind}.String()
	IntegrationRuntimeSelfHosted_KindAPIVersion   = IntegrationRuntimeSelfHosted_Kind + "." + CRDGroupVersion.String()
	IntegrationRuntimeSelfHosted_GroupVersionKind = CRDGroupVersion.WithKind(IntegrationRuntimeSelfHosted_Kind)
)

func init() {
	SchemeBuilder.Register(&IntegrationRuntimeSelfHosted{}, &IntegrationRuntimeSelfHostedList{})
}
