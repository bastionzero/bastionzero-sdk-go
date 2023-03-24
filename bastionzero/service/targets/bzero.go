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

func (t *BzeroTarget) GetTargetType() targettype.TargetType { return targettype.Bzero }
