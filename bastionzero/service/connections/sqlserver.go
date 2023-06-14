package connections

import (
	"context"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	sqlServerBasePath   = connectionsBasePath + "/sqlserver"
	sqlServerSinglePath = sqlServerBasePath + "/%s"
)

// SQLServerConnection is a connection to an SQL Server
type SQLServerConnection struct {
	Connection

	RemoteHost string `json:"remoteHost"`
	RemotePort int    `json:"remotePort"`
	TargetName string `json:"targetName"`
}

// ListSQLServerConnections lists all SQL Server connections.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/connections/sqlserver
func (s *ConnectionsService) ListSQLServerConnections(ctx context.Context, opts *ListConnectionOptions) ([]SQLServerConnection, *http.Response, error) {
	u := sqlServerBasePath
	u, err := client.AddOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	connList := new([]SQLServerConnection)
	resp, err := s.Client.Do(req, connList)
	if err != nil {
		return nil, resp, err
	}

	return *connList, resp, nil
}

// Ensure SQLServerConnection implementation satisfies the expected interfaces.
var (
	// SQLServerConnection implements ConnectionInterface
	_ ConnectionInterface = &SQLServerConnection{}
)
