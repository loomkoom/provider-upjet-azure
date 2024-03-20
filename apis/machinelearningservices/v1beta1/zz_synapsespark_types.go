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

type SynapseSparkIdentityInitParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Machine Learning Synapse Spark. Changing this forces a new resource to be created.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Machine Learning Synapse Spark. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both). Changing this forces a new resource to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SynapseSparkIdentityObservation struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Machine Learning Synapse Spark. Changing this forces a new resource to be created.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Principal ID for the Service Principal associated with the Managed Service Identity of this Machine Learning Synapse Spark.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID for the Service Principal associated with the Managed Service Identity of this Machine Learning Synapse Spark.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Machine Learning Synapse Spark. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both). Changing this forces a new resource to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SynapseSparkIdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Machine Learning Synapse Spark. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Machine Learning Synapse Spark. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both). Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type SynapseSparkInitParameters struct {

	// The description of the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An identity block as defined below. Changing this forces a new Machine Learning Synapse Spark to be created.
	Identity []SynapseSparkIdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Whether local authentication methods is enabled. Defaults to true. Changing this forces a new Machine Learning Synapse Spark to be created.
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// The Azure Region where the Machine Learning Synapse Spark should exist. Changing this forces a new Machine Learning Synapse Spark to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the linked Synapse Spark Pool. Changing this forces a new Machine Learning Synapse Spark to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/synapse/v1beta1.SparkPool
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SynapseSparkPoolID *string `json:"synapseSparkPoolId,omitempty" tf:"synapse_spark_pool_id,omitempty"`

	// Reference to a SparkPool in synapse to populate synapseSparkPoolId.
	// +kubebuilder:validation:Optional
	SynapseSparkPoolIDRef *v1.Reference `json:"synapseSparkPoolIdRef,omitempty" tf:"-"`

	// Selector for a SparkPool in synapse to populate synapseSparkPoolId.
	// +kubebuilder:validation:Optional
	SynapseSparkPoolIDSelector *v1.Selector `json:"synapseSparkPoolIdSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SynapseSparkObservation struct {

	// The description of the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Machine Learning Synapse Spark.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below. Changing this forces a new Machine Learning Synapse Spark to be created.
	Identity []SynapseSparkIdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// Whether local authentication methods is enabled. Defaults to true. Changing this forces a new Machine Learning Synapse Spark to be created.
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// The Azure Region where the Machine Learning Synapse Spark should exist. Changing this forces a new Machine Learning Synapse Spark to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Synapse Spark to be created.
	MachineLearningWorkspaceID *string `json:"machineLearningWorkspaceId,omitempty" tf:"machine_learning_workspace_id,omitempty"`

	// The ID of the linked Synapse Spark Pool. Changing this forces a new Machine Learning Synapse Spark to be created.
	SynapseSparkPoolID *string `json:"synapseSparkPoolId,omitempty" tf:"synapse_spark_pool_id,omitempty"`

	// A mapping of tags which should be assigned to the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SynapseSparkParameters struct {

	// The description of the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An identity block as defined below. Changing this forces a new Machine Learning Synapse Spark to be created.
	// +kubebuilder:validation:Optional
	Identity []SynapseSparkIdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Whether local authentication methods is enabled. Defaults to true. Changing this forces a new Machine Learning Synapse Spark to be created.
	// +kubebuilder:validation:Optional
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// The Azure Region where the Machine Learning Synapse Spark should exist. Changing this forces a new Machine Learning Synapse Spark to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Synapse Spark to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/machinelearningservices/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	MachineLearningWorkspaceID *string `json:"machineLearningWorkspaceId,omitempty" tf:"machine_learning_workspace_id,omitempty"`

	// Reference to a Workspace in machinelearningservices to populate machineLearningWorkspaceId.
	// +kubebuilder:validation:Optional
	MachineLearningWorkspaceIDRef *v1.Reference `json:"machineLearningWorkspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in machinelearningservices to populate machineLearningWorkspaceId.
	// +kubebuilder:validation:Optional
	MachineLearningWorkspaceIDSelector *v1.Selector `json:"machineLearningWorkspaceIdSelector,omitempty" tf:"-"`

	// The ID of the linked Synapse Spark Pool. Changing this forces a new Machine Learning Synapse Spark to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/synapse/v1beta1.SparkPool
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SynapseSparkPoolID *string `json:"synapseSparkPoolId,omitempty" tf:"synapse_spark_pool_id,omitempty"`

	// Reference to a SparkPool in synapse to populate synapseSparkPoolId.
	// +kubebuilder:validation:Optional
	SynapseSparkPoolIDRef *v1.Reference `json:"synapseSparkPoolIdRef,omitempty" tf:"-"`

	// Selector for a SparkPool in synapse to populate synapseSparkPoolId.
	// +kubebuilder:validation:Optional
	SynapseSparkPoolIDSelector *v1.Selector `json:"synapseSparkPoolIdSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SynapseSparkSpec defines the desired state of SynapseSpark
type SynapseSparkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SynapseSparkParameters `json:"forProvider"`
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
	InitProvider SynapseSparkInitParameters `json:"initProvider,omitempty"`
}

// SynapseSparkStatus defines the observed state of SynapseSpark.
type SynapseSparkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SynapseSparkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SynapseSpark is the Schema for the SynapseSparks API. Manages the linked service to link an Azure Machine learning workspace to an Azure Synapse workspace.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SynapseSpark struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   SynapseSparkSpec   `json:"spec"`
	Status SynapseSparkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SynapseSparkList contains a list of SynapseSparks
type SynapseSparkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SynapseSpark `json:"items"`
}

// Repository type metadata.
var (
	SynapseSpark_Kind             = "SynapseSpark"
	SynapseSpark_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SynapseSpark_Kind}.String()
	SynapseSpark_KindAPIVersion   = SynapseSpark_Kind + "." + CRDGroupVersion.String()
	SynapseSpark_GroupVersionKind = CRDGroupVersion.WithKind(SynapseSpark_Kind)
)

func init() {
	SchemeBuilder.Register(&SynapseSpark{}, &SynapseSparkList{})
}
