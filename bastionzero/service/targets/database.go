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
	databaseBasePath   = targetsBasePath + "/database"
	databaseSinglePath = databaseBasePath + "/%s"
)

// ModifyDatabaseTargetRequest is used to modify a Database target
type ModifyDatabaseTargetRequest struct {
	TargetName    *string `json:"targetName,omitempty"`
	ProxyTargetID *string `json:"proxyTargetId,omitempty"`
	RemoteHost    *string `json:"remoteHost,omitempty"`
	RemotePort    *Port   `json:"remotePort,omitempty"`
	LocalPort     *Port   `json:"localPort,omitempty"`
	LocalHost     *string `json:"localHost,omitempty"`
	IsSplitCert   *bool   `json:"splitCert,omitempty"`
	DatabaseType  *string `json:"databaseType,omitempty"`
	EnvironmentID *string `json:"environmentId,omitempty"`
}

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

// ModifyDatabaseTarget updates a Database target.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#patch-/api/v2/targets/database/-id-
func (s *TargetsService) ModifyDatabaseTarget(ctx context.Context, targetID string, request *ModifyDatabaseTargetRequest) (*DatabaseTarget, *http.Response, error) {
	u := fmt.Sprintf(databaseSinglePath, targetID)
	req, err := s.Client.NewRequest(ctx, http.MethodPatch, u, request)
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

// Ensure DatabaseTarget implementation satisfies the expected interfaces.
var (
	// DatabaseTarget implements TargetInterface
	_ TargetInterface = &DatabaseTarget{}
	// DatabaseTarget implements VirtualTargetInterface
	_ VirtualTargetInterface = &DatabaseTarget{}
)

func (t *DatabaseTarget) GetID() string                        { return t.ID }
func (t *DatabaseTarget) GetName() string                      { return t.Name }
func (t *DatabaseTarget) GetStatus() targetstatus.TargetStatus { return t.Status }
func (t *DatabaseTarget) GetEnvironmentID() string             { return t.EnvironmentID }
func (t *DatabaseTarget) GetLastAgentUpdate() *types.Timestamp { return t.LastAgentUpdate }
func (t *DatabaseTarget) GetAgentVersion() string              { return t.AgentVersion }
func (t *DatabaseTarget) GetRegion() string                    { return t.Region }
func (t *DatabaseTarget) GetAgentPublicKey() string            { return t.AgentPublicKey }
func (t *DatabaseTarget) GetTargetType() targettype.TargetType { return targettype.Db }

func (t *DatabaseTarget) GetProxyTargetID() string { return t.ProxyTargetID }
func (t *DatabaseTarget) GetRemoteHost() string    { return t.RemoteHost }
func (t *DatabaseTarget) GetRemotePort() Port      { return t.RemotePort }
func (t *DatabaseTarget) GetLocalPort() Port       { return t.LocalPort }
