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
	ControlChannelID string           `json:"controlChannelId"`
	ConnectionNodeID string           `json:"connectionNodeId"`
	StartTime        types.Timestamp  `json:"startTime"`
	EndTime          *types.Timestamp `json:"endTime"`
}

// Port describes a port number
type Port struct {
	Value *int `json:"value,omitempty"`
}

// TargetInterface lets you work with common target attributes from any kind of
// BastionZero target (excluding DAC targets)
type TargetInterface interface {
	// GetID returns the target's unique ID.
	GetID() string
	// GetName returns the the target's name.
	GetName() string
	// GetStatus returns the the target's status.
	GetStatus() targetstatus.TargetStatus
	// GetEnvironmentID returns the the target's environment's ID.
	GetEnvironmentID() string
	// GetLastAgentUpdate returns the target's last agent update if set.
	// Otherwise, returns nil.
	GetLastAgentUpdate() *types.Timestamp
	// GetAgentVersion returns the the target's agent version.
	GetAgentVersion() string
	// GetRegion returns the target's region.
	GetRegion() string
	// GetAgentPublicKey returns the the target's agent public key.
	GetAgentPublicKey() string
	// GetTargetType returns the target's type.
	GetTargetType() targettype.TargetType
}
