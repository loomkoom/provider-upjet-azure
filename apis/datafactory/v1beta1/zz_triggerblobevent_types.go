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

type TriggerBlobEventObservation struct {

	// The ID of the Data Factory Blob Event Trigger.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TriggerBlobEventParameters struct {

	// Specifies if the Data Factory Blob Event Trigger is activated. Defaults to true.
	// +kubebuilder:validation:Optional
	Activated *bool `json:"activated,omitempty" tf:"activated,omitempty"`

	// A map of additional properties to associate with the Data Factory Blob Event Trigger.
	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Blob Event Trigger.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The pattern that blob path starts with for trigger to fire.
	// +kubebuilder:validation:Optional
	BlobPathBeginsWith *string `json:"blobPathBeginsWith,omitempty" tf:"blob_path_begins_with,omitempty"`

	// The pattern that blob path ends with for trigger to fire.
	// +kubebuilder:validation:Optional
	BlobPathEndsWith *string `json:"blobPathEndsWith,omitempty" tf:"blob_path_ends_with,omitempty"`

	// The ID of Data Factory in which to associate the Trigger with. Changing this forces a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Factory
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Reference to a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDRef *v1.Reference `json:"dataFactoryIdRef,omitempty" tf:"-"`

	// Selector for a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDSelector *v1.Selector `json:"dataFactoryIdSelector,omitempty" tf:"-"`

	// The description for the Data Factory Blob Event Trigger.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// List of events that will fire this trigger. Possible values are Microsoft.Storage.BlobCreated and Microsoft.Storage.BlobDeleted.
	// +kubebuilder:validation:Required
	Events []*string `json:"events" tf:"events,omitempty"`

	// are blobs with zero bytes ignored?
	// +kubebuilder:validation:Optional
	IgnoreEmptyBlobs *bool `json:"ignoreEmptyBlobs,omitempty" tf:"ignore_empty_blobs,omitempty"`

	// One or more pipeline blocks as defined below.
	// +kubebuilder:validation:Required
	Pipeline []TriggerBlobEventPipelineParameters `json:"pipeline" tf:"pipeline,omitempty"`

	// The ID of Storage Account in which blob event will be listened. Changing this forces a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// Reference to a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDRef *v1.Reference `json:"storageAccountIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDSelector *v1.Selector `json:"storageAccountIdSelector,omitempty" tf:"-"`
}

type TriggerBlobEventPipelineObservation struct {
}

type TriggerBlobEventPipelineParameters struct {

	// The Data Factory Pipeline name that the trigger will act on.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Pipeline
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a Pipeline in datafactory to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a Pipeline in datafactory to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// The Data Factory Pipeline parameters that the trigger will act on.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

// TriggerBlobEventSpec defines the desired state of TriggerBlobEvent
type TriggerBlobEventSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TriggerBlobEventParameters `json:"forProvider"`
}

// TriggerBlobEventStatus defines the observed state of TriggerBlobEvent.
type TriggerBlobEventStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TriggerBlobEventObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TriggerBlobEvent is the Schema for the TriggerBlobEvents API. Manages a Blob Event Trigger inside an Azure Data Factory.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type TriggerBlobEvent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TriggerBlobEventSpec   `json:"spec"`
	Status            TriggerBlobEventStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TriggerBlobEventList contains a list of TriggerBlobEvents
type TriggerBlobEventList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TriggerBlobEvent `json:"items"`
}

// Repository type metadata.
var (
	TriggerBlobEvent_Kind             = "TriggerBlobEvent"
	TriggerBlobEvent_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TriggerBlobEvent_Kind}.String()
	TriggerBlobEvent_KindAPIVersion   = TriggerBlobEvent_Kind + "." + CRDGroupVersion.String()
	TriggerBlobEvent_GroupVersionKind = CRDGroupVersion.WithKind(TriggerBlobEvent_Kind)
)

func init() {
	SchemeBuilder.Register(&TriggerBlobEvent{}, &TriggerBlobEventList{})
}
