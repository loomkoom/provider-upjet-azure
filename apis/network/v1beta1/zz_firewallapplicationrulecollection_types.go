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

type FirewallApplicationRuleCollectionInitParameters struct {

	// Specifies the action the rule will apply to matching traffic. Possible values are Allow and Deny.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Specifies the priority of the rule collection. Possible values are between 100 - 65000.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// One or more rule blocks as defined below.
	Rule []RuleInitParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type FirewallApplicationRuleCollectionObservation struct {

	// Specifies the action the rule will apply to matching traffic. Possible values are Allow and Deny.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Specifies the name of the Firewall in which the Application Rule Collection should be created. Changing this forces a new resource to be created.
	AzureFirewallName *string `json:"azureFirewallName,omitempty" tf:"azure_firewall_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the priority of the rule collection. Possible values are between 100 - 65000.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// One or more rule blocks as defined below.
	Rule []RuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`
}

type FirewallApplicationRuleCollectionParameters struct {

	// Specifies the action the rule will apply to matching traffic. Possible values are Allow and Deny.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Specifies the name of the Firewall in which the Application Rule Collection should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Firewall
	// +kubebuilder:validation:Optional
	AzureFirewallName *string `json:"azureFirewallName,omitempty" tf:"azure_firewall_name,omitempty"`

	// Reference to a Firewall to populate azureFirewallName.
	// +kubebuilder:validation:Optional
	AzureFirewallNameRef *v1.Reference `json:"azureFirewallNameRef,omitempty" tf:"-"`

	// Selector for a Firewall to populate azureFirewallName.
	// +kubebuilder:validation:Optional
	AzureFirewallNameSelector *v1.Selector `json:"azureFirewallNameSelector,omitempty" tf:"-"`

	// Specifies the priority of the rule collection. Possible values are between 100 - 65000.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// One or more rule blocks as defined below.
	// +kubebuilder:validation:Optional
	Rule []RuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type ProtocolInitParameters struct {

	// Specify a port for the connection.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Specifies the type of connection. Possible values are Http, Https and Mssql.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ProtocolObservation struct {

	// Specify a port for the connection.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Specifies the type of connection. Possible values are Http, Https and Mssql.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ProtocolParameters struct {

	// Specify a port for the connection.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port" tf:"port,omitempty"`

	// Specifies the type of connection. Possible values are Http, Https and Mssql.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type RuleInitParameters struct {

	// Specifies a description for the rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A list of FQDN tags. Possible values are AppServiceEnvironment, AzureBackup, AzureKubernetesService, HDInsight, MicrosoftActiveProtectionService, WindowsDiagnostics, WindowsUpdate and WindowsVirtualDesktop.
	FqdnTags []*string `json:"fqdnTags,omitempty" tf:"fqdn_tags,omitempty"`

	// Specifies the name of the rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more protocol blocks as defined below.
	Protocol []ProtocolInitParameters `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// A list of source IP addresses and/or IP ranges.
	SourceAddresses []*string `json:"sourceAddresses,omitempty" tf:"source_addresses,omitempty"`

	// A list of source IP Group IDs for the rule.
	SourceIPGroups []*string `json:"sourceIpGroups,omitempty" tf:"source_ip_groups,omitempty"`

	// A list of FQDNs.
	TargetFqdns []*string `json:"targetFqdns,omitempty" tf:"target_fqdns,omitempty"`
}

type RuleObservation struct {

	// Specifies a description for the rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A list of FQDN tags. Possible values are AppServiceEnvironment, AzureBackup, AzureKubernetesService, HDInsight, MicrosoftActiveProtectionService, WindowsDiagnostics, WindowsUpdate and WindowsVirtualDesktop.
	FqdnTags []*string `json:"fqdnTags,omitempty" tf:"fqdn_tags,omitempty"`

	// Specifies the name of the rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more protocol blocks as defined below.
	Protocol []ProtocolObservation `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// A list of source IP addresses and/or IP ranges.
	SourceAddresses []*string `json:"sourceAddresses,omitempty" tf:"source_addresses,omitempty"`

	// A list of source IP Group IDs for the rule.
	SourceIPGroups []*string `json:"sourceIpGroups,omitempty" tf:"source_ip_groups,omitempty"`

	// A list of FQDNs.
	TargetFqdns []*string `json:"targetFqdns,omitempty" tf:"target_fqdns,omitempty"`
}

type RuleParameters struct {

	// Specifies a description for the rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A list of FQDN tags. Possible values are AppServiceEnvironment, AzureBackup, AzureKubernetesService, HDInsight, MicrosoftActiveProtectionService, WindowsDiagnostics, WindowsUpdate and WindowsVirtualDesktop.
	// +kubebuilder:validation:Optional
	FqdnTags []*string `json:"fqdnTags,omitempty" tf:"fqdn_tags,omitempty"`

	// Specifies the name of the rule.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// One or more protocol blocks as defined below.
	// +kubebuilder:validation:Optional
	Protocol []ProtocolParameters `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// A list of source IP addresses and/or IP ranges.
	// +kubebuilder:validation:Optional
	SourceAddresses []*string `json:"sourceAddresses,omitempty" tf:"source_addresses,omitempty"`

	// A list of source IP Group IDs for the rule.
	// +kubebuilder:validation:Optional
	SourceIPGroups []*string `json:"sourceIpGroups,omitempty" tf:"source_ip_groups,omitempty"`

	// A list of FQDNs.
	// +kubebuilder:validation:Optional
	TargetFqdns []*string `json:"targetFqdns,omitempty" tf:"target_fqdns,omitempty"`
}

// FirewallApplicationRuleCollectionSpec defines the desired state of FirewallApplicationRuleCollection
type FirewallApplicationRuleCollectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallApplicationRuleCollectionParameters `json:"forProvider"`
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
	InitProvider FirewallApplicationRuleCollectionInitParameters `json:"initProvider,omitempty"`
}

// FirewallApplicationRuleCollectionStatus defines the observed state of FirewallApplicationRuleCollection.
type FirewallApplicationRuleCollectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallApplicationRuleCollectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// FirewallApplicationRuleCollection is the Schema for the FirewallApplicationRuleCollections API. Manages an Application Rule Collection within an Azure Firewall.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FirewallApplicationRuleCollection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.action) || (has(self.initProvider) && has(self.initProvider.action))",message="spec.forProvider.action is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.priority) || (has(self.initProvider) && has(self.initProvider.priority))",message="spec.forProvider.priority is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.rule) || (has(self.initProvider) && has(self.initProvider.rule))",message="spec.forProvider.rule is a required parameter"
	Spec   FirewallApplicationRuleCollectionSpec   `json:"spec"`
	Status FirewallApplicationRuleCollectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallApplicationRuleCollectionList contains a list of FirewallApplicationRuleCollections
type FirewallApplicationRuleCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallApplicationRuleCollection `json:"items"`
}

// Repository type metadata.
var (
	FirewallApplicationRuleCollection_Kind             = "FirewallApplicationRuleCollection"
	FirewallApplicationRuleCollection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FirewallApplicationRuleCollection_Kind}.String()
	FirewallApplicationRuleCollection_KindAPIVersion   = FirewallApplicationRuleCollection_Kind + "." + CRDGroupVersion.String()
	FirewallApplicationRuleCollection_GroupVersionKind = CRDGroupVersion.WithKind(FirewallApplicationRuleCollection_Kind)
)

func init() {
	SchemeBuilder.Register(&FirewallApplicationRuleCollection{}, &FirewallApplicationRuleCollectionList{})
}
