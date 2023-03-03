package targets

import (
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service"
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
	ControlChannelID string             `json:"controlChannelId"`
	ConnectionNodeID string             `json:"connectionNodeId"`
	StartTime        service.Timestamp  `json:"startTime"`
	EndTime          *service.Timestamp `json:"endTime"`
}

// Port describes a port number
type Port struct {
	Value *int `json:"value,omitempty"`
}

// TargetStatus represents the status of a target
type TargetStatus string

const (
	// NotActivated represents a target that has not been activated
	NotActivated TargetStatus = "NotActivated"
	// Offline represents a target that is offline
	Offline TargetStatus = "Offline"
	// Online represents a target that is online
	Online TargetStatus = "Online"
	// Terminated represents a target that has been deleted
	Terminated TargetStatus = "Terminated"
	// Error represents a target that has entered an errored state
	Error TargetStatus = "Error"
	// Restarting represents a target that is currently restarting due to
	// receiving a restart request issued via the BastionZero API
	Restarting TargetStatus = "Restarting"
)
