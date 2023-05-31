package connections

import (
	"context"
	"net/http"
)

const (
	webBasePath   = connectionsBasePath + "/web"
	webSinglePath = webBasePath + "/%s"
)

// CreateWebConnectionRequest is used to create a web connection
type CreateWebConnectionRequest struct {
	TargetID string `json:"targetId"`
}

// CreateWebConnection creates a new web connection.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/connections/web
func (s *ConnectionsService) CreateWebConnection(ctx context.Context, request *CreateWebConnectionRequest) (*CreateConnectionResponse, *http.Response, error) {
	u := webBasePath
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
