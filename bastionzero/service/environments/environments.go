package environments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types/targettype"
	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	environmentsBasePath   = "api/v2/environments"
	environmentsSinglePath = environmentsBasePath + "/%s"
)

// EnvironmentsService handles communication with the environments endpoints of
// the BastionZero API.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#tag--Environments
type EnvironmentsService client.Service

// CreateEnvironmentRequest is used to create a new environment
type CreateEnvironmentRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	// OfflineCleanupTimeoutHours is the amount of time (in hours) to wait until
	// offline targets are automatically removed by BastionZero.
	OfflineCleanupTimeoutHours uint `json:"offlineCleanupTimeoutHours"`
}

// CreateEnvironmentResponse is the response returned if an environment is
// successfully created
type CreateEnvironmentResponse struct {
	ID string `json:"id"`
}

// ModifyEnvironmentRequest is used to modify an environment
type ModifyEnvironmentRequest struct {
	Description *string `json:"description,omitempty"`
	// OfflineCleanupTimeoutHours is the amount of time (in hours) to wait until
	// offline targets are automatically removed by BastionZero.
	OfflineCleanupTimeoutHours *uint `json:"offlineCleanupTimeoutHours,omitempty"`
}

// TargetSummary describes a target associated with an environment
type TargetSummary struct {
	ID   string                `json:"id"`
	Type targettype.TargetType `json:"targetType"`
}

// Environment is a collection of targets
type Environment struct {
	ID                         string          `json:"id"`
	OrganizationID             string          `json:"organizationId"`
	IsDefault                  bool            `json:"isDefault"`
	Name                       string          `json:"name"`
	Description                *string         `json:"description"`
	TimeCreated                types.Timestamp `json:"timeCreated"`
	OfflineCleanupTimeoutHours uint            `json:"offlineCleanupTimeoutHours"`
	Targets                    []TargetSummary `json:"targets"`
}

// ListEnvironments lists all environments.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/environments
func (s *EnvironmentsService) ListEnvironments(ctx context.Context) ([]Environment, *http.Response, error) {
	u := environmentsBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	environmentList := new([]Environment)
	resp, err := s.Client.Do(req, environmentList)
	if err != nil {
		return nil, resp, err
	}

	return *environmentList, resp, nil
}

// CreateEnvironment creates a new environment.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/environments
func (s *EnvironmentsService) CreateEnvironment(ctx context.Context, request *CreateEnvironmentRequest) (*CreateEnvironmentResponse, *http.Response, error) {
	u := environmentsBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, request)
	if err != nil {
		return nil, nil, err
	}

	createEnvResponse := new(CreateEnvironmentResponse)
	resp, err := s.Client.Do(req, createEnvResponse)
	if err != nil {
		return nil, resp, err
	}

	return createEnvResponse, resp, nil
}

// GetEnvironment fetches the specified environment.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/environments/-id-
func (s *EnvironmentsService) GetEnvironment(ctx context.Context, environmentID string) (*Environment, *http.Response, error) {
	u := fmt.Sprintf(environmentsSinglePath, environmentID)
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	environment := new(Environment)
	resp, err := s.Client.Do(req, environment)
	if err != nil {
		return nil, resp, err
	}

	return environment, resp, nil
}

// DeleteEnvironment deletes the specified environment.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#delete-/api/v2/environments/-id-
func (s *EnvironmentsService) DeleteEnvironment(ctx context.Context, environmentID string) (*http.Response, error) {
	u := fmt.Sprintf(environmentsSinglePath, environmentID)
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

// ModifyEnvironment updates an environment.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#patch-/api/v2/environments/-id-
func (s *EnvironmentsService) ModifyEnvironment(ctx context.Context, environmentID string, request *ModifyEnvironmentRequest) (*http.Response, error) {
	u := fmt.Sprintf(environmentsSinglePath, environmentID)
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

func (e *Environment) GetDescription() string {
	if e.Description == nil {
		return ""
	}
	return *e.Description
}
