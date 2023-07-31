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

type BotChannelSlackInitParameters struct {

	// The Client ID that will be used to authenticate with Slack.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The Slack Landing Page URL.
	LandingPageURL *string `json:"landingPageUrl,omitempty" tf:"landing_page_url,omitempty"`

	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`
}

type BotChannelSlackObservation struct {

	// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
	BotName *string `json:"botName,omitempty" tf:"bot_name,omitempty"`

	// The Client ID that will be used to authenticate with Slack.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The ID of the Slack Integration for a Bot Channel.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Slack Landing Page URL.
	LandingPageURL *string `json:"landingPageUrl,omitempty" tf:"landing_page_url,omitempty"`

	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
}

type BotChannelSlackParameters struct {

	// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/botservice/v1beta1.BotChannelsRegistration
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	BotName *string `json:"botName,omitempty" tf:"bot_name,omitempty"`

	// Reference to a BotChannelsRegistration in botservice to populate botName.
	// +kubebuilder:validation:Optional
	BotNameRef *v1.Reference `json:"botNameRef,omitempty" tf:"-"`

	// Selector for a BotChannelsRegistration in botservice to populate botName.
	// +kubebuilder:validation:Optional
	BotNameSelector *v1.Selector `json:"botNameSelector,omitempty" tf:"-"`

	// The Client ID that will be used to authenticate with Slack.
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The Client Secret that will be used to authenticate with Slack.
	// +kubebuilder:validation:Optional
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// The Slack Landing Page URL.
	// +kubebuilder:validation:Optional
	LandingPageURL *string `json:"landingPageUrl,omitempty" tf:"landing_page_url,omitempty"`

	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The Signing Secret that will be used to sign the requests.
	// +kubebuilder:validation:Optional
	SigningSecretSecretRef *v1.SecretKeySelector `json:"signingSecretSecretRef,omitempty" tf:"-"`

	// The Verification Token that will be used to authenticate with Slack.
	// +kubebuilder:validation:Optional
	VerificationTokenSecretRef v1.SecretKeySelector `json:"verificationTokenSecretRef" tf:"-"`
}

// BotChannelSlackSpec defines the desired state of BotChannelSlack
type BotChannelSlackSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BotChannelSlackParameters `json:"forProvider"`
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
	InitProvider BotChannelSlackInitParameters `json:"initProvider,omitempty"`
}

// BotChannelSlackStatus defines the observed state of BotChannelSlack.
type BotChannelSlackStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BotChannelSlackObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BotChannelSlack is the Schema for the BotChannelSlacks API. Manages a Slack integration for a Bot Channel
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BotChannelSlack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientId) || has(self.initProvider.clientId)",message="clientId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientSecretSecretRef)",message="clientSecretSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.verificationTokenSecretRef)",message="verificationTokenSecretRef is a required parameter"
	Spec   BotChannelSlackSpec   `json:"spec"`
	Status BotChannelSlackStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BotChannelSlackList contains a list of BotChannelSlacks
type BotChannelSlackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BotChannelSlack `json:"items"`
}

// Repository type metadata.
var (
	BotChannelSlack_Kind             = "BotChannelSlack"
	BotChannelSlack_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BotChannelSlack_Kind}.String()
	BotChannelSlack_KindAPIVersion   = BotChannelSlack_Kind + "." + CRDGroupVersion.String()
	BotChannelSlack_GroupVersionKind = CRDGroupVersion.WithKind(BotChannelSlack_Kind)
)

func init() {
	SchemeBuilder.Register(&BotChannelSlack{}, &BotChannelSlackList{})
}
