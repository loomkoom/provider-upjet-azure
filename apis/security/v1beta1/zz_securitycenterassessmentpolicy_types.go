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

type SecurityCenterAssessmentPolicyObservation struct {

	// A list of the categories of resource that is at risk when the Security Center Assessment is unhealthy. Possible values are Unknown, Compute, Data, IdentityAndAccess, IoT and Networking.
	Categories []*string `json:"categories,omitempty" tf:"categories,omitempty"`

	// The description of the Security Center Assessment.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The user-friendly display name of the Security Center Assessment.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The ID of the Security Center Assessment Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The implementation effort which is used to remediate the Security Center Assessment. Possible values are Low, Moderate and High.
	ImplementationEffort *string `json:"implementationEffort,omitempty" tf:"implementation_effort,omitempty"`

	// The GUID as the name of the Security Center Assessment Policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The description which is used to mitigate the security issue.
	RemediationDescription *string `json:"remediationDescription,omitempty" tf:"remediation_description,omitempty"`

	// The severity level of the Security Center Assessment. Possible values are Low, Medium and High. Defaults to Medium.
	Severity *string `json:"severity,omitempty" tf:"severity,omitempty"`

	// A list of the threat impacts for the Security Center Assessment. Possible values are AccountBreach, DataExfiltration, DataSpillage, DenialOfService, ElevationOfPrivilege, MaliciousInsider, MissingCoverage and ThreatResistance.
	Threats []*string `json:"threats,omitempty" tf:"threats,omitempty"`

	// The user impact of the Security Center Assessment. Possible values are Low, Moderate and High.
	UserImpact *string `json:"userImpact,omitempty" tf:"user_impact,omitempty"`
}

type SecurityCenterAssessmentPolicyParameters struct {

	// A list of the categories of resource that is at risk when the Security Center Assessment is unhealthy. Possible values are Unknown, Compute, Data, IdentityAndAccess, IoT and Networking.
	// +kubebuilder:validation:Optional
	Categories []*string `json:"categories,omitempty" tf:"categories,omitempty"`

	// The description of the Security Center Assessment.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The user-friendly display name of the Security Center Assessment.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The implementation effort which is used to remediate the Security Center Assessment. Possible values are Low, Moderate and High.
	// +kubebuilder:validation:Optional
	ImplementationEffort *string `json:"implementationEffort,omitempty" tf:"implementation_effort,omitempty"`

	// The description which is used to mitigate the security issue.
	// +kubebuilder:validation:Optional
	RemediationDescription *string `json:"remediationDescription,omitempty" tf:"remediation_description,omitempty"`

	// The severity level of the Security Center Assessment. Possible values are Low, Medium and High. Defaults to Medium.
	// +kubebuilder:validation:Optional
	Severity *string `json:"severity,omitempty" tf:"severity,omitempty"`

	// A list of the threat impacts for the Security Center Assessment. Possible values are AccountBreach, DataExfiltration, DataSpillage, DenialOfService, ElevationOfPrivilege, MaliciousInsider, MissingCoverage and ThreatResistance.
	// +kubebuilder:validation:Optional
	Threats []*string `json:"threats,omitempty" tf:"threats,omitempty"`

	// The user impact of the Security Center Assessment. Possible values are Low, Moderate and High.
	// +kubebuilder:validation:Optional
	UserImpact *string `json:"userImpact,omitempty" tf:"user_impact,omitempty"`
}

// SecurityCenterAssessmentPolicySpec defines the desired state of SecurityCenterAssessmentPolicy
type SecurityCenterAssessmentPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityCenterAssessmentPolicyParameters `json:"forProvider"`
}

// SecurityCenterAssessmentPolicyStatus defines the observed state of SecurityCenterAssessmentPolicy.
type SecurityCenterAssessmentPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityCenterAssessmentPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityCenterAssessmentPolicy is the Schema for the SecurityCenterAssessmentPolicys API. Manages the Security Center Assessment Metadata for Azure Security Center.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SecurityCenterAssessmentPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.description)",message="description is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName)",message="displayName is a required parameter"
	Spec   SecurityCenterAssessmentPolicySpec   `json:"spec"`
	Status SecurityCenterAssessmentPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityCenterAssessmentPolicyList contains a list of SecurityCenterAssessmentPolicys
type SecurityCenterAssessmentPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityCenterAssessmentPolicy `json:"items"`
}

// Repository type metadata.
var (
	SecurityCenterAssessmentPolicy_Kind             = "SecurityCenterAssessmentPolicy"
	SecurityCenterAssessmentPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityCenterAssessmentPolicy_Kind}.String()
	SecurityCenterAssessmentPolicy_KindAPIVersion   = SecurityCenterAssessmentPolicy_Kind + "." + CRDGroupVersion.String()
	SecurityCenterAssessmentPolicy_GroupVersionKind = CRDGroupVersion.WithKind(SecurityCenterAssessmentPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityCenterAssessmentPolicy{}, &SecurityCenterAssessmentPolicyList{})
}
