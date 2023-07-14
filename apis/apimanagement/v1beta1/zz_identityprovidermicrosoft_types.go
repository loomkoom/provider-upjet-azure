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

type IdentityProviderMicrosoftObservation struct {

	// The Name of the API Management Service where this Microsoft Identity Provider should be created. Changing this forces a new resource to be created.
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Client Id of the Azure AD Application.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The ID of the API Management Microsoft Identity Provider.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
}

type IdentityProviderMicrosoftParameters struct {

	// The Name of the API Management Service where this Microsoft Identity Provider should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Management
	// +kubebuilder:validation:Optional
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameRef *v1.Reference `json:"apiManagementNameRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameSelector *v1.Selector `json:"apiManagementNameSelector,omitempty" tf:"-"`

	// Client Id of the Azure AD Application.
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Client secret of the Azure AD Application.
	// +kubebuilder:validation:Optional
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// IdentityProviderMicrosoftSpec defines the desired state of IdentityProviderMicrosoft
type IdentityProviderMicrosoftSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IdentityProviderMicrosoftParameters `json:"forProvider"`
}

// IdentityProviderMicrosoftStatus defines the observed state of IdentityProviderMicrosoft.
type IdentityProviderMicrosoftStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IdentityProviderMicrosoftObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityProviderMicrosoft is the Schema for the IdentityProviderMicrosofts API. Manages an API Management Microsoft Identity Provider.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IdentityProviderMicrosoft struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientId)",message="clientId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientSecretSecretRef)",message="clientSecretSecretRef is a required parameter"
	Spec   IdentityProviderMicrosoftSpec   `json:"spec"`
	Status IdentityProviderMicrosoftStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityProviderMicrosoftList contains a list of IdentityProviderMicrosofts
type IdentityProviderMicrosoftList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IdentityProviderMicrosoft `json:"items"`
}

// Repository type metadata.
var (
	IdentityProviderMicrosoft_Kind             = "IdentityProviderMicrosoft"
	IdentityProviderMicrosoft_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IdentityProviderMicrosoft_Kind}.String()
	IdentityProviderMicrosoft_KindAPIVersion   = IdentityProviderMicrosoft_Kind + "." + CRDGroupVersion.String()
	IdentityProviderMicrosoft_GroupVersionKind = CRDGroupVersion.WithKind(IdentityProviderMicrosoft_Kind)
)

func init() {
	SchemeBuilder.Register(&IdentityProviderMicrosoft{}, &IdentityProviderMicrosoftList{})
}
