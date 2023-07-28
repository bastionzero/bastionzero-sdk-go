package agents

import (
	"context"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/agents/agentstatus"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/agents/agenttype"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/targets"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types"
	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
	"net/http"
)

const (
	agentsBasePath = "api/v2/agents"
)

// AgentsService handles communication with the agents endpoints of the
// BastionZero API.
//
// BastionZero API docs:
//
// https://cloud.bastionzero.com/api/#tag--Agents

type AgentsService client.Service

type AgentDetails struct {
	// Agent type.
	AgentType agenttype.AgentType `json:"type"`
	// Last reported status of the agent.
	AgentStatus agentstatus.AgentStatus `json:"status"`
	// UTC timestamp indicating when the last agent status change event occurred.
	LastStatusUpdate types.Timestamp `json:"lastStatusUpdate"`
	// Version of the agent.
	Version string `json:"version"`
	// The region this agent is registered in.
	Region string `json:"region"`
	// Agent's base64-encoded public key.
	PublicKey string `json:"publicKey"`
	// Information related to the agent's current control channel.
	ControlChannel *targets.ControlChannelSummary `json:"controlChannel"`
	// Unique identifier for the agent. Currently this is null for terminated agents.
	Id string `json:"id"`
	// Name of the agent. Currently this is null for terminated agents.
	Name string `json:"name"`
	// Unique identifier for the agent's environment. Currently this is null for terminated agents.
	EnvironmentId string `json:"environmentId"`
	// Name of the agent's environment. Currently this is null for terminated agents.
	EnvironmentName string `json:"environmentName"`
}

type ListAgentsOptions struct {
	// Only return agents with these types. Defaults to all.
	AgentTypes []agenttype.AgentType `url:"agentTypes,omitempty"`
	// Only return agents with these statuses. Defaults to all except 'Terminated.'
	AgentStatuses []agentstatus.AgentStatus `url:"agentStatuses,omitempty"`
	// Only return agents whose environment contains this substring. Note that terminated agents have no environment.
	EnvironmentName string `url:"environmentName,omitempty"`
	// Only return agents whose name contains this substring. Note that terminated agents have no name.
	Name string `url:"name,omitempty"`
}

// ListAgents
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/agents
func (s *AgentsService) ListAgents(ctx context.Context, opts *ListAgentsOptions) ([]AgentDetails, *http.Response, error) {
	u := agentsBasePath
	u, err := client.AddOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	agentList := new([]AgentDetails)
	resp, err := s.Client.Do(req, agentList)
	if err != nil {
		return nil, resp, err
	}

	return *agentList, resp, nil
}
