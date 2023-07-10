package connections

import (
	"context"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	rdpBasePath   = connectionsBasePath + "/rdp"
	rdpSinglePath = rdpBasePath + "/%s"
)

// RDPConnection is a connection to an RDP target
type RDPConnection struct {
	Connection

	RemoteHost string `json:"remoteHost"`
	RemotePort int    `json:"remotePort"`
	TargetName string `json:"targetName"`
}

// ListRDPConnections lists all RDP connections.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/connections/rdp
func (s *ConnectionsService) ListRDPConnections(ctx context.Context, opts *ListConnectionOptions) ([]RDPConnection, *http.Response, error) {
	u := rdpBasePath
	u, err := client.AddOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	connList := new([]RDPConnection)
	resp, err := s.Client.Do(req, connList)
	if err != nil {
		return nil, resp, err
	}

	return *connList, resp, nil
}

// Ensure RDPConnection implementation satisfies the expected interfaces.
var (
	// RDPConnection implements ConnectionInterface
	_ ConnectionInterface = &RDPConnection{}
)
