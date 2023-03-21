package targets

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/targets/targetstatus"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types/targettype"
)

const (
	kubeBasePath   = targetsBasePath + "/kube"
	kubeSinglePath = kubeBasePath + "/%s"
)

// ModifyClusterTargetRequest is used to modify a Cluster target
type ModifyClusterTargetRequest struct {
	TargetName    *string `json:"name,omitempty"`
	EnvironmentID *string `json:"environmentId,omitempty"`
}

// ModifyClusterTargetResponse is the response returned if a Cluster target is
// successfully modified
type ModifyClusterTargetResponse struct {
	ID            string `json:"id"`
	TargetName    string `json:"name"`
	EnvironmentID string `json:"environmentId"`
}

// ClusterTarget is a target running the Bctl agent within a Kubernetes cluster
type ClusterTarget struct {
	ID                   string                    `json:"id"`
	Name                 string                    `json:"name"`
	Status               targetstatus.TargetStatus `json:"status"`
	EnvironmentID        string                    `json:"environmentId"`
	LastAgentUpdate      *types.Timestamp          `json:"lastAgentUpdate"`
	AgentVersion         string                    `json:"agentVersion"`
	Region               string                    `json:"region"`
	AgentPublicKey       string                    `json:"agentPublicKey"`
	AllowedClusterUsers  []string                  `json:"allowedClusterUsers"`
	AllowedClusterGroups []string                  `json:"allowedClusterGroups"`
	ValidClusterUsers    []string                  `json:"validClusterUsers"`
	ControlChannel       *ControlChannelSummary    `json:"controlChannel"`
}

// ListClusterTargets lists all Cluster targets.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/targets/kube
func (s *TargetsService) ListClusterTargets(ctx context.Context) ([]ClusterTarget, *http.Response, error) {
	u := kubeBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	targetList := new([]ClusterTarget)
	resp, err := s.Client.Do(req, targetList)
	if err != nil {
		return nil, resp, err
	}

	return *targetList, resp, nil
}

// GetClusterTarget fetches the specified Cluster target.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/targets/kube/-id-
func (s *TargetsService) GetClusterTarget(ctx context.Context, targetID string) (*ClusterTarget, *http.Response, error) {
	u := fmt.Sprintf(kubeSinglePath, targetID)
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	target := new(ClusterTarget)
	resp, err := s.Client.Do(req, target)
	if err != nil {
		return nil, resp, err
	}

	return target, resp, nil
}

// ModifyClusterTarget updates a Cluster target.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#patch-/api/v2/targets/kube/-id-
func (s *TargetsService) ModifyClusterTarget(ctx context.Context, targetID string, request *ModifyClusterTargetRequest) (*ModifyClusterTargetResponse, *http.Response, error) {
	u := fmt.Sprintf(kubeSinglePath, targetID)
	req, err := s.Client.NewRequest(ctx, http.MethodPatch, u, request)
	if err != nil {
		return nil, nil, err
	}

	modifyResp := new(ModifyClusterTargetResponse)
	resp, err := s.Client.Do(req, modifyResp)
	if err != nil {
		return nil, resp, err
	}

	return modifyResp, resp, nil
}

// Ensure ClusterTarget implementation satisfies the expected interfaces.
var (
	// ClusterTarget implements TargetInterface
	_ TargetInterface = &ClusterTarget{}
)

func (t *ClusterTarget) GetID() string                        { return t.ID }
func (t *ClusterTarget) GetName() string                      { return t.Name }
func (t *ClusterTarget) GetStatus() targetstatus.TargetStatus { return t.Status }
func (t *ClusterTarget) GetEnvironmentID() string             { return t.EnvironmentID }
func (t *ClusterTarget) GetLastAgentUpdate() *types.Timestamp { return t.LastAgentUpdate }
func (t *ClusterTarget) GetAgentVersion() string              { return t.AgentVersion }
func (t *ClusterTarget) GetRegion() string                    { return t.Region }
func (t *ClusterTarget) GetAgentPublicKey() string            { return t.AgentPublicKey }
func (t *ClusterTarget) GetTargetType() targettype.TargetType { return targettype.Cluster }
