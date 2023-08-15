package targets_disambiguated

import (
	"context"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/agents"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/connections/connectionstate"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/targets/targetstatus"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types/targettype"
	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	targetsBasePath = "api/v2/targets"
)

// AllTargetsService handles communication with the disambiguated targets endpoints
// of the BastionZero API.
//
// BastionZero API docs:
//
// - https://cloud.bastionzero.com/api

type AllTargetsService client.Service

type AccessDetails struct {
	JIT                  bool            `json:"jit"`
	AccessExpirationTime types.Timestamp `json:"accessExpirationTime"`
}

// Port describes a port number
type Port struct {
	Value *int `json:"value,omitempty"`
}

// TargetInterface lets you work with common attributes from any kind of
// BastionZero target
type TargetInterface interface {
	// GetID returns the target's unique ID.
	GetID() string
	// GetName returns the target's name.
	GetName() string
	// GetStatus returns the target's status.
	GetStatus() targetstatus.TargetStatus
	// GetEnvironmentID returns the target's environment's ID.
	GetEnvironmentID() string
	// GetLastAgentUpdate returns the target's last agent update if set.
	// Otherwise, returns nil.
	GetLastAgentUpdate() *types.Timestamp
	// GetAgentVersion returns the target's agent version.
	GetAgentVersion() string
	// GetRegion returns the target's region.
	GetRegion() string
	// GetAgentPublicKey returns the target's agent public key.
	GetAgentPublicKey() string
	// GetTargetType returns the target's type.
	GetTargetType() targettype.TargetType
}

// Target abstracts common attributes from any kind of BastionZero target
type Target struct {
	// ID is the unique ID of the target
	ID string `json:"id"`
	// Name is the name of the target
	Name string `json:"name"`
	// Status is the condition of the target
	Status targetstatus.TargetStatus `json:"status"`
	// EnvironmentID is the ID of the environment the target belongs to
	EnvironmentID string `json:"environmentId"`
	// EnvironmentName is the name of the environment the target belongs to
	EnvironmentName string `json:"environmentName"`
	// Agent is the agent associated with this target
	Agent *agents.AgentSummary `json:"agent"`
	// AccessDetails describes how the target can be accessed
	AccessDetails *AccessDetails `json:"accessDetails"`
}

func (t *Target) GetID() string                        { return t.ID }
func (t *Target) GetName() string                      { return t.Name }
func (t *Target) GetStatus() targetstatus.TargetStatus { return t.Status }
func (t *Target) GetEnvironmentID() string             { return t.EnvironmentID }
func (t *Target) GetEnvironmentName() string           { return t.EnvironmentName }
func (t *Target) GetAgent() *agents.AgentSummary       { return t.Agent }
func (t *Target) GetAccessDetails() *AccessDetails     { return t.AccessDetails }

type ListAllTargetsOptions struct {
	// Each target's list of connections will include connections with only these states. Defaults to only Open.
	ConnectionStates connectionstate.ConnectionState `url:"connectionStates,omitempty"`
	// If true, returns all targets accessible to your organization (available to admins only). Defaults to false.
	AllTargetsInOrg bool `url:"allTargetsInOrg"`
	// If included in an admin's request, returns targets accessible to this user.
	UserEmail string `url:"userEmail"`
}

// AllTargetsResponse contains lists of each type of target
type AllTargetsResponse struct {
	Db           []DatabaseTarget
	Kubernetes   []KubeTarget
	FileTransfer []FileTransferTarget
	Rdp          []RDPTarget
	Shell        []ShellTarget
	Ssh          []SSHTarget
	SqlServer    []SQLServerTarget
	Web          []WebTarget
}

// ListAllTargets lists all targets that are accessible to you based on the currently configured policies.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/targets
func (s *AllTargetsService) ListAllTargets(ctx context.Context, opts *ListAllTargetsOptions) (*AllTargetsResponse, *http.Response, error) {
	u := targetsBasePath

	u, err := client.AddOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	allTargets := new(AllTargetsResponse)
	resp, err := s.Client.Do(req, allTargets)
	if err != nil {
		return nil, resp, err
	}

	return allTargets, resp, nil
}
