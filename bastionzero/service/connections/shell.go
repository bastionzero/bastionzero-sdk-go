package connections

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types/targettype"
)

const (
	shellBasePath   = connectionsBasePath + "/shell"
	shellSinglePath = shellBasePath + "/%s"
)

// CreateShellConnectionRequest is used to create a shell connection
type CreateShellConnectionRequest struct {
	SpaceID    string                `json:"spaceId"`
	TargetID   string                `json:"targetId"`
	TargetType targettype.TargetType `json:"targetType"`
	TargetUser string                `json:"targetUser"`
}

// ShellConnection is a shell connection to a Bzero target
type ShellConnection struct {
	Connection

	SpaceID                   string `json:"spaceId"`
	TargetUser                string `json:"targetUser"`
	SessionRecordingAvailable bool   `json:"sessionRecordingAvailable"`
	SessionRecording          bool   `json:"sessionRecording"`
	InputRecording            bool   `json:"inputRecording"`
}

// CreateShellConnection creates a new shell connection.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/connections/shell
func (s *ConnectionsService) CreateShellConnection(ctx context.Context, request *CreateShellConnectionRequest) (*CreateConnectionResponse, *http.Response, error) {
	u := shellBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, request)
	if err != nil {
		return nil, nil, err
	}

	createConnResponse := new(CreateConnectionResponse)
	resp, err := s.Client.Do(req, createConnResponse)
	if err != nil {
		return nil, resp, err
	}

	return createConnResponse, resp, nil
}

// GetShellConnection fetches the specified shell connection.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/connections/shell/-id-
func (s *ConnectionsService) GetShellConnection(ctx context.Context, connectionID string) (*ShellConnection, *http.Response, error) {
	u := fmt.Sprintf(shellSinglePath, connectionID)
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	conn := new(ShellConnection)
	resp, err := s.Client.Do(req, conn)
	if err != nil {
		return nil, resp, err
	}

	return conn, resp, nil
}

// Ensure ShellConnection implementation satisfies the expected interfaces.
var (
	// ShellConnection implements ConnectionInterface
	_ ConnectionInterface = &ShellConnection{}
)
