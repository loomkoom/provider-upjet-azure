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

type EncryptionObservation struct {

	// The current key used to encrypt the Media Services Account, including the key version.
	CurrentKeyIdentifier *string `json:"currentKeyIdentifier,omitempty" tf:"current_key_identifier,omitempty"`

	// Specifies the URI of the Key Vault Key used to encrypt data. The key may either be versioned (for example https://vault/keys/mykey/version1) or reference a key without a version (for example https://vault/keys/mykey).
	KeyVaultKeyIdentifier *string `json:"keyVaultKeyIdentifier,omitempty" tf:"key_vault_key_identifier,omitempty"`

	// A managed_identity block as defined below.
	ManagedIdentity []ManagedIdentityObservation `json:"managedIdentity,omitempty" tf:"managed_identity,omitempty"`

	// Specifies the type of key used to encrypt the account data. Possible values are SystemKey and CustomerKey.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type EncryptionParameters struct {

	// Specifies the URI of the Key Vault Key used to encrypt data. The key may either be versioned (for example https://vault/keys/mykey/version1) or reference a key without a version (for example https://vault/keys/mykey).
	// +kubebuilder:validation:Optional
	KeyVaultKeyIdentifier *string `json:"keyVaultKeyIdentifier,omitempty" tf:"key_vault_key_identifier,omitempty"`

	// A managed_identity block as defined below.
	// +kubebuilder:validation:Optional
	ManagedIdentity []ManagedIdentityParameters `json:"managedIdentity,omitempty" tf:"managed_identity,omitempty"`

	// Specifies the type of key used to encrypt the account data. Possible values are SystemKey and CustomerKey.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityObservation struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Media Services Account.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Principal ID associated with this Managed Service Identity.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID associated with this Managed Service Identity.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Media Services Account. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Media Services Account.
	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Media Services Account. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type KeyDeliveryAccessControlObservation struct {

	// The Default Action to use when no rules match from ip_allow_list. Possible values are Allow and Deny.
	DefaultAction *string `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// One or more IP Addresses, or CIDR Blocks which should be able to access the Key Delivery.
	IPAllowList []*string `json:"ipAllowList,omitempty" tf:"ip_allow_list,omitempty"`
}

type KeyDeliveryAccessControlParameters struct {

	// The Default Action to use when no rules match from ip_allow_list. Possible values are Allow and Deny.
	// +kubebuilder:validation:Optional
	DefaultAction *string `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// One or more IP Addresses, or CIDR Blocks which should be able to access the Key Delivery.
	// +kubebuilder:validation:Optional
	IPAllowList []*string `json:"ipAllowList,omitempty" tf:"ip_allow_list,omitempty"`
}

type ManagedIdentityObservation struct {

	// Whether to use System Assigned Identity. Possible Values are true and false.
	UseSystemAssignedIdentity *bool `json:"useSystemAssignedIdentity,omitempty" tf:"use_system_assigned_identity,omitempty"`

	// The ID of the User Assigned Identity. This value can only be set when use_system_assigned_identity is false
	UserAssignedIdentityID *string `json:"userAssignedIdentityId,omitempty" tf:"user_assigned_identity_id,omitempty"`
}

type ManagedIdentityParameters struct {

	// Whether to use System Assigned Identity. Possible Values are true and false.
	// +kubebuilder:validation:Optional
	UseSystemAssignedIdentity *bool `json:"useSystemAssignedIdentity,omitempty" tf:"use_system_assigned_identity,omitempty"`

	// The ID of the User Assigned Identity. This value can only be set when use_system_assigned_identity is false
	// +kubebuilder:validation:Optional
	UserAssignedIdentityID *string `json:"userAssignedIdentityId,omitempty" tf:"user_assigned_identity_id,omitempty"`
}

type ServicesAccountObservation struct {

	// An encryption block as defined below.
	Encryption []EncryptionObservation `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// The ID of the Media Services Account.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// A key_delivery_access_control block as defined below.
	KeyDeliveryAccessControl []KeyDeliveryAccessControlObservation `json:"keyDeliveryAccessControl,omitempty" tf:"key_delivery_access_control,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Whether public network access is allowed for this server. Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the resource group in which to create the Media Services Account. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// One or more storage_account blocks as defined below.
	StorageAccount []StorageAccountObservation `json:"storageAccount,omitempty" tf:"storage_account,omitempty"`

	// Specifies the storage authentication type. Possible value is ManagedIdentity or System.
	StorageAuthenticationType *string `json:"storageAuthenticationType,omitempty" tf:"storage_authentication_type,omitempty"`

	// A mapping of tags assigned to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ServicesAccountParameters struct {

	// An encryption block as defined below.
	// +kubebuilder:validation:Optional
	Encryption []EncryptionParameters `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// A key_delivery_access_control block as defined below.
	// +kubebuilder:validation:Optional
	KeyDeliveryAccessControl []KeyDeliveryAccessControlParameters `json:"keyDeliveryAccessControl,omitempty" tf:"key_delivery_access_control,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Whether public network access is allowed for this server. Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the resource group in which to create the Media Services Account. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// One or more storage_account blocks as defined below.
	// +kubebuilder:validation:Optional
	StorageAccount []StorageAccountParameters `json:"storageAccount,omitempty" tf:"storage_account,omitempty"`

	// Specifies the storage authentication type. Possible value is ManagedIdentity or System.
	// +kubebuilder:validation:Optional
	StorageAuthenticationType *string `json:"storageAuthenticationType,omitempty" tf:"storage_authentication_type,omitempty"`

	// A mapping of tags assigned to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type StorageAccountManagedIdentityObservation struct {

	// Whether to use System Assigned Identity. Possible Values are true and false.
	UseSystemAssignedIdentity *bool `json:"useSystemAssignedIdentity,omitempty" tf:"use_system_assigned_identity,omitempty"`

	// The ID of the User Assigned Identity. This value can only be set when use_system_assigned_identity is false
	UserAssignedIdentityID *string `json:"userAssignedIdentityId,omitempty" tf:"user_assigned_identity_id,omitempty"`
}

type StorageAccountManagedIdentityParameters struct {

	// Whether to use System Assigned Identity. Possible Values are true and false.
	// +kubebuilder:validation:Optional
	UseSystemAssignedIdentity *bool `json:"useSystemAssignedIdentity,omitempty" tf:"use_system_assigned_identity,omitempty"`

	// The ID of the User Assigned Identity. This value can only be set when use_system_assigned_identity is false
	// +kubebuilder:validation:Optional
	UserAssignedIdentityID *string `json:"userAssignedIdentityId,omitempty" tf:"user_assigned_identity_id,omitempty"`
}

type StorageAccountObservation struct {

	// Specifies the ID of the Storage Account that will be associated with the Media Services instance.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies whether the storage account should be the primary account or not. Defaults to false.
	IsPrimary *bool `json:"isPrimary,omitempty" tf:"is_primary,omitempty"`

	// A managed_identity block as defined below.
	ManagedIdentity []StorageAccountManagedIdentityObservation `json:"managedIdentity,omitempty" tf:"managed_identity,omitempty"`
}

type StorageAccountParameters struct {

	// Specifies the ID of the Storage Account that will be associated with the Media Services instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a Account in storage to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`

	// Specifies whether the storage account should be the primary account or not. Defaults to false.
	// +kubebuilder:validation:Optional
	IsPrimary *bool `json:"isPrimary,omitempty" tf:"is_primary,omitempty"`

	// A managed_identity block as defined below.
	// +kubebuilder:validation:Optional
	ManagedIdentity []StorageAccountManagedIdentityParameters `json:"managedIdentity,omitempty" tf:"managed_identity,omitempty"`
}

// ServicesAccountSpec defines the desired state of ServicesAccount
type ServicesAccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServicesAccountParameters `json:"forProvider"`
}

// ServicesAccountStatus defines the observed state of ServicesAccount.
type ServicesAccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServicesAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServicesAccount is the Schema for the ServicesAccounts API. Manages a Media Services Account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ServicesAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storageAccount)",message="storageAccount is a required parameter"
	Spec   ServicesAccountSpec   `json:"spec"`
	Status ServicesAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServicesAccountList contains a list of ServicesAccounts
type ServicesAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServicesAccount `json:"items"`
}

// Repository type metadata.
var (
	ServicesAccount_Kind             = "ServicesAccount"
	ServicesAccount_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServicesAccount_Kind}.String()
	ServicesAccount_KindAPIVersion   = ServicesAccount_Kind + "." + CRDGroupVersion.String()
	ServicesAccount_GroupVersionKind = CRDGroupVersion.WithKind(ServicesAccount_Kind)
)

func init() {
	SchemeBuilder.Register(&ServicesAccount{}, &ServicesAccountList{})
}
