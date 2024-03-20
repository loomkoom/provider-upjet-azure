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

type MonitorDataCollectionEndpointInitParameters struct {

	// Specifies a description for the Data Collection Endpoint.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The kind of the Data Collection Endpoint. Possible values are Linux and Windows.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// The Azure Region where the Data Collection Endpoint should exist. Changing this forces a new Data Collection Endpoint to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Whether network access from public internet to the Data Collection Endpoint are allowed. Possible values are true and false. Default to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// A mapping of tags which should be assigned to the Data Collection Endpoint.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MonitorDataCollectionEndpointObservation struct {

	// The endpoint used for accessing configuration, e.g., https://mydce-abcd.eastus-1.control.monitor.azure.com.
	ConfigurationAccessEndpoint *string `json:"configurationAccessEndpoint,omitempty" tf:"configuration_access_endpoint,omitempty"`

	// Specifies a description for the Data Collection Endpoint.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Data Collection Endpoint.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The kind of the Data Collection Endpoint. Possible values are Linux and Windows.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// The Azure Region where the Data Collection Endpoint should exist. Changing this forces a new Data Collection Endpoint to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The endpoint used for ingesting logs, e.g., https://mydce-abcd.eastus-1.ingest.monitor.azure.com.
	LogsIngestionEndpoint *string `json:"logsIngestionEndpoint,omitempty" tf:"logs_ingestion_endpoint,omitempty"`

	// Whether network access from public internet to the Data Collection Endpoint are allowed. Possible values are true and false. Default to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the Resource Group where the Data Collection Endpoint should exist. Changing this forces a new Data Collection Endpoint to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags which should be assigned to the Data Collection Endpoint.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MonitorDataCollectionEndpointParameters struct {

	// Specifies a description for the Data Collection Endpoint.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The kind of the Data Collection Endpoint. Possible values are Linux and Windows.
	// +kubebuilder:validation:Optional
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// The Azure Region where the Data Collection Endpoint should exist. Changing this forces a new Data Collection Endpoint to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Whether network access from public internet to the Data Collection Endpoint are allowed. Possible values are true and false. Default to true.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the Resource Group where the Data Collection Endpoint should exist. Changing this forces a new Data Collection Endpoint to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Data Collection Endpoint.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// MonitorDataCollectionEndpointSpec defines the desired state of MonitorDataCollectionEndpoint
type MonitorDataCollectionEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitorDataCollectionEndpointParameters `json:"forProvider"`
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
	InitProvider MonitorDataCollectionEndpointInitParameters `json:"initProvider,omitempty"`
}

// MonitorDataCollectionEndpointStatus defines the observed state of MonitorDataCollectionEndpoint.
type MonitorDataCollectionEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitorDataCollectionEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MonitorDataCollectionEndpoint is the Schema for the MonitorDataCollectionEndpoints API. Manages a Data Collection Endpoint.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MonitorDataCollectionEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   MonitorDataCollectionEndpointSpec   `json:"spec"`
	Status MonitorDataCollectionEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorDataCollectionEndpointList contains a list of MonitorDataCollectionEndpoints
type MonitorDataCollectionEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitorDataCollectionEndpoint `json:"items"`
}

// Repository type metadata.
var (
	MonitorDataCollectionEndpoint_Kind             = "MonitorDataCollectionEndpoint"
	MonitorDataCollectionEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MonitorDataCollectionEndpoint_Kind}.String()
	MonitorDataCollectionEndpoint_KindAPIVersion   = MonitorDataCollectionEndpoint_Kind + "." + CRDGroupVersion.String()
	MonitorDataCollectionEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(MonitorDataCollectionEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&MonitorDataCollectionEndpoint{}, &MonitorDataCollectionEndpointList{})
}
