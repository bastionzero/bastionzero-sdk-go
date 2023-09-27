package connections

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/connections/connectionstate"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/policies/verbtype"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types/targettype"
	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	connectionsBasePath = "api/v2/connections"
)

// ConnectionsService handles communication with the connections endpoints of
// the BastionZero API.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#tag--Connections
type ConnectionsService client.Service

// ListConnectionOptions specifies the optional parameters to various List
// methods involving connections
type ListConnectionOptions struct {
	// ConnectionState filters the state of the connections listed. Defaults to
	// all connection states if not specified.
	ConnectionState connectionstate.ConnectionState `url:"connectionState,omitempty"`

	// UserEmail filters for a specific user's connections. Must be an admin
	// user to use this query parameter.
	UserEmail string `url:"userEmail,omitempty"`
}

// CreateConnectionResponse is the response returned if a connection is
// successfully created when using the target-type specific connect endpoints
type CreateConnectionResponse struct {
	ConnectionID string `json:"connectionId"`
}

// ConnectionInterface lets you work with common attributes from any kind of
// BastionZero connection (excluding DAC connections)
type ConnectionInterface interface {
	// GetID returns the target's unique ID.
	GetID() string
	// GetTimeCreated returns the timestamp of when the connection was created.
	GetTimeCreated() types.Timestamp
	// GetState returns the connection's state.
	GetState() connectionstate.ConnectionState
	// GetTargetID returns the ID of the target connected to.
	GetTargetID() string
	// GetTargetType returns the type of the target connected to.
	GetTargetType() targettype.TargetType
	// GetSubjectID returns the ID of the subject that is using the connection.
	GetSubjectID() string
}

// Connection abstracts common attributes from any kind of BastionZero
// connection (excluding DAC connections)
type Connection struct {
	// ID is the unique ID of the target
	ID string `json:"id"`
	// TimeCreated is the timestamp of when the connection was created
	TimeCreated types.Timestamp `json:"timeCreated"`
	// State is the current state of the connection
	State connectionstate.ConnectionState `json:"state"`
	// TargetID is the ID of the target connected to
	TargetID string `json:"targetID"`
	// TargetType is the type of the target connected to
	TargetType targettype.TargetType `json:"targetType"`
	// SubjectID is the ID of the subject that is using the connection
	SubjectID string `json:"subjectId"`
}

func (c *Connection) GetID() string                             { return c.ID }
func (c *Connection) GetTimeCreated() types.Timestamp           { return c.TimeCreated }
func (c *Connection) GetState() connectionstate.ConnectionState { return c.State }
func (c *Connection) GetTargetID() string                       { return c.TargetID }
func (c *Connection) GetTargetType() targettype.TargetType      { return c.TargetType }
func (c *Connection) GetSubjectID() string                      { return c.SubjectID }

const (
	universalBasePath = connectionsBasePath + "/universal"
)

// CreateUniversalConnectionRequest is used to create a connection to any type
// of target
type CreateUniversalConnectionRequest struct {
	TargetID        string                `json:"targetId,omitempty"`
	TargetName      string                `json:"targetName,omitempty"`
	TargetUser      string                `json:"targetUser,omitempty"`
	EnvironmentID   string                `json:"envId,omitempty"`
	EnvironmentName string                `json:"envName,omitempty"`
	TargetGroups    []string              `json:"targetGroups,omitempty"`
	TargetType      targettype.TargetType `json:"targetType,omitempty"`
	VerbType        verbtype.VerbType     `json:"verbType,omitempty"`
}

// CreateUniversalSshConnectionRequest is used to create a connection to any
// type of target that supports ssh
type CreateUniversalSshConnectionRequest struct {
	TargetID        string `json:"targetId,omitempty"`
	TargetName      string `json:"targetName,omitempty"`
	TargetUser      string `json:"targetUser"`
	RemoteHost      string `json:"remoteHost"`
	RemotePort      int    `json:"remotePort"`
	EnvironmentName string `json:"environmentName,omitempty"`
}

// CreateUniversalConnectionResponse is the response returned if a connection is
// successfully created when using the universal connect endpoints
type CreateUniversalConnectionResponse struct {
	ConnectionID          string                `json:"connectionId"`
	TargetId              string                `json:"targetId"`
	TargetName            string                `json:"targetName"`
	TargetUser            string                `json:"targetUser"`
	TargetType            targettype.TargetType `json:"targetType"`
	VerbType              verbtype.VerbType     `json:"verbType"`
	AgentPublicKey        string                `json:"agentPublicKey"`
	AgentVersion          string                `json:"agentVersion"`
	ConnectionAuthDetails ConnectionAuthDetails `json:"connectionAuthDetails"`
	SshScpOnly            bool                  `json:"sshScpOnly"`
	SplitCert             bool                  `json:"splitCert"`
}

// ConnectionAuthDetails contains details needed to connect to a connection node
// once a connection has been created
type ConnectionAuthDetails struct {
	ConnectionNodeId     string `json:"connectionNodeId"`
	AuthToken            string `json:"authToken"`
	ConnectionServiceUrl string `json:"connectionServiceUrl"`
	Region               string `json:"region"`
}

// CreateUniversalConnection creates a new, non-ssh connection to any type of target.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/connections/universal
func (s *ConnectionsService) CreateUniversalConnection(ctx context.Context, request *CreateUniversalConnectionRequest) (*CreateUniversalConnectionResponse, *http.Response, error) {
	u := universalBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, request)
	if err != nil {
		return nil, nil, err
	}

	createConnResponse := new(CreateUniversalConnectionResponse)
	resp, err := s.Client.Do(req, createConnResponse)
	if err != nil {
		return nil, resp, err
	}

	return createConnResponse, resp, nil
}

// CreateUniversalSshConnection creates a new ssh connection to any type of target that permits ssh.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/connections/universal/ssh
func (s *ConnectionsService) CreateUniversalSshConnection(ctx context.Context, request *CreateUniversalSshConnectionRequest) (*CreateUniversalConnectionResponse, *http.Response, error) {
	u := universalBasePath + "/ssh"
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, request)
	if err != nil {
		return nil, nil, err
	}

	createConnResponse := new(CreateUniversalConnectionResponse)
	resp, err := s.Client.Do(req, createConnResponse)
	if err != nil {
		return nil, resp, err
	}

	return createConnResponse, resp, nil
}

// CloseConnection closes a connection.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#patch-/api/v2/connections/-id-/close
func (s *ConnectionsService) CloseConnection(ctx context.Context, connectionID string) (*http.Response, error) {
	u := connectionsBasePath + fmt.Sprintf("/%s/close", connectionID)
	req, err := s.Client.NewRequest(ctx, http.MethodPatch, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
