// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FunctionAppHybridConnectionInitParameters struct {

	// The ID of the Function App for this Hybrid Connection. Changing this forces a new resource to be created.
	// The ID of the Function App for this Hybrid Connection.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/web/v1beta1.WindowsFunctionApp
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	FunctionAppID *string `json:"functionAppId,omitempty" tf:"function_app_id,omitempty"`

	// Reference to a WindowsFunctionApp in web to populate functionAppId.
	// +kubebuilder:validation:Optional
	FunctionAppIDRef *v1.Reference `json:"functionAppIdRef,omitempty" tf:"-"`

	// Selector for a WindowsFunctionApp in web to populate functionAppId.
	// +kubebuilder:validation:Optional
	FunctionAppIDSelector *v1.Selector `json:"functionAppIdSelector,omitempty" tf:"-"`

	// The hostname of the endpoint.
	// The hostname of the endpoint.
	HostName *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// The port to use for the endpoint
	// The port to use for the endpoint
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the Relay Hybrid Connection to use. Changing this forces a new resource to be created.
	// The ID of the Relay Hybrid Connection to use.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/relay/v1beta1.HybridConnection
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	RelayID *string `json:"relayId,omitempty" tf:"relay_id,omitempty"`

	// Reference to a HybridConnection in relay to populate relayId.
	// +kubebuilder:validation:Optional
	RelayIDRef *v1.Reference `json:"relayIdRef,omitempty" tf:"-"`

	// Selector for a HybridConnection in relay to populate relayId.
	// +kubebuilder:validation:Optional
	RelayIDSelector *v1.Selector `json:"relayIdSelector,omitempty" tf:"-"`

	// The name of the Relay key with Send permission to use. Defaults to RootManageSharedAccessKey
	// The name of the Relay key with `Send` permission to use. Defaults to `RootManageSharedAccessKey`
	SendKeyName *string `json:"sendKeyName,omitempty" tf:"send_key_name,omitempty"`
}

type FunctionAppHybridConnectionObservation struct {

	// The ID of the Function App for this Hybrid Connection. Changing this forces a new resource to be created.
	// The ID of the Function App for this Hybrid Connection.
	FunctionAppID *string `json:"functionAppId,omitempty" tf:"function_app_id,omitempty"`

	// The hostname of the endpoint.
	// The hostname of the endpoint.
	HostName *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// The ID of the Function App Hybrid Connection
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Relay Namespace.
	// The name of the Relay Namespace.
	NamespaceName *string `json:"namespaceName,omitempty" tf:"namespace_name,omitempty"`

	// The port to use for the endpoint
	// The port to use for the endpoint
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the Relay Hybrid Connection to use. Changing this forces a new resource to be created.
	// The ID of the Relay Hybrid Connection to use.
	RelayID *string `json:"relayId,omitempty" tf:"relay_id,omitempty"`

	// The name of the Relay in use.
	// The name of the Relay in use.
	RelayName *string `json:"relayName,omitempty" tf:"relay_name,omitempty"`

	// The name of the Relay key with Send permission to use. Defaults to RootManageSharedAccessKey
	// The name of the Relay key with `Send` permission to use. Defaults to `RootManageSharedAccessKey`
	SendKeyName *string `json:"sendKeyName,omitempty" tf:"send_key_name,omitempty"`

	// The Service Bus Namespace.
	// The Service Bus Namespace.
	ServiceBusNamespace *string `json:"serviceBusNamespace,omitempty" tf:"service_bus_namespace,omitempty"`

	// The suffix for the endpoint.
	// The suffix for the endpoint.
	ServiceBusSuffix *string `json:"serviceBusSuffix,omitempty" tf:"service_bus_suffix,omitempty"`
}

type FunctionAppHybridConnectionParameters struct {

	// The ID of the Function App for this Hybrid Connection. Changing this forces a new resource to be created.
	// The ID of the Function App for this Hybrid Connection.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/web/v1beta1.WindowsFunctionApp
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	FunctionAppID *string `json:"functionAppId,omitempty" tf:"function_app_id,omitempty"`

	// Reference to a WindowsFunctionApp in web to populate functionAppId.
	// +kubebuilder:validation:Optional
	FunctionAppIDRef *v1.Reference `json:"functionAppIdRef,omitempty" tf:"-"`

	// Selector for a WindowsFunctionApp in web to populate functionAppId.
	// +kubebuilder:validation:Optional
	FunctionAppIDSelector *v1.Selector `json:"functionAppIdSelector,omitempty" tf:"-"`

	// The hostname of the endpoint.
	// The hostname of the endpoint.
	// +kubebuilder:validation:Optional
	HostName *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// The port to use for the endpoint
	// The port to use for the endpoint
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the Relay Hybrid Connection to use. Changing this forces a new resource to be created.
	// The ID of the Relay Hybrid Connection to use.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/relay/v1beta1.HybridConnection
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RelayID *string `json:"relayId,omitempty" tf:"relay_id,omitempty"`

	// Reference to a HybridConnection in relay to populate relayId.
	// +kubebuilder:validation:Optional
	RelayIDRef *v1.Reference `json:"relayIdRef,omitempty" tf:"-"`

	// Selector for a HybridConnection in relay to populate relayId.
	// +kubebuilder:validation:Optional
	RelayIDSelector *v1.Selector `json:"relayIdSelector,omitempty" tf:"-"`

	// The name of the Relay key with Send permission to use. Defaults to RootManageSharedAccessKey
	// The name of the Relay key with `Send` permission to use. Defaults to `RootManageSharedAccessKey`
	// +kubebuilder:validation:Optional
	SendKeyName *string `json:"sendKeyName,omitempty" tf:"send_key_name,omitempty"`
}

// FunctionAppHybridConnectionSpec defines the desired state of FunctionAppHybridConnection
type FunctionAppHybridConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FunctionAppHybridConnectionParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider FunctionAppHybridConnectionInitParameters `json:"initProvider,omitempty"`
}

// FunctionAppHybridConnectionStatus defines the observed state of FunctionAppHybridConnection.
type FunctionAppHybridConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FunctionAppHybridConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// FunctionAppHybridConnection is the Schema for the FunctionAppHybridConnections API. Manages a Function App Hybrid Connection.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FunctionAppHybridConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hostname) || (has(self.initProvider) && has(self.initProvider.hostname))",message="spec.forProvider.hostname is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.port) || (has(self.initProvider) && has(self.initProvider.port))",message="spec.forProvider.port is a required parameter"
	Spec   FunctionAppHybridConnectionSpec   `json:"spec"`
	Status FunctionAppHybridConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionAppHybridConnectionList contains a list of FunctionAppHybridConnections
type FunctionAppHybridConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FunctionAppHybridConnection `json:"items"`
}

// Repository type metadata.
var (
	FunctionAppHybridConnection_Kind             = "FunctionAppHybridConnection"
	FunctionAppHybridConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FunctionAppHybridConnection_Kind}.String()
	FunctionAppHybridConnection_KindAPIVersion   = FunctionAppHybridConnection_Kind + "." + CRDGroupVersion.String()
	FunctionAppHybridConnection_GroupVersionKind = CRDGroupVersion.WithKind(FunctionAppHybridConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&FunctionAppHybridConnection{}, &FunctionAppHybridConnectionList{})
}
