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

type ServiceInitParameters struct {

	// The location where the Communication service stores its data at rest. Possible values are Africa, Asia Pacific, Australia, Brazil, Canada, Europe, France, Germany, India, Japan, Korea, Norway, Switzerland, UAE, UK and United States. Defaults to United States. Changing this forces a new Communication Service to be created.
	DataLocation *string `json:"dataLocation,omitempty" tf:"data_location,omitempty"`

	// A mapping of tags which should be assigned to the Communication Service.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ServiceObservation struct {

	// The location where the Communication service stores its data at rest. Possible values are Africa, Asia Pacific, Australia, Brazil, Canada, Europe, France, Germany, India, Japan, Korea, Norway, Switzerland, UAE, UK and United States. Defaults to United States. Changing this forces a new Communication Service to be created.
	DataLocation *string `json:"dataLocation,omitempty" tf:"data_location,omitempty"`

	// The ID of the Communication Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The primary connection string of the Communication Service.
	PrimaryConnectionString *string `json:"primaryConnectionString,omitempty" tf:"primary_connection_string,omitempty"`

	// The primary key of the Communication Service.
	PrimaryKey *string `json:"primaryKey,omitempty" tf:"primary_key,omitempty"`

	// The name of the Resource Group where the Communication Service should exist. Changing this forces a new Communication Service to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The secondary connection string of the Communication Service.
	SecondaryConnectionString *string `json:"secondaryConnectionString,omitempty" tf:"secondary_connection_string,omitempty"`

	// The secondary key of the Communication Service.
	SecondaryKey *string `json:"secondaryKey,omitempty" tf:"secondary_key,omitempty"`

	// A mapping of tags which should be assigned to the Communication Service.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ServiceParameters struct {

	// The location where the Communication service stores its data at rest. Possible values are Africa, Asia Pacific, Australia, Brazil, Canada, Europe, France, Germany, India, Japan, Korea, Norway, Switzerland, UAE, UK and United States. Defaults to United States. Changing this forces a new Communication Service to be created.
	// +kubebuilder:validation:Optional
	DataLocation *string `json:"dataLocation,omitempty" tf:"data_location,omitempty"`

	// The name of the Resource Group where the Communication Service should exist. Changing this forces a new Communication Service to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Communication Service.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ServiceSpec defines the desired state of Service
type ServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceParameters `json:"forProvider"`
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
	InitProvider ServiceInitParameters `json:"initProvider,omitempty"`
}

// ServiceStatus defines the observed state of Service.
type ServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Service is the Schema for the Services API. Manages a Communication Service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Service struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceSpec   `json:"spec"`
	Status            ServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceList contains a list of Services
type ServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Service `json:"items"`
}

// Repository type metadata.
var (
	Service_Kind             = "Service"
	Service_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Service_Kind}.String()
	Service_KindAPIVersion   = Service_Kind + "." + CRDGroupVersion.String()
	Service_GroupVersionKind = CRDGroupVersion.WithKind(Service_Kind)
)

func init() {
	SchemeBuilder.Register(&Service{}, &ServiceList{})
}
