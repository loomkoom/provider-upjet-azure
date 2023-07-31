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

type ManagementGroupInitParameters struct {

	// A friendly name for this Management Group. If not specified, this will be the same as the name.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// A list of Subscription GUIDs which should be assigned to the Management Group.
	SubscriptionIds []*string `json:"subscriptionIds,omitempty" tf:"subscription_ids,omitempty"`
}

type ManagementGroupObservation struct {

	// A friendly name for this Management Group. If not specified, this will be the same as the name.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The ID of the Management Group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Parent Management Group.
	ParentManagementGroupID *string `json:"parentManagementGroupId,omitempty" tf:"parent_management_group_id,omitempty"`

	// A list of Subscription GUIDs which should be assigned to the Management Group.
	SubscriptionIds []*string `json:"subscriptionIds,omitempty" tf:"subscription_ids,omitempty"`
}

type ManagementGroupParameters struct {

	// A friendly name for this Management Group. If not specified, this will be the same as the name.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The ID of the Parent Management Group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/management/v1beta1.ManagementGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ParentManagementGroupID *string `json:"parentManagementGroupId,omitempty" tf:"parent_management_group_id,omitempty"`

	// Reference to a ManagementGroup in management to populate parentManagementGroupId.
	// +kubebuilder:validation:Optional
	ParentManagementGroupIDRef *v1.Reference `json:"parentManagementGroupIdRef,omitempty" tf:"-"`

	// Selector for a ManagementGroup in management to populate parentManagementGroupId.
	// +kubebuilder:validation:Optional
	ParentManagementGroupIDSelector *v1.Selector `json:"parentManagementGroupIdSelector,omitempty" tf:"-"`

	// A list of Subscription GUIDs which should be assigned to the Management Group.
	// +kubebuilder:validation:Optional
	SubscriptionIds []*string `json:"subscriptionIds,omitempty" tf:"subscription_ids,omitempty"`
}

// ManagementGroupSpec defines the desired state of ManagementGroup
type ManagementGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagementGroupParameters `json:"forProvider"`
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
	InitProvider ManagementGroupInitParameters `json:"initProvider,omitempty"`
}

// ManagementGroupStatus defines the observed state of ManagementGroup.
type ManagementGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagementGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementGroup is the Schema for the ManagementGroups API. Manages a Management Group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ManagementGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagementGroupSpec   `json:"spec"`
	Status            ManagementGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementGroupList contains a list of ManagementGroups
type ManagementGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagementGroup `json:"items"`
}

// Repository type metadata.
var (
	ManagementGroup_Kind             = "ManagementGroup"
	ManagementGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagementGroup_Kind}.String()
	ManagementGroup_KindAPIVersion   = ManagementGroup_Kind + "." + CRDGroupVersion.String()
	ManagementGroup_GroupVersionKind = CRDGroupVersion.WithKind(ManagementGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagementGroup{}, &ManagementGroupList{})
}
