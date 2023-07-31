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

type ACLInitParameters struct {

	// An access_policy block as defined below.
	AccessPolicy []AccessPolicyInitParameters `json:"accessPolicy,omitempty" tf:"access_policy,omitempty"`

	// The ID which should be used for this Shared Identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ACLObservation struct {

	// An access_policy block as defined below.
	AccessPolicy []AccessPolicyObservation `json:"accessPolicy,omitempty" tf:"access_policy,omitempty"`

	// The ID which should be used for this Shared Identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ACLParameters struct {

	// An access_policy block as defined below.
	// +kubebuilder:validation:Optional
	AccessPolicy []AccessPolicyParameters `json:"accessPolicy,omitempty" tf:"access_policy,omitempty"`

	// The ID which should be used for this Shared Identifier.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AccessPolicyInitParameters struct {

	// The time at which this Access Policy should be valid until, in ISO8601 format.
	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`

	// The permissions which should be associated with this Shared Identifier. Possible value is combination of r (read), w (write), d (delete), and l (list).
	Permissions *string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// The time at which this Access Policy should be valid from, in ISO8601 format.
	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type AccessPolicyObservation struct {

	// The time at which this Access Policy should be valid until, in ISO8601 format.
	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`

	// The permissions which should be associated with this Shared Identifier. Possible value is combination of r (read), w (write), d (delete), and l (list).
	Permissions *string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// The time at which this Access Policy should be valid from, in ISO8601 format.
	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type AccessPolicyParameters struct {

	// The time at which this Access Policy should be valid until, in ISO8601 format.
	// +kubebuilder:validation:Optional
	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`

	// The permissions which should be associated with this Shared Identifier. Possible value is combination of r (read), w (write), d (delete), and l (list).
	// +kubebuilder:validation:Optional
	Permissions *string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// The time at which this Access Policy should be valid from, in ISO8601 format.
	// +kubebuilder:validation:Optional
	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type ShareInitParameters struct {

	// One or more acl blocks as defined below.
	ACL []ACLInitParameters `json:"acl,omitempty" tf:"acl,omitempty"`

	// The access tier of the File Share. Possible values are Hot, Cool and TransactionOptimized, Premium.
	AccessTier *string `json:"accessTier,omitempty" tf:"access_tier,omitempty"`

	// The protocol used for the share. Possible values are SMB and NFS. The SMB indicates the share can be accessed by SMBv3.0, SMBv2.1 and REST. The NFS indicates the share can be accessed by NFSv4.1. Defaults to SMB. Changing this forces a new resource to be created.
	EnabledProtocol *string `json:"enabledProtocol,omitempty" tf:"enabled_protocol,omitempty"`

	// A mapping of MetaData for this File Share.
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The maximum size of the share, in gigabytes. For Standard storage accounts, this must be 1GB (or higher) and at most 5120 GB (5 TB). For Premium FileStorage storage accounts, this must be greater than 100 GB and at most 102400 GB (100 TB).
	Quota *float64 `json:"quota,omitempty" tf:"quota,omitempty"`
}

type ShareObservation struct {

	// One or more acl blocks as defined below.
	ACL []ACLObservation `json:"acl,omitempty" tf:"acl,omitempty"`

	// The access tier of the File Share. Possible values are Hot, Cool and TransactionOptimized, Premium.
	AccessTier *string `json:"accessTier,omitempty" tf:"access_tier,omitempty"`

	// The protocol used for the share. Possible values are SMB and NFS. The SMB indicates the share can be accessed by SMBv3.0, SMBv2.1 and REST. The NFS indicates the share can be accessed by NFSv4.1. Defaults to SMB. Changing this forces a new resource to be created.
	EnabledProtocol *string `json:"enabledProtocol,omitempty" tf:"enabled_protocol,omitempty"`

	// The ID of the File Share.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A mapping of MetaData for this File Share.
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The maximum size of the share, in gigabytes. For Standard storage accounts, this must be 1GB (or higher) and at most 5120 GB (5 TB). For Premium FileStorage storage accounts, this must be greater than 100 GB and at most 102400 GB (100 TB).
	Quota *float64 `json:"quota,omitempty" tf:"quota,omitempty"`

	// The Resource Manager ID of this File Share.
	ResourceManagerID *string `json:"resourceManagerId,omitempty" tf:"resource_manager_id,omitempty"`

	// Specifies the storage account in which to create the share. Changing this forces a new resource to be created.
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// The URL of the File Share
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type ShareParameters struct {

	// One or more acl blocks as defined below.
	// +kubebuilder:validation:Optional
	ACL []ACLParameters `json:"acl,omitempty" tf:"acl,omitempty"`

	// The access tier of the File Share. Possible values are Hot, Cool and TransactionOptimized, Premium.
	// +kubebuilder:validation:Optional
	AccessTier *string `json:"accessTier,omitempty" tf:"access_tier,omitempty"`

	// The protocol used for the share. Possible values are SMB and NFS. The SMB indicates the share can be accessed by SMBv3.0, SMBv2.1 and REST. The NFS indicates the share can be accessed by NFSv4.1. Defaults to SMB. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	EnabledProtocol *string `json:"enabledProtocol,omitempty" tf:"enabled_protocol,omitempty"`

	// A mapping of MetaData for this File Share.
	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The maximum size of the share, in gigabytes. For Standard storage accounts, this must be 1GB (or higher) and at most 5120 GB (5 TB). For Premium FileStorage storage accounts, this must be greater than 100 GB and at most 102400 GB (100 TB).
	// +kubebuilder:validation:Optional
	Quota *float64 `json:"quota,omitempty" tf:"quota,omitempty"`

	// Specifies the storage account in which to create the share. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +kubebuilder:validation:Optional
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// Reference to a Account in storage to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameRef *v1.Reference `json:"storageAccountNameRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameSelector *v1.Selector `json:"storageAccountNameSelector,omitempty" tf:"-"`
}

// ShareSpec defines the desired state of Share
type ShareSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ShareParameters `json:"forProvider"`
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
	InitProvider ShareInitParameters `json:"initProvider,omitempty"`
}

// ShareStatus defines the observed state of Share.
type ShareStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ShareObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Share is the Schema for the Shares API. Manages a File Share within Azure Storage.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Share struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.quota) || has(self.initProvider.quota)",message="quota is a required parameter"
	Spec   ShareSpec   `json:"spec"`
	Status ShareStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ShareList contains a list of Shares
type ShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Share `json:"items"`
}

// Repository type metadata.
var (
	Share_Kind             = "Share"
	Share_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Share_Kind}.String()
	Share_KindAPIVersion   = Share_Kind + "." + CRDGroupVersion.String()
	Share_GroupVersionKind = CRDGroupVersion.WithKind(Share_Kind)
)

func init() {
	SchemeBuilder.Register(&Share{}, &ShareList{})
}
