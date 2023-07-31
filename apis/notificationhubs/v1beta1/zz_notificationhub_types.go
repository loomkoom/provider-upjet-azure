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

type APNSCredentialInitParameters struct {

	// The Application Mode which defines which server the APNS Messages should be sent to. Possible values are Production and Sandbox.
	ApplicationMode *string `json:"applicationMode,omitempty" tf:"application_mode,omitempty"`

	// The Bundle ID of the iOS/macOS application to send push notifications for, such as com.hashicorp.example.
	BundleID *string `json:"bundleId,omitempty" tf:"bundle_id,omitempty"`

	// The Apple Push Notifications Service (APNS) Key.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// The ID of the team the Token.
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`
}

type APNSCredentialObservation struct {

	// The Application Mode which defines which server the APNS Messages should be sent to. Possible values are Production and Sandbox.
	ApplicationMode *string `json:"applicationMode,omitempty" tf:"application_mode,omitempty"`

	// The Bundle ID of the iOS/macOS application to send push notifications for, such as com.hashicorp.example.
	BundleID *string `json:"bundleId,omitempty" tf:"bundle_id,omitempty"`

	// The Apple Push Notifications Service (APNS) Key.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// The ID of the team the Token.
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`
}

type APNSCredentialParameters struct {

	// The Application Mode which defines which server the APNS Messages should be sent to. Possible values are Production and Sandbox.
	// +kubebuilder:validation:Optional
	ApplicationMode *string `json:"applicationMode,omitempty" tf:"application_mode,omitempty"`

	// The Bundle ID of the iOS/macOS application to send push notifications for, such as com.hashicorp.example.
	// +kubebuilder:validation:Optional
	BundleID *string `json:"bundleId,omitempty" tf:"bundle_id,omitempty"`

	// The Apple Push Notifications Service (APNS) Key.
	// +kubebuilder:validation:Optional
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// The ID of the team the Token.
	// +kubebuilder:validation:Optional
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// The Push Token associated with the Apple Developer Account. This is the contents of the key downloaded from the Apple Developer Portal between the -----BEGIN PRIVATE KEY----- and -----END PRIVATE KEY----- blocks.
	// +kubebuilder:validation:Required
	TokenSecretRef v1.SecretKeySelector `json:"tokenSecretRef" tf:"-"`
}

type GCMCredentialInitParameters struct {
}

type GCMCredentialObservation struct {
}

type GCMCredentialParameters struct {

	// The API Key associated with the Google Cloud Messaging service.
	// +kubebuilder:validation:Required
	APIKeySecretRef v1.SecretKeySelector `json:"apiKeySecretRef" tf:"-"`
}

type NotificationHubInitParameters struct {

	// A apns_credential block as defined below.
	APNSCredential []APNSCredentialInitParameters `json:"apnsCredential,omitempty" tf:"apns_credential,omitempty"`

	// A gcm_credential block as defined below.
	GCMCredential []GCMCredentialInitParameters `json:"gcmCredential,omitempty" tf:"gcm_credential,omitempty"`

	// The Azure Region in which this Notification Hub Namespace exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type NotificationHubObservation struct {

	// A apns_credential block as defined below.
	APNSCredential []APNSCredentialObservation `json:"apnsCredential,omitempty" tf:"apns_credential,omitempty"`

	// A gcm_credential block as defined below.
	GCMCredential []GCMCredentialParameters `json:"gcmCredential,omitempty" tf:"gcm_credential,omitempty"`

	// The ID of the Notification Hub.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Azure Region in which this Notification Hub Namespace exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Notification Hub Namespace in which to create this Notification Hub. Changing this forces a new resource to be created.
	NamespaceName *string `json:"namespaceName,omitempty" tf:"namespace_name,omitempty"`

	// The name of the Resource Group in which the Notification Hub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type NotificationHubParameters struct {

	// A apns_credential block as defined below.
	// +kubebuilder:validation:Optional
	APNSCredential []APNSCredentialParameters `json:"apnsCredential,omitempty" tf:"apns_credential,omitempty"`

	// A gcm_credential block as defined below.
	// +kubebuilder:validation:Optional
	GCMCredential []GCMCredentialParameters `json:"gcmCredential,omitempty" tf:"gcm_credential,omitempty"`

	// The Azure Region in which this Notification Hub Namespace exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Notification Hub Namespace in which to create this Notification Hub. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/notificationhubs/v1beta1.NotificationHubNamespace
	// +kubebuilder:validation:Optional
	NamespaceName *string `json:"namespaceName,omitempty" tf:"namespace_name,omitempty"`

	// Reference to a NotificationHubNamespace in notificationhubs to populate namespaceName.
	// +kubebuilder:validation:Optional
	NamespaceNameRef *v1.Reference `json:"namespaceNameRef,omitempty" tf:"-"`

	// Selector for a NotificationHubNamespace in notificationhubs to populate namespaceName.
	// +kubebuilder:validation:Optional
	NamespaceNameSelector *v1.Selector `json:"namespaceNameSelector,omitempty" tf:"-"`

	// The name of the Resource Group in which the Notification Hub Namespace exists. Changing this forces a new resource to be created.
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

// NotificationHubSpec defines the desired state of NotificationHub
type NotificationHubSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NotificationHubParameters `json:"forProvider"`
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
	InitProvider NotificationHubInitParameters `json:"initProvider,omitempty"`
}

// NotificationHubStatus defines the observed state of NotificationHub.
type NotificationHubStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NotificationHubObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NotificationHub is the Schema for the NotificationHubs API. Manages a Notification Hub within a Notification Hub Namespace.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type NotificationHub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	Spec   NotificationHubSpec   `json:"spec"`
	Status NotificationHubStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NotificationHubList contains a list of NotificationHubs
type NotificationHubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotificationHub `json:"items"`
}

// Repository type metadata.
var (
	NotificationHub_Kind             = "NotificationHub"
	NotificationHub_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NotificationHub_Kind}.String()
	NotificationHub_KindAPIVersion   = NotificationHub_Kind + "." + CRDGroupVersion.String()
	NotificationHub_GroupVersionKind = CRDGroupVersion.WithKind(NotificationHub_Kind)
)

func init() {
	SchemeBuilder.Register(&NotificationHub{}, &NotificationHubList{})
}
