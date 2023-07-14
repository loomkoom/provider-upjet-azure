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

type ActiveDirectoryAdministratorObservation struct {

	// The ID of the PostgreSQL Active Directory Administrator.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The login name of the principal to set as the server administrator
	Login *string `json:"login,omitempty" tf:"login,omitempty"`

	// The ID of the principal to set as the server administrator. For a managed identity this should be the Client ID of the identity.
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The name of the resource group for the PostgreSQL server. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The name of the PostgreSQL Server on which to set the administrator. Changing this forces a new resource to be created.
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`

	// The Azure Tenant ID
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type ActiveDirectoryAdministratorParameters struct {

	// The login name of the principal to set as the server administrator
	// +kubebuilder:validation:Optional
	Login *string `json:"login,omitempty" tf:"login,omitempty"`

	// The ID of the principal to set as the server administrator. For a managed identity this should be the Client ID of the identity.
	// +kubebuilder:validation:Optional
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The name of the resource group for the PostgreSQL server. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The name of the PostgreSQL Server on which to set the administrator. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Server
	// +kubebuilder:validation:Optional
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`

	// Reference to a Server to populate serverName.
	// +kubebuilder:validation:Optional
	ServerNameRef *v1.Reference `json:"serverNameRef,omitempty" tf:"-"`

	// Selector for a Server to populate serverName.
	// +kubebuilder:validation:Optional
	ServerNameSelector *v1.Selector `json:"serverNameSelector,omitempty" tf:"-"`

	// The Azure Tenant ID
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

// ActiveDirectoryAdministratorSpec defines the desired state of ActiveDirectoryAdministrator
type ActiveDirectoryAdministratorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ActiveDirectoryAdministratorParameters `json:"forProvider"`
}

// ActiveDirectoryAdministratorStatus defines the observed state of ActiveDirectoryAdministrator.
type ActiveDirectoryAdministratorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ActiveDirectoryAdministratorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ActiveDirectoryAdministrator is the Schema for the ActiveDirectoryAdministrators API. Manages an Active Directory administrator on a PostgreSQL server
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ActiveDirectoryAdministrator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.login)",message="login is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.objectId)",message="objectId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tenantId)",message="tenantId is a required parameter"
	Spec   ActiveDirectoryAdministratorSpec   `json:"spec"`
	Status ActiveDirectoryAdministratorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ActiveDirectoryAdministratorList contains a list of ActiveDirectoryAdministrators
type ActiveDirectoryAdministratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ActiveDirectoryAdministrator `json:"items"`
}

// Repository type metadata.
var (
	ActiveDirectoryAdministrator_Kind             = "ActiveDirectoryAdministrator"
	ActiveDirectoryAdministrator_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ActiveDirectoryAdministrator_Kind}.String()
	ActiveDirectoryAdministrator_KindAPIVersion   = ActiveDirectoryAdministrator_Kind + "." + CRDGroupVersion.String()
	ActiveDirectoryAdministrator_GroupVersionKind = CRDGroupVersion.WithKind(ActiveDirectoryAdministrator_Kind)
)

func init() {
	SchemeBuilder.Register(&ActiveDirectoryAdministrator{}, &ActiveDirectoryAdministratorList{})
}
