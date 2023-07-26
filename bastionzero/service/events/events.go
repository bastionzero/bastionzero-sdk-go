package events

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/events/connectioneventtype"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types/subjecttype"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types/targettype"
	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	eventsBasePath = "api/v2/events"
)

// EventsService handles communication with the events endpoints of the
// BastionZero API.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#tag--Events
type EventsService client.Service

// SubjectEventOptions specifies the optional parameters to filter the subject
// events returned when listing subject events
type SubjectEventOptions struct {
	// SubjectIDs filters the list of events to only those that contain the
	// specified subject IDs.
	SubjectIDs []string `url:"subjectIds,omitempty"`
	// SubjectTypes filters the list of events to only those that contain the
	// specified subject types.
	SubjectTypes []subjecttype.SubjectType `url:"subjectTypes,omitempty"`
	// SubjectNames filters the list of events to only those that contain the
	// specified subject names.
	SubjectNames []string `url:"subjectNames,omitempty"`
	// IsAdmin filters the list of events to only those that were initiated by
	// admins. Otherwise, if false only events initiated by non-admin subjects
	// are returned.
	IsAdmin *bool `url:"isAdmin,omitempty"`
	// IPAddresses filters the list of events to only those that contain the
	// specified IP addresses.
	IpAddresses []string `url:"ipAddresses,omitempty"`
	// StartTimestamp filters the list of events to only those that occurred
	// after the specified time.
	StartTimestamp *types.Timestamp `url:"startTimestamp,omitempty"`
	// EndTimestamp filters the list of events to only those that occurred
	// before the specified time.
	EndTimestamp *types.Timestamp `url:"endTimestamp,omitempty"`
	// EventCount limits the number of events returned to this value.
	EventCount int `url:"eventCount,omitempty"`
	// ServiceActions filters the list of events to only those that contain the
	// specified service actions.
	ServiceActions []string `url:"serviceActions,omitempty"`
	// Evaluation filters the list of events to only those that were approved by
	// policy. Otherwise, if false only events for actions that were not
	// approved by policy are returned.
	Evaluation *bool `url:"evaluation,omitempty"`
}

// SubjectEvent represents a subject event. A subject event describes a service
// action executed by a subject at a specific point in time.
type SubjectEvent struct {
	ID             string                  `json:"id"`
	OrganizationID string                  `json:"organizationId"`
	SubjectID      string                  `json:"subjectId"`
	SubjectType    subjecttype.SubjectType `json:"subjectType"`
	SubjectName    string                  `json:"subjectName"`
	IsAdmin        bool                    `json:"isAdmin"`
	ServiceAction  string                  `json:"serviceAction"`
	Resource       string                  `json:"resource"`
	Evaluation     bool                    `json:"evaluation"`
	Timestamp      types.Timestamp         `json:"timestamp"`
	IPAddress      string                  `json:"ipAddress"`
	Context        string                  `json:"context"`
}

