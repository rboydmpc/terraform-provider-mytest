// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type ZoneVcenterConfig struct {
	APIURL       types.String `tfsdk:"api_url"`
	ApplianceURL types.String `tfsdk:"appliance_url"`
	Datacenter   types.String `tfsdk:"datacenter"`
}
