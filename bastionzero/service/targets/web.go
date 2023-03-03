package targets

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service"
)

const (
	webBasePath   = targetsBasePath + "/web"
	webSinglePath = webBasePath + "/%s"
)

// WebTarget is a virtual target that provides HTTP(S) access to a remote web
// server. The connection is proxied by a Bzero or Cluster target.
type WebTarget struct {
	ID              string             `json:"id"`
	Name            string             `json:"name"`
	Status          TargetStatus       `json:"status"`
	ProxyTargetID   string             `json:"proxyTargetId"`
	LastAgentUpdate *service.Timestamp `json:"lastAgentUpdate"`
	AgentVersion    string             `json:"agentVersion"`
	RemoteHost      string             `json:"remoteHost"`
	RemotePort      Port               `json:"remotePort"`
	LocalPort       Port               `json:"localPort"`
	LocalHost       string             `json:"localHost"`
	EnvironmentID   string             `json:"environmentId"`
	Region          string             `json:"region"`
	AgentPublicKey  string             `json:"agentPublicKey"`
}

// ListWebTargets lists all Web targets.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/targets/web
func (s *TargetsService) ListWebTargets(ctx context.Context) ([]WebTarget, *http.Response, error) {
	u := webBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	targetList := new([]WebTarget)
	resp, err := s.Client.Do(req, targetList)
	if err != nil {
		return nil, resp, err
	}

	return *targetList, resp, nil
}

// GetWebTarget fetches the specified Web target.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/targets/web/-id-
func (s *TargetsService) GetWebTarget(ctx context.Context, targetID string) (*WebTarget, *http.Response, error) {
	u := fmt.Sprintf(webSinglePath, targetID)
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	target := new(WebTarget)
	resp, err := s.Client.Do(req, target)
	if err != nil {
		return nil, resp, err
	}

	return target, resp, nil
}
