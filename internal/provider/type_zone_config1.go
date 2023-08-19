// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type ZoneConfig1 struct {
	EnableNetworkTypeSelection types.String             `tfsdk:"enable_network_type_selection"`
	UseHostCredentials         types.String             `tfsdk:"use_host_credentials"`
	AccessKey                  types.String             `tfsdk:"access_key"`
	AccountType                types.String             `tfsdk:"account_type"`
	APIURL                     types.String             `tfsdk:"api_url"`
	APIVersion                 types.String             `tfsdk:"api_version"`
	ApplianceURL               types.String             `tfsdk:"appliance_url"`
	AzureCostingMode           types.String             `tfsdk:"azure_costing_mode"`
	BackupMode                 types.String             `tfsdk:"backup_mode"`
	CertificateProvider        types.String             `tfsdk:"certificate_provider"`
	ClientEmail                types.String             `tfsdk:"client_email"`
	ClientID                   types.String             `tfsdk:"client_id"`
	ClientSecret               types.String             `tfsdk:"client_secret"`
	ClientSecretHash           types.String             `tfsdk:"client_secret_hash"`
	CloudType                  types.String             `tfsdk:"cloud_type"`
	Cluster                    types.String             `tfsdk:"cluster"`
	ConfigCmID                 types.String             `tfsdk:"config_cm_id"`
	ConfigCmdbDiscovery        types.Bool               `tfsdk:"config_cmdb_discovery"`
	ConfigCmdbID               types.String             `tfsdk:"config_cmdb_id"`
	ConfigManagementID         types.String             `tfsdk:"config_management_id"`
	CostingAccessKey           types.String             `tfsdk:"costing_access_key"`
	CostingBucket              types.String             `tfsdk:"costing_bucket"`
	CostingBucketName          types.String             `tfsdk:"costing_bucket_name"`
	CostingFolder              types.String             `tfsdk:"costing_folder"`
	CostingRegion              types.String             `tfsdk:"costing_region"`
	CostingReport              types.String             `tfsdk:"costing_report"`
	CostingReportName          types.String             `tfsdk:"costing_report_name"`
	CostingSecretKey           types.String             `tfsdk:"costing_secret_key"`
	CostingSecretKeyHash       types.String             `tfsdk:"costing_secret_key_hash"`
	CspClientID                types.String             `tfsdk:"csp_client_id"`
	CspClientSecret            types.String             `tfsdk:"csp_client_secret"`
	CspClientSecretHash        types.String             `tfsdk:"csp_client_secret_hash"`
	CspCustomer                types.String             `tfsdk:"csp_customer"`
	CspTenantID                types.String             `tfsdk:"csp_tenant_id"`
	Datacenter                 types.String             `tfsdk:"datacenter"`
	DatacenterID               types.String             `tfsdk:"datacenter_id"`
	DatacenterName             types.String             `tfsdk:"datacenter_name"`
	DiskEncryption             types.String             `tfsdk:"disk_encryption"`
	DiskStorageType            types.String             `tfsdk:"disk_storage_type"`
	DistributedWorkerID        types.String             `tfsdk:"distributed_worker_id"`
	DNSIntegrationID           types.String             `tfsdk:"dns_integration_id"`
	EbsEncryption              types.String             `tfsdk:"ebs_encryption"`
	EnableDiskTypeSelection    types.String             `tfsdk:"enable_disk_type_selection"`
	EnableVnc                  types.String             `tfsdk:"enable_vnc"`
	EncryptionSet              types.String             `tfsdk:"encryption_set"`
	Endpoint                   types.String             `tfsdk:"endpoint"`
	GoogleRegionID             types.String             `tfsdk:"google_region_id"`
	HideHostSelection          types.String             `tfsdk:"hide_host_selection"`
	ImageStoreID               types.String             `tfsdk:"image_store_id"`
	ImportExisting             types.String             `tfsdk:"import_existing"`
	InventoryLevel             types.String             `tfsdk:"inventory_level"`
	IsVpc                      types.String             `tfsdk:"is_vpc"`
	KubeURL                    types.String             `tfsdk:"kube_url"`
	NetworkServer              *ZoneConfigNetworkServer `tfsdk:"network_server"`
	NetworkServerID            types.String             `tfsdk:"network_server_id"`
	Password                   types.String             `tfsdk:"password"`
	PasswordHash               types.String             `tfsdk:"password_hash"`
	PrivateKey                 types.String             `tfsdk:"private_key"`
	PrivateKeyHash             types.String             `tfsdk:"private_key_hash"`
	ProjectID                  types.String             `tfsdk:"project_id"`
	ReplicationMode            types.String             `tfsdk:"replication_mode"`
	ResourceGroup              types.String             `tfsdk:"resource_group"`
	ResourcePool               types.String             `tfsdk:"resource_pool"`
	ResourcePoolID             types.String             `tfsdk:"resource_pool_id"`
	RPCMode                    types.String             `tfsdk:"rpc_mode"`
	SecretKey                  types.String             `tfsdk:"secret_key"`
	SecretKeyHash              types.String             `tfsdk:"secret_key_hash"`
	SecurityMode               types.String             `tfsdk:"security_mode"`
	SecurityServer             types.String             `tfsdk:"security_server"`
	ServiceRegistryID          types.String             `tfsdk:"service_registry_id"`
	StsAssumeRole              types.String             `tfsdk:"sts_assume_role"`
	SubscriberID               types.String             `tfsdk:"subscriber_id"`
	TenantID                   types.String             `tfsdk:"tenant_id"`
	Username                   types.String             `tfsdk:"username"`
	Vpc                        types.String             `tfsdk:"vpc"`
}