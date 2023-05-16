package connections

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/connections/connectionstate"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/connections/datstate"
)

const (
	dynamicAccessBasePath   = connectionsBasePath + "/dynamic-access"
	dynamicAccessSinglePath = dynamicAccessBasePath + "/%s"
)

// DynamicAccessConnection is a dynamic access connection to a Bzero target
type DynamicAccessConnection struct {
	ID                             string                          `json:"id"`
	ConnectionState                connectionstate.ConnectionState `json:"connectionState"`
	DynamicAccessTargetState       datstate.DATState               `json:"dynamicAccessTargetState"`
	ProvisioningServerUniqueId     string                          `json:"provisioningServerUniqueId"`
	ProvisioningServerErrorMessage string                          `json:"provisioningServerErrorMessage"`
}

// GetDynamicAccessConnection fetches the specified dynamic access connection.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/connections/dynamic-access/-id-
func (s *ConnectionsService) GetDynamicAccessConnection(ctx context.Context, connectionID string) (*DynamicAccessConnection, *http.Response, error) {
	u := fmt.Sprintf(dynamicAccessSinglePath, connectionID)
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	conn := new(DynamicAccessConnection)
	resp, err := s.Client.Do(req, conn)
	if err != nil {
		return nil, resp, err
	}

	return conn, resp, nil
}
