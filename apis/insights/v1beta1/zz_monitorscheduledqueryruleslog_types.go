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

type MonitorScheduledQueryRulesLogCriteriaDimensionInitParameters struct {

	// Name of the dimension.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Operator for dimension values, - 'Include'. Defaults to Include.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// List of dimension values.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MonitorScheduledQueryRulesLogCriteriaDimensionObservation struct {

	// Name of the dimension.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Operator for dimension values, - 'Include'. Defaults to Include.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// List of dimension values.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MonitorScheduledQueryRulesLogCriteriaDimensionParameters struct {

	// Name of the dimension.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Operator for dimension values, - 'Include'. Defaults to Include.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// List of dimension values.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type MonitorScheduledQueryRulesLogCriteriaInitParameters struct {

	// A dimension block as defined below.
	Dimension []MonitorScheduledQueryRulesLogCriteriaDimensionInitParameters `json:"dimension,omitempty" tf:"dimension,omitempty"`

	// Name of the metric. Supported metrics are listed in the Azure Monitor Microsoft.OperationalInsights/workspaces metrics namespace.
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`
}

type MonitorScheduledQueryRulesLogCriteriaObservation struct {

	// A dimension block as defined below.
	Dimension []MonitorScheduledQueryRulesLogCriteriaDimensionObservation `json:"dimension,omitempty" tf:"dimension,omitempty"`

	// Name of the metric. Supported metrics are listed in the Azure Monitor Microsoft.OperationalInsights/workspaces metrics namespace.
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`
}

type MonitorScheduledQueryRulesLogCriteriaParameters struct {

	// A dimension block as defined below.
	// +kubebuilder:validation:Optional
	Dimension []MonitorScheduledQueryRulesLogCriteriaDimensionParameters `json:"dimension" tf:"dimension,omitempty"`

	// Name of the metric. Supported metrics are listed in the Azure Monitor Microsoft.OperationalInsights/workspaces metrics namespace.
	// +kubebuilder:validation:Optional
	MetricName *string `json:"metricName" tf:"metric_name,omitempty"`
}

type MonitorScheduledQueryRulesLogInitParameters struct {

	// A list of IDs of Resources referred into query.
	// +listType=set
	AuthorizedResourceIds []*string `json:"authorizedResourceIds,omitempty" tf:"authorized_resource_ids,omitempty"`

	// A criteria block as defined below.
	Criteria []MonitorScheduledQueryRulesLogCriteriaInitParameters `json:"criteria,omitempty" tf:"criteria,omitempty"`

	// The resource URI over which log search query is to be run. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/operationalinsights/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	DataSourceID *string `json:"dataSourceId,omitempty" tf:"data_source_id,omitempty"`

	// Reference to a Workspace in operationalinsights to populate dataSourceId.
	// +kubebuilder:validation:Optional
	DataSourceIDRef *v1.Reference `json:"dataSourceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in operationalinsights to populate dataSourceId.
	// +kubebuilder:validation:Optional
	DataSourceIDSelector *v1.Selector `json:"dataSourceIdSelector,omitempty" tf:"-"`

	// The description of the scheduled query rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether this scheduled query rule is enabled. Default is true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies the Azure Region where the resource should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the scheduled query rule. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the resource group in which to create the scheduled query rule instance. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MonitorScheduledQueryRulesLogObservation struct {

	// A list of IDs of Resources referred into query.
	// +listType=set
	AuthorizedResourceIds []*string `json:"authorizedResourceIds,omitempty" tf:"authorized_resource_ids,omitempty"`

	// A criteria block as defined below.
	Criteria []MonitorScheduledQueryRulesLogCriteriaObservation `json:"criteria,omitempty" tf:"criteria,omitempty"`

	// The resource URI over which log search query is to be run. Changing this forces a new resource to be created.
	DataSourceID *string `json:"dataSourceId,omitempty" tf:"data_source_id,omitempty"`

	// The description of the scheduled query rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether this scheduled query rule is enabled. Default is true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The ID of the scheduled query rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the Azure Region where the resource should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the scheduled query rule. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the resource group in which to create the scheduled query rule instance. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MonitorScheduledQueryRulesLogParameters struct {

	// A list of IDs of Resources referred into query.
	// +kubebuilder:validation:Optional
	// +listType=set
	AuthorizedResourceIds []*string `json:"authorizedResourceIds,omitempty" tf:"authorized_resource_ids,omitempty"`

	// A criteria block as defined below.
	// +kubebuilder:validation:Optional
	Criteria []MonitorScheduledQueryRulesLogCriteriaParameters `json:"criteria,omitempty" tf:"criteria,omitempty"`

	// The resource URI over which log search query is to be run. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/operationalinsights/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataSourceID *string `json:"dataSourceId,omitempty" tf:"data_source_id,omitempty"`

	// Reference to a Workspace in operationalinsights to populate dataSourceId.
	// +kubebuilder:validation:Optional
	DataSourceIDRef *v1.Reference `json:"dataSourceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in operationalinsights to populate dataSourceId.
	// +kubebuilder:validation:Optional
	DataSourceIDSelector *v1.Selector `json:"dataSourceIdSelector,omitempty" tf:"-"`

	// The description of the scheduled query rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether this scheduled query rule is enabled. Default is true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies the Azure Region where the resource should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the scheduled query rule. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the resource group in which to create the scheduled query rule instance. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
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
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// MonitorScheduledQueryRulesLogSpec defines the desired state of MonitorScheduledQueryRulesLog
type MonitorScheduledQueryRulesLogSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitorScheduledQueryRulesLogParameters `json:"forProvider"`
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
	InitProvider MonitorScheduledQueryRulesLogInitParameters `json:"initProvider,omitempty"`
}

// MonitorScheduledQueryRulesLogStatus defines the observed state of MonitorScheduledQueryRulesLog.
type MonitorScheduledQueryRulesLogStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitorScheduledQueryRulesLogObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MonitorScheduledQueryRulesLog is the Schema for the MonitorScheduledQueryRulesLogs API. Manages a LogToMetricAction Scheduled Query Rules resources within Azure Monitor
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MonitorScheduledQueryRulesLog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.criteria) || (has(self.initProvider) && has(self.initProvider.criteria))",message="spec.forProvider.criteria is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   MonitorScheduledQueryRulesLogSpec   `json:"spec"`
	Status MonitorScheduledQueryRulesLogStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorScheduledQueryRulesLogList contains a list of MonitorScheduledQueryRulesLogs
type MonitorScheduledQueryRulesLogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitorScheduledQueryRulesLog `json:"items"`
}

// Repository type metadata.
var (
	MonitorScheduledQueryRulesLog_Kind             = "MonitorScheduledQueryRulesLog"
	MonitorScheduledQueryRulesLog_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MonitorScheduledQueryRulesLog_Kind}.String()
	MonitorScheduledQueryRulesLog_KindAPIVersion   = MonitorScheduledQueryRulesLog_Kind + "." + CRDGroupVersion.String()
	MonitorScheduledQueryRulesLog_GroupVersionKind = CRDGroupVersion.WithKind(MonitorScheduledQueryRulesLog_Kind)
)

func init() {
	SchemeBuilder.Register(&MonitorScheduledQueryRulesLog{}, &MonitorScheduledQueryRulesLogList{})
}
