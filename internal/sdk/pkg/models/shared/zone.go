// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
)

type ZoneConfigType string

const (
	ZoneConfigTypeZoneVcenterConfig ZoneConfigType = "zoneVcenterConfig"
)

type ZoneConfig struct {
	ZoneVcenterConfig *ZoneVcenterConfig

	Type ZoneConfigType
}

func CreateZoneConfigZoneVcenterConfig(zoneVcenterConfig ZoneVcenterConfig) ZoneConfig {
	typ := ZoneConfigTypeZoneVcenterConfig

	return ZoneConfig{
		ZoneVcenterConfig: &zoneVcenterConfig,
		Type:              typ,
	}
}

func (u *ZoneConfig) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	zoneVcenterConfig := new(ZoneVcenterConfig)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&zoneVcenterConfig); err == nil {
		u.ZoneVcenterConfig = zoneVcenterConfig
		u.Type = ZoneConfigTypeZoneVcenterConfig
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u ZoneConfig) MarshalJSON() ([]byte, error) {
	if u.ZoneVcenterConfig != nil {
		return json.Marshal(u.ZoneVcenterConfig)
	}

	return nil, nil
}

type ZoneCredential struct {
	ID *int64 `json:"id,omitempty"`
}

type ZoneGroups struct {
	AccountID *int64  `json:"accountId,omitempty"`
	ID        *int64  `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
}

// ZoneZoneType - Map containing the Cloud (zone) code name. See the zone-types API to fetch a list of all available Cloud (zone) types and their codes.
type ZoneZoneType struct {
	Code *string `json:"code,omitempty"`
}

type Zone struct {
	AccountID     *int64          `json:"accountId,omitempty"`
	Code          *string         `json:"code,omitempty"`
	Config        *ZoneConfig     `json:"config,omitempty"`
	Credential    *ZoneCredential `json:"credential,omitempty"`
	Enabled       *bool           `json:"enabled,omitempty"`
	Groups        []ZoneGroups    `json:"groups,omitempty"`
	ID            *int64          `json:"id,omitempty"`
	Name          *string         `json:"name,omitempty"`
	ScalePriority *int64          `json:"scalePriority,omitempty"`
	Visibility    *string         `json:"visibility,omitempty"`
	// Map containing the Cloud (zone) code name. See the zone-types API to fetch a list of all available Cloud (zone) types and their codes.
	ZoneType *ZoneZoneType `json:"zoneType,omitempty"`
}
