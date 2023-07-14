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

type CapacityObservation struct {

	// The number of instances that are available for scaling if metrics are not available for evaluation. The default is only used if the current instance count is lower than the default. Valid values are between 0 and 1000.
	Default *float64 `json:"default,omitempty" tf:"default,omitempty"`

	// The maximum number of instances for this resource. Valid values are between 0 and 1000.
	Maximum *float64 `json:"maximum,omitempty" tf:"maximum,omitempty"`

	// The minimum number of instances for this resource. Valid values are between 0 and 1000.
	Minimum *float64 `json:"minimum,omitempty" tf:"minimum,omitempty"`
}

type CapacityParameters struct {

	// The number of instances that are available for scaling if metrics are not available for evaluation. The default is only used if the current instance count is lower than the default. Valid values are between 0 and 1000.
	// +kubebuilder:validation:Required
	Default *float64 `json:"default" tf:"default,omitempty"`

	// The maximum number of instances for this resource. Valid values are between 0 and 1000.
	// +kubebuilder:validation:Required
	Maximum *float64 `json:"maximum" tf:"maximum,omitempty"`

	// The minimum number of instances for this resource. Valid values are between 0 and 1000.
	// +kubebuilder:validation:Required
	Minimum *float64 `json:"minimum" tf:"minimum,omitempty"`
}

type DimensionsObservation struct {

	// Specifies the name of the profile.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the operator used to compare the metric data and threshold. Possible values are: Equals, NotEquals, GreaterThan, GreaterThanOrEqual, LessThan, LessThanOrEqual.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of dimension values.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type DimensionsParameters struct {

	// Specifies the name of the profile.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the operator used to compare the metric data and threshold. Possible values are: Equals, NotEquals, GreaterThan, GreaterThanOrEqual, LessThan, LessThanOrEqual.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of dimension values.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type EmailObservation struct {

	// Specifies a list of custom email addresses to which the email notifications will be sent.
	CustomEmails []*string `json:"customEmails,omitempty" tf:"custom_emails,omitempty"`

	// Should email notifications be sent to the subscription administrator? Defaults to false.
	SendToSubscriptionAdministrator *bool `json:"sendToSubscriptionAdministrator,omitempty" tf:"send_to_subscription_administrator,omitempty"`

	// Should email notifications be sent to the subscription co-administrator? Defaults to false.
	SendToSubscriptionCoAdministrator *bool `json:"sendToSubscriptionCoAdministrator,omitempty" tf:"send_to_subscription_co_administrator,omitempty"`
}

type EmailParameters struct {

	// Specifies a list of custom email addresses to which the email notifications will be sent.
	// +kubebuilder:validation:Optional
	CustomEmails []*string `json:"customEmails,omitempty" tf:"custom_emails,omitempty"`

	// Should email notifications be sent to the subscription administrator? Defaults to false.
	// +kubebuilder:validation:Optional
	SendToSubscriptionAdministrator *bool `json:"sendToSubscriptionAdministrator,omitempty" tf:"send_to_subscription_administrator,omitempty"`

	// Should email notifications be sent to the subscription co-administrator? Defaults to false.
	// +kubebuilder:validation:Optional
	SendToSubscriptionCoAdministrator *bool `json:"sendToSubscriptionCoAdministrator,omitempty" tf:"send_to_subscription_co_administrator,omitempty"`
}

