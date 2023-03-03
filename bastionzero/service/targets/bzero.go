package targets

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/policies"
)

const (
	bzeroBasePath   = targetsBasePath + "/bzero"
	bzeroSinglePath = bzeroBasePath + "/%s"
)

// BzeroTarget is a target running the Bzero agent
type BzeroTarget struct {
	ID                 string                      `json:"id"`
	Name               string                      `json:"name"`
	Status             TargetStatus                `json:"status"`
	EnvironmentID      string                      `json:"environmentId"`
	LastAgentUpdate    *service.Timestamp          `json:"lastAgentUpdate"`
	AgentVersion       string                      `json:"agentVersion"`
	Region             string                      `json:"region"`
	AgentPublicKey     string                      `json:"agentPublicKey"`
	AllowedTargetUsers []policies.PolicyTargetUser `json:"allowedTargetUsers"`
	AllowedVerbs       []policies.Verb             `json:"allowedVerbs"`
	ControlChannel     *ControlChannelSummary      `json:"controlChannel"`
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
