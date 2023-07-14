/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PermissionsObservation struct {

	// One or more Allowed Actions, such as *, Microsoft.Resources/subscriptions/resourceGroups/read. See 'Azure Resource Manager resource provider operations' for details.
	Actions []*string `json:"actions,omitempty" tf:"actions,omitempty"`

	// One or more Allowed Data Actions, such as *, Microsoft.Storage/storageAccounts/blobServices/containers/blobs/read. See 'Azure Resource Manager resource provider operations' for details.
	DataActions []*string `json:"dataActions,omitempty" tf:"data_actions,omitempty"`

	// One or more Disallowed Actions, such as *, Microsoft.Resources/subscriptions/resourceGroups/read. See 'Azure Resource Manager resource provider operations' for details.
	NotActions []*string `json:"notActions,omitempty" tf:"not_actions,omitempty"`

	// One or more Disallowed Data Actions, such as *, Microsoft.Resources/subscriptions/resourceGroups/read. See 'Azure Resource Manager resource provider operations' for details.
	NotDataActions []*string `json:"notDataActions,omitempty" tf:"not_data_actions,omitempty"`
}

type PermissionsParameters struct {

	// One or more Allowed Actions, such as *, Microsoft.Resources/subscriptions/resourceGroups/read. See 'Azure Resource Manager resource provider operations' for details.
	// +kubebuilder:validation:Optional
	Actions []*string `json:"actions,omitempty" tf:"actions,omitempty"`

	// One or more Allowed Data Actions, such as *, Microsoft.Storage/storageAccounts/blobServices/containers/blobs/read. See 'Azure Resource Manager resource provider operations' for details.
	// +kubebuilder:validation:Optional
	DataActions []*string `json:"dataActions,omitempty" tf:"data_actions,omitempty"`

	// One or more Disallowed Actions, such as *, Microsoft.Resources/subscriptions/resourceGroups/read. See 'Azure Resource Manager resource provider operations' for details.
	// +kubebuilder:validation:Optional
	NotActions []*string `json:"notActions,omitempty" tf:"not_actions,omitempty"`

	// One or more Disallowed Data Actions, such as *, Microsoft.Resources/subscriptions/resourceGroups/read. See 'Azure Resource Manager resource provider operations' for details.
	// +kubebuilder:validation:Optional
	NotDataActions []*string `json:"notDataActions,omitempty" tf:"not_data_actions,omitempty"`
}

type RoleDefinitionObservation struct {

	// One or more assignable scopes for this Role Definition, such as /subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333, /subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup, or /subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM.
	AssignableScopes []*string `json:"assignableScopes,omitempty" tf:"assignable_scopes,omitempty"`

	// A description of the Role Definition.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Role Definition.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A permissions block as defined below.
	Permissions []PermissionsObservation `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// A unique UUID/GUID which identifies this role - one will be generated if not specified. Changing this forces a new resource to be created.
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty" tf:"role_definition_id,omitempty"`

	// The Azure Resource Manager ID for the resource.
	RoleDefinitionResourceID *string `json:"roleDefinitionResourceId,omitempty" tf:"role_definition_resource_id,omitempty"`

	// The scope at which the Role Definition applies to, such as /subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333, /subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup, or /subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM. It is recommended to use the first entry of the assignable_scopes. Changing this forces a new resource to be created.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type RoleDefinitionParameters struct {

	// One or more assignable scopes for this Role Definition, such as /subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333, /subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup, or /subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM.
	// +kubebuilder:validation:Optional
	AssignableScopes []*string `json:"assignableScopes,omitempty" tf:"assignable_scopes,omitempty"`

	// A description of the Role Definition.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the Role Definition.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A permissions block as defined below.
	// +kubebuilder:validation:Optional
	Permissions []PermissionsParameters `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// A unique UUID/GUID which identifies this role - one will be generated if not specified. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty" tf:"role_definition_id,omitempty"`

	// The scope at which the Role Definition applies to, such as /subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333, /subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup, or /subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM. It is recommended to use the first entry of the assignable_scopes. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

// RoleDefinitionSpec defines the desired state of RoleDefinition
type RoleDefinitionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleDefinitionParameters `json:"forProvider"`
}

// RoleDefinitionStatus defines the observed state of RoleDefinition.
type RoleDefinitionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleDefinitionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RoleDefinition is the Schema for the RoleDefinitions API. Manages a custom Role Definition.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type RoleDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scope)",message="scope is a required parameter"
	Spec   RoleDefinitionSpec   `json:"spec"`
	Status RoleDefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleDefinitionList contains a list of RoleDefinitions
type RoleDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleDefinition `json:"items"`
}

// Repository type metadata.
var (
	RoleDefinition_Kind             = "RoleDefinition"
	RoleDefinition_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RoleDefinition_Kind}.String()
	RoleDefinition_KindAPIVersion   = RoleDefinition_Kind + "." + CRDGroupVersion.String()
	RoleDefinition_GroupVersionKind = CRDGroupVersion.WithKind(RoleDefinition_Kind)
)

func init() {
	SchemeBuilder.Register(&RoleDefinition{}, &RoleDefinitionList{})
}
