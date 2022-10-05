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

type PrivateDNSZoneVirtualNetworkLinkObservation struct {

	// The ID of the Private DNS Zone Virtual Network Link.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PrivateDNSZoneVirtualNetworkLinkParameters struct {

	// The name of the Private DNS zone (without a terminating dot). Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=PrivateDNSZone
	// +kubebuilder:validation:Optional
	PrivateDNSZoneName *string `json:"privateDnsZoneName,omitempty" tf:"private_dns_zone_name,omitempty"`

	// Reference to a PrivateDNSZone to populate privateDnsZoneName.
	// +kubebuilder:validation:Optional
	PrivateDNSZoneNameRef *v1.Reference `json:"privateDnsZoneNameRef,omitempty" tf:"-"`

	// Selector for a PrivateDNSZone to populate privateDnsZoneName.
	// +kubebuilder:validation:Optional
	PrivateDNSZoneNameSelector *v1.Selector `json:"privateDnsZoneNameSelector,omitempty" tf:"-"`

	// Is auto-registration of virtual machine records in the virtual network in the Private DNS zone enabled? Defaults to false.
	// +kubebuilder:validation:Optional
	RegistrationEnabled *bool `json:"registrationEnabled,omitempty" tf:"registration_enabled,omitempty"`

	// Specifies the resource group where the Private DNS Zone exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the Virtual Network that should be linked to the DNS Zone. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=VirtualNetwork
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VirtualNetworkID *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id,omitempty"`

	// Reference to a VirtualNetwork to populate virtualNetworkId.
	// +kubebuilder:validation:Optional
	VirtualNetworkIDRef *v1.Reference `json:"virtualNetworkIdRef,omitempty" tf:"-"`

	// Selector for a VirtualNetwork to populate virtualNetworkId.
	// +kubebuilder:validation:Optional
	VirtualNetworkIDSelector *v1.Selector `json:"virtualNetworkIdSelector,omitempty" tf:"-"`
}

// PrivateDNSZoneVirtualNetworkLinkSpec defines the desired state of PrivateDNSZoneVirtualNetworkLink
type PrivateDNSZoneVirtualNetworkLinkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateDNSZoneVirtualNetworkLinkParameters `json:"forProvider"`
}

// PrivateDNSZoneVirtualNetworkLinkStatus defines the observed state of PrivateDNSZoneVirtualNetworkLink.
type PrivateDNSZoneVirtualNetworkLinkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateDNSZoneVirtualNetworkLinkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDNSZoneVirtualNetworkLink is the Schema for the PrivateDNSZoneVirtualNetworkLinks API. Manages a Private DNS Zone Virtual Network Link.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PrivateDNSZoneVirtualNetworkLink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateDNSZoneVirtualNetworkLinkSpec   `json:"spec"`
	Status            PrivateDNSZoneVirtualNetworkLinkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDNSZoneVirtualNetworkLinkList contains a list of PrivateDNSZoneVirtualNetworkLinks
type PrivateDNSZoneVirtualNetworkLinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateDNSZoneVirtualNetworkLink `json:"items"`
}

// Repository type metadata.
var (
	PrivateDNSZoneVirtualNetworkLink_Kind             = "PrivateDNSZoneVirtualNetworkLink"
	PrivateDNSZoneVirtualNetworkLink_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivateDNSZoneVirtualNetworkLink_Kind}.String()
	PrivateDNSZoneVirtualNetworkLink_KindAPIVersion   = PrivateDNSZoneVirtualNetworkLink_Kind + "." + CRDGroupVersion.String()
	PrivateDNSZoneVirtualNetworkLink_GroupVersionKind = CRDGroupVersion.WithKind(PrivateDNSZoneVirtualNetworkLink_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivateDNSZoneVirtualNetworkLink{}, &PrivateDNSZoneVirtualNetworkLinkList{})
}