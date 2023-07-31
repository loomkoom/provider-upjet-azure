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

type QueueInitParameters struct {

	// The ISO 8601 timespan duration of the idle interval after which the Queue is automatically deleted, minimum of 5 minutes.
	AutoDeleteOnIdle *string `json:"autoDeleteOnIdle,omitempty" tf:"auto_delete_on_idle,omitempty"`

	// Boolean flag which controls whether the Queue has dead letter support when a message expires. Defaults to false.
	DeadLetteringOnMessageExpiration *bool `json:"deadLetteringOnMessageExpiration,omitempty" tf:"dead_lettering_on_message_expiration,omitempty"`

	// The ISO 8601 timespan duration of the TTL of messages sent to this queue. This is the default value used when TTL is not set on message itself.
	DefaultMessageTTL *string `json:"defaultMessageTtl,omitempty" tf:"default_message_ttl,omitempty"`

	// The ISO 8601 timespan duration during which duplicates can be detected. Defaults to 10 minutes (PT10M).
	DuplicateDetectionHistoryTimeWindow *string `json:"duplicateDetectionHistoryTimeWindow,omitempty" tf:"duplicate_detection_history_time_window,omitempty"`

	// Boolean flag which controls whether server-side batched operations are enabled. Defaults to true.
	EnableBatchedOperations *bool `json:"enableBatchedOperations,omitempty" tf:"enable_batched_operations,omitempty"`

	// Boolean flag which controls whether Express Entities are enabled. An express queue holds a message in memory temporarily before writing it to persistent storage. Defaults to false for Basic and Standard. For Premium, it MUST be set to false.
	EnableExpress *bool `json:"enableExpress,omitempty" tf:"enable_express,omitempty"`

	// Boolean flag which controls whether to enable the queue to be partitioned across multiple message brokers. Changing this forces a new resource to be created. Defaults to false for Basic and Standard.
	EnablePartitioning *bool `json:"enablePartitioning,omitempty" tf:"enable_partitioning,omitempty"`

	// The name of a Queue or Topic to automatically forward dead lettered messages to.
	ForwardDeadLetteredMessagesTo *string `json:"forwardDeadLetteredMessagesTo,omitempty" tf:"forward_dead_lettered_messages_to,omitempty"`

	// The name of a Queue or Topic to automatically forward messages to. Please see the documentation for more information.
	ForwardTo *string `json:"forwardTo,omitempty" tf:"forward_to,omitempty"`

	// The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. Maximum value is 5 minutes. Defaults to 1 minute (PT1M).
	LockDuration *string `json:"lockDuration,omitempty" tf:"lock_duration,omitempty"`

	// Integer value which controls when a message is automatically dead lettered. Defaults to 10.
	MaxDeliveryCount *float64 `json:"maxDeliveryCount,omitempty" tf:"max_delivery_count,omitempty"`

	// Integer value which controls the maximum size of a message allowed on the queue for Premium SKU. For supported values see the "Large messages support" section of this document.
	MaxMessageSizeInKilobytes *float64 `json:"maxMessageSizeInKilobytes,omitempty" tf:"max_message_size_in_kilobytes,omitempty"`

	// Integer value which controls the size of memory allocated for the queue. For supported values see the "Queue or topic size" section of Service Bus Quotas. Defaults to 1024.
	MaxSizeInMegabytes *float64 `json:"maxSizeInMegabytes,omitempty" tf:"max_size_in_megabytes,omitempty"`

	// Boolean flag which controls whether the Queue requires duplicate detection. Changing this forces a new resource to be created. Defaults to false.
	RequiresDuplicateDetection *bool `json:"requiresDuplicateDetection,omitempty" tf:"requires_duplicate_detection,omitempty"`

	// Boolean flag which controls whether the Queue requires sessions. This will allow ordered handling of unbounded sequences of related messages. With sessions enabled a queue can guarantee first-in-first-out delivery of messages. Changing this forces a new resource to be created. Defaults to false.
	RequiresSession *bool `json:"requiresSession,omitempty" tf:"requires_session,omitempty"`

	// The status of the Queue. Possible values are Active, Creating, Deleting, Disabled, ReceiveDisabled, Renaming, SendDisabled, Unknown. Note that Restoring is not accepted. Defaults to Active.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type QueueObservation struct {

	// The ISO 8601 timespan duration of the idle interval after which the Queue is automatically deleted, minimum of 5 minutes.
	AutoDeleteOnIdle *string `json:"autoDeleteOnIdle,omitempty" tf:"auto_delete_on_idle,omitempty"`

	// Boolean flag which controls whether the Queue has dead letter support when a message expires. Defaults to false.
	DeadLetteringOnMessageExpiration *bool `json:"deadLetteringOnMessageExpiration,omitempty" tf:"dead_lettering_on_message_expiration,omitempty"`

	// The ISO 8601 timespan duration of the TTL of messages sent to this queue. This is the default value used when TTL is not set on message itself.
	DefaultMessageTTL *string `json:"defaultMessageTtl,omitempty" tf:"default_message_ttl,omitempty"`

	// The ISO 8601 timespan duration during which duplicates can be detected. Defaults to 10 minutes (PT10M).
	DuplicateDetectionHistoryTimeWindow *string `json:"duplicateDetectionHistoryTimeWindow,omitempty" tf:"duplicate_detection_history_time_window,omitempty"`

	// Boolean flag which controls whether server-side batched operations are enabled. Defaults to true.
	EnableBatchedOperations *bool `json:"enableBatchedOperations,omitempty" tf:"enable_batched_operations,omitempty"`

	// Boolean flag which controls whether Express Entities are enabled. An express queue holds a message in memory temporarily before writing it to persistent storage. Defaults to false for Basic and Standard. For Premium, it MUST be set to false.
	EnableExpress *bool `json:"enableExpress,omitempty" tf:"enable_express,omitempty"`

	// Boolean flag which controls whether to enable the queue to be partitioned across multiple message brokers. Changing this forces a new resource to be created. Defaults to false for Basic and Standard.
	EnablePartitioning *bool `json:"enablePartitioning,omitempty" tf:"enable_partitioning,omitempty"`

	// The name of a Queue or Topic to automatically forward dead lettered messages to.
	ForwardDeadLetteredMessagesTo *string `json:"forwardDeadLetteredMessagesTo,omitempty" tf:"forward_dead_lettered_messages_to,omitempty"`

	// The name of a Queue or Topic to automatically forward messages to. Please see the documentation for more information.
	ForwardTo *string `json:"forwardTo,omitempty" tf:"forward_to,omitempty"`

	// The ServiceBus Queue ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. Maximum value is 5 minutes. Defaults to 1 minute (PT1M).
	LockDuration *string `json:"lockDuration,omitempty" tf:"lock_duration,omitempty"`

	// Integer value which controls when a message is automatically dead lettered. Defaults to 10.
	MaxDeliveryCount *float64 `json:"maxDeliveryCount,omitempty" tf:"max_delivery_count,omitempty"`

	// Integer value which controls the maximum size of a message allowed on the queue for Premium SKU. For supported values see the "Large messages support" section of this document.
	MaxMessageSizeInKilobytes *float64 `json:"maxMessageSizeInKilobytes,omitempty" tf:"max_message_size_in_kilobytes,omitempty"`

	// Integer value which controls the size of memory allocated for the queue. For supported values see the "Queue or topic size" section of Service Bus Quotas. Defaults to 1024.
	MaxSizeInMegabytes *float64 `json:"maxSizeInMegabytes,omitempty" tf:"max_size_in_megabytes,omitempty"`

	// The ID of the ServiceBus Namespace to create this queue in. Changing this forces a new resource to be created.
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// Boolean flag which controls whether the Queue requires duplicate detection. Changing this forces a new resource to be created. Defaults to false.
	RequiresDuplicateDetection *bool `json:"requiresDuplicateDetection,omitempty" tf:"requires_duplicate_detection,omitempty"`

	// Boolean flag which controls whether the Queue requires sessions. This will allow ordered handling of unbounded sequences of related messages. With sessions enabled a queue can guarantee first-in-first-out delivery of messages. Changing this forces a new resource to be created. Defaults to false.
	RequiresSession *bool `json:"requiresSession,omitempty" tf:"requires_session,omitempty"`

	// The status of the Queue. Possible values are Active, Creating, Deleting, Disabled, ReceiveDisabled, Renaming, SendDisabled, Unknown. Note that Restoring is not accepted. Defaults to Active.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type QueueParameters struct {

	// The ISO 8601 timespan duration of the idle interval after which the Queue is automatically deleted, minimum of 5 minutes.
	// +kubebuilder:validation:Optional
	AutoDeleteOnIdle *string `json:"autoDeleteOnIdle,omitempty" tf:"auto_delete_on_idle,omitempty"`

	// Boolean flag which controls whether the Queue has dead letter support when a message expires. Defaults to false.
	// +kubebuilder:validation:Optional
	DeadLetteringOnMessageExpiration *bool `json:"deadLetteringOnMessageExpiration,omitempty" tf:"dead_lettering_on_message_expiration,omitempty"`

	// The ISO 8601 timespan duration of the TTL of messages sent to this queue. This is the default value used when TTL is not set on message itself.
	// +kubebuilder:validation:Optional
	DefaultMessageTTL *string `json:"defaultMessageTtl,omitempty" tf:"default_message_ttl,omitempty"`

	// The ISO 8601 timespan duration during which duplicates can be detected. Defaults to 10 minutes (PT10M).
	// +kubebuilder:validation:Optional
	DuplicateDetectionHistoryTimeWindow *string `json:"duplicateDetectionHistoryTimeWindow,omitempty" tf:"duplicate_detection_history_time_window,omitempty"`

	// Boolean flag which controls whether server-side batched operations are enabled. Defaults to true.
	// +kubebuilder:validation:Optional
	EnableBatchedOperations *bool `json:"enableBatchedOperations,omitempty" tf:"enable_batched_operations,omitempty"`

	// Boolean flag which controls whether Express Entities are enabled. An express queue holds a message in memory temporarily before writing it to persistent storage. Defaults to false for Basic and Standard. For Premium, it MUST be set to false.
	// +kubebuilder:validation:Optional
	EnableExpress *bool `json:"enableExpress,omitempty" tf:"enable_express,omitempty"`

	// Boolean flag which controls whether to enable the queue to be partitioned across multiple message brokers. Changing this forces a new resource to be created. Defaults to false for Basic and Standard.
	// +kubebuilder:validation:Optional
	EnablePartitioning *bool `json:"enablePartitioning,omitempty" tf:"enable_partitioning,omitempty"`

	// The name of a Queue or Topic to automatically forward dead lettered messages to.
	// +kubebuilder:validation:Optional
	ForwardDeadLetteredMessagesTo *string `json:"forwardDeadLetteredMessagesTo,omitempty" tf:"forward_dead_lettered_messages_to,omitempty"`

	// The name of a Queue or Topic to automatically forward messages to. Please see the documentation for more information.
	// +kubebuilder:validation:Optional
	ForwardTo *string `json:"forwardTo,omitempty" tf:"forward_to,omitempty"`

	// The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. Maximum value is 5 minutes. Defaults to 1 minute (PT1M).
	// +kubebuilder:validation:Optional
	LockDuration *string `json:"lockDuration,omitempty" tf:"lock_duration,omitempty"`

	// Integer value which controls when a message is automatically dead lettered. Defaults to 10.
	// +kubebuilder:validation:Optional
	MaxDeliveryCount *float64 `json:"maxDeliveryCount,omitempty" tf:"max_delivery_count,omitempty"`

	// Integer value which controls the maximum size of a message allowed on the queue for Premium SKU. For supported values see the "Large messages support" section of this document.
	// +kubebuilder:validation:Optional
	MaxMessageSizeInKilobytes *float64 `json:"maxMessageSizeInKilobytes,omitempty" tf:"max_message_size_in_kilobytes,omitempty"`

	// Integer value which controls the size of memory allocated for the queue. For supported values see the "Queue or topic size" section of Service Bus Quotas. Defaults to 1024.
	// +kubebuilder:validation:Optional
	MaxSizeInMegabytes *float64 `json:"maxSizeInMegabytes,omitempty" tf:"max_size_in_megabytes,omitempty"`

	// The ID of the ServiceBus Namespace to create this queue in. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/servicebus/v1beta1.ServiceBusNamespace
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// Reference to a ServiceBusNamespace in servicebus to populate namespaceId.
	// +kubebuilder:validation:Optional
	NamespaceIDRef *v1.Reference `json:"namespaceIdRef,omitempty" tf:"-"`

	// Selector for a ServiceBusNamespace in servicebus to populate namespaceId.
	// +kubebuilder:validation:Optional
	NamespaceIDSelector *v1.Selector `json:"namespaceIdSelector,omitempty" tf:"-"`

	// Boolean flag which controls whether the Queue requires duplicate detection. Changing this forces a new resource to be created. Defaults to false.
	// +kubebuilder:validation:Optional
	RequiresDuplicateDetection *bool `json:"requiresDuplicateDetection,omitempty" tf:"requires_duplicate_detection,omitempty"`

	// Boolean flag which controls whether the Queue requires sessions. This will allow ordered handling of unbounded sequences of related messages. With sessions enabled a queue can guarantee first-in-first-out delivery of messages. Changing this forces a new resource to be created. Defaults to false.
	// +kubebuilder:validation:Optional
	RequiresSession *bool `json:"requiresSession,omitempty" tf:"requires_session,omitempty"`

	// The status of the Queue. Possible values are Active, Creating, Deleting, Disabled, ReceiveDisabled, Renaming, SendDisabled, Unknown. Note that Restoring is not accepted. Defaults to Active.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

// QueueSpec defines the desired state of Queue
type QueueSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QueueParameters `json:"forProvider"`
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
	InitProvider QueueInitParameters `json:"initProvider,omitempty"`
}

// QueueStatus defines the observed state of Queue.
type QueueStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QueueObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Queue is the Schema for the Queues API. Manages a ServiceBus Queue.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Queue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              QueueSpec   `json:"spec"`
	Status            QueueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QueueList contains a list of Queues
type QueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Queue `json:"items"`
}

// Repository type metadata.
var (
	Queue_Kind             = "Queue"
	Queue_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Queue_Kind}.String()
	Queue_KindAPIVersion   = Queue_Kind + "." + CRDGroupVersion.String()
	Queue_GroupVersionKind = CRDGroupVersion.WithKind(Queue_Kind)
)

func init() {
	SchemeBuilder.Register(&Queue{}, &QueueList{})
}
