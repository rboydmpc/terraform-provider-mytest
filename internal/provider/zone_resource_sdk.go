// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"MyTest/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *ZoneResourceModel) ToCreateSDKType() *shared.ZoneCreate {
	accountID := new(int64)
	if !r.AccountID.IsUnknown() && !r.AccountID.IsNull() {
		*accountID = r.AccountID.ValueInt64()
	} else {
		accountID = nil
	}
	code := new(string)
	if !r.Code.IsUnknown() && !r.Code.IsNull() {
		*code = r.Code.ValueString()
	} else {
		code = nil
	}
	var config *shared.ZoneVcenterConfig
	if r.Config != nil {
		apiURL := new(string)
		if !r.Config.APIURL.IsUnknown() && !r.Config.APIURL.IsNull() {
			*apiURL = r.Config.APIURL.ValueString()
		} else {
			apiURL = nil
		}
		applianceURL := new(string)
		if !r.Config.ApplianceURL.IsUnknown() && !r.Config.ApplianceURL.IsNull() {
			*applianceURL = r.Config.ApplianceURL.ValueString()
		} else {
			applianceURL = nil
		}
		datacenter := new(string)
		if !r.Config.Datacenter.IsUnknown() && !r.Config.Datacenter.IsNull() {
			*datacenter = r.Config.Datacenter.ValueString()
		} else {
			datacenter = nil
		}
		username := new(string)
		if !r.Config.Username.IsUnknown() && !r.Config.Username.IsNull() {
			*username = r.Config.Username.ValueString()
		} else {
			username = nil
		}
		config = &shared.ZoneVcenterConfig{
			APIURL:       apiURL,
			ApplianceURL: applianceURL,
			Datacenter:   datacenter,
			Username:     username,
		}
	}
	var credential *shared.ZoneCreateCredential
	if r.Credential != nil {
		typeVar := new(string)
		if !r.Credential.Type.IsUnknown() && !r.Credential.Type.IsNull() {
			*typeVar = r.Credential.Type.ValueString()
		} else {
			typeVar = nil
		}
		credential = &shared.ZoneCreateCredential{
			Type: typeVar,
		}
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	groupID := r.GroupID.ValueInt64()
	name := r.Name.ValueString()
	scalePriority := new(int64)
	if !r.ScalePriority.IsUnknown() && !r.ScalePriority.IsNull() {
		*scalePriority = r.ScalePriority.ValueInt64()
	} else {
		scalePriority = nil
	}
	visibility := new(shared.ZoneCreateVisibility)
	if !r.Visibility.IsUnknown() && !r.Visibility.IsNull() {
		*visibility = shared.ZoneCreateVisibility(r.Visibility.ValueString())
	} else {
		visibility = nil
	}
	code1 := new(string)
	if !r.ZoneType.Code.IsUnknown() && !r.ZoneType.Code.IsNull() {
		*code1 = r.ZoneType.Code.ValueString()
	} else {
		code1 = nil
	}
	zoneType := shared.ZoneCreateZoneType{
		Code: code1,
	}
	out := shared.ZoneCreate{
		AccountID:     accountID,
		Code:          code,
		Config:        config,
		Credential:    credential,
		Description:   description,
		Enabled:       enabled,
		GroupID:       groupID,
		Name:          name,
		ScalePriority: scalePriority,
		Visibility:    visibility,
		ZoneType:      zoneType,
	}
	return &out
}

func (r *ZoneResourceModel) ToGetSDKType() *shared.ZoneCreate {
	out := r.ToCreateSDKType()
	return out
}

