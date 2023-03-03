package organization

import (
	"context"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	organizationBasePath = "api/v2/organization"
)

// OrganizationService handles communication with the organization endpoints of
// the BastionZero API.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#tag--Organization
type OrganizationService client.Service

// Group is an Identity provider (IdP) group synced to BastionZero. A group
// contains users from an organization.
type Group struct {
	ID   string `json:"idPGroupId"`
	Name string `json:"name"`
}

// ListGroups lists the groups configured (synced from IdP) for your organization.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/organization/groups
func (s *OrganizationService) ListGroups(ctx context.Context) ([]Group, *http.Response, error) {
	u := organizationBasePath + "/groups"
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	groupList := new([]Group)
	resp, err := s.Client.Do(req, groupList)
	if err != nil {
		return nil, resp, err
	}

	return *groupList, resp, nil
}
