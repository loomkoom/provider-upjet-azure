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

type APISchemaInitParameters struct {

	// Types definitions. Used for Swagger/OpenAPI v2/v3 schemas only.
	Components *string `json:"components,omitempty" tf:"components,omitempty"`

	// The content type of the API Schema.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Types definitions. Used for Swagger/OpenAPI v1 schemas only.
	Definitions *string `json:"definitions,omitempty" tf:"definitions,omitempty"`

	// The JSON escaped string defining the document representing the Schema.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type APISchemaObservation struct {

	// The Name of the API Management Service where the API exists. Changing this forces a new resource to be created.
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// The name of the API within the API Management Service where this API Schema should be created. Changing this forces a new resource to be created.
	APIName *string `json:"apiName,omitempty" tf:"api_name,omitempty"`

	// Types definitions. Used for Swagger/OpenAPI v2/v3 schemas only.
	Components *string `json:"components,omitempty" tf:"components,omitempty"`

	// The content type of the API Schema.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Types definitions. Used for Swagger/OpenAPI v1 schemas only.
	Definitions *string `json:"definitions,omitempty" tf:"definitions,omitempty"`

	// The ID of the API Management API Schema.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The JSON escaped string defining the document representing the Schema.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type APISchemaParameters struct {

	// The Name of the API Management Service where the API exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Management
	// +kubebuilder:validation:Optional
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Reference to a Management to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameRef *v1.Reference `json:"apiManagementNameRef,omitempty" tf:"-"`

	// Selector for a Management to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameSelector *v1.Selector `json:"apiManagementNameSelector,omitempty" tf:"-"`

	// The name of the API within the API Management Service where this API Schema should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=API
	// +kubebuilder:validation:Optional
	APIName *string `json:"apiName,omitempty" tf:"api_name,omitempty"`

	// Reference to a API to populate apiName.
	// +kubebuilder:validation:Optional
	APINameRef *v1.Reference `json:"apiNameRef,omitempty" tf:"-"`

	// Selector for a API to populate apiName.
	// +kubebuilder:validation:Optional
	APINameSelector *v1.Selector `json:"apiNameSelector,omitempty" tf:"-"`

	// Types definitions. Used for Swagger/OpenAPI v2/v3 schemas only.
	// +kubebuilder:validation:Optional
	Components *string `json:"components,omitempty" tf:"components,omitempty"`

	// The content type of the API Schema.
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Types definitions. Used for Swagger/OpenAPI v1 schemas only.
	// +kubebuilder:validation:Optional
	Definitions *string `json:"definitions,omitempty" tf:"definitions,omitempty"`

	// The Name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The JSON escaped string defining the document representing the Schema.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// APISchemaSpec defines the desired state of APISchema
type APISchemaSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APISchemaParameters `json:"forProvider"`
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
	InitProvider APISchemaInitParameters `json:"initProvider,omitempty"`
}

// APISchemaStatus defines the observed state of APISchema.
type APISchemaStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APISchemaObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// APISchema is the Schema for the APISchemas API. Manages an API Schema within an API Management Service.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type APISchema struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.contentType) || (has(self.initProvider) && has(self.initProvider.contentType))",message="spec.forProvider.contentType is a required parameter"
	Spec   APISchemaSpec   `json:"spec"`
	Status APISchemaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APISchemaList contains a list of APISchemas
type APISchemaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APISchema `json:"items"`
}

// Repository type metadata.
var (
	APISchema_Kind             = "APISchema"
	APISchema_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: APISchema_Kind}.String()
	APISchema_KindAPIVersion   = APISchema_Kind + "." + CRDGroupVersion.String()
	APISchema_GroupVersionKind = CRDGroupVersion.WithKind(APISchema_Kind)
)

func init() {
	SchemeBuilder.Register(&APISchema{}, &APISchemaList{})
}
