package targets

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/policies"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types/targettype"
)

const (
	bzeroBasePath   = targetsBasePath + "/bzero"
	bzeroSinglePath = bzeroBasePath + "/%s"
)

// ModifyBzeroTargetRequest is used to modify a Bzero target
type ModifyBzeroTargetRequest struct {
	TargetName    *string `json:"targetName,omitempty"`
	EnvironmentID *string `json:"environmentId,omitempty"`
}

// RestartBzeroTargetRequest is used to restart a Bzero target
type RestartBzeroTargetRequest struct {
	TargetID        string `json:"targetId,omitempty"`
	TargetName      string `json:"targetName,omitempty"`
	EnvironmentID   string `json:"envId,omitempty"`
	EnvironmentName string `json:"envName,omitempty"`
}

// RequestBzeroAgentLogsRequest is used to request the Bzero agent to post its
// logs to BastionZero
type RequestBzeroAgentLogsRequest struct {
	TargetID            string `json:"targetId,omitempty"`
	TargetName          string `json:"targetName,omitempty"`
	EnvironmentID       string `json:"envId,omitempty"`
	EnvironmentName     string `json:"envName,omitempty"`
	UploadLogsRequestId string `json:"uploadLogsRequestId"`
}

// UpdateAgentConfigRequest is used to update the Bzero agent config
type UpdateAgentConfigRequest struct {
	TargetID        string `json:"targetId,omitempty"`
	TargetName      string `json:"targetName,omitempty"`
	EnvironmentID   string `json:"envId,omitempty"`
	EnvironmentName string `json:"envName,omitempty"`
	Key             string `json:"key"`
	Value           string `json:"value"`
}

// BzeroTarget is a target running the Bzero agent
type BzeroTarget struct {
	Target

	AllowedTargetUsers []policies.TargetUser  `json:"allowedTargetUsers"`
	AllowedVerbs       []policies.Verb        `json:"allowedVerbs"`
	ControlChannel     *ControlChannelSummary `json:"controlChannel"`
}

// ListBzeroTargets lists all Bzero targets.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/targets/bzero
func (s *TargetsService) ListBzeroTargets(ctx context.Context) ([]BzeroTarget, *http.Response, error) {
	u := bzeroBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	targetList := new([]BzeroTarget)
	resp, err := s.Client.Do(req, targetList)
	if err != nil {
		return nil, resp, err
	}

	return *targetList, resp, nil
}

// GetBzeroTarget fetches the specified Bzero target.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/targets/bzero/-id-
func (s *TargetsService) GetBzeroTarget(ctx context.Context, targetID string) (*BzeroTarget, *http.Response, error) {
	u := fmt.Sprintf(bzeroSinglePath, targetID)
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	target := new(BzeroTarget)
	resp, err := s.Client.Do(req, target)
	if err != nil {
		return nil, resp, err
	}

	return target, resp, nil
}

// DeleteBzeroTarget deletes the specified Bzero target.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#delete-/api/v2/targets/bzero/-id-
func (s *TargetsService) DeleteBzeroTarget(ctx context.Context, targetID string) (*http.Response, error) {
	u := fmt.Sprintf(bzeroSinglePath, targetID)
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

// ModifyBzeroTarget updates a Bzero target.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#patch-/api/v2/targets/bzero/-id-
func (s *TargetsService) ModifyBzeroTarget(ctx context.Context, targetID string, request *ModifyBzeroTargetRequest) (*BzeroTarget, *http.Response, error) {
	u := fmt.Sprintf(bzeroSinglePath, targetID)
	req, err := s.Client.NewRequest(ctx, http.MethodPatch, u, request)
	if err != nil {
		return nil, nil, err
	}

	target := new(BzeroTarget)
	resp, err := s.Client.Do(req, target)
	if err != nil {
		return nil, resp, err
	}

	return target, resp, nil
}

// RestartBzeroTarget restarts a Bzero target.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/targets/bzero/restart
func (s *TargetsService) RestartBzeroTarget(ctx context.Context, request *RestartBzeroTargetRequest) (*http.Response, error) {
	u := bzeroBasePath + "/restart"
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, request)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// RequestBzeroTargetLogs requests that a Bzero target's agent post its logs to BastionZero.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/targets/bzero/retrieve-logs
func (s *TargetsService) RequestBzeroTargetLogs(ctx context.Context, request *RequestBzeroAgentLogsRequest) (*http.Response, error) {
	u := bzeroBasePath + "/retrieve-logs"
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, request)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// UpdateAgentConfig requests that the Bzero agent's config is updated.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/targets/bzero/update-config
func (s *TargetsService) UpdateAgentConfig(ctx context.Context, request *UpdateAgentConfigRequest) (*http.Response, error) {
	u := bzeroBasePath + "/update-config"
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, request)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Ensure BzeroTarget implementation satisfies the expected interfaces.
var (
	// BzeroTarget implements TargetInterface
	_ TargetInterface = &BzeroTarget{}
)

func (t *BzeroTarget) GetTargetType() targettype.TargetType { return targettype.Bzero }
