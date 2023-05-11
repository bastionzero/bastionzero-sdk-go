package targets

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/policies"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types/targettype"
)

const (
	databaseBasePath   = targetsBasePath + "/database"
	databaseSinglePath = databaseBasePath + "/%s"
)

// CreateDatabaseTargetRequest is used to create a new Database target
type CreateDatabaseTargetRequest struct {
	TargetName      string `json:"targetName"`
	ProxyTargetID   string `json:"proxyTargetId"`
	RemoteHost      string `json:"remoteHost"`
	RemotePort      Port   `json:"remotePort"`
	LocalPort       *Port  `json:"localPort,omitempty"`
	LocalHost       string `json:"localHost,omitempty"`
	IsSplitCert     bool   `json:"splitCert,omitempty"`
	DatabaseType    string `json:"databaseType,omitempty"`
	EnvironmentID   string `json:"environmentId,omitempty"`
	EnvironmentName string `json:"environmentName,omitempty"`
}

// CreateDatabaseTargetRequest is the response returned if a Database target is
// successfully created
type CreateDatabaseTargetResponse struct {
	TargetId string `json:"targetId"`
}

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
	VirtualTarget

	IsSplitCert        bool                  `json:"splitCert"`
	DatabaseType       *string               `json:"databaseType"`
	AllowedTargetUsers []policies.TargetUser `json:"allowedTargetUsers"`
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

// CreateDatabaseTarget creates a new Database target.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/targets/database
func (s *TargetsService) CreateDatabaseTarget(ctx context.Context, request *CreateDatabaseTargetRequest) (*CreateDatabaseTargetResponse, *http.Response, error) {
	u := databaseBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, request)
	if err != nil {
		return nil, nil, err
	}

	createTargetResponse := new(CreateDatabaseTargetResponse)
	resp, err := s.Client.Do(req, createTargetResponse)
	if err != nil {
		return nil, resp, err
	}

	return createTargetResponse, resp, nil
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

// DeleteDatabaseTarget deletes the specified Database target.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#delete-/api/v2/targets/database/-id-
func (s *TargetsService) DeleteDatabaseTarget(ctx context.Context, targetID string) (*http.Response, error) {
	u := fmt.Sprintf(databaseSinglePath, targetID)
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
	// DatabaseTarget implements VirtualTargetInterface
	_ VirtualTargetInterface = &DatabaseTarget{}
)

func (t *DatabaseTarget) GetTargetType() targettype.TargetType { return targettype.Db }
