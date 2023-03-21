package targets

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/policies"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/targets/targetstatus"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types"
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

// BzeroTarget is a target running the Bzero agent
type BzeroTarget struct {
	ID                 string                    `json:"id"`
	Name               string                    `json:"name"`
	Status             targetstatus.TargetStatus `json:"status"`
	EnvironmentID      string                    `json:"environmentId"`
	LastAgentUpdate    *types.Timestamp          `json:"lastAgentUpdate"`
	AgentVersion       string                    `json:"agentVersion"`
	Region             string                    `json:"region"`
	AgentPublicKey     string                    `json:"agentPublicKey"`
	AllowedTargetUsers []policies.TargetUser     `json:"allowedTargetUsers"`
	AllowedVerbs       []policies.Verb           `json:"allowedVerbs"`
	ControlChannel     *ControlChannelSummary    `json:"controlChannel"`
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

// Ensure BzeroTarget implementation satisfies the expected interfaces.
var (
	// BzeroTarget implements TargetInterface
	_ TargetInterface = &BzeroTarget{}
)

func (t *BzeroTarget) GetID() string                        { return t.ID }
func (t *BzeroTarget) GetName() string                      { return t.Name }
func (t *BzeroTarget) GetStatus() targetstatus.TargetStatus { return t.Status }
func (t *BzeroTarget) GetEnvironmentID() string             { return t.EnvironmentID }
func (t *BzeroTarget) GetLastAgentUpdate() *types.Timestamp { return t.LastAgentUpdate }
func (t *BzeroTarget) GetAgentVersion() string              { return t.AgentVersion }
func (t *BzeroTarget) GetRegion() string                    { return t.Region }
func (t *BzeroTarget) GetAgentPublicKey() string            { return t.AgentPublicKey }
func (t *BzeroTarget) GetTargetType() targettype.TargetType { return targettype.Bzero }
