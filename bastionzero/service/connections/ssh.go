package connections

import (
	"context"
	"net/http"
)

const (
	sshBasePath   = connectionsBasePath + "/ssh"
	sshSinglePath = sshBasePath + "/%s"
)

// CreateSshConnectionRequest is used to create an ssh connection
type CreateSshConnectionRequest struct {
	TargetID   string `json:"targetId"`
	TargetUser string `json:"targetUser"`
	RemoteHost string `json:"remoteHost"`
	RemotePort int    `json:"remotePort"`
	ScpOnly    bool   `json:"scpOnly"`
}

// CreateSshConnection creates a new ssh connection.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/connections/ssh
func (s *ConnectionsService) CreateSshConnection(ctx context.Context, request *CreateSshConnectionRequest) (*CreateConnectionResponse, *http.Response, error) {
	u := sshBasePath
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
