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

type BotChannelWebChatInitParameters struct {

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A site represents a client application that you want to connect to your bot. One or more site blocks as defined below.
	Site []BotChannelWebChatSiteInitParameters `json:"site,omitempty" tf:"site,omitempty"`

	// A list of Web Chat Site names.
	// +listType=set
	SiteNames []*string `json:"siteNames,omitempty" tf:"site_names,omitempty"`
}

type BotChannelWebChatObservation struct {

	// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
	BotName *string `json:"botName,omitempty" tf:"bot_name,omitempty"`

	// The ID of the Web Chat Channel.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group where the Web Chat Channel should be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A site represents a client application that you want to connect to your bot. One or more site blocks as defined below.
	Site []BotChannelWebChatSiteObservation `json:"site,omitempty" tf:"site,omitempty"`

	// A list of Web Chat Site names.
	// +listType=set
	SiteNames []*string `json:"siteNames,omitempty" tf:"site_names,omitempty"`
}

type BotChannelWebChatParameters struct {

	// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/botservice/v1beta1.BotChannelsRegistration
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	BotName *string `json:"botName,omitempty" tf:"bot_name,omitempty"`

	// Reference to a BotChannelsRegistration in botservice to populate botName.
	// +kubebuilder:validation:Optional
	BotNameRef *v1.Reference `json:"botNameRef,omitempty" tf:"-"`

	// Selector for a BotChannelsRegistration in botservice to populate botName.
	// +kubebuilder:validation:Optional
	BotNameSelector *v1.Selector `json:"botNameSelector,omitempty" tf:"-"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group where the Web Chat Channel should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A site represents a client application that you want to connect to your bot. One or more site blocks as defined below.
	// +kubebuilder:validation:Optional
	Site []BotChannelWebChatSiteParameters `json:"site,omitempty" tf:"site,omitempty"`

	// A list of Web Chat Site names.
	// +kubebuilder:validation:Optional
	// +listType=set
	SiteNames []*string `json:"siteNames,omitempty" tf:"site_names,omitempty"`
}

type BotChannelWebChatSiteInitParameters struct {

	// Is the endpoint parameters enabled for this site?
	EndpointParametersEnabled *bool `json:"endpointParametersEnabled,omitempty" tf:"endpoint_parameters_enabled,omitempty"`

	// The name of the site.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Is the storage site enabled for detailed logging? Defaults to true.
	StorageEnabled *bool `json:"storageEnabled,omitempty" tf:"storage_enabled,omitempty"`

	// Is the user upload enabled for this site? Defaults to true.
	UserUploadEnabled *bool `json:"userUploadEnabled,omitempty" tf:"user_upload_enabled,omitempty"`
}

type BotChannelWebChatSiteObservation struct {

	// Is the endpoint parameters enabled for this site?
	EndpointParametersEnabled *bool `json:"endpointParametersEnabled,omitempty" tf:"endpoint_parameters_enabled,omitempty"`

	// The name of the site.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Is the storage site enabled for detailed logging? Defaults to true.
	StorageEnabled *bool `json:"storageEnabled,omitempty" tf:"storage_enabled,omitempty"`

	// Is the user upload enabled for this site? Defaults to true.
	UserUploadEnabled *bool `json:"userUploadEnabled,omitempty" tf:"user_upload_enabled,omitempty"`
}

type BotChannelWebChatSiteParameters struct {

	// Is the endpoint parameters enabled for this site?
	// +kubebuilder:validation:Optional
	EndpointParametersEnabled *bool `json:"endpointParametersEnabled,omitempty" tf:"endpoint_parameters_enabled,omitempty"`

	// The name of the site.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Is the storage site enabled for detailed logging? Defaults to true.
	// +kubebuilder:validation:Optional
	StorageEnabled *bool `json:"storageEnabled,omitempty" tf:"storage_enabled,omitempty"`

	// Is the user upload enabled for this site? Defaults to true.
	// +kubebuilder:validation:Optional
	UserUploadEnabled *bool `json:"userUploadEnabled,omitempty" tf:"user_upload_enabled,omitempty"`
}

// BotChannelWebChatSpec defines the desired state of BotChannelWebChat
type BotChannelWebChatSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BotChannelWebChatParameters `json:"forProvider"`
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
	InitProvider BotChannelWebChatInitParameters `json:"initProvider,omitempty"`
}

// BotChannelWebChatStatus defines the observed state of BotChannelWebChat.
type BotChannelWebChatStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BotChannelWebChatObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BotChannelWebChat is the Schema for the BotChannelWebChats API. Manages a Web Chat integration for a Bot Channel
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BotChannelWebChat struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   BotChannelWebChatSpec   `json:"spec"`
	Status BotChannelWebChatStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BotChannelWebChatList contains a list of BotChannelWebChats
type BotChannelWebChatList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BotChannelWebChat `json:"items"`
}

// Repository type metadata.
var (
	BotChannelWebChat_Kind             = "BotChannelWebChat"
	BotChannelWebChat_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BotChannelWebChat_Kind}.String()
	BotChannelWebChat_KindAPIVersion   = BotChannelWebChat_Kind + "." + CRDGroupVersion.String()
	BotChannelWebChat_GroupVersionKind = CRDGroupVersion.WithKind(BotChannelWebChat_Kind)
)

func init() {
	SchemeBuilder.Register(&BotChannelWebChat{}, &BotChannelWebChatList{})
}
