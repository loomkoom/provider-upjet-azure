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

type WorkspaceRootDbfsCustomerManagedKeyInitParameters struct {

	// The resource ID of the Key Vault Key to be used.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id,omitempty"`

	// Reference to a Key in keyvault to populate keyVaultKeyId.
	// +kubebuilder:validation:Optional
	KeyVaultKeyIDRef *v1.Reference `json:"keyVaultKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in keyvault to populate keyVaultKeyId.
	// +kubebuilder:validation:Optional
	KeyVaultKeyIDSelector *v1.Selector `json:"keyVaultKeyIdSelector,omitempty" tf:"-"`
}

type WorkspaceRootDbfsCustomerManagedKeyObservation struct {

	// The ID of the Databricks Workspace.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The resource ID of the Key Vault Key to be used.
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id,omitempty"`

	// The resource ID of the Databricks Workspace.
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

type WorkspaceRootDbfsCustomerManagedKeyParameters struct {

	// The resource ID of the Key Vault Key to be used.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id,omitempty"`

	// Reference to a Key in keyvault to populate keyVaultKeyId.
	// +kubebuilder:validation:Optional
	KeyVaultKeyIDRef *v1.Reference `json:"keyVaultKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in keyvault to populate keyVaultKeyId.
	// +kubebuilder:validation:Optional
	KeyVaultKeyIDSelector *v1.Selector `json:"keyVaultKeyIdSelector,omitempty" tf:"-"`

	// The resource ID of the Databricks Workspace.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/databricks/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`

	// Reference to a Workspace in databricks to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDRef *v1.Reference `json:"workspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in databricks to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDSelector *v1.Selector `json:"workspaceIdSelector,omitempty" tf:"-"`
}

// WorkspaceRootDbfsCustomerManagedKeySpec defines the desired state of WorkspaceRootDbfsCustomerManagedKey
type WorkspaceRootDbfsCustomerManagedKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkspaceRootDbfsCustomerManagedKeyParameters `json:"forProvider"`
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
	InitProvider WorkspaceRootDbfsCustomerManagedKeyInitParameters `json:"initProvider,omitempty"`
}

// WorkspaceRootDbfsCustomerManagedKeyStatus defines the observed state of WorkspaceRootDbfsCustomerManagedKey.
type WorkspaceRootDbfsCustomerManagedKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkspaceRootDbfsCustomerManagedKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// WorkspaceRootDbfsCustomerManagedKey is the Schema for the WorkspaceRootDbfsCustomerManagedKeys API. Manages a Customer Managed Key for the Databricks Workspaces root Databricks File System(DBFS)
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type WorkspaceRootDbfsCustomerManagedKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkspaceRootDbfsCustomerManagedKeySpec   `json:"spec"`
	Status            WorkspaceRootDbfsCustomerManagedKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceRootDbfsCustomerManagedKeyList contains a list of WorkspaceRootDbfsCustomerManagedKeys
type WorkspaceRootDbfsCustomerManagedKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkspaceRootDbfsCustomerManagedKey `json:"items"`
}

// Repository type metadata.
var (
	WorkspaceRootDbfsCustomerManagedKey_Kind             = "WorkspaceRootDbfsCustomerManagedKey"
	WorkspaceRootDbfsCustomerManagedKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WorkspaceRootDbfsCustomerManagedKey_Kind}.String()
	WorkspaceRootDbfsCustomerManagedKey_KindAPIVersion   = WorkspaceRootDbfsCustomerManagedKey_Kind + "." + CRDGroupVersion.String()
	WorkspaceRootDbfsCustomerManagedKey_GroupVersionKind = CRDGroupVersion.WithKind(WorkspaceRootDbfsCustomerManagedKey_Kind)
)

func init() {
	SchemeBuilder.Register(&WorkspaceRootDbfsCustomerManagedKey{}, &WorkspaceRootDbfsCustomerManagedKeyList{})
}
