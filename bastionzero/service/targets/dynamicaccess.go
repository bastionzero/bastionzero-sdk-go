package targets

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/policies"
)

const (
	dynamicAccessBasePath   = targetsBasePath + "/dynamic-access"
	dynamicAccessSinglePath = dynamicAccessBasePath + "/%s"
)

// CreateDynamicAccessConfigurationRequest is used to create a new dynamic
// access configuration
type CreateDynamicAccessConfigurationRequest struct {
	Name          string  `json:"name"`
	StartWebhook  string  `json:"startWebhook"`
	StopWebhook   string  `json:"stopWebhook"`
	HealthWebhook string  `json:"healthWebhook"`
	EnvironmentId string  `json:"environmentId"`
	SharedSecret  *string `json:"sharedSecret,omitempty"`
}

// CreateDynamicAccessResponse is the response returned if a dynamic access
// configuration is successfully created
type CreateDynamicAccessResponse struct {
	ID string `json:"id"`
}

// ModifyDynamicAccessConfigurationRequest is used to modify a dynamic access
// configuration
type ModifyDynamicAccessConfigurationRequest struct {
	Name          *string `json:"name,omitempty"`
	StartWebhook  *string `json:"startWebhook,omitempty"`
	StopWebhook   *string `json:"stopWebhook,omitempty"`
	HealthWebhook *string `json:"healthWebhook,omitempty"`
	SharedSecret  *string `json:"sharedSecret,omitempty"`
}

// DynamicAccessConfigurationStatus represents the status of a dynamic access
// configuration
type DynamicAccessConfigurationStatus string

const (
	// DACOffline indicates the health webhook of the DAC is responding
	// unhealthy
	DACOffline DynamicAccessConfigurationStatus = "Offline"
	// DACOnline indicates the health webhook of the DAC is responding healthy
	DACOnline DynamicAccessConfigurationStatus = "Online"
)

// DynamicAccessConfiguration configures dynamic access targets
type DynamicAccessConfiguration struct {
	ID                 string                           `json:"id"`
	Name               string                           `json:"name"`
	EnvironmentId      string                           `json:"environmentId"`
	StartWebhook       string                           `json:"startWebhook"`
	StopWebhook        string                           `json:"stopWebhook"`
	HealthWebhook      string                           `json:"healthWebhook"`
	AllowedTargetUsers []policies.PolicyTargetUser      `json:"allowedTargetUsers"`
	AllowedVerbs       []policies.Verb                  `json:"allowedVerbs"`
	Status             DynamicAccessConfigurationStatus `json:"status"`
}

// ListDynamicAccessConfigurations lists all dynamic access configurations.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/targets/dynamic-access
func (s *TargetsService) ListDynamicAccessConfigurations(ctx context.Context) ([]DynamicAccessConfiguration, *http.Response, error) {
	u := dynamicAccessBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	dacList := new([]DynamicAccessConfiguration)
	resp, err := s.Client.Do(req, dacList)
	if err != nil {
		return nil, resp, err
	}

	return *dacList, resp, nil
}

// CreateDynamicAccessConfiguration creates a new dynamic access configuration.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/targets/dynamic-access
func (s *TargetsService) CreateDynamicAccessConfiguration(ctx context.Context, request *CreateDynamicAccessConfigurationRequest) (*CreateDynamicAccessResponse, *http.Response, error) {
	u := dynamicAccessBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, request)
	if err != nil {
		return nil, nil, err
	}

	createDACResponse := new(CreateDynamicAccessResponse)
	resp, err := s.Client.Do(req, createDACResponse)
	if err != nil {
		return nil, resp, err
	}

	return createDACResponse, resp, nil
}

// GetDynamicAccessConfiguration fetches the specified dynamic access configuration.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/targets/dynamic-access/-id-
func (s *TargetsService) GetDynamicAccessConfiguration(ctx context.Context, dacID string) (*DynamicAccessConfiguration, *http.Response, error) {
	u := fmt.Sprintf(dynamicAccessSinglePath, dacID)
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	dac := new(DynamicAccessConfiguration)
	resp, err := s.Client.Do(req, dac)
	if err != nil {
		return nil, resp, err
	}

	return dac, resp, nil
}

// DeleteDynamicAccessConfiguration deletes the specified dynamic access configuration.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#delete-/api/v2/targets/dynamic-access/-id-
func (s *TargetsService) DeleteDynamicAccessConfiguration(ctx context.Context, dacID string) (*http.Response, error) {
	u := fmt.Sprintf(dynamicAccessSinglePath, dacID)
	req, err := s.Client.NewRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ModifyDynamicAccessConfiguration updates a dynamic access configuration.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#patch-/api/v2/targets/dynamic-access/-id-
func (s *TargetsService) ModifyDynamicAccessConfiguration(ctx context.Context, dacID string, request *ModifyDynamicAccessConfigurationRequest) (*http.Response, error) {
	u := fmt.Sprintf(dynamicAccessSinglePath, dacID)
	req, err := s.Client.NewRequest(ctx, http.MethodPatch, u, request)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
