package targets

import (
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/targets/targetstatus"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types/targettype"
	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	targetsBasePath = "api/v2/targets"
)

// TargetsService handles communication with the targets endpoints of the
// BastionZero API.
//
// BastionZero API docs:
//
// - https://cloud.bastionzero.com/api/#tag--Targets---Bzero-Agent
//
// - https://cloud.bastionzero.com/api/#tag--Targets---Database
//
// - https://cloud.bastionzero.com/api/#tag--Targets---Dynamic
//
// - https://cloud.bastionzero.com/api/#tag--Targets---Kubernetes
//
// - https://cloud.bastionzero.com/api/#tag--Targets---Web
type TargetsService client.Service

// ControlChannelSummary describes a target's currently active control channel
type ControlChannelSummary struct {
	// ControlChannelID is a unique ID that identifies the control channel
	ControlChannelID string `json:"controlChannelId"`
	// ConnectionNodeID is the ID of the connection node the control channel
	// connected to
	ConnectionNodeID string `json:"connectionNodeId"`
	// StartTime is the time the control channel connection was established
	StartTime types.Timestamp `json:"startTime"`
	// EndTime is the time the control channel connection finished. Null if the
	// connection is still active
	EndTime *types.Timestamp `json:"endTime"`
}

// Port describes a port number
type Port struct {
	Value *int `json:"value,omitempty"`
}

// TargetInterface lets you work with common attributes from any kind of
// BastionZero target (excluding DAC targets)
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

// VirtualTargetInterface lets you work with common attributes from any kind of
// BastionZero virtual target (e.g. Db, Web)
type VirtualTargetInterface interface {
	TargetInterface

	// GetProxyTargetID returns the virtual target's proxy target's ID.
	GetProxyTargetID() string
	// GetRemoteHost returns the virtual target's remote host.
	GetRemoteHost() string
	// GetRemotePort returns the virtual target's remote port.
	GetRemotePort() Port
	// GetLocalPort returns the virtual target's local port.
	GetLocalPort() Port
}

// Target abstracts common attributes from any kind of BastionZero target
// (excluding DAC targets)
type Target struct {
	// ID is the unique ID of the target
	ID string `json:"id"`
	// Name is the name of the target
	Name string `json:"name"`
	// Status is the condition of the target
	Status targetstatus.TargetStatus `json:"status"`
	// EnvironmentID is the ID of the environment the target belongs to
	EnvironmentID string `json:"environmentId"`
	// LastAgentUpdate is the timestamp of the last transition change in the
	// target's Status
	LastAgentUpdate *types.Timestamp `json:"lastAgentUpdate"`
	// AgentVersion is the version of the agent running on the target
	AgentVersion string `json:"agentVersion"`
	// Region is the BastionZero region that this target has connected to
	// (follows same naming convention as AWS regions)
	Region string `json:"region"`
	// AgentPublicKey is the public key this target's agent uses when running
	// the MrTAP protocol
	AgentPublicKey string `json:"agentPublicKey"`
}

func (t *Target) GetID() string                        { return t.ID }
func (t *Target) GetName() string                      { return t.Name }
func (t *Target) GetStatus() targetstatus.TargetStatus { return t.Status }
func (t *Target) GetEnvironmentID() string             { return t.EnvironmentID }
func (t *Target) GetLastAgentUpdate() *types.Timestamp { return t.LastAgentUpdate }
func (t *Target) GetAgentVersion() string              { return t.AgentVersion }
func (t *Target) GetRegion() string                    { return t.Region }
func (t *Target) GetAgentPublicKey() string            { return t.AgentPublicKey }

// VirtualTarget abstracts common attributes from any kind of BastionZero
// virtual target (e.g. DB, Web)
type VirtualTarget struct {
	Target

	// ProxyTargetID is the ID of the target that proxies connections made to
	// this virtual target
	ProxyTargetID string `json:"proxyTargetId"`
	// RemoteHost is the IP address of the remote server that is connected to
	// when a connection is made to this virtual target
	RemoteHost string `json:"remoteHost"`
	// RemotePort is the port of the remote server that is connected to when a
	// connection is made to this virtual target
	RemotePort Port `json:"remotePort"`
	// LocalPort is the port of the daemon's localhost server that is spawned on
	// the user's machine on connect. If null, then the zli chooses the port for
	// the user on connect.
	LocalPort Port   `json:"localPort"`
	LocalHost string `json:"localHost"`
}

func (t *VirtualTarget) GetProxyTargetID() string { return t.ProxyTargetID }
func (t *VirtualTarget) GetRemoteHost() string    { return t.RemoteHost }
func (t *VirtualTarget) GetRemotePort() Port      { return t.RemotePort }
func (t *VirtualTarget) GetLocalPort() Port       { return t.LocalPort }
