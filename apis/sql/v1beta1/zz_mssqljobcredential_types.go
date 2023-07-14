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

type MSSQLJobCredentialObservation struct {

	// The ID of the Elastic Job Credential.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Elastic Job Agent. Changing this forces a new Elastic Job Credential to be created.
	JobAgentID *string `json:"jobAgentId,omitempty" tf:"job_agent_id,omitempty"`

	// The username part of the credential.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type MSSQLJobCredentialParameters struct {

	// The ID of the Elastic Job Agent. Changing this forces a new Elastic Job Credential to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/sql/v1beta1.MSSQLJobAgent
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	JobAgentID *string `json:"jobAgentId,omitempty" tf:"job_agent_id,omitempty"`

	// Reference to a MSSQLJobAgent in sql to populate jobAgentId.
	// +kubebuilder:validation:Optional
	JobAgentIDRef *v1.Reference `json:"jobAgentIdRef,omitempty" tf:"-"`

	// Selector for a MSSQLJobAgent in sql to populate jobAgentId.
	// +kubebuilder:validation:Optional
	JobAgentIDSelector *v1.Selector `json:"jobAgentIdSelector,omitempty" tf:"-"`

	// The password part of the credential.
	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The username part of the credential.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// MSSQLJobCredentialSpec defines the desired state of MSSQLJobCredential
type MSSQLJobCredentialSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MSSQLJobCredentialParameters `json:"forProvider"`
}

// MSSQLJobCredentialStatus defines the observed state of MSSQLJobCredential.
type MSSQLJobCredentialStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MSSQLJobCredentialObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLJobCredential is the Schema for the MSSQLJobCredentials API. Manages an Elastic Job Credential.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MSSQLJobCredential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.passwordSecretRef)",message="passwordSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.username)",message="username is a required parameter"
	Spec   MSSQLJobCredentialSpec   `json:"spec"`
	Status MSSQLJobCredentialStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLJobCredentialList contains a list of MSSQLJobCredentials
type MSSQLJobCredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MSSQLJobCredential `json:"items"`
}

// Repository type metadata.
var (
	MSSQLJobCredential_Kind             = "MSSQLJobCredential"
	MSSQLJobCredential_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MSSQLJobCredential_Kind}.String()
	MSSQLJobCredential_KindAPIVersion   = MSSQLJobCredential_Kind + "." + CRDGroupVersion.String()
	MSSQLJobCredential_GroupVersionKind = CRDGroupVersion.WithKind(MSSQLJobCredential_Kind)
)

func init() {
	SchemeBuilder.Register(&MSSQLJobCredential{}, &MSSQLJobCredentialList{})
}