func (r *ZoneResourceModel) ToUpdateSDKType() *shared.Zone {
	accountID := new(int64)
	if !r.AccountID.IsUnknown() && !r.AccountID.IsNull() {
		*accountID = r.AccountID.ValueInt64()
	} else {
		accountID = nil
	}
	code := new(string)
	if !r.Code.IsUnknown() && !r.Code.IsNull() {
		*code = r.Code.ValueString()
	} else {
		code = nil
	}
	var config *shared.ZoneVcenterConfig2
	if r.Config != nil {
		apiURL := new(string)
		if !r.Config.APIURL.IsUnknown() && !r.Config.APIURL.IsNull() {
			*apiURL = r.Config.APIURL.ValueString()
		} else {
			apiURL = nil
		}
		applianceURL := new(string)
		if !r.Config.ApplianceURL.IsUnknown() && !r.Config.ApplianceURL.IsNull() {
			*applianceURL = r.Config.ApplianceURL.ValueString()
		} else {
			applianceURL = nil
		}
		datacenter := new(string)
		if !r.Config.Datacenter.IsUnknown() && !r.Config.Datacenter.IsNull() {
			*datacenter = r.Config.Datacenter.ValueString()
		} else {
			datacenter = nil
		}
		username := new(string)
		if !r.Config.Username.IsUnknown() && !r.Config.Username.IsNull() {
			*username = r.Config.Username.ValueString()
		} else {
			username = nil
		}
		config = &shared.ZoneVcenterConfig2{
			APIURL:       apiURL,
			ApplianceURL: applianceURL,
			Datacenter:   datacenter,
			Username:     username,
		}
	}
	var credential *shared.ZoneCredential
	if r.Credential != nil {
		typeVar := new(string)
		if !r.Credential.Type.IsUnknown() && !r.Credential.Type.IsNull() {
			*typeVar = r.Credential.Type.ValueString()
		} else {
			typeVar = nil
		}
		credential = &shared.ZoneCredential{
			Type: typeVar,
		}
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	var groups []shared.ZoneGroups = nil
	for _, groupsItem := range r.Groups {
		accountId1 := new(int64)
		if !groupsItem.AccountID.IsUnknown() && !groupsItem.AccountID.IsNull() {
			*accountId1 = groupsItem.AccountID.ValueInt64()
		} else {
			accountId1 = nil
		}
		id := new(int64)
		if !groupsItem.ID.IsUnknown() && !groupsItem.ID.IsNull() {
			*id = groupsItem.ID.ValueInt64()
		} else {
			id = nil
		}
		name := new(string)
		if !groupsItem.Name.IsUnknown() && !groupsItem.Name.IsNull() {
			*name = groupsItem.Name.ValueString()
		} else {
			name = nil
		}
		groups = append(groups, shared.ZoneGroups{
			AccountID: accountId1,
			ID:        id,
			Name:      name,
		})
	}
	id1 := new(int64)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id1 = r.ID.ValueInt64()
	} else {
		id1 = nil
	}
	name1 := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name1 = r.Name.ValueString()
	} else {
		name1 = nil
	}
	scalePriority := new(int64)
	if !r.ScalePriority.IsUnknown() && !r.ScalePriority.IsNull() {
		*scalePriority = r.ScalePriority.ValueInt64()
	} else {
		scalePriority = nil
	}
	visibility := new(string)
	if !r.Visibility.IsUnknown() && !r.Visibility.IsNull() {
		*visibility = r.Visibility.ValueString()
	} else {
		visibility = nil
	}
	var zoneType *shared.ZoneZoneType
	if r.ZoneType != nil {
		code1 := new(string)
		if !r.ZoneType.Code.IsUnknown() && !r.ZoneType.Code.IsNull() {
			*code1 = r.ZoneType.Code.ValueString()
		} else {
			code1 = nil
		}
		zoneType = &shared.ZoneZoneType{
			Code: code1,
		}
	}
	out := shared.Zone{
		AccountID:     accountID,
		Code:          code,
		Config:        config,
		Credential:    credential,
		Enabled:       enabled,
		Groups:        groups,
		ID:            id1,
		Name:          name1,
		ScalePriority: scalePriority,
		Visibility:    visibility,
		ZoneType:      zoneType,
	}
	return &out
}

func (r *ZoneResourceModel) ToDeleteSDKType() *shared.ZoneCreate {
	out := r.ToCreateSDKType()
	return out
}