type FixedDateObservation struct {

	// Specifies the end date for the profile, formatted as an RFC3339 date string.
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	// Specifies the start date for the profile, formatted as an RFC3339 date string.
	Start *string `json:"start,omitempty" tf:"start,omitempty"`

	// The Time Zone used for the hours field. A list of possible values can be found here. Defaults to UTC.
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

type FixedDateParameters struct {

	// Specifies the end date for the profile, formatted as an RFC3339 date string.
	// +kubebuilder:validation:Required
	End *string `json:"end" tf:"end,omitempty"`

	// Specifies the start date for the profile, formatted as an RFC3339 date string.
	// +kubebuilder:validation:Required
	Start *string `json:"start" tf:"start,omitempty"`

	// The Time Zone used for the hours field. A list of possible values can be found here. Defaults to UTC.
	// +kubebuilder:validation:Optional
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

type MetricTriggerObservation struct {

	// One or more dimensions block as defined below.
	Dimensions []DimensionsObservation `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// Whether to enable metric divide by instance count.
	DivideByInstanceCount *bool `json:"divideByInstanceCount,omitempty" tf:"divide_by_instance_count,omitempty"`

	// The name of the metric that defines what the rule monitors, such as Percentage CPU for Virtual Machine Scale Sets and CpuPercentage for App Service Plan.
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// The namespace of the metric that defines what the rule monitors, such as microsoft.compute/virtualmachinescalesets for Virtual Machine Scale Sets.
	MetricNamespace *string `json:"metricNamespace,omitempty" tf:"metric_namespace,omitempty"`

	// The ID of the Resource which the Rule monitors.
	MetricResourceID *string `json:"metricResourceId,omitempty" tf:"metric_resource_id,omitempty"`

	// Specifies the operator used to compare the metric data and threshold. Possible values are: Equals, NotEquals, GreaterThan, GreaterThanOrEqual, LessThan, LessThanOrEqual.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies how the metrics from multiple instances are combined. Possible values are Average, Max, Min and Sum.
	Statistic *string `json:"statistic,omitempty" tf:"statistic,omitempty"`

	// Specifies the threshold of the metric that triggers the scale action.
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// Specifies how the data that's collected should be combined over time. Possible values include Average, Count, Maximum, Minimum, Last and Total.
	TimeAggregation *string `json:"timeAggregation,omitempty" tf:"time_aggregation,omitempty"`

	// Specifies the granularity of metrics that the rule monitors, which must be one of the pre-defined values returned from the metric definitions for the metric. This value must be between 1 minute and 12 hours an be formatted as an ISO 8601 string.
	TimeGrain *string `json:"timeGrain,omitempty" tf:"time_grain,omitempty"`

	// Specifies the time range for which data is collected, which must be greater than the delay in metric collection (which varies from resource to resource). This value must be between 5 minutes and 12 hours and be formatted as an ISO 8601 string.
	TimeWindow *string `json:"timeWindow,omitempty" tf:"time_window,omitempty"`
}

type MetricTriggerParameters struct {

	// One or more dimensions block as defined below.
	// +kubebuilder:validation:Optional
	Dimensions []DimensionsParameters `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// Whether to enable metric divide by instance count.
	// +kubebuilder:validation:Optional
	DivideByInstanceCount *bool `json:"divideByInstanceCount,omitempty" tf:"divide_by_instance_count,omitempty"`

	// The name of the metric that defines what the rule monitors, such as Percentage CPU for Virtual Machine Scale Sets and CpuPercentage for App Service Plan.
	// +kubebuilder:validation:Required
	MetricName *string `json:"metricName" tf:"metric_name,omitempty"`

	// The namespace of the metric that defines what the rule monitors, such as microsoft.compute/virtualmachinescalesets for Virtual Machine Scale Sets.
	// +kubebuilder:validation:Optional
	MetricNamespace *string `json:"metricNamespace,omitempty" tf:"metric_namespace,omitempty"`

	// The ID of the Resource which the Rule monitors.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/compute/v1beta1.LinuxVirtualMachineScaleSet
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	MetricResourceID *string `json:"metricResourceId,omitempty" tf:"metric_resource_id,omitempty"`

	// Reference to a LinuxVirtualMachineScaleSet in compute to populate metricResourceId.
	// +kubebuilder:validation:Optional
	MetricResourceIDRef *v1.Reference `json:"metricResourceIdRef,omitempty" tf:"-"`

	// Selector for a LinuxVirtualMachineScaleSet in compute to populate metricResourceId.
	// +kubebuilder:validation:Optional
	MetricResourceIDSelector *v1.Selector `json:"metricResourceIdSelector,omitempty" tf:"-"`

	// Specifies the operator used to compare the metric data and threshold. Possible values are: Equals, NotEquals, GreaterThan, GreaterThanOrEqual, LessThan, LessThanOrEqual.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// Specifies how the metrics from multiple instances are combined. Possible values are Average, Max, Min and Sum.
	// +kubebuilder:validation:Required
	Statistic *string `json:"statistic" tf:"statistic,omitempty"`

	// Specifies the threshold of the metric that triggers the scale action.
	// +kubebuilder:validation:Required
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`

	// Specifies how the data that's collected should be combined over time. Possible values include Average, Count, Maximum, Minimum, Last and Total.
	// +kubebuilder:validation:Required
	TimeAggregation *string `json:"timeAggregation" tf:"time_aggregation,omitempty"`

	// Specifies the granularity of metrics that the rule monitors, which must be one of the pre-defined values returned from the metric definitions for the metric. This value must be between 1 minute and 12 hours an be formatted as an ISO 8601 string.
	// +kubebuilder:validation:Required
	TimeGrain *string `json:"timeGrain" tf:"time_grain,omitempty"`

	// Specifies the time range for which data is collected, which must be greater than the delay in metric collection (which varies from resource to resource). This value must be between 5 minutes and 12 hours and be formatted as an ISO 8601 string.
	// +kubebuilder:validation:Required
	TimeWindow *string `json:"timeWindow" tf:"time_window,omitempty"`
}

type MonitorAutoscaleSettingObservation struct {

	// Specifies whether automatic scaling is enabled for the target resource. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The ID of the AutoScale Setting.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure location where the AutoScale Setting should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the AutoScale Setting. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies a notification block as defined below.
	Notification []NotificationObservation `json:"notification,omitempty" tf:"notification,omitempty"`

	// Specifies one or more (up to 20) profile blocks as defined below.
	Profile []ProfileObservation `json:"profile,omitempty" tf:"profile,omitempty"`

	// The name of the Resource Group in the AutoScale Setting should be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the resource ID of the resource that the autoscale setting should be added to. Changing this forces a new resource to be created.
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`
}

type MonitorAutoscaleSettingParameters struct {

	// Specifies whether automatic scaling is enabled for the target resource. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies the supported Azure location where the AutoScale Setting should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the AutoScale Setting. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies a notification block as defined below.
	// +kubebuilder:validation:Optional
	Notification []NotificationParameters `json:"notification,omitempty" tf:"notification,omitempty"`

	// Specifies one or more (up to 20) profile blocks as defined below.
	// +kubebuilder:validation:Optional
	Profile []ProfileParameters `json:"profile,omitempty" tf:"profile,omitempty"`

	// The name of the Resource Group in the AutoScale Setting should be created. Changing this forces a new resource to be created.
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
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the resource ID of the resource that the autoscale setting should be added to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/compute/v1beta1.LinuxVirtualMachineScaleSet
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`

	// Reference to a LinuxVirtualMachineScaleSet in compute to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDRef *v1.Reference `json:"targetResourceIdRef,omitempty" tf:"-"`

	// Selector for a LinuxVirtualMachineScaleSet in compute to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDSelector *v1.Selector `json:"targetResourceIdSelector,omitempty" tf:"-"`
}

type NotificationObservation struct {

	// A email block as defined below.
	Email []EmailObservation `json:"email,omitempty" tf:"email,omitempty"`

	// One or more webhook blocks as defined below.
	Webhook []WebhookObservation `json:"webhook,omitempty" tf:"webhook,omitempty"`
}

type NotificationParameters struct {

	// A email block as defined below.
	// +kubebuilder:validation:Optional
	Email []EmailParameters `json:"email,omitempty" tf:"email,omitempty"`

	// One or more webhook blocks as defined below.
	// +kubebuilder:validation:Optional
	Webhook []WebhookParameters `json:"webhook,omitempty" tf:"webhook,omitempty"`
}

type ProfileObservation struct {

	// A capacity block as defined below.
	Capacity []CapacityObservation `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// A fixed_date block as defined below. This cannot be specified if a recurrence block is specified.
	FixedDate []FixedDateObservation `json:"fixedDate,omitempty" tf:"fixed_date,omitempty"`

	// Specifies the name of the profile.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A recurrence block as defined below. This cannot be specified if a fixed_date block is specified.
	Recurrence []RecurrenceObservation `json:"recurrence,omitempty" tf:"recurrence,omitempty"`

	// One or more (up to 10) rule blocks as defined below.
	Rule []RuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`
}

type ProfileParameters struct {

	// A capacity block as defined below.
	// +kubebuilder:validation:Required
	Capacity []CapacityParameters `json:"capacity" tf:"capacity,omitempty"`

	// A fixed_date block as defined below. This cannot be specified if a recurrence block is specified.
	// +kubebuilder:validation:Optional
	FixedDate []FixedDateParameters `json:"fixedDate,omitempty" tf:"fixed_date,omitempty"`

	// Specifies the name of the profile.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// A recurrence block as defined below. This cannot be specified if a fixed_date block is specified.
	// +kubebuilder:validation:Optional
	Recurrence []RecurrenceParameters `json:"recurrence,omitempty" tf:"recurrence,omitempty"`

	// One or more (up to 10) rule blocks as defined below.
	// +kubebuilder:validation:Optional
	Rule []RuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type RecurrenceObservation struct {

	// A list of days that this profile takes effect on. Possible values include Monday, Tuesday, Wednesday, Thursday, Friday, Saturday and Sunday.
	Days []*string `json:"days,omitempty" tf:"days,omitempty"`

	// A list containing a single item, which specifies the Hour interval at which this recurrence should be triggered (in 24-hour time). Possible values are from 0 to 23.
	Hours []*float64 `json:"hours,omitempty" tf:"hours,omitempty"`

	// A list containing a single item which specifies the Minute interval at which this recurrence should be triggered.
	Minutes []*float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// The Time Zone used for the hours field. A list of possible values can be found here. Defaults to UTC.
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

type RecurrenceParameters struct {

	// A list of days that this profile takes effect on. Possible values include Monday, Tuesday, Wednesday, Thursday, Friday, Saturday and Sunday.
	// +kubebuilder:validation:Required
	Days []*string `json:"days" tf:"days,omitempty"`

	// A list containing a single item, which specifies the Hour interval at which this recurrence should be triggered (in 24-hour time). Possible values are from 0 to 23.
	// +kubebuilder:validation:Required
	Hours []*float64 `json:"hours" tf:"hours,omitempty"`

	// A list containing a single item which specifies the Minute interval at which this recurrence should be triggered.
	// +kubebuilder:validation:Required
	Minutes []*float64 `json:"minutes" tf:"minutes,omitempty"`

	// The Time Zone used for the hours field. A list of possible values can be found here. Defaults to UTC.
	// +kubebuilder:validation:Optional
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

type RuleObservation struct {

	// A metric_trigger block as defined below.
	MetricTrigger []MetricTriggerObservation `json:"metricTrigger,omitempty" tf:"metric_trigger,omitempty"`

	// A scale_action block as defined below.
	ScaleAction []ScaleActionObservation `json:"scaleAction,omitempty" tf:"scale_action,omitempty"`
}

type RuleParameters struct {

	// A metric_trigger block as defined below.
	// +kubebuilder:validation:Required
	MetricTrigger []MetricTriggerParameters `json:"metricTrigger" tf:"metric_trigger,omitempty"`

	// A scale_action block as defined below.
	// +kubebuilder:validation:Required
	ScaleAction []ScaleActionParameters `json:"scaleAction" tf:"scale_action,omitempty"`
}

type ScaleActionObservation struct {

	// The amount of time to wait since the last scaling action before this action occurs. Must be between 1 minute and 1 week and formatted as a ISO 8601 string.
	Cooldown *string `json:"cooldown,omitempty" tf:"cooldown,omitempty"`

	// The scale direction. Possible values are Increase and Decrease.
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// The type of action that should occur. Possible values are ChangeCount, ExactCount, PercentChangeCount and ServiceAllowedNextValue.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The number of instances involved in the scaling action.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type ScaleActionParameters struct {

	// The amount of time to wait since the last scaling action before this action occurs. Must be between 1 minute and 1 week and formatted as a ISO 8601 string.
	// +kubebuilder:validation:Required
	Cooldown *string `json:"cooldown" tf:"cooldown,omitempty"`

	// The scale direction. Possible values are Increase and Decrease.
	// +kubebuilder:validation:Required
	Direction *string `json:"direction" tf:"direction,omitempty"`

	// The type of action that should occur. Possible values are ChangeCount, ExactCount, PercentChangeCount and ServiceAllowedNextValue.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// The number of instances involved in the scaling action.
	// +kubebuilder:validation:Required
	Value *float64 `json:"value" tf:"value,omitempty"`
}

type WebhookObservation struct {

	// A map of settings.
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// The HTTPS URI which should receive scale notifications.
	ServiceURI *string `json:"serviceUri,omitempty" tf:"service_uri,omitempty"`
}

type WebhookParameters struct {

	// A map of settings.
	// +kubebuilder:validation:Optional
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// The HTTPS URI which should receive scale notifications.
	// +kubebuilder:validation:Required
	ServiceURI *string `json:"serviceUri" tf:"service_uri,omitempty"`
}

// MonitorAutoscaleSettingSpec defines the desired state of MonitorAutoscaleSetting
type MonitorAutoscaleSettingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitorAutoscaleSettingParameters `json:"forProvider"`
}

// MonitorAutoscaleSettingStatus defines the observed state of MonitorAutoscaleSetting.
type MonitorAutoscaleSettingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitorAutoscaleSettingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorAutoscaleSetting is the Schema for the MonitorAutoscaleSettings API. Manages an AutoScale Setting which can be applied to Virtual Machine Scale Sets, App Services and other scalable resources.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MonitorAutoscaleSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.profile)",message="profile is a required parameter"
	Spec   MonitorAutoscaleSettingSpec   `json:"spec"`
	Status MonitorAutoscaleSettingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorAutoscaleSettingList contains a list of MonitorAutoscaleSettings
type MonitorAutoscaleSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitorAutoscaleSetting `json:"items"`
}

// Repository type metadata.
var (
	MonitorAutoscaleSetting_Kind             = "MonitorAutoscaleSetting"
	MonitorAutoscaleSetting_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MonitorAutoscaleSetting_Kind}.String()
	MonitorAutoscaleSetting_KindAPIVersion   = MonitorAutoscaleSetting_Kind + "." + CRDGroupVersion.String()
	MonitorAutoscaleSetting_GroupVersionKind = CRDGroupVersion.WithKind(MonitorAutoscaleSetting_Kind)
)

func init() {
	SchemeBuilder.Register(&MonitorAutoscaleSetting{}, &MonitorAutoscaleSettingList{})
}
