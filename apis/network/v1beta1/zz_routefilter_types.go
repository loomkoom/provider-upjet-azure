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

type RouteFilterInitParameters struct {

	// The Azure Region where the Route Filter should exist. Changing this forces a new Route Filter to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A rule block as defined below.
	Rule []RouteFilterRuleInitParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// A mapping of tags which should be assigned to the Route Filter.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RouteFilterObservation struct {

	// The ID of the Route Filter.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Azure Region where the Route Filter should exist. Changing this forces a new Route Filter to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Resource Group where the Route Filter should exist. Changing this forces a new Route Filter to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A rule block as defined below.
	Rule []RouteFilterRuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`

	// A mapping of tags which should be assigned to the Route Filter.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RouteFilterParameters struct {

	// The Azure Region where the Route Filter should exist. Changing this forces a new Route Filter to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Resource Group where the Route Filter should exist. Changing this forces a new Route Filter to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A rule block as defined below.
	// +kubebuilder:validation:Optional
	Rule []RouteFilterRuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// A mapping of tags which should be assigned to the Route Filter.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RouteFilterRuleInitParameters struct {

	// The access type of the rule. The only possible value is Allow.
	Access *string `json:"access,omitempty" tf:"access"`

	// The collection for bgp community values to filter on. e.g. ['12076:5010','12076:5020'].
	Communities []*string `json:"communities,omitempty" tf:"communities"`

	// The name of the route filter rule.
	Name *string `json:"name,omitempty" tf:"name"`

	// The rule type of the rule. The only possible value is Community.
	RuleType *string `json:"ruleType,omitempty" tf:"rule_type"`
}

type RouteFilterRuleObservation struct {

	// The access type of the rule. The only possible value is Allow.
	Access *string `json:"access,omitempty" tf:"access,omitempty"`

	// The collection for bgp community values to filter on. e.g. ['12076:5010','12076:5020'].
	Communities []*string `json:"communities,omitempty" tf:"communities,omitempty"`

	// The name of the route filter rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The rule type of the rule. The only possible value is Community.
	RuleType *string `json:"ruleType,omitempty" tf:"rule_type,omitempty"`
}

type RouteFilterRuleParameters struct {

	// The access type of the rule. The only possible value is Allow.
	// +kubebuilder:validation:Optional
	Access *string `json:"access,omitempty" tf:"access"`

	// The collection for bgp community values to filter on. e.g. ['12076:5010','12076:5020'].
	// +kubebuilder:validation:Optional
	Communities []*string `json:"communities,omitempty" tf:"communities"`

	// The name of the route filter rule.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`

	// The rule type of the rule. The only possible value is Community.
	// +kubebuilder:validation:Optional
	RuleType *string `json:"ruleType,omitempty" tf:"rule_type"`
}

// RouteFilterSpec defines the desired state of RouteFilter
type RouteFilterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteFilterParameters `json:"forProvider"`
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
	InitProvider RouteFilterInitParameters `json:"initProvider,omitempty"`
}

// RouteFilterStatus defines the observed state of RouteFilter.
type RouteFilterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteFilterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RouteFilter is the Schema for the RouteFilters API. Manages a Route Filter.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type RouteFilter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	Spec   RouteFilterSpec   `json:"spec"`
	Status RouteFilterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteFilterList contains a list of RouteFilters
type RouteFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteFilter `json:"items"`
}

// Repository type metadata.
var (
	RouteFilter_Kind             = "RouteFilter"
	RouteFilter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouteFilter_Kind}.String()
	RouteFilter_KindAPIVersion   = RouteFilter_Kind + "." + CRDGroupVersion.String()
	RouteFilter_GroupVersionKind = CRDGroupVersion.WithKind(RouteFilter_Kind)
)

func init() {
	SchemeBuilder.Register(&RouteFilter{}, &RouteFilterList{})
}
