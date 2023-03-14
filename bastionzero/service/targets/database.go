package targets

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/policies"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/targets/targetstatus"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types"
)

const (
	databaseBasePath   = targetsBasePath + "/database"
	databaseSinglePath = databaseBasePath + "/%s"
)

// DatabaseTarget is a virtual target that provides DB access to a remote
// database. The connection is proxied by a Bzero or Cluster target. The remote
// server doesn't necessarily have to be a database as under the hood it is a
// proxied TCP connection.
type DatabaseTarget struct {
	ID                 string                      `json:"id"`
	Name               string                      `json:"name"`
	Status             targetstatus.TargetStatus   `json:"status"`
	ProxyTargetID      string                      `json:"proxyTargetId"`
	LastAgentUpdate    *types.Timestamp            `json:"lastAgentUpdate"`
	AgentVersion       string                      `json:"agentVersion"`
	RemoteHost         string                      `json:"remoteHost"`
	RemotePort         Port                        `json:"remotePort"`
	LocalPort          Port                        `json:"localPort"`
	LocalHost          string                      `json:"localHost"`
	IsSplitCert        bool                        `json:"splitCert"`
	DatabaseType       *string                     `json:"databaseType"`
	EnvironmentID      string                      `json:"environmentId"`
	Region             string                      `json:"region"`
	AgentPublicKey     string                      `json:"agentPublicKey"`
	AllowedTargetUsers []policies.PolicyTargetUser `json:"allowedTargetUsers"`
}

// ListDatabaseTargets lists all Database targets.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/targets/database
func (s *TargetsService) ListDatabaseTargets(ctx context.Context) ([]DatabaseTarget, *http.Response, error) {
	u := databaseBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	targetList := new([]DatabaseTarget)
	resp, err := s.Client.Do(req, targetList)
	if err != nil {
		return nil, resp, err
	}

	return *targetList, resp, nil
}

// GetDatabaseTarget fetches the specified Database target.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/targets/database/-id-
func (s *TargetsService) GetDatabaseTarget(ctx context.Context, targetID string) (*DatabaseTarget, *http.Response, error) {
	u := fmt.Sprintf(databaseSinglePath, targetID)
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	target := new(DatabaseTarget)
	resp, err := s.Client.Do(req, target)
	if err != nil {
		return nil, resp, err
	}

	return target, resp, nil
}
