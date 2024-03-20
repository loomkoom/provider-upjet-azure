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

type CrossTenantScopesInitParameters struct {
}

type CrossTenantScopesObservation struct {

	// List of management groups.
	ManagementGroups []*string `json:"managementGroups,omitempty" tf:"management_groups,omitempty"`

	// List of subscriptions.
	Subscriptions []*string `json:"subscriptions,omitempty" tf:"subscriptions,omitempty"`

	// Tenant ID.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type CrossTenantScopesParameters struct {
}

type ManagerInitParameters struct {

	// A description of the network manager.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the Azure Region where the Network Managers should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A scope block as defined below.
	Scope []ScopeInitParameters `json:"scope,omitempty" tf:"scope,omitempty"`

	// A list of configuration deployment type. Possible values are Connectivity and SecurityAdmin, corresponds to if Connectivity Configuration and Security Admin Configuration is allowed for the Network Manager.
	ScopeAccesses []*string `json:"scopeAccesses,omitempty" tf:"scope_accesses,omitempty"`

	// A mapping of tags which should be assigned to the Network Managers.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ManagerObservation struct {

	// One or more cross_tenant_scopes blocks as defined below.
	CrossTenantScopes []CrossTenantScopesObservation `json:"crossTenantScopes,omitempty" tf:"cross_tenant_scopes,omitempty"`

	// A description of the network manager.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Network Managers.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the Azure Region where the Network Managers should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name of the Resource Group where the Network Managers should exist. Changing this forces a new Network Managers to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A scope block as defined below.
	Scope []ScopeObservation `json:"scope,omitempty" tf:"scope,omitempty"`

	// A list of configuration deployment type. Possible values are Connectivity and SecurityAdmin, corresponds to if Connectivity Configuration and Security Admin Configuration is allowed for the Network Manager.
	ScopeAccesses []*string `json:"scopeAccesses,omitempty" tf:"scope_accesses,omitempty"`

	// A mapping of tags which should be assigned to the Network Managers.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ManagerParameters struct {

	// A description of the network manager.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the Azure Region where the Network Managers should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name of the Resource Group where the Network Managers should exist. Changing this forces a new Network Managers to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A scope block as defined below.
	// +kubebuilder:validation:Optional
	Scope []ScopeParameters `json:"scope,omitempty" tf:"scope,omitempty"`

	// A list of configuration deployment type. Possible values are Connectivity and SecurityAdmin, corresponds to if Connectivity Configuration and Security Admin Configuration is allowed for the Network Manager.
	// +kubebuilder:validation:Optional
	ScopeAccesses []*string `json:"scopeAccesses,omitempty" tf:"scope_accesses,omitempty"`

	// A mapping of tags which should be assigned to the Network Managers.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ScopeInitParameters struct {

	// A list of management group IDs.
	ManagementGroupIds []*string `json:"managementGroupIds,omitempty" tf:"management_group_ids,omitempty"`

	// A list of subscription IDs.
	SubscriptionIds []*string `json:"subscriptionIds,omitempty" tf:"subscription_ids,omitempty"`
}

type ScopeObservation struct {

	// A list of management group IDs.
	ManagementGroupIds []*string `json:"managementGroupIds,omitempty" tf:"management_group_ids,omitempty"`

	// A list of subscription IDs.
	SubscriptionIds []*string `json:"subscriptionIds,omitempty" tf:"subscription_ids,omitempty"`
}

type ScopeParameters struct {

	// A list of management group IDs.
	// +kubebuilder:validation:Optional
	ManagementGroupIds []*string `json:"managementGroupIds,omitempty" tf:"management_group_ids,omitempty"`

	// A list of subscription IDs.
	// +kubebuilder:validation:Optional
	SubscriptionIds []*string `json:"subscriptionIds,omitempty" tf:"subscription_ids,omitempty"`
}

// ManagerSpec defines the desired state of Manager
type ManagerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagerParameters `json:"forProvider"`
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
	InitProvider ManagerInitParameters `json:"initProvider,omitempty"`
}

// ManagerStatus defines the observed state of Manager.
type ManagerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Manager is the Schema for the Managers API. Manages a Network Managers.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Manager struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scope) || (has(self.initProvider) && has(self.initProvider.scope))",message="spec.forProvider.scope is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scopeAccesses) || (has(self.initProvider) && has(self.initProvider.scopeAccesses))",message="spec.forProvider.scopeAccesses is a required parameter"
	Spec   ManagerSpec   `json:"spec"`
	Status ManagerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagerList contains a list of Managers
type ManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Manager `json:"items"`
}

// Repository type metadata.
var (
	Manager_Kind             = "Manager"
	Manager_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Manager_Kind}.String()
	Manager_KindAPIVersion   = Manager_Kind + "." + CRDGroupVersion.String()
	Manager_GroupVersionKind = CRDGroupVersion.WithKind(Manager_Kind)
)

func init() {
	SchemeBuilder.Register(&Manager{}, &ManagerList{})
}
