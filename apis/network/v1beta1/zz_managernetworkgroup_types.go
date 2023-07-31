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

type ManagerNetworkGroupInitParameters struct {

	// A description of the Network Manager Network Group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`
}

type ManagerNetworkGroupObservation struct {

	// A description of the Network Manager Network Group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Network Manager Network Group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the ID of the Network Manager. Changing this forces a new Network Manager Network Group to be created.
	NetworkManagerID *string `json:"networkManagerId,omitempty" tf:"network_manager_id,omitempty"`
}

type ManagerNetworkGroupParameters struct {

	// A description of the Network Manager Network Group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the ID of the Network Manager. Changing this forces a new Network Manager Network Group to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Manager
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NetworkManagerID *string `json:"networkManagerId,omitempty" tf:"network_manager_id,omitempty"`

	// Reference to a Manager in network to populate networkManagerId.
	// +kubebuilder:validation:Optional
	NetworkManagerIDRef *v1.Reference `json:"networkManagerIdRef,omitempty" tf:"-"`

	// Selector for a Manager in network to populate networkManagerId.
	// +kubebuilder:validation:Optional
	NetworkManagerIDSelector *v1.Selector `json:"networkManagerIdSelector,omitempty" tf:"-"`
}

// ManagerNetworkGroupSpec defines the desired state of ManagerNetworkGroup
type ManagerNetworkGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagerNetworkGroupParameters `json:"forProvider"`
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
	InitProvider ManagerNetworkGroupInitParameters `json:"initProvider,omitempty"`
}

// ManagerNetworkGroupStatus defines the observed state of ManagerNetworkGroup.
type ManagerNetworkGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagerNetworkGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagerNetworkGroup is the Schema for the ManagerNetworkGroups API. Manages a Network Manager Network Group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ManagerNetworkGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagerNetworkGroupSpec   `json:"spec"`
	Status            ManagerNetworkGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagerNetworkGroupList contains a list of ManagerNetworkGroups
type ManagerNetworkGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagerNetworkGroup `json:"items"`
}

// Repository type metadata.
var (
	ManagerNetworkGroup_Kind             = "ManagerNetworkGroup"
	ManagerNetworkGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagerNetworkGroup_Kind}.String()
	ManagerNetworkGroup_KindAPIVersion   = ManagerNetworkGroup_Kind + "." + CRDGroupVersion.String()
	ManagerNetworkGroup_GroupVersionKind = CRDGroupVersion.WithKind(ManagerNetworkGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagerNetworkGroup{}, &ManagerNetworkGroupList{})
}
