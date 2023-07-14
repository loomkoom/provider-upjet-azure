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

type SecurityCenterSubscriptionPricingObservation struct {

	// The subscription pricing ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The resource type this setting affects. Possible values are AppServices, ContainerRegistry, KeyVaults, KubernetesService, SqlServers, SqlServerVirtualMachines, StorageAccounts, VirtualMachines, Arm, Dns, OpenSourceRelationalDatabases, Containers, CosmosDbs and CloudPosture. Defaults to VirtualMachines
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// Resource type pricing subplan. Contact your MSFT representative for possible values.
	Subplan *string `json:"subplan,omitempty" tf:"subplan,omitempty"`

	// The pricing tier to use. Possible values are Free and Standard.
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`
}

type SecurityCenterSubscriptionPricingParameters struct {

	// The resource type this setting affects. Possible values are AppServices, ContainerRegistry, KeyVaults, KubernetesService, SqlServers, SqlServerVirtualMachines, StorageAccounts, VirtualMachines, Arm, Dns, OpenSourceRelationalDatabases, Containers, CosmosDbs and CloudPosture. Defaults to VirtualMachines
	// +kubebuilder:validation:Optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// Resource type pricing subplan. Contact your MSFT representative for possible values.
	// +kubebuilder:validation:Optional
	Subplan *string `json:"subplan,omitempty" tf:"subplan,omitempty"`

	// The pricing tier to use. Possible values are Free and Standard.
	// +kubebuilder:validation:Optional
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`
}

// SecurityCenterSubscriptionPricingSpec defines the desired state of SecurityCenterSubscriptionPricing
type SecurityCenterSubscriptionPricingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityCenterSubscriptionPricingParameters `json:"forProvider"`
}

// SecurityCenterSubscriptionPricingStatus defines the observed state of SecurityCenterSubscriptionPricing.
type SecurityCenterSubscriptionPricingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityCenterSubscriptionPricingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityCenterSubscriptionPricing is the Schema for the SecurityCenterSubscriptionPricings API. Manages the Pricing Tier for Azure Security Center in the current subscription.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SecurityCenterSubscriptionPricing struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tier)",message="tier is a required parameter"
	Spec   SecurityCenterSubscriptionPricingSpec   `json:"spec"`
	Status SecurityCenterSubscriptionPricingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityCenterSubscriptionPricingList contains a list of SecurityCenterSubscriptionPricings
type SecurityCenterSubscriptionPricingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityCenterSubscriptionPricing `json:"items"`
}

// Repository type metadata.
var (
	SecurityCenterSubscriptionPricing_Kind             = "SecurityCenterSubscriptionPricing"
	SecurityCenterSubscriptionPricing_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityCenterSubscriptionPricing_Kind}.String()
	SecurityCenterSubscriptionPricing_KindAPIVersion   = SecurityCenterSubscriptionPricing_Kind + "." + CRDGroupVersion.String()
	SecurityCenterSubscriptionPricing_GroupVersionKind = CRDGroupVersion.WithKind(SecurityCenterSubscriptionPricing_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityCenterSubscriptionPricing{}, &SecurityCenterSubscriptionPricingList{})
}
