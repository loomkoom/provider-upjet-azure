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

type ComputeClusterInitParameters struct {

	// The description of the Machine Learning compute. Changing this forces a new Machine Learning Compute Cluster to be created.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An identity block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
	Identity []IdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Whether local authentication methods is enabled. Defaults to true. Changing this forces a new Machine Learning Compute Cluster to be created.
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// The Azure Region where the Machine Learning Compute Cluster should exist. Changing this forces a new Machine Learning Compute Cluster to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Compute Cluster to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/machinelearningservices/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	MachineLearningWorkspaceID *string `json:"machineLearningWorkspaceId,omitempty" tf:"machine_learning_workspace_id,omitempty"`

	// Reference to a Workspace in machinelearningservices to populate machineLearningWorkspaceId.
	// +kubebuilder:validation:Optional
	MachineLearningWorkspaceIDRef *v1.Reference `json:"machineLearningWorkspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in machinelearningservices to populate machineLearningWorkspaceId.
	// +kubebuilder:validation:Optional
	MachineLearningWorkspaceIDSelector *v1.Selector `json:"machineLearningWorkspaceIdSelector,omitempty" tf:"-"`

	// The name which should be used for this Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Whether the compute cluster will have a public ip. To set this to false a subnet_resource_id needs to be set. Defaults to true. Changing this forces a new Machine Learning Compute Cluster to be created.
	NodePublicIPEnabled *bool `json:"nodePublicIpEnabled,omitempty" tf:"node_public_ip_enabled,omitempty"`

	// Credentials for an administrator user account that will be created on each compute node. A ssh block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
	SSH []SSHInitParameters `json:"ssh,omitempty" tf:"ssh,omitempty"`

	// A boolean value indicating whether enable the public SSH port. Changing this forces a new Machine Learning Compute Cluster to be created.
	SSHPublicAccessEnabled *bool `json:"sshPublicAccessEnabled,omitempty" tf:"ssh_public_access_enabled,omitempty"`

	// A scale_settings block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
	ScaleSettings []ScaleSettingsInitParameters `json:"scaleSettings,omitempty" tf:"scale_settings,omitempty"`

	// The ID of the Subnet that the Compute Cluster should reside in. Changing this forces a new Machine Learning Compute Cluster to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SubnetResourceID *string `json:"subnetResourceId,omitempty" tf:"subnet_resource_id,omitempty"`

	// Reference to a Subnet in network to populate subnetResourceId.
	// +kubebuilder:validation:Optional
	SubnetResourceIDRef *v1.Reference `json:"subnetResourceIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetResourceId.
	// +kubebuilder:validation:Optional
	SubnetResourceIDSelector *v1.Selector `json:"subnetResourceIdSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The priority of the VM. Changing this forces a new Machine Learning Compute Cluster to be created. Accepted values are Dedicated and LowPriority.
	VMPriority *string `json:"vmPriority,omitempty" tf:"vm_priority,omitempty"`

	// The size of the VM. Changing this forces a new Machine Learning Compute Cluster to be created.
	VMSize *string `json:"vmSize,omitempty" tf:"vm_size,omitempty"`
}

type ComputeClusterObservation struct {

	// The description of the Machine Learning compute. Changing this forces a new Machine Learning Compute Cluster to be created.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Machine Learning Compute Cluster.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// Whether local authentication methods is enabled. Defaults to true. Changing this forces a new Machine Learning Compute Cluster to be created.
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// The Azure Region where the Machine Learning Compute Cluster should exist. Changing this forces a new Machine Learning Compute Cluster to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Compute Cluster to be created.
	MachineLearningWorkspaceID *string `json:"machineLearningWorkspaceId,omitempty" tf:"machine_learning_workspace_id,omitempty"`

	// The name which should be used for this Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Whether the compute cluster will have a public ip. To set this to false a subnet_resource_id needs to be set. Defaults to true. Changing this forces a new Machine Learning Compute Cluster to be created.
	NodePublicIPEnabled *bool `json:"nodePublicIpEnabled,omitempty" tf:"node_public_ip_enabled,omitempty"`

	// Credentials for an administrator user account that will be created on each compute node. A ssh block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
	SSH []SSHObservation `json:"ssh,omitempty" tf:"ssh,omitempty"`

	// A boolean value indicating whether enable the public SSH port. Changing this forces a new Machine Learning Compute Cluster to be created.
	SSHPublicAccessEnabled *bool `json:"sshPublicAccessEnabled,omitempty" tf:"ssh_public_access_enabled,omitempty"`

	// A scale_settings block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
	ScaleSettings []ScaleSettingsObservation `json:"scaleSettings,omitempty" tf:"scale_settings,omitempty"`

	// The ID of the Subnet that the Compute Cluster should reside in. Changing this forces a new Machine Learning Compute Cluster to be created.
	SubnetResourceID *string `json:"subnetResourceId,omitempty" tf:"subnet_resource_id,omitempty"`

	// A mapping of tags which should be assigned to the Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The priority of the VM. Changing this forces a new Machine Learning Compute Cluster to be created. Accepted values are Dedicated and LowPriority.
	VMPriority *string `json:"vmPriority,omitempty" tf:"vm_priority,omitempty"`

	// The size of the VM. Changing this forces a new Machine Learning Compute Cluster to be created.
	VMSize *string `json:"vmSize,omitempty" tf:"vm_size,omitempty"`
}

type ComputeClusterParameters struct {

	// The description of the Machine Learning compute. Changing this forces a new Machine Learning Compute Cluster to be created.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An identity block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Whether local authentication methods is enabled. Defaults to true. Changing this forces a new Machine Learning Compute Cluster to be created.
	// +kubebuilder:validation:Optional
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// The Azure Region where the Machine Learning Compute Cluster should exist. Changing this forces a new Machine Learning Compute Cluster to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Compute Cluster to be created.
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

	// The name which should be used for this Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Whether the compute cluster will have a public ip. To set this to false a subnet_resource_id needs to be set. Defaults to true. Changing this forces a new Machine Learning Compute Cluster to be created.
	// +kubebuilder:validation:Optional
	NodePublicIPEnabled *bool `json:"nodePublicIpEnabled,omitempty" tf:"node_public_ip_enabled,omitempty"`

	// Credentials for an administrator user account that will be created on each compute node. A ssh block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
	// +kubebuilder:validation:Optional
	SSH []SSHParameters `json:"ssh,omitempty" tf:"ssh,omitempty"`

	// A boolean value indicating whether enable the public SSH port. Changing this forces a new Machine Learning Compute Cluster to be created.
	// +kubebuilder:validation:Optional
	SSHPublicAccessEnabled *bool `json:"sshPublicAccessEnabled,omitempty" tf:"ssh_public_access_enabled,omitempty"`

	// A scale_settings block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
	// +kubebuilder:validation:Optional
	ScaleSettings []ScaleSettingsParameters `json:"scaleSettings,omitempty" tf:"scale_settings,omitempty"`

	// The ID of the Subnet that the Compute Cluster should reside in. Changing this forces a new Machine Learning Compute Cluster to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetResourceID *string `json:"subnetResourceId,omitempty" tf:"subnet_resource_id,omitempty"`

	// Reference to a Subnet in network to populate subnetResourceId.
	// +kubebuilder:validation:Optional
	SubnetResourceIDRef *v1.Reference `json:"subnetResourceIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetResourceId.
	// +kubebuilder:validation:Optional
	SubnetResourceIDSelector *v1.Selector `json:"subnetResourceIdSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The priority of the VM. Changing this forces a new Machine Learning Compute Cluster to be created. Accepted values are Dedicated and LowPriority.
	// +kubebuilder:validation:Optional
	VMPriority *string `json:"vmPriority,omitempty" tf:"vm_priority,omitempty"`

	// The size of the VM. Changing this forces a new Machine Learning Compute Cluster to be created.
	// +kubebuilder:validation:Optional
	VMSize *string `json:"vmSize,omitempty" tf:"vm_size,omitempty"`
}

type IdentityInitParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Machine Learning Compute Cluster. Changing this forces a new resource to be created.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Machine Learning Compute Cluster. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both). Changing this forces a new resource to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityObservation struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Machine Learning Compute Cluster. Changing this forces a new resource to be created.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Principal ID for the Service Principal associated with the Managed Service Identity of this Machine Learning Compute Cluster.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID for the Service Principal associated with the Managed Service Identity of this Machine Learning Compute Cluster.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Machine Learning Compute Cluster. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both). Changing this forces a new resource to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Machine Learning Compute Cluster. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Machine Learning Compute Cluster. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both). Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type SSHInitParameters struct {

	// Password of the administrator user account. Changing this forces a new Machine Learning Compute Cluster to be created.
	AdminPassword *string `json:"adminPassword,omitempty" tf:"admin_password,omitempty"`

	// Name of the administrator user account which can be used to SSH to nodes. Changing this forces a new Machine Learning Compute Cluster to be created.
	AdminUsername *string `json:"adminUsername,omitempty" tf:"admin_username,omitempty"`

	// SSH public key of the administrator user account. Changing this forces a new Machine Learning Compute Cluster to be created.
	KeyValue *string `json:"keyValue,omitempty" tf:"key_value,omitempty"`
}

type SSHObservation struct {

	// Password of the administrator user account. Changing this forces a new Machine Learning Compute Cluster to be created.
	AdminPassword *string `json:"adminPassword,omitempty" tf:"admin_password,omitempty"`

	// Name of the administrator user account which can be used to SSH to nodes. Changing this forces a new Machine Learning Compute Cluster to be created.
	AdminUsername *string `json:"adminUsername,omitempty" tf:"admin_username,omitempty"`

	// SSH public key of the administrator user account. Changing this forces a new Machine Learning Compute Cluster to be created.
	KeyValue *string `json:"keyValue,omitempty" tf:"key_value,omitempty"`
}

type SSHParameters struct {

	// Password of the administrator user account. Changing this forces a new Machine Learning Compute Cluster to be created.
	// +kubebuilder:validation:Optional
	AdminPassword *string `json:"adminPassword,omitempty" tf:"admin_password,omitempty"`

	// Name of the administrator user account which can be used to SSH to nodes. Changing this forces a new Machine Learning Compute Cluster to be created.
	// +kubebuilder:validation:Optional
	AdminUsername *string `json:"adminUsername" tf:"admin_username,omitempty"`

	// SSH public key of the administrator user account. Changing this forces a new Machine Learning Compute Cluster to be created.
	// +kubebuilder:validation:Optional
	KeyValue *string `json:"keyValue,omitempty" tf:"key_value,omitempty"`
}

type ScaleSettingsInitParameters struct {

	// Maximum node count. Changing this forces a new Machine Learning Compute Cluster to be created.
	MaxNodeCount *float64 `json:"maxNodeCount,omitempty" tf:"max_node_count,omitempty"`

	// Minimal node count. Changing this forces a new Machine Learning Compute Cluster to be created.
	MinNodeCount *float64 `json:"minNodeCount,omitempty" tf:"min_node_count,omitempty"`

	// Node Idle Time Before Scale Down: defines the time until the compute is shutdown when it has gone into Idle state. Is defined according to W3C XML schema standard for duration. Changing this forces a new Machine Learning Compute Cluster to be created.
	ScaleDownNodesAfterIdleDuration *string `json:"scaleDownNodesAfterIdleDuration,omitempty" tf:"scale_down_nodes_after_idle_duration,omitempty"`
}

type ScaleSettingsObservation struct {

	// Maximum node count. Changing this forces a new Machine Learning Compute Cluster to be created.
	MaxNodeCount *float64 `json:"maxNodeCount,omitempty" tf:"max_node_count,omitempty"`

	// Minimal node count. Changing this forces a new Machine Learning Compute Cluster to be created.
	MinNodeCount *float64 `json:"minNodeCount,omitempty" tf:"min_node_count,omitempty"`

	// Node Idle Time Before Scale Down: defines the time until the compute is shutdown when it has gone into Idle state. Is defined according to W3C XML schema standard for duration. Changing this forces a new Machine Learning Compute Cluster to be created.
	ScaleDownNodesAfterIdleDuration *string `json:"scaleDownNodesAfterIdleDuration,omitempty" tf:"scale_down_nodes_after_idle_duration,omitempty"`
}

type ScaleSettingsParameters struct {

	// Maximum node count. Changing this forces a new Machine Learning Compute Cluster to be created.
	// +kubebuilder:validation:Optional
	MaxNodeCount *float64 `json:"maxNodeCount" tf:"max_node_count,omitempty"`

	// Minimal node count. Changing this forces a new Machine Learning Compute Cluster to be created.
	// +kubebuilder:validation:Optional
	MinNodeCount *float64 `json:"minNodeCount" tf:"min_node_count,omitempty"`

	// Node Idle Time Before Scale Down: defines the time until the compute is shutdown when it has gone into Idle state. Is defined according to W3C XML schema standard for duration. Changing this forces a new Machine Learning Compute Cluster to be created.
	// +kubebuilder:validation:Optional
	ScaleDownNodesAfterIdleDuration *string `json:"scaleDownNodesAfterIdleDuration" tf:"scale_down_nodes_after_idle_duration,omitempty"`
}

// ComputeClusterSpec defines the desired state of ComputeCluster
type ComputeClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ComputeClusterParameters `json:"forProvider"`
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
	InitProvider ComputeClusterInitParameters `json:"initProvider,omitempty"`
}

// ComputeClusterStatus defines the observed state of ComputeCluster.
type ComputeClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ComputeClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ComputeCluster is the Schema for the ComputeClusters API. Manages a Machine Learning Compute Cluster.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ComputeCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scaleSettings) || (has(self.initProvider) && has(self.initProvider.scaleSettings))",message="spec.forProvider.scaleSettings is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vmPriority) || (has(self.initProvider) && has(self.initProvider.vmPriority))",message="spec.forProvider.vmPriority is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vmSize) || (has(self.initProvider) && has(self.initProvider.vmSize))",message="spec.forProvider.vmSize is a required parameter"
	Spec   ComputeClusterSpec   `json:"spec"`
	Status ComputeClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ComputeClusterList contains a list of ComputeClusters
type ComputeClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeCluster `json:"items"`
}

// Repository type metadata.
var (
	ComputeCluster_Kind             = "ComputeCluster"
	ComputeCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ComputeCluster_Kind}.String()
	ComputeCluster_KindAPIVersion   = ComputeCluster_Kind + "." + CRDGroupVersion.String()
	ComputeCluster_GroupVersionKind = CRDGroupVersion.WithKind(ComputeCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&ComputeCluster{}, &ComputeClusterList{})
}
