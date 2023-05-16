package connections

import (
	"context"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	dbBasePath   = connectionsBasePath + "/db"
	dbSinglePath = dbBasePath + "/%s"
)

// CreateDbConnectionRequest is used to create a db connection
type CreateDbConnectionRequest struct {
	TargetID   string `json:"targetId"`
	TargetUser string `json:"targetUser,omitempty"`
}

// DbConnection is a connection to a db target
type DbConnection struct {
	Connection

	RemoteHost string `json:"remoteHost"`
	RemotePort int    `json:"remotePort"`
	TargetName string `json:"targetName"`
	TargetUser string `json:"targetUser"`
}

// CreateDbConnection creates a new db connection.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/connections/db
func (s *ConnectionsService) CreateDbConnection(ctx context.Context, request *CreateDbConnectionRequest) (*CreateConnectionResponse, *http.Response, error) {
	u := dbBasePath
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

// ListDbConnections lists all db connections.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/connections/db
func (s *ConnectionsService) ListDbConnections(ctx context.Context, opts *ListConnectionOptions) ([]DbConnection, *http.Response, error) {
	u := dbBasePath
	u, err := client.AddOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	connList := new([]DbConnection)
	resp, err := s.Client.Do(req, connList)
	if err != nil {
		return nil, resp, err
	}

	return *connList, resp, nil
}

// Ensure DbConnection implementation satisfies the expected interfaces.
var (
	// DbConnection implements ConnectionInterface
	_ ConnectionInterface = &DbConnection{}
)
