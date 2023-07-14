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

type TopicIdentityObservation struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Event Grid Topic.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Principal ID associated with this Managed Service Identity.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID associated with this Managed Service Identity.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Event Grid Topic. Possible values are SystemAssigned, UserAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TopicIdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Event Grid Topic.
	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Event Grid Topic. Possible values are SystemAssigned, UserAssigned.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type TopicInboundIPRuleObservation struct {

	// The action to take when the rule is matched. Possible values are Allow.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The IP mask (CIDR) to match on.
	IPMask *string `json:"ipMask,omitempty" tf:"ip_mask,omitempty"`
}

type TopicInboundIPRuleParameters struct {

	// The action to take when the rule is matched. Possible values are Allow.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action"`

	// The IP mask (CIDR) to match on.
	// +kubebuilder:validation:Optional
	IPMask *string `json:"ipMask,omitempty" tf:"ip_mask"`
}

type TopicInputMappingDefaultValuesObservation struct {

	// Specifies the default data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
	DataVersion *string `json:"dataVersion,omitempty" tf:"data_version,omitempty"`

	// Specifies the default event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
	EventType *string `json:"eventType,omitempty" tf:"event_type,omitempty"`

	// Specifies the default subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`
}

type TopicInputMappingDefaultValuesParameters struct {

	// Specifies the default data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DataVersion *string `json:"dataVersion,omitempty" tf:"data_version,omitempty"`

	// Specifies the default event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	EventType *string `json:"eventType,omitempty" tf:"event_type,omitempty"`

	// Specifies the default subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`
}

type TopicInputMappingFieldsObservation struct {

	// Specifies the data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
	DataVersion *string `json:"dataVersion,omitempty" tf:"data_version,omitempty"`

	// Specifies the event time of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
	EventTime *string `json:"eventTime,omitempty" tf:"event_time,omitempty"`

	// Specifies the event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
	EventType *string `json:"eventType,omitempty" tf:"event_type,omitempty"`

	// Specifies the id of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// Specifies the topic of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`
}

type TopicInputMappingFieldsParameters struct {

	// Specifies the data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DataVersion *string `json:"dataVersion,omitempty" tf:"data_version,omitempty"`

	// Specifies the event time of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	EventTime *string `json:"eventTime,omitempty" tf:"event_time,omitempty"`

	// Specifies the event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	EventType *string `json:"eventType,omitempty" tf:"event_type,omitempty"`

	// Specifies the id of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// Specifies the topic of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`
}

type TopicObservation struct {

	// The Endpoint associated with the EventGrid Topic.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The EventGrid Topic ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity []TopicIdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// One or more inbound_ip_rule blocks as defined below.
	InboundIPRule []TopicInboundIPRuleObservation `json:"inboundIpRule,omitempty" tf:"inbound_ip_rule,omitempty"`

	// A input_mapping_default_values block as defined below. Changing this forces a new resource to be created.
	InputMappingDefaultValues []TopicInputMappingDefaultValuesObservation `json:"inputMappingDefaultValues,omitempty" tf:"input_mapping_default_values,omitempty"`

	// A input_mapping_fields block as defined below. Changing this forces a new resource to be created.
	InputMappingFields []TopicInputMappingFieldsObservation `json:"inputMappingFields,omitempty" tf:"input_mapping_fields,omitempty"`

	// Specifies the schema in which incoming events will be published to this domain. Allowed values are CloudEventSchemaV1_0, CustomEventSchema, or EventGridSchema. Defaults to EventGridSchema. Changing this forces a new resource to be created.
	InputSchema *string `json:"inputSchema,omitempty" tf:"input_schema,omitempty"`

	// Whether local authentication methods is enabled for the EventGrid Topic. Defaults to true.
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Whether or not public network access is allowed for this server. Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type TopicParameters struct {

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []TopicIdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// One or more inbound_ip_rule blocks as defined below.
	// +kubebuilder:validation:Optional
	InboundIPRule []TopicInboundIPRuleParameters `json:"inboundIpRule,omitempty" tf:"inbound_ip_rule,omitempty"`

	// A input_mapping_default_values block as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	InputMappingDefaultValues []TopicInputMappingDefaultValuesParameters `json:"inputMappingDefaultValues,omitempty" tf:"input_mapping_default_values,omitempty"`

	// A input_mapping_fields block as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	InputMappingFields []TopicInputMappingFieldsParameters `json:"inputMappingFields,omitempty" tf:"input_mapping_fields,omitempty"`

	// Specifies the schema in which incoming events will be published to this domain. Allowed values are CloudEventSchemaV1_0, CustomEventSchema, or EventGridSchema. Defaults to EventGridSchema. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	InputSchema *string `json:"inputSchema,omitempty" tf:"input_schema,omitempty"`

	// Whether local authentication methods is enabled for the EventGrid Topic. Defaults to true.
	// +kubebuilder:validation:Optional
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Whether or not public network access is allowed for this server. Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.
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
}

// TopicSpec defines the desired state of Topic
type TopicSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TopicParameters `json:"forProvider"`
}

// TopicStatus defines the observed state of Topic.
type TopicStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TopicObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Topic is the Schema for the Topics API. Manages an EventGrid Topic
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Topic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location)",message="location is a required parameter"
	Spec   TopicSpec   `json:"spec"`
	Status TopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TopicList contains a list of Topics
type TopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Topic `json:"items"`
}

// Repository type metadata.
var (
	Topic_Kind             = "Topic"
	Topic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Topic_Kind}.String()
	Topic_KindAPIVersion   = Topic_Kind + "." + CRDGroupVersion.String()
	Topic_GroupVersionKind = CRDGroupVersion.WithKind(Topic_Kind)
)

func init() {
	SchemeBuilder.Register(&Topic{}, &TopicList{})
}