func (r *ZoneResourceModel) RefreshFromGetResponse(resp *shared.Zone) {
	if resp.AccountID != nil {
		r.AccountID = types.Int64Value(*resp.AccountID)
	} else {
		r.AccountID = types.Int64Null()
	}
	if resp.Code != nil {
		r.Code = types.StringValue(*resp.Code)
	} else {
		r.Code = types.StringNull()
	}
	if r.Config == nil {
		r.Config = &ZoneVcenterConfig{}
	}
	if resp.Config == nil {
		r.Config = nil
	} else {
		r.Config = &ZoneVcenterConfig{}
		if resp.Config.APIURL != nil {
			r.Config.APIURL = types.StringValue(*resp.Config.APIURL)
		} else {
			r.Config.APIURL = types.StringNull()
		}
		if resp.Config.ApplianceURL != nil {
			r.Config.ApplianceURL = types.StringValue(*resp.Config.ApplianceURL)
		} else {
			r.Config.ApplianceURL = types.StringNull()
		}
		if resp.Config.Datacenter != nil {
			r.Config.Datacenter = types.StringValue(*resp.Config.Datacenter)
		} else {
			r.Config.Datacenter = types.StringNull()
		}
		if resp.Config.Username != nil {
			r.Config.Username = types.StringValue(*resp.Config.Username)
		} else {
			r.Config.Username = types.StringNull()
		}
	}
	if r.Credential == nil {
		r.Credential = &ZoneCredential{}
	}
	if resp.Credential == nil {
		r.Credential = nil
	} else {
		r.Credential = &ZoneCredential{}
		if resp.Credential.Type != nil {
			r.Credential.Type = types.StringValue(*resp.Credential.Type)
		} else {
			r.Credential.Type = types.StringNull()
		}
	}
	if resp.Enabled != nil {
		r.Enabled = types.BoolValue(*resp.Enabled)
	} else {
		r.Enabled = types.BoolNull()
	}
	r.Groups = nil
	for _, groupsItem := range resp.Groups {
		var groups1 ZoneGroups
		if groupsItem.AccountID != nil {
			groups1.AccountID = types.Int64Value(*groupsItem.AccountID)
		} else {
			groups1.AccountID = types.Int64Null()
		}
		if groupsItem.ID != nil {
			groups1.ID = types.Int64Value(*groupsItem.ID)
		} else {
			groups1.ID = types.Int64Null()
		}
		if groupsItem.Name != nil {
			groups1.Name = types.StringValue(*groupsItem.Name)
		} else {
			groups1.Name = types.StringNull()
		}
		r.Groups = append(r.Groups, groups1)
	}
	if resp.ID != nil {
		r.ID = types.Int64Value(*resp.ID)
	} else {
		r.ID = types.Int64Null()
	}
	if resp.Name != nil {
		r.Name = types.StringValue(*resp.Name)
	} else {
		r.Name = types.StringNull()
	}
	if resp.ScalePriority != nil {
		r.ScalePriority = types.Int64Value(*resp.ScalePriority)
	} else {
		r.ScalePriority = types.Int64Null()
	}
	if resp.Visibility != nil {
		r.Visibility = types.StringValue(*resp.Visibility)
	} else {
		r.Visibility = types.StringNull()
	}
	if r.ZoneType == nil {
		r.ZoneType = &ZoneZoneType{}
	}
	if resp.ZoneType == nil {
		r.ZoneType = nil
	} else {
		r.ZoneType = &ZoneZoneType{}
		if resp.ZoneType.Code != nil {
			r.ZoneType.Code = types.StringValue(*resp.ZoneType.Code)
		} else {
			r.ZoneType.Code = types.StringNull()
		}
	}
}

func (r *ZoneResourceModel) RefreshFromCreateResponse(resp *shared.Zone) {
	r.RefreshFromGetResponse(resp)
}

func (r *ZoneResourceModel) RefreshFromUpdateResponse(resp *shared.Zone) {
	r.RefreshFromGetResponse(resp)
}
