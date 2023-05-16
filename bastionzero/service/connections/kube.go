package connections

import (
	"context"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	kubeBasePath   = connectionsBasePath + "/kube"
	kubeSinglePath = kubeBasePath + "/%s"
)

// CreateKubeConnectionRequest is used to create a kube connection
type CreateKubeConnectionRequest struct {
	TargetUser   string   `json:"targetUser"`
	TargetGroups []string `json:"targetGroups"`
	TargetID     string   `json:"targetId"`
}

// KubeConnection is a connection to a Kubernetes target
type KubeConnection struct {
	Connection

	TargetUser   string   `json:"targetUser"`
	TargetGroups []string `json:"targetGroups"`
	TargetName   string   `json:"targetName"`
}

// CreateKubeConnection creates a new kube connection.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/connections/kube
func (s *ConnectionsService) CreateKubeConnection(ctx context.Context, request *CreateKubeConnectionRequest) (*CreateConnectionResponse, *http.Response, error) {
	u := kubeBasePath
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

// ListKubeConnections lists all kube connections.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/connections/kube
func (s *ConnectionsService) ListKubeConnections(ctx context.Context, opts *ListConnectionOptions) ([]KubeConnection, *http.Response, error) {
	u := kubeBasePath
	u, err := client.AddOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	connList := new([]KubeConnection)
	resp, err := s.Client.Do(req, connList)
	if err != nil {
		return nil, resp, err
	}

	return *connList, resp, nil
}

// Ensure KubeConnection implementation satisfies the expected interfaces.
var (
	// KubeConnection implements ConnectionInterface
	_ ConnectionInterface = &KubeConnection{}
)