// ConnectionEventOptions specifies the optional parameters to filter the
// connection events returned when listing connection events
type ConnectionEventOptions struct {
	// SubjectIDs filters the list of events to only those that contain the
	// specified subject IDs.
	SubjectIDs []string `url:"subjectIds,omitempty"`
	// SubjectTypes filters the list of events to only those that contain the
	// specified subject types.
	SubjectTypes []subjecttype.SubjectType `url:"subjectTypes,omitempty"`
	// SubjectNames filters the list of events to only those that contain the
	// specified subject names.
	SubjectNames []string `url:"userNames,omitempty"`
	// IsAdmin filters the list of events to only those that were initiated by
	// admins. Otherwise, if false only events initiated by non-admin subjects
	// are returned.
	IsAdmin *bool `url:"isAdmin,omitempty"`
	// IPAddresses filters the list of events to only those that contain the
	// specified IP addresses.
	IpAddresses []string `url:"ipAddresses,omitempty"`
	// StartTimestamp filters the list of events to only those that occurred
	// after the specified time.
	StartTimestamp *types.Timestamp `url:"startTimestamp,omitempty"`
	// EndTimestamp filters the list of events to only those that occurred
	// before the specified time.
	EndTimestamp *types.Timestamp `url:"endTimestamp,omitempty"`
	// EventCount limits the number of events returned to this value.
	EventCount int `url:"eventCount,omitempty"`
	// ConnectionIDs filters the list of events to only those that contain the
	// specified connection IDs.
	ConnectionIDs []string `url:"connectionIds,omitempty"`
	// SpaceIDs filters the list of events to only those that contain the
	// specified space IDs.
	SpaceIDs []string `url:"spaceIds,omitempty"`
	// SpaceNames filters the list of events to only those that contain the
	// specified space names.
	SpaceNames []string `url:"spaceNames,omitempty"`
	// TargetIDs filters the list of events to only those that contain the
	// specified target IDs.
	TargetIDs []string `url:"targetIds,omitempty"`
	// TargetNames filters the list of events to only those that contain the
	// specified target names.
	TargetNames []string `url:"targetNames,omitempty"`
	// TargetTypes filters the list of events to only those that contain the
	// specified target types.
	TargetTypes []targettype.TargetType `url:"targetTypes,omitempty"`
	// TargetUsers filters the list of events to only those that contain the
	// specified target users.
	TargetUsers []string `url:"targetUsers,omitempty"`
	// EnvironmentIDs filters the list of events to only those that contain the
	// specified target environment IDs.
	EnvironmentIDs []string `url:"environmentIds,omitempty"`
	// EnvironmentNames filters the list of events to only those that contain
	// the specified target environment names.
	EnvironmentNames []string `url:"environmentNames,omitempty"`
	// ConnectionEventTypes filters the list of events to only those that
	// contain the specified event types.
	ConnectionEventTypes []connectioneventtype.ConnectionEventType `url:"connectionEventTypes,omitempty"`
	// ConnectionNodeIds filters the list of events to only those that
	// contain the specified connection node ID.
	ConnectionNodeIds []string `url:"connectionNodeIds,omitempty"`
}

// ConnectionEvent represents a connection event. A connection event describes a
// connection-related event against a target executed by a subject at a specific
// point in time.
type ConnectionEvent struct {
	ID             string                  `json:"id"`
	ConnectionID   string                  `json:"connectionId"`
	SubjectID      string                  `json:"subjectId"`
	SubjectType    subjecttype.SubjectType `json:"subjectType"`
	SubjectName    string                  `json:"subjectName"`
	OrganizationID string                  `json:"organizationId"`
	// SessionID is the ID of the space this connection belongs to
	SessionID string `json:"sessionId"`
	// SessionName is the name of the space this connection belongs to
	SessionName string `json:"sessionName"`
	TargetID    string `json:"targetId"`
	// TargetType is the type of connection to this target
	TargetType          string                                  `json:"targetType"`
	TargetName          string                                  `json:"targetName"`
	TargetUser          string                                  `json:"targetUser"`
	EnvironmentID       string                                  `json:"environmentId"`
	EnvironmentName     string                                  `json:"environmentName"`
	Timestamp           types.Timestamp                         `json:"timestamp"`
	ConnectionEventType connectioneventtype.ConnectionEventType `json:"connectionEventType"`
	Reason              string                                  `json:"reason"`
	ConnectionNodeId    string                                  `json:"connectionNodeId"`
}

