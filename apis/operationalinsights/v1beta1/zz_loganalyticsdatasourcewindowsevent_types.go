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

type LogAnalyticsDataSourceWindowsEventObservation struct {

	// Specifies the name of the Windows Event Log to collect events from.
	EventLogName *string `json:"eventLogName,omitempty" tf:"event_log_name,omitempty"`

	// Specifies an array of event types applied to the specified event log. Possible values include Error, Warning and Information.
	EventTypes []*string `json:"eventTypes,omitempty" tf:"event_types,omitempty"`

	// The ID of the Log Analytics Windows Event DataSource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Resource Group where the Log Analytics Windows Event DataSource should exist. Changing this forces a new Log Analytics Windows Event DataSource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The name of the Log Analytics Workspace where the Log Analytics Windows Event DataSource should exist. Changing this forces a new Log Analytics Windows Event DataSource to be created.
	WorkspaceName *string `json:"workspaceName,omitempty" tf:"workspace_name,omitempty"`
}

type LogAnalyticsDataSourceWindowsEventParameters struct {

	// Specifies the name of the Windows Event Log to collect events from.
	// +kubebuilder:validation:Optional
	EventLogName *string `json:"eventLogName,omitempty" tf:"event_log_name,omitempty"`

	// Specifies an array of event types applied to the specified event log. Possible values include Error, Warning and Information.
	// +kubebuilder:validation:Optional
	EventTypes []*string `json:"eventTypes,omitempty" tf:"event_types,omitempty"`

	// The name of the Resource Group where the Log Analytics Windows Event DataSource should exist. Changing this forces a new Log Analytics Windows Event DataSource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The name of the Log Analytics Workspace where the Log Analytics Windows Event DataSource should exist. Changing this forces a new Log Analytics Windows Event DataSource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/operationalinsights/v1beta1.Workspace
	// +kubebuilder:validation:Optional
	WorkspaceName *string `json:"workspaceName,omitempty" tf:"workspace_name,omitempty"`

	// Reference to a Workspace in operationalinsights to populate workspaceName.
	// +kubebuilder:validation:Optional
	WorkspaceNameRef *v1.Reference `json:"workspaceNameRef,omitempty" tf:"-"`

	// Selector for a Workspace in operationalinsights to populate workspaceName.
	// +kubebuilder:validation:Optional
	WorkspaceNameSelector *v1.Selector `json:"workspaceNameSelector,omitempty" tf:"-"`
}

// LogAnalyticsDataSourceWindowsEventSpec defines the desired state of LogAnalyticsDataSourceWindowsEvent
type LogAnalyticsDataSourceWindowsEventSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogAnalyticsDataSourceWindowsEventParameters `json:"forProvider"`
}

// LogAnalyticsDataSourceWindowsEventStatus defines the observed state of LogAnalyticsDataSourceWindowsEvent.
type LogAnalyticsDataSourceWindowsEventStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogAnalyticsDataSourceWindowsEventObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LogAnalyticsDataSourceWindowsEvent is the Schema for the LogAnalyticsDataSourceWindowsEvents API. Manages a Log Analytics Windows Event DataSource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LogAnalyticsDataSourceWindowsEvent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.eventLogName)",message="eventLogName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.eventTypes)",message="eventTypes is a required parameter"
	Spec   LogAnalyticsDataSourceWindowsEventSpec   `json:"spec"`
	Status LogAnalyticsDataSourceWindowsEventStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogAnalyticsDataSourceWindowsEventList contains a list of LogAnalyticsDataSourceWindowsEvents
type LogAnalyticsDataSourceWindowsEventList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogAnalyticsDataSourceWindowsEvent `json:"items"`
}

// Repository type metadata.
var (
	LogAnalyticsDataSourceWindowsEvent_Kind             = "LogAnalyticsDataSourceWindowsEvent"
	LogAnalyticsDataSourceWindowsEvent_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LogAnalyticsDataSourceWindowsEvent_Kind}.String()
	LogAnalyticsDataSourceWindowsEvent_KindAPIVersion   = LogAnalyticsDataSourceWindowsEvent_Kind + "." + CRDGroupVersion.String()
	LogAnalyticsDataSourceWindowsEvent_GroupVersionKind = CRDGroupVersion.WithKind(LogAnalyticsDataSourceWindowsEvent_Kind)
)

func init() {
	SchemeBuilder.Register(&LogAnalyticsDataSourceWindowsEvent{}, &LogAnalyticsDataSourceWindowsEventList{})
}
