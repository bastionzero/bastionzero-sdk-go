package targets

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service"
)

const (
	kubeBasePath   = targetsBasePath + "/kube"
	kubeSinglePath = kubeBasePath + "/%s"
)

// ClusterTarget is a target running the Bctl agent within a Kubernetes cluster
type ClusterTarget struct {
	ID                   string                 `json:"id"`
	Name                 string                 `json:"name"`
	Status               TargetStatus           `json:"status"`
	EnvironmentID        string                 `json:"environmentId"`
	LastAgentUpdate      *service.Timestamp     `json:"lastAgentUpdate"`
	AgentVersion         string                 `json:"agentVersion"`
	Region               string                 `json:"region"`
	AgentPublicKey       string                 `json:"agentPublicKey"`
	AllowedClusterUsers  []string               `json:"allowedClusterUsers"`
	AllowedClusterGroups []string               `json:"allowedClusterGroups"`
	ValidClusterUsers    []string               `json:"validClusterUsers"`
	ControlChannel       *ControlChannelSummary `json:"controlChannel"`
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
