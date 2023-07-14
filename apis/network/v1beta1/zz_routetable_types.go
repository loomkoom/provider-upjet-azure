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

type RouteTableObservation struct {

	// Boolean flag which controls propagation of routes learned by BGP on that route table. True means disable.
	DisableBGPRoutePropagation *bool `json:"disableBgpRoutePropagation,omitempty" tf:"disable_bgp_route_propagation,omitempty"`

	// The Route Table ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group in which to create the route table. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// List of objects representing routes. Each object accepts the arguments documented below.
	Route []RouteTableRouteObservation `json:"route,omitempty" tf:"route,omitempty"`

	// The collection of Subnets associated with this route table.
	Subnets []*string `json:"subnets,omitempty" tf:"subnets,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RouteTableParameters struct {

	// Boolean flag which controls propagation of routes learned by BGP on that route table. True means disable.
	// +kubebuilder:validation:Optional
	DisableBGPRoutePropagation *bool `json:"disableBgpRoutePropagation,omitempty" tf:"disable_bgp_route_propagation,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group in which to create the route table. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// List of objects representing routes. Each object accepts the arguments documented below.
	// +kubebuilder:validation:Optional
	Route []RouteTableRouteParameters `json:"route,omitempty" tf:"route,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RouteTableRouteObservation struct {

	// The destination to which the route applies. Can be CIDR (such as 10.1.0.0/16) or Azure Service Tag (such as ApiManagement, AzureBackup or AzureMonitor) format.
	AddressPrefix *string `json:"addressPrefix,omitempty" tf:"address_prefix,omitempty"`

	// The name of the route.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Contains the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
	NextHopInIPAddress *string `json:"nextHopInIpAddress,omitempty" tf:"next_hop_in_ip_address,omitempty"`

	// The type of Azure hop the packet should be sent to. Possible values are VirtualNetworkGateway, VnetLocal, Internet, VirtualAppliance and None.
	NextHopType *string `json:"nextHopType,omitempty" tf:"next_hop_type,omitempty"`
}

type RouteTableRouteParameters struct {

	// The destination to which the route applies. Can be CIDR (such as 10.1.0.0/16) or Azure Service Tag (such as ApiManagement, AzureBackup or AzureMonitor) format.
	// +kubebuilder:validation:Optional
	AddressPrefix *string `json:"addressPrefix,omitempty" tf:"address_prefix"`

	// The name of the route.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`

	// Contains the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
	// +kubebuilder:validation:Optional
	NextHopInIPAddress *string `json:"nextHopInIpAddress,omitempty" tf:"next_hop_in_ip_address"`

	// The type of Azure hop the packet should be sent to. Possible values are VirtualNetworkGateway, VnetLocal, Internet, VirtualAppliance and None.
	// +kubebuilder:validation:Optional
	NextHopType *string `json:"nextHopType,omitempty" tf:"next_hop_type"`
}

// RouteTableSpec defines the desired state of RouteTable
type RouteTableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteTableParameters `json:"forProvider"`
}

// RouteTableStatus defines the observed state of RouteTable.
type RouteTableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteTableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RouteTable is the Schema for the RouteTables API. Manages a Route Table
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type RouteTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location)",message="location is a required parameter"
	Spec   RouteTableSpec   `json:"spec"`
	Status RouteTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteTableList contains a list of RouteTables
type RouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTable `json:"items"`
}

// Repository type metadata.
var (
	RouteTable_Kind             = "RouteTable"
	RouteTable_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouteTable_Kind}.String()
	RouteTable_KindAPIVersion   = RouteTable_Kind + "." + CRDGroupVersion.String()
	RouteTable_GroupVersionKind = CRDGroupVersion.WithKind(RouteTable_Kind)
)

func init() {
	SchemeBuilder.Register(&RouteTable{}, &RouteTableList{})
}
