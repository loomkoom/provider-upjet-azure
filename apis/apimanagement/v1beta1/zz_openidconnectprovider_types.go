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

type OpenIDConnectProviderObservation struct {

	// The name of the API Management Service in which this OpenID Connect Provider should be created. Changing this forces a new resource to be created.
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// A description of this OpenID Connect Provider.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A user-friendly name for this OpenID Connect Provider.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The ID of the API Management OpenID Connect Provider.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The URI of the Metadata endpoint.
	MetadataEndpoint *string `json:"metadataEndpoint,omitempty" tf:"metadata_endpoint,omitempty"`

	// The name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
}

type OpenIDConnectProviderParameters struct {

	// The name of the API Management Service in which this OpenID Connect Provider should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Management
	// +kubebuilder:validation:Optional
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameRef *v1.Reference `json:"apiManagementNameRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameSelector *v1.Selector `json:"apiManagementNameSelector,omitempty" tf:"-"`

	// The Client ID used for the Client Application.
	// +kubebuilder:validation:Optional
	ClientIDSecretRef v1.SecretKeySelector `json:"clientIdSecretRef" tf:"-"`

	// The Client Secret used for the Client Application.
	// +kubebuilder:validation:Optional
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// A description of this OpenID Connect Provider.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A user-friendly name for this OpenID Connect Provider.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The URI of the Metadata endpoint.
	// +kubebuilder:validation:Optional
	MetadataEndpoint *string `json:"metadataEndpoint,omitempty" tf:"metadata_endpoint,omitempty"`

	// The name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
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

// OpenIDConnectProviderSpec defines the desired state of OpenIDConnectProvider
type OpenIDConnectProviderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OpenIDConnectProviderParameters `json:"forProvider"`
}

// OpenIDConnectProviderStatus defines the observed state of OpenIDConnectProvider.
type OpenIDConnectProviderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OpenIDConnectProviderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OpenIDConnectProvider is the Schema for the OpenIDConnectProviders API. Manages an OpenID Connect Provider within a API Management Service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type OpenIDConnectProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientIdSecretRef)",message="clientIdSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientSecretSecretRef)",message="clientSecretSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName)",message="displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.metadataEndpoint)",message="metadataEndpoint is a required parameter"
	Spec   OpenIDConnectProviderSpec   `json:"spec"`
	Status OpenIDConnectProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpenIDConnectProviderList contains a list of OpenIDConnectProviders
type OpenIDConnectProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpenIDConnectProvider `json:"items"`
}

// Repository type metadata.
var (
	OpenIDConnectProvider_Kind             = "OpenIDConnectProvider"
	OpenIDConnectProvider_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OpenIDConnectProvider_Kind}.String()
	OpenIDConnectProvider_KindAPIVersion   = OpenIDConnectProvider_Kind + "." + CRDGroupVersion.String()
	OpenIDConnectProvider_GroupVersionKind = CRDGroupVersion.WithKind(OpenIDConnectProvider_Kind)
)

func init() {
	SchemeBuilder.Register(&OpenIDConnectProvider{}, &OpenIDConnectProviderList{})
}
