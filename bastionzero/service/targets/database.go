package targets

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/policies"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/targets/dbauthconfig"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types/targettype"
	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	databaseBasePath   = targetsBasePath + "/database"
	databaseSinglePath = databaseBasePath + "/%s"
)

// CreateDatabaseTargetRequest is used to create a new Database target
type CreateDatabaseTargetRequest struct {
	TargetName    string `json:"targetName"`
	ProxyTargetID string `json:"proxyTargetId"`
	RemoteHost    string `json:"remoteHost"`
	// RemotePort is required for all databases; however, for GCP-hosted databases, the
	// value specified for Port.Value will be ignored when connecting to the database.
	// TODO: To match REST API, change to: RemotePort *Port  `json:"remotePort,omitempty"` (to be combined with other breaking changes)
	RemotePort Port   `json:"remotePort"`
	LocalPort  *Port  `json:"localPort,omitempty"`
	LocalHost  string `json:"localHost,omitempty"`
	// Deprecated: IsSplitCert exists for historical compatibility and should not be used.
	// Set AuthenticationType in DatabaseAuthenticationConfig appropriately instead.
	IsSplitCert bool `json:"splitCert,omitempty"`
	// Deprecated: DatabaseType exists for historical compatibility and should not be used.
	// Set Database in DatabaseAuthenticationConfig appropriately instead.
	DatabaseType                 string                                     `json:"databaseType,omitempty"`
	EnvironmentID                string                                     `json:"environmentId,omitempty"`
	EnvironmentName              string                                     `json:"environmentName,omitempty"`
	DatabaseAuthenticationConfig *dbauthconfig.DatabaseAuthenticationConfig `json:"databaseAuthenticationConfig,omitempty"`
}

// CreateDatabaseTargetResponse is the response returned if a Database target is
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
	// Deprecated: IsSplitCert exists for historical compatibility and should not be used.
	// Set AuthenticationType in DatabaseAuthenticationConfig appropriately instead.
	IsSplitCert *bool `json:"splitCert,omitempty"`
	// Deprecated: DatabaseType exists for historical compatibility and should not be used.
	// Set Database in DatabaseAuthenticationConfig appropriately instead.
	DatabaseType                 *string                                    `json:"databaseType,omitempty"`
	EnvironmentID                *string                                    `json:"environmentId,omitempty"`
	DatabaseAuthenticationConfig *dbauthconfig.DatabaseAuthenticationConfig `json:"databaseAuthenticationConfig,omitempty"`
}

// ListDatabaseTargetsOptions specifies the optional parameters when querying
// for a list of database targets
type ListDatabaseTargetsOptions struct {
	// TargetNames filters the list of database targets to only those that
	// contain the specified target names.
	TargetNames []string `url:"targetNames,omitempty"`
	// TargetIDs filters the list of database targets to only those that contain the
	// specified target IDs.
	TargetIDs []string `url:"targetIds,omitempty"`
	// EnvironmentName disambiguates conflicting targets that have the same
	// target name when using the TargetNames option by filtering the list of
	// database targets to those that are part of the specified environment (by
	// name).
	EnvironmentName string `url:"envName,omitempty"`
	// EnvironmentID disambiguates conflicting targets that have the same target
	// name when using the TargetNames option by filtering the list of database
	// targets to those that are part of the specified environment (by id).
	EnvironmentID string `url:"envId,omitempty"`
}

// ListSplitCertDatabaseTypesResponse is the response returned when querying for
// a list of databases types that have SplitCert support.
type ListSplitCertDatabaseTypesResponse struct {
	Databases []string `json:"databases"`
}

// DatabaseTarget is a virtual target that provides DB access to a remote
// database. The connection is proxied by a Bzero or Cluster target. The remote
// server doesn't necessarily have to be a database as under the hood it is a
// proxied TCP connection.
type DatabaseTarget struct {
	VirtualTarget

	// Deprecated: IsSplitCert exists for historical compatibility and should not be used.
	// Set AuthenticationType in DatabaseAuthenticationConfig appropriately instead.
	IsSplitCert bool `json:"splitCert"`
	// Deprecated: DatabaseType exists for historical compatibility and should not be used.
	// Set Database in DatabaseAuthenticationConfig appropriately instead.
	DatabaseType                 *string                                   `json:"databaseType"`
	AllowedTargetUsers           []policies.TargetUser                     `json:"allowedTargetUsers"`
	DatabaseAuthenticationConfig dbauthconfig.DatabaseAuthenticationConfig `json:"databaseAuthenticationConfig"`
}

// ListDatabaseTargets lists all Database targets.
//
// Use ListDatabaseTargetsWithFilter() to pass optional filters when querying
// for the list of targets.
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

// TODO-Yuval: Remove ListDatabaseTargetsWithFilter() and add opts parameter to
// ListDatabaseTargets() in a batched breaking changes release.

// ListDatabaseTargetsWithFilter lists all Database targets. Pass in a non-nil
// ListDatabaseTargetsOptions struct to filter the list of database targets
// returned.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/targets/database
func (s *TargetsService) ListDatabaseTargetsWithFilter(ctx context.Context, opts *ListDatabaseTargetsOptions) ([]DatabaseTarget, *http.Response, error) {
	u := databaseBasePath
	u, err := client.AddOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

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

// ListSplitCertDatabaseTypes lists all Database types for which SplitCert
// access is supported.
//
// Deprecated: Use ListDatabaseAuthenticationConfigs
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/targets/database/supported-databases
func (s *TargetsService) ListSplitCertDatabaseTypes(ctx context.Context) (*ListSplitCertDatabaseTypesResponse, *http.Response, error) {
	u := databaseBasePath + "/supported-databases"
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	listResp := new(ListSplitCertDatabaseTypesResponse)
	resp, err := s.Client.Do(req, listResp)
	if err != nil {
		return nil, resp, err
	}

	return listResp, resp, nil
}

// ListDatabaseAuthenticationConfigs lists all database authentication configurations supported by BasionZero.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/targets/database/supported-database-configs
func (s *TargetsService) ListDatabaseAuthenticationConfigs(ctx context.Context) ([]dbauthconfig.DatabaseAuthenticationConfig, *http.Response, error) {
	u := databaseBasePath + "/supported-database-configs"
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	dbAuthConfigList := new([]dbauthconfig.DatabaseAuthenticationConfig)
	resp, err := s.Client.Do(req, dbAuthConfigList)
	if err != nil {
		return nil, resp, err
	}

	return *dbAuthConfigList, resp, nil
}

// Ensure DatabaseTarget implementation satisfies the expected interfaces.
var (
	// DatabaseTarget implements VirtualTargetInterface
	_ VirtualTargetInterface = &DatabaseTarget{}
)

func (t *DatabaseTarget) GetTargetType() targettype.TargetType { return targettype.Db }
