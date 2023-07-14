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

type IdentityObservation struct {

	// A list of User Assigned Managed Identity IDs to be assigned to this Redis Cluster.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Route ID.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Route ID.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Redis Cluster. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// A list of User Assigned Managed Identity IDs to be assigned to this Redis Cluster.
	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Redis Cluster. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type PatchScheduleObservation struct {

	// the Weekday name - possible values include Monday, Tuesday, Wednesday etc.
	DayOfWeek *string `json:"dayOfWeek,omitempty" tf:"day_of_week,omitempty"`

	// The ISO 8601 timespan which specifies the amount of time the Redis Cache can be updated. Defaults to PT5H.
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// the Start Hour for maintenance in UTC - possible values range from 0 - 23.
	StartHourUtc *float64 `json:"startHourUtc,omitempty" tf:"start_hour_utc,omitempty"`
}

type PatchScheduleParameters struct {

	// the Weekday name - possible values include Monday, Tuesday, Wednesday etc.
	// +kubebuilder:validation:Required
	DayOfWeek *string `json:"dayOfWeek" tf:"day_of_week,omitempty"`

	// The ISO 8601 timespan which specifies the amount of time the Redis Cache can be updated. Defaults to PT5H.
	// +kubebuilder:validation:Optional
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// the Start Hour for maintenance in UTC - possible values range from 0 - 23.
	// +kubebuilder:validation:Optional
	StartHourUtc *float64 `json:"startHourUtc,omitempty" tf:"start_hour_utc,omitempty"`
}

type RedisCacheObservation struct {

	// The size of the Redis cache to deploy. Valid values for a SKU family of C (Basic/Standard) are 0, 1, 2, 3, 4, 5, 6, and for P (Premium) family are 1, 2, 3, 4, 5.
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// Enable the non-SSL port (6379) - disabled by default.
	EnableNonSSLPort *bool `json:"enableNonSslPort,omitempty" tf:"enable_non_ssl_port,omitempty"`

	// The SKU family/pricing group to use. Valid values are C (for Basic/Standard SKU family) and P (for Premium)
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// The Hostname of the Redis Instance
	HostName *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// The Route ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// The location of the resource group. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The minimum TLS version. Possible values are 1.0, 1.1 and 1.2. Defaults to 1.0.
	MinimumTLSVersion *string `json:"minimumTlsVersion,omitempty" tf:"minimum_tls_version,omitempty"`

	// A list of patch_schedule blocks as defined below.
	PatchSchedule []PatchScheduleObservation `json:"patchSchedule,omitempty" tf:"patch_schedule,omitempty"`

	// The non-SSL Port of the Redis Instance
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The Static IP Address to assign to the Redis Cache when hosted inside the Virtual Network. This argument implies the use of subnet_id. Changing this forces a new resource to be created.
	PrivateStaticIPAddress *string `json:"privateStaticIpAddress,omitempty" tf:"private_static_ip_address,omitempty"`

	// Whether or not public network access is allowed for this Redis Cache. true means this resource could be accessed by both public and private endpoint. false means only private endpoint access is allowed. Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// A redis_configuration as defined below - with some limitations by SKU - defaults/details are shown below.
	RedisConfiguration []RedisConfigurationObservation `json:"redisConfiguration,omitempty" tf:"redis_configuration,omitempty"`

	// Redis version. Only major version needed. Valid values: 4, 6.
	RedisVersion *string `json:"redisVersion,omitempty" tf:"redis_version,omitempty"`

	// Amount of replicas to create per master for this Redis Cache.
	ReplicasPerMaster *float64 `json:"replicasPerMaster,omitempty" tf:"replicas_per_master,omitempty"`

	// Amount of replicas to create per primary for this Redis Cache. If both replicas_per_primary and replicas_per_master are set, they need to be equal.
	ReplicasPerPrimary *float64 `json:"replicasPerPrimary,omitempty" tf:"replicas_per_primary,omitempty"`

	// The name of the resource group in which to create the Redis instance. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The SSL Port of the Redis Instance
	SSLPort *float64 `json:"sslPort,omitempty" tf:"ssl_port,omitempty"`

	// Only available when using the Premium SKU The number of Shards to create on the Redis Cluster.
	ShardCount *float64 `json:"shardCount,omitempty" tf:"shard_count,omitempty"`

	// The SKU of Redis to use. Possible values are Basic, Standard and Premium.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// Only available when using the Premium SKU The ID of the Subnet within which the Redis Cache should be deployed. This Subnet must only contain Azure Cache for Redis instances without any other type of resources. Changing this forces a new resource to be created.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A mapping of tenant settings to assign to the resource.
	TenantSettings map[string]*string `json:"tenantSettings,omitempty" tf:"tenant_settings,omitempty"`

	// Specifies a list of Availability Zones in which this Redis Cache should be located. Changing this forces a new Redis Cache to be created.
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type RedisCacheParameters struct {

	// The size of the Redis cache to deploy. Valid values for a SKU family of C (Basic/Standard) are 0, 1, 2, 3, 4, 5, 6, and for P (Premium) family are 1, 2, 3, 4, 5.
	// +kubebuilder:validation:Optional
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// Enable the non-SSL port (6379) - disabled by default.
	// +kubebuilder:validation:Optional
	EnableNonSSLPort *bool `json:"enableNonSslPort,omitempty" tf:"enable_non_ssl_port,omitempty"`

	// The SKU family/pricing group to use. Valid values are C (for Basic/Standard SKU family) and P (for Premium)
	// +kubebuilder:validation:Optional
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// The location of the resource group. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The minimum TLS version. Possible values are 1.0, 1.1 and 1.2. Defaults to 1.0.
	// +kubebuilder:validation:Optional
	MinimumTLSVersion *string `json:"minimumTlsVersion,omitempty" tf:"minimum_tls_version,omitempty"`

	// A list of patch_schedule blocks as defined below.
	// +kubebuilder:validation:Optional
	PatchSchedule []PatchScheduleParameters `json:"patchSchedule,omitempty" tf:"patch_schedule,omitempty"`

	// The Static IP Address to assign to the Redis Cache when hosted inside the Virtual Network. This argument implies the use of subnet_id. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PrivateStaticIPAddress *string `json:"privateStaticIpAddress,omitempty" tf:"private_static_ip_address,omitempty"`

	// Whether or not public network access is allowed for this Redis Cache. true means this resource could be accessed by both public and private endpoint. false means only private endpoint access is allowed. Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// A redis_configuration as defined below - with some limitations by SKU - defaults/details are shown below.
	// +kubebuilder:validation:Optional
	RedisConfiguration []RedisConfigurationParameters `json:"redisConfiguration,omitempty" tf:"redis_configuration,omitempty"`

	// Redis version. Only major version needed. Valid values: 4, 6.
	// +kubebuilder:validation:Optional
	RedisVersion *string `json:"redisVersion,omitempty" tf:"redis_version,omitempty"`

	// Amount of replicas to create per master for this Redis Cache.
	// +kubebuilder:validation:Optional
	ReplicasPerMaster *float64 `json:"replicasPerMaster,omitempty" tf:"replicas_per_master,omitempty"`

	// Amount of replicas to create per primary for this Redis Cache. If both replicas_per_primary and replicas_per_master are set, they need to be equal.
	// +kubebuilder:validation:Optional
	ReplicasPerPrimary *float64 `json:"replicasPerPrimary,omitempty" tf:"replicas_per_primary,omitempty"`

	// The name of the resource group in which to create the Redis instance. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Only available when using the Premium SKU The number of Shards to create on the Redis Cluster.
	// +kubebuilder:validation:Optional
	ShardCount *float64 `json:"shardCount,omitempty" tf:"shard_count,omitempty"`

	// The SKU of Redis to use. Possible values are Basic, Standard and Premium.
	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// Only available when using the Premium SKU The ID of the Subnet within which the Redis Cache should be deployed. This Subnet must only contain Azure Cache for Redis instances without any other type of resources. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A mapping of tenant settings to assign to the resource.
	// +kubebuilder:validation:Optional
	TenantSettings map[string]*string `json:"tenantSettings,omitempty" tf:"tenant_settings,omitempty"`

	// Specifies a list of Availability Zones in which this Redis Cache should be located. Changing this forces a new Redis Cache to be created.
	// +kubebuilder:validation:Optional
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type RedisConfigurationObservation struct {

	// Enable or disable AOF persistence for this Redis Cache. Defaults to false.
	AofBackupEnabled *bool `json:"aofBackupEnabled,omitempty" tf:"aof_backup_enabled,omitempty"`

	// If set to false, the Redis instance will be accessible without authentication. Defaults to true.
	EnableAuthentication *bool `json:"enableAuthentication,omitempty" tf:"enable_authentication,omitempty"`

	// Returns the max number of connected clients at the same time.
	Maxclients *float64 `json:"maxclients,omitempty" tf:"maxclients,omitempty"`

	// Value in megabytes reserved to accommodate for memory fragmentation. Defaults are shown below.
	MaxfragmentationmemoryReserved *float64 `json:"maxfragmentationmemoryReserved,omitempty" tf:"maxfragmentationmemory_reserved,omitempty"`

	// The max-memory delta for this Redis instance. Defaults are shown below.
	MaxmemoryDelta *float64 `json:"maxmemoryDelta,omitempty" tf:"maxmemory_delta,omitempty"`

	// How Redis will select what to remove when maxmemory is reached. Defaults are shown below. Defaults to volatile-lru.
	MaxmemoryPolicy *string `json:"maxmemoryPolicy,omitempty" tf:"maxmemory_policy,omitempty"`

	// Value in megabytes reserved for non-cache usage e.g. failover. Defaults are shown below.
	MaxmemoryReserved *float64 `json:"maxmemoryReserved,omitempty" tf:"maxmemory_reserved,omitempty"`

	// Keyspace notifications allows clients to subscribe to Pub/Sub channels in order to receive events affecting the Redis data set in some way. Reference
	NotifyKeySpaceEvents *string `json:"notifyKeyspaceEvents,omitempty" tf:"notify_keyspace_events,omitempty"`

	// Is Backup Enabled? Only supported on Premium SKUs. Defaults to false.
	RdbBackupEnabled *bool `json:"rdbBackupEnabled,omitempty" tf:"rdb_backup_enabled,omitempty"`

	// The Backup Frequency in Minutes. Only supported on Premium SKUs. Possible values are: 15, 30, 60, 360, 720 and 1440.
	RdbBackupFrequency *float64 `json:"rdbBackupFrequency,omitempty" tf:"rdb_backup_frequency,omitempty"`

	// The maximum number of snapshots to create as a backup. Only supported for Premium SKUs.
	RdbBackupMaxSnapshotCount *float64 `json:"rdbBackupMaxSnapshotCount,omitempty" tf:"rdb_backup_max_snapshot_count,omitempty"`
}

type RedisConfigurationParameters struct {

	// Enable or disable AOF persistence for this Redis Cache. Defaults to false.
	// +kubebuilder:validation:Optional
	AofBackupEnabled *bool `json:"aofBackupEnabled,omitempty" tf:"aof_backup_enabled,omitempty"`

	// First Storage Account connection string for AOF persistence.
	// +kubebuilder:validation:Optional
	AofStorageConnectionString0SecretRef *v1.SecretKeySelector `json:"aofStorageConnectionString0SecretRef,omitempty" tf:"-"`

	// Second Storage Account connection string for AOF persistence.
	// +kubebuilder:validation:Optional
	AofStorageConnectionString1SecretRef *v1.SecretKeySelector `json:"aofStorageConnectionString1SecretRef,omitempty" tf:"-"`

	// If set to false, the Redis instance will be accessible without authentication. Defaults to true.
	// +kubebuilder:validation:Optional
	EnableAuthentication *bool `json:"enableAuthentication,omitempty" tf:"enable_authentication,omitempty"`

	// Value in megabytes reserved to accommodate for memory fragmentation. Defaults are shown below.
	// +kubebuilder:validation:Optional
	MaxfragmentationmemoryReserved *float64 `json:"maxfragmentationmemoryReserved,omitempty" tf:"maxfragmentationmemory_reserved,omitempty"`

	// The max-memory delta for this Redis instance. Defaults are shown below.
	// +kubebuilder:validation:Optional
	MaxmemoryDelta *float64 `json:"maxmemoryDelta,omitempty" tf:"maxmemory_delta,omitempty"`

	// How Redis will select what to remove when maxmemory is reached. Defaults are shown below. Defaults to volatile-lru.
	// +kubebuilder:validation:Optional
	MaxmemoryPolicy *string `json:"maxmemoryPolicy,omitempty" tf:"maxmemory_policy,omitempty"`

	// Value in megabytes reserved for non-cache usage e.g. failover. Defaults are shown below.
	// +kubebuilder:validation:Optional
	MaxmemoryReserved *float64 `json:"maxmemoryReserved,omitempty" tf:"maxmemory_reserved,omitempty"`

	// Keyspace notifications allows clients to subscribe to Pub/Sub channels in order to receive events affecting the Redis data set in some way. Reference
	// +kubebuilder:validation:Optional
	NotifyKeySpaceEvents *string `json:"notifyKeyspaceEvents,omitempty" tf:"notify_keyspace_events,omitempty"`

	// Is Backup Enabled? Only supported on Premium SKUs. Defaults to false.
	// +kubebuilder:validation:Optional
	RdbBackupEnabled *bool `json:"rdbBackupEnabled,omitempty" tf:"rdb_backup_enabled,omitempty"`

	// The Backup Frequency in Minutes. Only supported on Premium SKUs. Possible values are: 15, 30, 60, 360, 720 and 1440.
	// +kubebuilder:validation:Optional
	RdbBackupFrequency *float64 `json:"rdbBackupFrequency,omitempty" tf:"rdb_backup_frequency,omitempty"`

	// The maximum number of snapshots to create as a backup. Only supported for Premium SKUs.
	// +kubebuilder:validation:Optional
	RdbBackupMaxSnapshotCount *float64 `json:"rdbBackupMaxSnapshotCount,omitempty" tf:"rdb_backup_max_snapshot_count,omitempty"`

	// The Connection String to the Storage Account. Only supported for Premium SKUs. In the format: DefaultEndpointsProtocol=https;BlobEndpoint=${azurerm_storage_account.example.primary_blob_endpoint};AccountName=${azurerm_storage_account.example.name};AccountKey=${azurerm_storage_account.example.primary_access_key}.
	// +kubebuilder:validation:Optional
	RdbStorageConnectionStringSecretRef *v1.SecretKeySelector `json:"rdbStorageConnectionStringSecretRef,omitempty" tf:"-"`
}

// RedisCacheSpec defines the desired state of RedisCache
type RedisCacheSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RedisCacheParameters `json:"forProvider"`
}

// RedisCacheStatus defines the observed state of RedisCache.
type RedisCacheStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RedisCacheObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RedisCache is the Schema for the RedisCaches API. Manages a Redis Cache
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type RedisCache struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.capacity)",message="capacity is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.family)",message="family is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.redisVersion)",message="redisVersion is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.skuName)",message="skuName is a required parameter"
	Spec   RedisCacheSpec   `json:"spec"`
	Status RedisCacheStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedisCacheList contains a list of RedisCaches
type RedisCacheList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisCache `json:"items"`
}

// Repository type metadata.
var (
	RedisCache_Kind             = "RedisCache"
	RedisCache_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RedisCache_Kind}.String()
	RedisCache_KindAPIVersion   = RedisCache_Kind + "." + CRDGroupVersion.String()
	RedisCache_GroupVersionKind = CRDGroupVersion.WithKind(RedisCache_Kind)
)

func init() {
	SchemeBuilder.Register(&RedisCache{}, &RedisCacheList{})
}
