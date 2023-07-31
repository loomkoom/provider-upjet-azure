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

type OutputFunctionInitParameters struct {

	// The maximum number of events in each batch that's sent to the function. Defaults to 100.
	BatchMaxCount *float64 `json:"batchMaxCount,omitempty" tf:"batch_max_count,omitempty"`

	// The maximum batch size in bytes that's sent to the function. Defaults to 262144 (256 kB).
	BatchMaxInBytes *float64 `json:"batchMaxInBytes,omitempty" tf:"batch_max_in_bytes,omitempty"`

	// The name of the function in the Function App.
	FunctionName *string `json:"functionName,omitempty" tf:"function_name,omitempty"`
}

type OutputFunctionObservation struct {

	// The maximum number of events in each batch that's sent to the function. Defaults to 100.
	BatchMaxCount *float64 `json:"batchMaxCount,omitempty" tf:"batch_max_count,omitempty"`

	// The maximum batch size in bytes that's sent to the function. Defaults to 262144 (256 kB).
	BatchMaxInBytes *float64 `json:"batchMaxInBytes,omitempty" tf:"batch_max_in_bytes,omitempty"`

	// The name of the Function App.
	FunctionApp *string `json:"functionApp,omitempty" tf:"function_app,omitempty"`

	// The name of the function in the Function App.
	FunctionName *string `json:"functionName,omitempty" tf:"function_name,omitempty"`

	// The ID of the Stream Analytics Output Function.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Resource Group where the Stream Analytics Output should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName,omitempty" tf:"stream_analytics_job_name,omitempty"`
}

type OutputFunctionParameters struct {

	// The API key for the Function.
	// +kubebuilder:validation:Optional
	APIKeySecretRef v1.SecretKeySelector `json:"apiKeySecretRef" tf:"-"`

	// The maximum number of events in each batch that's sent to the function. Defaults to 100.
	// +kubebuilder:validation:Optional
	BatchMaxCount *float64 `json:"batchMaxCount,omitempty" tf:"batch_max_count,omitempty"`

	// The maximum batch size in bytes that's sent to the function. Defaults to 262144 (256 kB).
	// +kubebuilder:validation:Optional
	BatchMaxInBytes *float64 `json:"batchMaxInBytes,omitempty" tf:"batch_max_in_bytes,omitempty"`

	// The name of the Function App.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/web/v1beta1.FunctionApp
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	FunctionApp *string `json:"functionApp,omitempty" tf:"function_app,omitempty"`

	// Reference to a FunctionApp in web to populate functionApp.
	// +kubebuilder:validation:Optional
	FunctionAppRef *v1.Reference `json:"functionAppRef,omitempty" tf:"-"`

	// Selector for a FunctionApp in web to populate functionApp.
	// +kubebuilder:validation:Optional
	FunctionAppSelector *v1.Selector `json:"functionAppSelector,omitempty" tf:"-"`

	// The name of the function in the Function App.
	// +kubebuilder:validation:Optional
	FunctionName *string `json:"functionName,omitempty" tf:"function_name,omitempty"`

	// The name of the Resource Group where the Stream Analytics Output should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/streamanalytics/v1beta1.Job
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName,omitempty" tf:"stream_analytics_job_name,omitempty"`

	// Reference to a Job in streamanalytics to populate streamAnalyticsJobName.
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobNameRef *v1.Reference `json:"streamAnalyticsJobNameRef,omitempty" tf:"-"`

	// Selector for a Job in streamanalytics to populate streamAnalyticsJobName.
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobNameSelector *v1.Selector `json:"streamAnalyticsJobNameSelector,omitempty" tf:"-"`
}

// OutputFunctionSpec defines the desired state of OutputFunction
type OutputFunctionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OutputFunctionParameters `json:"forProvider"`
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
	InitProvider OutputFunctionInitParameters `json:"initProvider,omitempty"`
}

// OutputFunctionStatus defines the observed state of OutputFunction.
type OutputFunctionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OutputFunctionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OutputFunction is the Schema for the OutputFunctions API. Manages a Stream Analytics Output Function.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type OutputFunction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.apiKeySecretRef)",message="apiKeySecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.functionName) || has(self.initProvider.functionName)",message="functionName is a required parameter"
	Spec   OutputFunctionSpec   `json:"spec"`
	Status OutputFunctionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OutputFunctionList contains a list of OutputFunctions
type OutputFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OutputFunction `json:"items"`
}

// Repository type metadata.
var (
	OutputFunction_Kind             = "OutputFunction"
	OutputFunction_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OutputFunction_Kind}.String()
	OutputFunction_KindAPIVersion   = OutputFunction_Kind + "." + CRDGroupVersion.String()
	OutputFunction_GroupVersionKind = CRDGroupVersion.WithKind(OutputFunction_Kind)
)

func init() {
	SchemeBuilder.Register(&OutputFunction{}, &OutputFunctionList{})
}