// CommandEventOptions specifies the optional parameters to filter the command
// events returned when listing command events
type CommandEventOptions struct {
	// SubjectIDs filters the list of events to only those that contain the
	// specified subject IDs.
	SubjectIDs []string `url:"subjectIds,omitempty"`
	// SubjectTypes filters the list of events to only those that contain the
	// specified subject types.
	SubjectTypes []subjecttype.SubjectType `url:"subjectTypes,omitempty"`
	// SubjectNames filters the list of events to only those that contain the
	// specified subject names.
	SubjectNames []string `url:"userNames,omitempty"`
	// IsAdmin filters the list of events to only those that were initiated by
	// admins. Otherwise, if false only events initiated by non-admin subjects
	// are returned.
	IsAdmin *bool `url:"isAdmin,omitempty"`
	// IPAddresses filters the list of events to only those that contain the
	// specified IP addresses.
	IpAddresses []string `url:"ipAddresses,omitempty"`
	// StartTimestamp filters the list of events to only those that occurred
	// after the specified time.
	StartTimestamp *types.Timestamp `url:"startTimestamp,omitempty"`
	// EndTimestamp filters the list of events to only those that occurred
	// before the specified time.
	EndTimestamp *types.Timestamp `url:"endTimestamp,omitempty"`
	// EventCount limits the number of events returned to this value.
	EventCount int `url:"eventCount,omitempty"`
	// ConnectionIDs filters the list of events to only those that contain the
	// specified connection IDs.
	ConnectionIDs []string `url:"connectionIds,omitempty"`
	// SpaceIDs filters the list of events to only those that contain the
	// specified space IDs.
	SpaceIDs []string `url:"spaceIds,omitempty"`
	// SpaceNames filters the list of events to only those that contain the
	// specified space names.
	SpaceNames []string `url:"spaceNames,omitempty"`
	// TargetIDs filters the list of events to only those that contain the
	// specified target IDs.
	TargetIDs []string `url:"targetIds,omitempty"`
	// TargetNames filters the list of events to only those that contain the
	// specified target names.
	TargetNames []string `url:"targetNames,omitempty"`
	// TargetTypes filters the list of events to only those that contain the
	// specified target types.
	TargetTypes []targettype.TargetType `url:"targetTypes,omitempty"`
	// TargetUsers filters the list of events to only those that contain the
	// specified target users.
	TargetUsers []string `url:"targetUsers,omitempty"`
	// EnvironmentIDs filters the list of events to only those that contain the
	// specified target environment IDs.
	EnvironmentIDs []string `url:"environmentIds,omitempty"`
	// EnvironmentNames filters the list of events to only those that contain
	// the specified target environment names.
	EnvironmentNames []string `url:"environmentNames,omitempty"`
	// CommandSearch filters the list of events to only those that contain the
	// specified command text
	CommandSearch string `url:"commandSearch,omitempty"`
}

// CommandEvent represents a command event. A command event describes a command
// that was executed during a shell or kube exec connection
type CommandEvent struct {
	ID           string `json:"id"`
	ConnectionID string `json:"connectionId"`
	TargetID     string `json:"targetId"`
	// TargetType is the type of connection to this target
	TargetType      string                  `json:"targetType"`
	TargetName      string                  `json:"targetName"`
	SubjectID       string                  `json:"subjectId"`
	SubjectType     subjecttype.SubjectType `json:"subjectType"`
	SubjectName     string                  `json:"subjectName"`
	OrganizationID  string                  `json:"organizationId"`
	Timestamp       types.Timestamp         `json:"timestamp"`
	TargetUser      string                  `json:"targetUser"`
	EnvironmentID   string                  `json:"environmentId"`
	EnvironmentName string                  `json:"environmentName"`
	Command         string                  `json:"command"`
}

// KubernetesEvent represents a Kubernetes event. A Kubernetes event describes a
// Kubernetes REST API endpoint or Kube exec command that was executed during a
// kube connection
type KubernetesEvent struct {
	ID                 string                       `json:"id"`
	OrganizationID     string                       `json:"organizationId"`
	CreationDate       types.Timestamp              `json:"creationDate"`
	Role               string                       `json:"role"`
	TargetGroups       []string                     `json:"targetGroups"`
	Endpoints          []KubernetesEventEndpoint    `json:"endpoints"`
	ExecCommands       []KubernetesEventExecCommand `json:"execCommands"`
	KubeEnglishCommand string                       `json:"kubeEnglishCommand"`
	StatusCode         int                          `json:"statusCode"`
	UserID             string                       `json:"userId"`
	ClusterID          string                       `json:"clusterId"`
	TargetName         string                       `json:"targetName"`
	UserEmail          string                       `json:"userEmail"`
}

type KubernetesEventEndpoint struct {
	ID          string          `json:"id"`
	TimeCreated types.Timestamp `json:"timeCreated"`
	Event       string          `json:"event"`
}

type KubernetesEventExecCommand struct {
	ID          string          `json:"id"`
	TimeCreated types.Timestamp `json:"timeCreated"`
	Event       string          `json:"event"`
}

