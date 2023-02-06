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

type LogAnalyticsQueryPackObservation struct {

	// The ID of the Log Analytics Query Pack.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LogAnalyticsQueryPackParameters struct {

	// The Azure Region where the Log Analytics Query Pack should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The name of the Resource Group where the Log Analytics Query Pack should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Log Analytics Query Pack.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// LogAnalyticsQueryPackSpec defines the desired state of LogAnalyticsQueryPack
type LogAnalyticsQueryPackSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogAnalyticsQueryPackParameters `json:"forProvider"`
}

// LogAnalyticsQueryPackStatus defines the observed state of LogAnalyticsQueryPack.
type LogAnalyticsQueryPackStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogAnalyticsQueryPackObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LogAnalyticsQueryPack is the Schema for the LogAnalyticsQueryPacks API. Manages a Log Analytics Query Pack.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LogAnalyticsQueryPack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogAnalyticsQueryPackSpec   `json:"spec"`
	Status            LogAnalyticsQueryPackStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogAnalyticsQueryPackList contains a list of LogAnalyticsQueryPacks
type LogAnalyticsQueryPackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogAnalyticsQueryPack `json:"items"`
}

// Repository type metadata.
var (
	LogAnalyticsQueryPack_Kind             = "LogAnalyticsQueryPack"
	LogAnalyticsQueryPack_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LogAnalyticsQueryPack_Kind}.String()
	LogAnalyticsQueryPack_KindAPIVersion   = LogAnalyticsQueryPack_Kind + "." + CRDGroupVersion.String()
	LogAnalyticsQueryPack_GroupVersionKind = CRDGroupVersion.WithKind(LogAnalyticsQueryPack_Kind)
)

func init() {
	SchemeBuilder.Register(&LogAnalyticsQueryPack{}, &LogAnalyticsQueryPackList{})
}
