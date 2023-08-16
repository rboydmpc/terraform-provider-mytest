// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"MyTest/internal/sdk/pkg/models/operations"
	"MyTest/internal/sdk/pkg/models/shared"
	"MyTest/internal/sdk/pkg/utils"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// applianceSettings - Manage Appliance Settings
type applianceSettings struct {
	sdkConfiguration sdkConfiguration
}

func newApplianceSettings(sdkConfig sdkConfiguration) *applianceSettings {
	return &applianceSettings{
		sdkConfiguration: sdkConfig,
	}
}

// SetApplianceSettingsMaintenanceMode - Toggle Maintenance Mode
// This endpoint allows toggling the appliance maintenance mode.
func (s *applianceSettings) SetApplianceSettingsMaintenanceMode(ctx context.Context, request operations.SetApplianceSettingsMaintenanceModeRequest) (*operations.SetApplianceSettingsMaintenanceModeResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/appliance-settings/maintenance"

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.SetApplianceSettingsMaintenanceModeResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.TwoHundredSuccess
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.TwoHundredSuccess = out
		}
	}

	return res, nil
}
