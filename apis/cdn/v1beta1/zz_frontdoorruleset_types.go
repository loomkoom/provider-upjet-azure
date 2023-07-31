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

type FrontdoorRuleSetInitParameters struct {
}

type FrontdoorRuleSetObservation struct {

	// The ID of the Front Door Profile. Changing this forces a new Front Door Rule Set to be created.
	CdnFrontdoorProfileID *string `json:"cdnFrontdoorProfileId,omitempty" tf:"cdn_frontdoor_profile_id,omitempty"`

	// The ID of the Front Door Rule Set.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FrontdoorRuleSetParameters struct {

	// The ID of the Front Door Profile. Changing this forces a new Front Door Rule Set to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cdn/v1beta1.FrontdoorProfile
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CdnFrontdoorProfileID *string `json:"cdnFrontdoorProfileId,omitempty" tf:"cdn_frontdoor_profile_id,omitempty"`

	// Reference to a FrontdoorProfile in cdn to populate cdnFrontdoorProfileId.
	// +kubebuilder:validation:Optional
	CdnFrontdoorProfileIDRef *v1.Reference `json:"cdnFrontdoorProfileIdRef,omitempty" tf:"-"`

	// Selector for a FrontdoorProfile in cdn to populate cdnFrontdoorProfileId.
	// +kubebuilder:validation:Optional
	CdnFrontdoorProfileIDSelector *v1.Selector `json:"cdnFrontdoorProfileIdSelector,omitempty" tf:"-"`
}

// FrontdoorRuleSetSpec defines the desired state of FrontdoorRuleSet
type FrontdoorRuleSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FrontdoorRuleSetParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider FrontdoorRuleSetInitParameters `json:"initProvider,omitempty"`
}

// FrontdoorRuleSetStatus defines the observed state of FrontdoorRuleSet.
type FrontdoorRuleSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FrontdoorRuleSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FrontdoorRuleSet is the Schema for the FrontdoorRuleSets API. Manages a Front Door (standard/premium) Rule Set.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FrontdoorRuleSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FrontdoorRuleSetSpec   `json:"spec"`
	Status            FrontdoorRuleSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FrontdoorRuleSetList contains a list of FrontdoorRuleSets
type FrontdoorRuleSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FrontdoorRuleSet `json:"items"`
}

// Repository type metadata.
var (
	FrontdoorRuleSet_Kind             = "FrontdoorRuleSet"
	FrontdoorRuleSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FrontdoorRuleSet_Kind}.String()
	FrontdoorRuleSet_KindAPIVersion   = FrontdoorRuleSet_Kind + "." + CRDGroupVersion.String()
	FrontdoorRuleSet_GroupVersionKind = CRDGroupVersion.WithKind(FrontdoorRuleSet_Kind)
)

func init() {
	SchemeBuilder.Register(&FrontdoorRuleSet{}, &FrontdoorRuleSetList{})
}