// AgentStatusChangeEventOptions specifies the parameters to query for the agent
// status change events for a specific target
type AgentStatusChangeEventOptions struct {
	// TargetID is the ID of the target to get status change events for.
	// Required.
	TargetID string `url:"targetId"`
	// StartTimestamp filters the list of events to only those that occurred
	// after the specified time. Optional.
	StartTimestamp *types.Timestamp `url:"startTimestamp,omitempty"`
	// EndTimestamp filters the list of events to only those that occurred
	// before the specified time. Optional.
	EndTimestamp *types.Timestamp `url:"endTimestamp,omitempty"`
}

// AgentStatusChangeEvent represents an agent status change event. An agent
// status change event marks a specific point in time when the target's agent's
// status changed.
type AgentStatusChangeEvent struct {
	StatusChange   string          `json:"statusChange"`
	Timestamp      types.Timestamp `json:"timeStamp"`
	Reason         string          `json:"reason"`
	AgentPublicKey string          `json:"agentPublicKey"`
}

// ListSubjectEvents lists all subject events. The events are returned in
// reverse chronological order (most recent first). Pass in a non-nil
// SubjectEventOptions struct to filter what kind of events are returned.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/events/subject
func (s *EventsService) ListSubjectEvents(ctx context.Context, opts *SubjectEventOptions) ([]SubjectEvent, *http.Response, error) {
	u := eventsBasePath + "/subject"
	u, err := client.AddOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	eventList := new([]SubjectEvent)
	resp, err := s.Client.Do(req, eventList)
	if err != nil {
		return nil, resp, err
	}

	return *eventList, resp, nil
}

// ListConnectionEvents lists all connection events. The events are returned in
// reverse chronological order (most recent first). Pass in a non-nil
// ConnectionEventOptions struct to filter what kind of events are returned.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/events/connection
func (s *EventsService) ListConnectionEvents(ctx context.Context, opts *ConnectionEventOptions) ([]ConnectionEvent, *http.Response, error) {
	u := eventsBasePath + "/connection"
	u, err := client.AddOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	eventList := new([]ConnectionEvent)
	resp, err := s.Client.Do(req, eventList)
	if err != nil {
		return nil, resp, err
	}

	return *eventList, resp, nil
}

// ListCommandEvents lists all command events. The events are returned in
// reverse chronological order (most recent first). Pass in a non-nil
// CommandEventOptions struct to filter what kind of events are returned.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/events/command
func (s *EventsService) ListCommandEvents(ctx context.Context, opts *CommandEventOptions) ([]CommandEvent, *http.Response, error) {
	u := eventsBasePath + "/command"
	u, err := client.AddOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	eventList := new([]CommandEvent)
	resp, err := s.Client.Do(req, eventList)
	if err != nil {
		return nil, resp, err
	}

	return *eventList, resp, nil
}

// ListKubernetesEvents lists all Kubernetes events. The events are returned in
// reverse chronological order (most recent first).
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/events/kube
func (s *EventsService) ListKubernetesEvents(ctx context.Context) ([]KubernetesEvent, *http.Response, error) {
	u := eventsBasePath + "/kube"
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	eventList := new([]KubernetesEvent)
	resp, err := s.Client.Do(req, eventList)
	if err != nil {
		return nil, resp, err
	}

	return *eventList, resp, nil
}

// ListAgentStatusChangeEvents lists all agent status events. The events are returned in
// reverse chronological order (most recent first). The opts argument is required.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/events/agent-status-change
func (s *EventsService) ListAgentStatusChangeEvents(ctx context.Context, opts *AgentStatusChangeEventOptions) ([]AgentStatusChangeEvent, *http.Response, error) {
	if opts == nil {
		return nil, nil, fmt.Errorf("invalid arguments: opts is required")
	}
	if opts.TargetID == "" {
		return nil, nil, fmt.Errorf("opts: TargetID is required")
	}

	u := eventsBasePath + "/agent-status-change"
	u, err := client.AddOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	eventList := new([]AgentStatusChangeEvent)
	resp, err := s.Client.Do(req, eventList)
	if err != nil {
		return nil, resp, err
	}

	return *eventList, resp, nil
}
