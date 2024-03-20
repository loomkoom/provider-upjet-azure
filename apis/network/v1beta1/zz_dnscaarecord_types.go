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

type DNSCAARecordInitParameters struct {

	// A list of values that make up the CAA record. Each record block supports fields documented below.
	Record []RecordInitParameters `json:"record,omitempty" tf:"record,omitempty"`

	// The Time To Live (TTL) of the DNS record in seconds.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DNSCAARecordObservation struct {

	// The FQDN of the DNS CAA Record.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// The DNS CAA Record ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of values that make up the CAA record. Each record block supports fields documented below.
	Record []RecordObservation `json:"record,omitempty" tf:"record,omitempty"`

	// Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The Time To Live (TTL) of the DNS record in seconds.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName *string `json:"zoneName,omitempty" tf:"zone_name,omitempty"`
}

type DNSCAARecordParameters struct {

	// A list of values that make up the CAA record. Each record block supports fields documented below.
	// +kubebuilder:validation:Optional
	Record []RecordParameters `json:"record,omitempty" tf:"record,omitempty"`

	// Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The Time To Live (TTL) of the DNS record in seconds.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=DNSZone
	// +kubebuilder:validation:Optional
	ZoneName *string `json:"zoneName,omitempty" tf:"zone_name,omitempty"`

	// Reference to a DNSZone to populate zoneName.
	// +kubebuilder:validation:Optional
	ZoneNameRef *v1.Reference `json:"zoneNameRef,omitempty" tf:"-"`

	// Selector for a DNSZone to populate zoneName.
	// +kubebuilder:validation:Optional
	ZoneNameSelector *v1.Selector `json:"zoneNameSelector,omitempty" tf:"-"`
}

type RecordInitParameters struct {

	// Extensible CAA flags, currently only 1 is implemented to set the issuer critical flag.
	Flags *float64 `json:"flags,omitempty" tf:"flags,omitempty"`

	// A property tag, options are issue, issuewild and iodef.
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// A property value such as a registrar domain.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type RecordObservation struct {

	// Extensible CAA flags, currently only 1 is implemented to set the issuer critical flag.
	Flags *float64 `json:"flags,omitempty" tf:"flags,omitempty"`

	// A property tag, options are issue, issuewild and iodef.
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// A property value such as a registrar domain.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type RecordParameters struct {

	// Extensible CAA flags, currently only 1 is implemented to set the issuer critical flag.
	// +kubebuilder:validation:Optional
	Flags *float64 `json:"flags" tf:"flags,omitempty"`

	// A property tag, options are issue, issuewild and iodef.
	// +kubebuilder:validation:Optional
	Tag *string `json:"tag" tf:"tag,omitempty"`

	// A property value such as a registrar domain.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

// DNSCAARecordSpec defines the desired state of DNSCAARecord
type DNSCAARecordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DNSCAARecordParameters `json:"forProvider"`
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
	InitProvider DNSCAARecordInitParameters `json:"initProvider,omitempty"`
}

// DNSCAARecordStatus defines the observed state of DNSCAARecord.
type DNSCAARecordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DNSCAARecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DNSCAARecord is the Schema for the DNSCAARecords API. Manages a DNS CAA Record.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DNSCAARecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.record) || (has(self.initProvider) && has(self.initProvider.record))",message="spec.forProvider.record is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ttl) || (has(self.initProvider) && has(self.initProvider.ttl))",message="spec.forProvider.ttl is a required parameter"
	Spec   DNSCAARecordSpec   `json:"spec"`
	Status DNSCAARecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DNSCAARecordList contains a list of DNSCAARecords
type DNSCAARecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DNSCAARecord `json:"items"`
}

// Repository type metadata.
var (
	DNSCAARecord_Kind             = "DNSCAARecord"
	DNSCAARecord_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DNSCAARecord_Kind}.String()
	DNSCAARecord_KindAPIVersion   = DNSCAARecord_Kind + "." + CRDGroupVersion.String()
	DNSCAARecord_GroupVersionKind = CRDGroupVersion.WithKind(DNSCAARecord_Kind)
)

func init() {
	SchemeBuilder.Register(&DNSCAARecord{}, &DNSCAARecordList{})
}
