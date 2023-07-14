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

type VirtualHubObservation_2 struct {

	// The Address Prefix which should be used for this Virtual Hub. Changing this forces a new resource to be created. The address prefix subnet cannot be smaller than a .
	AddressPrefix *string `json:"addressPrefix,omitempty" tf:"address_prefix,omitempty"`

	// The ID of the default Route Table in the Virtual Hub.
	DefaultRouteTableID *string `json:"defaultRouteTableId,omitempty" tf:"default_route_table_id,omitempty"`

	// The hub routing preference. Possible values are ExpressRoute, ASPath and VpnGateway. Defaults to ExpressRoute.
	HubRoutingPreference *string `json:"hubRoutingPreference,omitempty" tf:"hub_routing_preference,omitempty"`

	// The ID of the Virtual Hub.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure location where the Virtual Hub should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name of the Resource Group where the Virtual Hub should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// One or more route blocks as defined below.
	Route []VirtualHubRouteObservation `json:"route,omitempty" tf:"route,omitempty"`

	// The SKU of the Virtual Hub. Possible values are Basic and Standard. Changing this forces a new resource to be created.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the Virtual Hub.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The Autonomous System Number of the Virtual Hub BGP router.
	VirtualRouterAsn *float64 `json:"virtualRouterAsn,omitempty" tf:"virtual_router_asn,omitempty"`

	// The IP addresses of the Virtual Hub BGP router.
	VirtualRouterIps []*string `json:"virtualRouterIps,omitempty" tf:"virtual_router_ips,omitempty"`

	// The ID of a Virtual WAN within which the Virtual Hub should be created. Changing this forces a new resource to be created.
	VirtualWanID *string `json:"virtualWanId,omitempty" tf:"virtual_wan_id,omitempty"`
}

type VirtualHubParameters_2 struct {

	// The Address Prefix which should be used for this Virtual Hub. Changing this forces a new resource to be created. The address prefix subnet cannot be smaller than a .
	// +kubebuilder:validation:Optional
	AddressPrefix *string `json:"addressPrefix,omitempty" tf:"address_prefix,omitempty"`

	// The hub routing preference. Possible values are ExpressRoute, ASPath and VpnGateway. Defaults to ExpressRoute.
	// +kubebuilder:validation:Optional
	HubRoutingPreference *string `json:"hubRoutingPreference,omitempty" tf:"hub_routing_preference,omitempty"`

	// Specifies the supported Azure location where the Virtual Hub should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name of the Resource Group where the Virtual Hub should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// One or more route blocks as defined below.
	// +kubebuilder:validation:Optional
	Route []VirtualHubRouteParameters `json:"route,omitempty" tf:"route,omitempty"`

	// The SKU of the Virtual Hub. Possible values are Basic and Standard. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the Virtual Hub.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of a Virtual WAN within which the Virtual Hub should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=VirtualWAN
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VirtualWanID *string `json:"virtualWanId,omitempty" tf:"virtual_wan_id,omitempty"`

	// Reference to a VirtualWAN to populate virtualWanId.
	// +kubebuilder:validation:Optional
	VirtualWanIDRef *v1.Reference `json:"virtualWanIdRef,omitempty" tf:"-"`

	// Selector for a VirtualWAN to populate virtualWanId.
	// +kubebuilder:validation:Optional
	VirtualWanIDSelector *v1.Selector `json:"virtualWanIdSelector,omitempty" tf:"-"`
}

type VirtualHubRouteObservation struct {

	// A list of Address Prefixes.
	AddressPrefixes []*string `json:"addressPrefixes,omitempty" tf:"address_prefixes,omitempty"`

	// The IP Address that Packets should be forwarded to as the Next Hop.
	NextHopIPAddress *string `json:"nextHopIpAddress,omitempty" tf:"next_hop_ip_address,omitempty"`
}

type VirtualHubRouteParameters struct {

	// A list of Address Prefixes.
	// +kubebuilder:validation:Required
	AddressPrefixes []*string `json:"addressPrefixes" tf:"address_prefixes,omitempty"`

	// The IP Address that Packets should be forwarded to as the Next Hop.
	// +kubebuilder:validation:Required
	NextHopIPAddress *string `json:"nextHopIpAddress" tf:"next_hop_ip_address,omitempty"`
}

// VirtualHubSpec defines the desired state of VirtualHub
type VirtualHubSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualHubParameters_2 `json:"forProvider"`
}

// VirtualHubStatus defines the observed state of VirtualHub.
type VirtualHubStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualHubObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualHub is the Schema for the VirtualHubs API. Manages a Virtual Hub within a Virtual WAN.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualHub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location)",message="location is a required parameter"
	Spec   VirtualHubSpec   `json:"spec"`
	Status VirtualHubStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualHubList contains a list of VirtualHubs
type VirtualHubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualHub `json:"items"`
}

// Repository type metadata.
var (
	VirtualHub_Kind             = "VirtualHub"
	VirtualHub_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualHub_Kind}.String()
	VirtualHub_KindAPIVersion   = VirtualHub_Kind + "." + CRDGroupVersion.String()
	VirtualHub_GroupVersionKind = CRDGroupVersion.WithKind(VirtualHub_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualHub{}, &VirtualHubList{})
}
